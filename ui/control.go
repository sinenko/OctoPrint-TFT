package ui

import (
	"fmt"
	"strings"
	"math"
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/mcuadros/go-octoprint"
	"github.com/mcuadros/OctoPrint-TFT/ui_lang"
	exe "github.com/mcuadros/OctoPrint-TFT/ui_exec"
)

var control = []*octoprint.ControlDefinition{{
	Name:    "Motor Off",
	Command: "M18",
}, {
	Name:    "LED On",
	Command: "M355 S1",
}, {
	Name:    "LED Off",
	Command: "M355 S0",
}}

var controlPanelInstance *controlPanel

type controlPanel struct {
	CommonPanel
	fanSpeed int
	isSpeedChanged bool
	speedLabel *LabelWithImage
}

func ControlPanel(ui *UI, parent Panel) Panel {
	if controlPanelInstance == nil {
		m := &controlPanel{CommonPanel: NewCommonPanel(ui, parent)}
		m.initialize()
		controlPanelInstance = m
	}
	return controlPanelInstance
}

func (m *controlPanel) initialize() {
	m.fanSpeed = 215
	m.isSpeedChanged = false
	defer m.Initialize()
		
	m.Grid().Attach(m.createCalibrateButton(), 1, 0, 1, 1)
	var i, k int
	i = 2
	k = 0
	for _, c := range m.getControl() {
		b := m.createControlButton(c)
		m.Grid().Attach(b, i, k, 1, 1)
		if i == 5 {
			i = 0
			k++
		}
		i++
	}
	
	m.Grid().Attach(m.createFanSlowButton(), 1, 1, 1, 1)
	m.Grid().Attach(m.createSpeedBox(), 2, 1, 1, 1)
	m.Grid().Attach(m.createFanFastButton(), 3, 1, 1, 1)
	m.Grid().Attach(m.createPressDownButton(), 1, 2, 1, 1)
	m.Grid().Attach(m.createPushUpButton(), 2, 2, 1, 1)
	
}

func (m *controlPanel) getControl() []*octoprint.ControlDefinition {
	control := control

	Logger.Info(ui_lang.Translate("Retrieving custom controls"))
	r, err := (&octoprint.CustomCommandsRequest{}).Do(m.UI.Printer)
	if err != nil {
		Logger.Error(err)
		return control
	}

	for _, c := range r.Controls {
		control = append(control, c.Children...)
	}

	return control
}

func (m *controlPanel) createCalibrateButton() gtk.IWidget {
	do := func() {
		if(exe.Vars.IsPrinting) { 
			Logger.Error(ui_lang.Translate("Error: printer is busy"))
			return 
		}
	
		r := &octoprint.CommandRequest{}
		r.Commands = exe.Conf.Calibrating
		
		if err := r.Do(m.UI.Printer); err != nil {
			Logger.Error(err)
			return
		}
	}

	return MustButtonImage(ui_lang.Translate("Calibrate"), "extruder.svg", do)
}

func (m *controlPanel) createPressDownButton() gtk.IWidget {
	do := func() {
	
		r := &octoprint.CommandRequest{}
		r.Commands = []string{"M290 Z-0.05", "M500"}
		
		if err := r.Do(m.UI.Printer); err != nil {
			Logger.Error(err)
			return
		}
	}

	return MustButtonImage(ui_lang.Translate("Press down"), "pressdown.svg", do)
}

func (m *controlPanel) createPushUpButton() gtk.IWidget {
	do := func() {
	
		r := &octoprint.CommandRequest{}
		r.Commands = []string{"M290 Z0.05", "M500"}
		
		if err := r.Do(m.UI.Printer); err != nil {
			Logger.Error(err)
			return
		}
	}

	return MustButtonImage(ui_lang.Translate("Push up"), "pushup.svg", do)
}


func (m *controlPanel) createSpeedBox() *gtk.Box {


	m.speedLabel = MustLabelWithImage("", "")
	
	if(m.isSpeedChanged){
		m.speedLabel.Label.SetLabel(fmt.Sprintf(ui_lang.Translate("Speed: %d%%"), int(math.Ceil (float64(m.fanSpeed) / 25) * 10)))
	}

	info := MustBox(gtk.ORIENTATION_VERTICAL, 0)
	info.SetHExpand(true)
	info.SetVExpand(true)
	
	info.Add(m.speedLabel)
	info.SetMarginTop(50)
	info.SetMarginStart(10)

	return info
}

