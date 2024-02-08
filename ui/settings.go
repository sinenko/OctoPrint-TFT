package ui

import (
	
	l "github.com/mcuadros/OctoPrint-TFT/ui_lang"
	"github.com/gotk3/gotk3/gtk"
	exe "github.com/mcuadros/OctoPrint-TFT/ui_exec"
)

var settingsPanelInstance *settingsPanel

type settingsPanel struct {
	CommonPanel
}

func SettingsPanel(ui *UI, parent Panel) Panel {
	if settingsPanelInstance == nil {
		m := &settingsPanel{CommonPanel: NewCommonPanel(ui, parent)}
		m.initialize()
		settingsPanelInstance = m
	}

	return settingsPanelInstance
}

func (m *settingsPanel) initialize() {
	defer m.Initialize()
	m.Grid().Attach(m.createMainBox(), 1, 0, 4, 3)
	m.Grid().Attach(MustButtonImage(l.Translate("Language"), "worldwide.svg", m.showSLanguages), 1, 0, 1, 1)
	m.Grid().Attach(MustButtonImage(l.Translate("Network"), "wifi.svg", m.showSNetwork), 2, 0, 1, 1)
	m.Grid().Attach(MustButtonImage(l.Translate("Update Software"), "softupdate.svg", m.showSUpdateSoftware), 1, 3, 2, 1)
}

func (m *settingsPanel) createMainBox() *gtk.Box {
	grid := MustGrid()
	grid.SetHExpand(true)

	box := MustBox(gtk.ORIENTATION_VERTICAL, 3)
	box.SetVAlign(gtk.ALIGN_CENTER)
	box.SetVExpand(true)
	box.Add(grid)

	return box
}

func (m *settingsPanel) showSLanguages() {
	m.UI.Add(SLanguagesPanel(m.UI, m.p))
}

func (m *settingsPanel) showSNetwork() {
	m.UI.Add(SNetworkPanel(m.UI, m.p))
}

func (m *settingsPanel) showSUpdateSoftware() {

	if(exe.Vars.IsUpdating) {
		return
	}

	status := exe.UpdateSoftware(Version)
	
	if(status == 1) {
		Logger.Warningf(l.Translate("Updates will be instaled within 1 minute"))
	} else if(status == -1) {
		Logger.Warningf(l.Translate("Current version is latest"))
	} else if(status == -2) {
		Logger.Warningf(l.Translate("Update error: md5 sum incorrect"))
	}	else if(status == 0){
		Logger.Error(l.Translate("Update error"))
	} 
	
}