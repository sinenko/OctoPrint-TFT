package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/mcuadros/go-octoprint"
	"github.com/mcuadros/OctoPrint-TFT/ui_lang"
	exe "github.com/mcuadros/OctoPrint-TFT/ui_exec"
)

var homePanelInstance *homePanel

type homePanel struct {
	CommonPanel
}

func HomePanel(ui *UI, parent Panel) Panel {
	if homePanelInstance == nil {
		m := &homePanel{CommonPanel: NewCommonPanel(ui, parent)}
		m.initialize()
		homePanelInstance = m
	}

	return homePanelInstance
}

func (m *homePanel) initialize() {
	defer m.Initialize()
	
	m.Grid().Attach(m.createMainBox(), 1, 0, 4, 3)
	m.Grid().Attach(m.createMoveButton(ui_lang.Translate("Home X"), "home-x.svg", octoprint.XAxis), 1, 0, 1, 1)
	m.Grid().Attach(m.createMoveButton(ui_lang.Translate("Home Y"), "home-y.svg", octoprint.YAxis), 2, 0, 1, 1)
	m.Grid().Attach(m.createHomeAllButton(), 1, 1, 2, 1)
}

func (m *homePanel) createMainBox() *gtk.Box {
	grid := MustGrid()
	grid.SetHExpand(true)

	box := MustBox(gtk.ORIENTATION_VERTICAL, 3)
	box.SetVAlign(gtk.ALIGN_CENTER)
	box.SetVExpand(true)
	box.Add(grid)

	return box
}

func (m *homePanel) createHomeAllButton() gtk.IWidget {
	do := func() {
		if(exe.Vars.IsPrinting)  { 
			Logger.Error(ui_lang.Translate("Error: printer is busy"))
			return 
		}
		r := &octoprint.CommandRequest{}
		r.Commands = exe.Conf.HomeAll
		
		if err := r.Do(m.UI.Printer); err != nil {
			Logger.Error(err)
			return
		}
	}

	return MustButtonImage(ui_lang.Translate("Home All"), "home.svg", do)
}

func (m *homePanel) createMoveButton(label, image string, axes ...octoprint.Axis) gtk.IWidget {
	return MustButtonImage(label, image, func() {
		if(exe.Vars.IsPrinting)  {
			Logger.Error(ui_lang.Translate("Error: printer is busy"))
			return 
		}
		cmd := &octoprint.PrintHeadHomeRequest{Axes: axes}
		Logger.Warningf(ui_lang.Translate("Homing the print head in %s axes"), axes)
		if err := cmd.Do(m.UI.Printer); err != nil {
			Logger.Error(err)
			return
		}
	})
}