func (m *controlPanel) createFanFastButton() gtk.IWidget {
	do := func() {
		m.isSpeedChanged = true
		
		if m.fanSpeed >= 255 {
			return
		}
	
		m.fanSpeed += 8
		if m.fanSpeed >= 250 {
			m.fanSpeed = 255
		}
		
		m.speedLabel.Label.SetLabel(fmt.Sprintf(ui_lang.Translate("Speed: %d%%"), int(math.Floor ((float64(m.fanSpeed)-175) / 8) * 10)))
		r := &octoprint.CommandRequest{}
		
		if(m.fanSpeed <= 183 || !m.isSpeedChanged) { 
			r.Commands = []string{"M106 S255",}
			if err := r.Do(m.UI.Printer); err != nil {
				Logger.Error(err)
				return
			}
			time.Sleep(time.Duration(400)*time.Millisecond)
			r.Commands = []string{fmt.Sprintf("M106 S%d", m.fanSpeed),}
			if err := r.Do(m.UI.Printer); err != nil {
				Logger.Error(err)
				return
			}
		
		} else {
			r.Commands = []string{fmt.Sprintf("M106 S%d", m.fanSpeed)}
			
		
			if err := r.Do(m.UI.Printer); err != nil {
				Logger.Error(err)
				return
			}
		}
	}

	return MustButtonImage(ui_lang.Translate("Fan +"), "fan-on.svg", do)
}



func (m *controlPanel) createFanSlowButton() gtk.IWidget {
	do := func() {
	
		m.isSpeedChanged = true
	
		if m.fanSpeed <= 175 {
			return
		}
	
		m.fanSpeed -= 8
		if m.fanSpeed <= 180 {
			m.fanSpeed = 175
		}
		m.speedLabel.Label.SetLabel(fmt.Sprintf(ui_lang.Translate("Speed: %d%%"), int(math.Ceil ((float64(m.fanSpeed)-175) / 8) * 10)))
		r := &octoprint.CommandRequest{}
		if(m.fanSpeed <= 175){
			r.Commands = []string{fmt.Sprintf("M106 S%d", 0)}
		} else {
			r.Commands = []string{fmt.Sprintf("M106 S%d", m.fanSpeed)}
		}
		
		
		if err := r.Do(m.UI.Printer); err != nil {
			Logger.Error(err)
			return
		}
	}

	return MustButtonImage(ui_lang.Translate("Fan -"), "fan-off.svg", do)
}

func (m *controlPanel) createControlButton(c *octoprint.ControlDefinition) gtk.IWidget {
	
	icon := strings.ToLower(strings.Replace(c.Name, " ", "-", -1))
	do := func() {
		if(exe.Vars.IsPrinting && c.Name == "Motor Off") { 
			Logger.Error(ui_lang.Translate("Error: printer is busy"))
			return 
		}
		r := &octoprint.CommandRequest{
			Commands: c.Commands,
		}

		if len(c.Command) != 0 {
			r.Commands = []string{c.Command}
		}

		Logger.Infof(ui_lang.Translate("Executing command %q"), ui_lang.Translate(c.Name))
		if err := r.Do(m.UI.Printer); err != nil {
			Logger.Error(err)
			return
		}
	}

	cb := do
	if len(c.Confirm) != 0 {
		cb = MustConfirmDialog(m.UI.w, c.Confirm, do)
	}

	return MustButtonImage(ui_lang.Translate(c.Name), icon+".svg", cb)
}

func (m *controlPanel) createCommandButton(c *octoprint.CommandDefinition) gtk.IWidget {
	do := func() {
		r := &octoprint.SystemExecuteCommandRequest{
			Source: octoprint.Custom,
			Action: c.Action,
		}

		if err := r.Do(m.UI.Printer); err != nil {
			Logger.Error(err)
			return
		}
	}

	cb := do
	if len(c.Confirm) != 0 {
		cb = MustConfirmDialog(m.UI.w, c.Confirm, do)
	}

	return MustButtonImage(ui_lang.Translate(c.Name), c.Action+".svg", cb)
}

// func (m *controlPanel) getCommands() []*octoprint.CommandDefinition {
	// Logger.Info(ui_lang.Translate("Retrieving custom commands"))
	// r, err := (&octoprint.SystemCommandsRequest{}).Do(m.UI.Printer)
	// if err != nil {
		// Logger.Error(err)
		// return nil
	// }

	// return r.Custom
// }
