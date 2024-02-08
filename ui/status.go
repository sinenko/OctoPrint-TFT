package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/mcuadros/go-octoprint"
	"github.com/mcuadros/OctoPrint-TFT/ui_lang"
	exe "github.com/mcuadros/OctoPrint-TFT/ui_exec"
)

var statusPanelInstance *statusPanel

type statusPanel struct {
	CommonPanel
	step *StepButton
	pb   *gtk.ProgressBar

	bed, tool0, tool1  *LabelWithImage
	file, left         *LabelWithImage
	speedLabel				 *LabelWithImage
	print, pause, stop *gtk.Button
	isPauseIcon bool
	currentSpeed int
}

func StatusPanel(ui *UI, parent Panel) Panel {
	if statusPanelInstance == nil {
		m := &statusPanel{CommonPanel: NewCommonPanel(ui, parent)}
		m.b = NewBackgroundTask(time.Second*5, m.update)
		m.initialize()

		statusPanelInstance = m
	}

	return statusPanelInstance
}

func (m *statusPanel) initialize() {
	m.currentSpeed = 100
	defer m.Initialize()

	m.Grid().Attach(m.createMainBox(), 1, 0, 5, 1)
	m.Grid().Attach(m.createPrintButton(), 1, 2, 1, 1)
	m.Grid().Attach(m.createPauseButton(), 2, 2, 1, 1)
	m.Grid().Attach(m.createStopButton(), 3, 2, 1, 1)
	m.Grid().Attach(m.createSlowButton(), 1, 3, 1, 1)
	m.Grid().Attach(m.createFastButton(), 3, 3, 1, 1)
	m.Grid().Attach(m.createSpeedBox(), 2, 3, 1, 1)
}

func (m *statusPanel) createSpeedBox() *gtk.Box {
	m.speedLabel = MustLabelWithImage("", "")
	m.speedLabel.Label.SetLabel(fmt.Sprintf(ui_lang.Translate("Speed: %d%%"), m.currentSpeed))

	info := MustBox(gtk.ORIENTATION_VERTICAL, 0)
	info.SetHExpand(true)
	info.SetVExpand(true)
	
	info.Add(m.speedLabel)
	info.SetMarginTop(50)
	info.SetMarginStart(10)

	return info
}

func (m *statusPanel) createProgressBar() *gtk.ProgressBar {
	m.pb = MustProgressBar()
	m.pb.SetShowText(true)
	m.pb.SetMarginTop(10)
	m.pb.SetMarginStart(10)
	m.pb.SetMarginEnd(10)

	return m.pb
}

func (m *statusPanel) createMainBox() *gtk.Box {
	grid := MustGrid()
	grid.SetHExpand(true)
	grid.Add(m.createInfoBox())
	grid.Add(m.createTemperatureBox())

	box := MustBox(gtk.ORIENTATION_VERTICAL, 5)
	box.SetVAlign(gtk.ALIGN_CENTER)
	box.SetVExpand(true)
	box.Add(grid)
	box.Add(m.createProgressBar())

	return box
}

func (m *statusPanel) createInfoBox() *gtk.Box {
	m.file = MustLabelWithImage("file.svg", "")
	m.left = MustLabelWithImage("speed-step.svg", "")

	info := MustBox(gtk.ORIENTATION_VERTICAL, 5)
	info.SetHAlign(gtk.ALIGN_START)
	info.SetHExpand(true)
	info.SetVExpand(true)
	info.Add(m.file)
	info.Add(m.left)
	info.SetMarginStart(10)

	return info
}

func (m *statusPanel) createTemperatureBox() *gtk.Box {
	m.bed = MustLabelWithImage("bed.svg", "")
	m.tool0 = MustLabelWithImage("extruder.svg", "")
	m.tool1 = MustLabelWithImage("extruder.svg", "")

	temp := MustBox(gtk.ORIENTATION_VERTICAL, 5)
	temp.SetHAlign(gtk.ALIGN_START)
	temp.SetHExpand(true)
	temp.SetVExpand(true)
	temp.Add(m.bed)
	temp.Add(m.tool0)
	temp.Add(m.tool1)

	return temp
}

func (m *statusPanel) createPrintButton() gtk.IWidget {
	m.print = MustButtonImage(ui_lang.Translate("Print"), "status.svg", func() {
		defer m.updateTemperature()

		Logger.Warning(ui_lang.Translate("Starting a new job"))
		if err := (&octoprint.StartRequest{}).Do(m.UI.Printer); err != nil {
			Logger.Error(err)
			return
		}
	})

	return m.print
}

func (m *statusPanel) createPauseButton() gtk.IWidget {
	m.pause = MustButtonImage(ui_lang.Translate("Pause"), "pause.svg", func() {
		defer m.updateTemperature()

		Logger.Warning(ui_lang.Translate("Pausing/Resuming job"))
		cmd := &octoprint.PauseRequest{Action: octoprint.Toggle}
		if err := cmd.Do(m.UI.Printer); err != nil {
			Logger.Error(err)
			return
		}
	})

	return m.pause
}

func (m *statusPanel) createStopButton() gtk.IWidget {
	m.stop = MustButtonImage(ui_lang.Translate("Stop"), "stop.svg", func() {
		defer m.updateTemperature()

		Logger.Warning(ui_lang.Translate("Stopping job"))
		if err := (&octoprint.CancelRequest{}).Do(m.UI.Printer); err != nil {
			Logger.Error(err)
			return
		}
	})

	return m.stop
}

