package ui

import (
	"strings"

	"github.com/gotk3/gotk3/gtk"
	"github.com/mcuadros/go-octoprint"
	"github.com/mcuadros/OctoPrint-TFT/ui_lang"
	exe "github.com/mcuadros/OctoPrint-TFT/ui_exec"
)

var movePanelInstance *movePanel

type movePanel struct {
	CommonPanel
	step *StepButton
}

func MovePanel(ui *UI, parent Panel) Panel {
	if movePanelInstance == nil {
		m := &movePanel{CommonPanel: NewCommonPanel(ui, parent)}
		m.initialize()
		movePanelInstance = m
	}

	return movePanelInstance
}

func (m *movePanel) initialize() {
	defer m.Initialize()
	
	
	m.Grid().Attach(m.createMainBox(), 1, 0, 4, 3)
	
	m.Grid().Attach(m.createMoveButton(ui_lang.Translate("X-"), "move-x-.svg", octoprint.XAxis, -1), 1, 1, 1, 1)
	
	m.Grid().Attach(m.createMoveButton(ui_lang.Translate("Y+"), "move-y+.svg", octoprint.YAxis, 1), 2, 0, 1, 1)
	m.Grid().Attach(m.createMoveButton(ui_lang.Translate("Y-"), "move-y-.svg", octoprint.YAxis, -1), 2, 2, 1, 1)
	
	
	m.Grid().Attach(m.createMoveButton(ui_lang.Translate("X+"), "move-x+.svg", octoprint.XAxis, 1), 3, 1, 1, 1)
	
	
	
	m.Grid().Attach(m.createMoveButton(ui_lang.Translate("Z-"), "move-z+.svg", octoprint.ZAxis, -1), 5, 0, 1, 1)
	m.Grid().Attach(m.createMoveButton(ui_lang.Translate("Z+"), "move-z-.svg", octoprint.ZAxis, 1), 5, 1, 1, 1)
	
	m.step = MustStepButton("move-step.svg",
		Step{ui_lang.Translate("5mm"), 5}, Step{ui_lang.Translate("10mm"), 10},  Step{ui_lang.Translate("50mm"), 50}, Step{ui_lang.Translate("1mm"), 1}, 
	)
	m.Grid().Attach(m.step, 2, 1, 1, 1)
}

func (m *movePanel) createMainBox() *gtk.Box {
	grid := MustGrid()
	grid.SetHExpand(true)

	box := MustBox(gtk.ORIENTATION_VERTICAL, 3)
	box.SetVAlign(gtk.ALIGN_CENTER)
	box.SetVExpand(true)
	box.Add(grid)

	return box
}

func (m *movePanel) createMoveButton(label, image string, a octoprint.Axis, dir int) gtk.IWidget {
	return MustButtonImage(label, image, func() {
		if(exe.Vars.IsPrinting)  { 
			Logger.Error(ui_lang.Translate("Error: printer is busy"))
			return 
		}
		distance := m.step.Value().(int) * dir

		cmd := &octoprint.PrintHeadJogRequest{}
		switch a {
		case octoprint.XAxis:
			cmd.X = distance
		case octoprint.YAxis:
			cmd.Y = distance
		case octoprint.ZAxis:
			cmd.Z = distance
		}

		Logger.Warningf(ui_lang.Translate("Jogging print head axis %s in %dmm"),
			strings.ToUpper(string(a)), distance,
		)

		if err := cmd.Do(m.UI.Printer); err != nil {
			Logger.Error(err)
			return
		}
	})
}
