package ui

import (
	l "github.com/mcuadros/OctoPrint-TFT/ui_lang"
	"github.com/gotk3/gotk3/gtk"
	"os"
	"fmt"
	exe "github.com/mcuadros/OctoPrint-TFT/ui_exec"
	
)


var slanguagesPanelInstance *slanguagesPanel

type slanguagesPanel struct {
	CommonPanel
	cb	*gtk.ComboBoxText
	// rb *gtk.RadioButton
	// buttons []gtk.IWidget
	// groupLang *glib.SList
}

func SLanguagesPanel(ui *UI, parent Panel) Panel {
	if slanguagesPanelInstance == nil {
		m := &slanguagesPanel{CommonPanel: NewCommonPanel(ui, parent)}
		m.initialize()
		slanguagesPanelInstance = m
		
	}
	
	return slanguagesPanelInstance
}


func (m *slanguagesPanel) initialize() {
	defer m.Initialize()

	var langs = l.GetLanguagesList()
	for i := 0; i < len(langs); i++ {
		m.Grid().Attach(m.createChangeLangButton(langs[i]), i, 0, 1, 1)
	}
}

func (m *slanguagesPanel) createChangeLangButton(lang string) *gtk.Button {
	return MustButtonImage(l.Translate(lang), fmt.Sprintf("%s.svg", lang), func() {
		l.CurrentLang = lang
		exe.Conf.Lang = l.CurrentLang
		Logger.Warningf("Selected language: %s", l.CurrentLang)
		exe.SaveConfig()
		os.Exit(3)
	
	})
}

func (m *slanguagesPanel) createMainBox() *gtk.Box {
	grid := MustGrid()
	grid.SetHExpand(true)

	box := MustBox(gtk.ORIENTATION_VERTICAL, 3)
	box.SetVAlign(gtk.ALIGN_CENTER)
	box.SetVExpand(true)
	box.Add(grid)
	// box.Add(m.createComboBox())

	return box
}





