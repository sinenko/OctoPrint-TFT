package ui

import (
	l "github.com/mcuadros/OctoPrint-TFT/ui_lang"
)

var defaultPanelInstance *defaultPanel

type defaultPanel struct {
	CommonPanel
}

func DefaultPanel(ui *UI) Panel {
	if defaultPanelInstance == nil {
		m := &defaultPanel{CommonPanel: NewCommonPanel(ui, nil)}
		m.initialize()
		defaultPanelInstance = m
	}

	return defaultPanelInstance
}

func (m *defaultPanel) initialize() {
	m.Grid().Attach(MustButtonImage(l.Translate("Status"), "status.svg", m.showStatus), 1, 0, 1, 1)
	m.Grid().Attach(MustButtonImage(l.Translate("Heat Up"), "heat-up.svg", m.showTemperature), 2, 0, 1, 1)
	m.Grid().Attach(MustButtonImage(l.Translate("Move"), "move.svg", m.showMove), 3, 0, 1, 1)
	m.Grid().Attach(MustButtonImage(l.Translate("Home"), "home.svg", m.showHome), 4, 0, 1, 1)
	m.Grid().Attach(MustButtonImage(l.Translate("Settings"), "settings.svg", m.showSettings), 5, 0, 1, 1)
	m.Grid().Attach(MustButtonImage(l.Translate("Filament"), "filament.svg", m.showFilament), 1, 1, 1, 1)
	m.Grid().Attach(MustButtonImage(l.Translate("Control"), "control.svg", m.showControl), 2, 1, 1, 1)
	m.Grid().Attach(MustButtonImage(l.Translate("Files"), "files.svg", m.showFiles), 3, 1, 1, 1)
	m.Grid().Attach(MustButtonImage(l.Translate("System"), "info.svg", m.showSystem), 4, 1, 1, 1)
}

func (m *defaultPanel) showStatus() {
	m.UI.Add(StatusPanel(m.UI, m))
}

func (m *defaultPanel) showHome() {
	m.UI.Add(HomePanel(m.UI, m))
}

func (m *defaultPanel) showTemperature() {
	m.UI.Add(TemperaturePanel(m.UI, m))
}

func (m *defaultPanel) showFilament() {
	m.UI.Add(FilamentPanel(m.UI, m))
}

func (m *defaultPanel) showMove() {
	m.UI.Add(MovePanel(m.UI, m))
}

func (m *defaultPanel) showControl() {
	m.UI.Add(ControlPanel(m.UI, m))
}

func (m *defaultPanel) showFiles() {
	m.UI.Add(FilesPanel(m.UI, m))
}

func (m *defaultPanel) showSystem() {
	m.UI.Add(SystemPanel(m.UI, m))
}

func (m *defaultPanel) showSettings() {
	m.UI.Add(SettingsPanel(m.UI, m))
}