func (m *statusPanel) createSlowButton() gtk.IWidget {
	do := func() {
	
		m.currentSpeed -= 2
		if m.currentSpeed < 50 {
			m.currentSpeed = 50
			return
		}
		m.speedLabel.Label.SetLabel(fmt.Sprintf(ui_lang.Translate("Speed: %d%%"), m.currentSpeed))
		r := &octoprint.CommandRequest{}
		r.Commands = []string{fmt.Sprintf("M220 S%d", m.currentSpeed)}
		
		if err := r.Do(m.UI.Printer); err != nil {
			Logger.Error(err)
			return
		}
	}

	return MustButtonImage(ui_lang.Translate("Slower"), "decrease.svg", do)
}

func (m *statusPanel) createFastButton() gtk.IWidget {
	do := func() {
	
		m.currentSpeed += 2
		if m.currentSpeed > 200 {
			m.currentSpeed = 200
			return
		}
		m.speedLabel.Label.SetLabel(fmt.Sprintf(ui_lang.Translate("Speed: %d%%"), m.currentSpeed))
		r := &octoprint.CommandRequest{}
		r.Commands = []string{fmt.Sprintf("M220 S%d", m.currentSpeed)}
		
		if err := r.Do(m.UI.Printer); err != nil {
			Logger.Error(err)
			return
		}
	}

	return MustButtonImage(ui_lang.Translate("Faster"), "increase.svg", do)
}





func (m *statusPanel) update() {
	m.updateTemperature()
	m.updateJob()
}

func (m *statusPanel) updateTemperature() {
	s, err := (&octoprint.StateRequest{Exclude: []string{"sd"}}).Do(m.UI.Printer)
	if err != nil {
		Logger.Error(err)
		return
	}

	m.doUpdateState(&s.State)

	m.tool1.Hide()
	for tool, s := range s.Temperature.Current {
		text := fmt.Sprintf("%s: %.0f°C / %.0f°C", ui_lang.Translate(strings.Title(tool)), s.Actual, s.Target)
		switch tool {
		case "bed":
			m.bed.Label.SetLabel(text)
		case "tool0":
			m.tool0.Label.SetLabel(text)
		case "tool1":
			m.tool1.Label.SetLabel(text)
			m.tool1.Show()
		}
	}
}

func (m *statusPanel) doUpdateState(s *octoprint.PrinterState) {
	switch {
	case s.Flags.Printing:
		m.print.SetSensitive(false)
		if(!m.isPauseIcon){
			m.pause.SetLabel(ui_lang.Translate("Pause"))
			m.pause.SetImage(MustImageFromFile("pause.svg"))
			m.isPauseIcon = true
		}
		if(!m.isPauseIcon && !exe.Vars.IsPrinting){
			exe.Vars.IsPrinting = true
		}
		m.pause.SetSensitive(true)
		m.stop.SetSensitive(true)
	case s.Flags.Paused:
		m.print.SetSensitive(false)
		if(m.isPauseIcon){
			m.pause.SetLabel(ui_lang.Translate("Resume"))
			m.pause.SetImage(MustImageFromFile("resume.svg"))
			m.isPauseIcon = false
		}
		if(m.isPauseIcon && exe.Vars.IsPrinting){
			exe.Vars.IsPrinting = false
		}
		m.pause.SetSensitive(true)
		m.stop.SetSensitive(true)
		return
	case s.Flags.Ready:
		m.print.SetSensitive(true)
		m.pause.SetSensitive(false)
		m.stop.SetSensitive(false)
		m.isPauseIcon = true
		exe.Vars.IsPrinting = false
	default:
		m.print.SetSensitive(false)
		m.pause.SetSensitive(false)
		m.stop.SetSensitive(false)
		m.isPauseIcon = true
		exe.Vars.IsPrinting = false
	}

}

func (m *statusPanel) updateJob() {
	s, err := (&octoprint.JobRequest{}).Do(m.UI.Printer)
	if err != nil {
		Logger.Error(err)
		return
	}

	file := ui_lang.Translate("<i>not-set</i>")
	if s.Job.File.Name != "" {
		file = filenameEllipsis(s.Job.File.Name)
	}

	m.file.Label.SetLabel(fmt.Sprintf(ui_lang.Translate("File: %s"), file))
	m.pb.SetFraction(s.Progress.Completion / 100)

	if m.UI.State.IsOperational() {
		m.left.Label.SetLabel(ui_lang.Translate("Printer is ready"))
		
		return
	}

	var text string
	switch s.Progress.Completion {
	case 100:
		text = fmt.Sprintf(ui_lang.Translate("Job Completed in %s"), time.Duration(int64(s.Job.LastPrintTime)*1e9))
	case 0:
		text = ui_lang.Translate("Warming up ...")
	default:
		e := time.Duration(int64(s.Progress.PrintTime) * 1e9)
		l := time.Duration(int64(s.Progress.PrintTimeLeft) * 1e9)
		text = fmt.Sprintf(ui_lang.Translate("Elapsed/Left: %s / %s"), e, l)
		if l == 0 {
			text = fmt.Sprintf(ui_lang.Translate("Elapsed: %s"), e)
		}
	}

	m.left.Label.SetLabel(text)
}

func filenameEllipsis(name string) string {
	if len(name) > 26 {
		return name[:23] + "..."
	}

	return name
}
