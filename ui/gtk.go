package ui

import (
	"fmt"
	"path/filepath"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	// "github.com/gotk3/gotk3/glib"
)

// MustWindow returns a new gtk.Window, if error panics.
func MustWindow(t gtk.WindowType) *gtk.Window {
	win, err := gtk.WindowNew(t)
	if err != nil {
		panic(err)
	}

	return win
}

// MustGrid returns a new gtk.Grid, if error panics.
func MustGrid() *gtk.Grid {
	grid, err := gtk.GridNew()
	if err != nil {
		panic(err)
	}

	return grid
}

// MustBox returns a new gtk.Box, with the given configuration, if err panics.
func MustBox(o gtk.Orientation, spacing int) *gtk.Box {
	box, err := gtk.BoxNew(o, spacing)
	if err != nil {
		panic(err)
	}

	return box
}

// MustProgressBar returns a new gtk.ProgressBar, if err panics.
func MustProgressBar() *gtk.ProgressBar {
	p, err := gtk.ProgressBarNew()
	if err != nil {
		panic(err)
	}
	ctx, _ := p.GetStyleContext()
	ctx.AddClass("labelbig")

	return p
}

// MustLabel returns a new gtk.Label, if err panics.
func MustLabel(label string, args ...interface{}) *gtk.Label {
	l, err := gtk.LabelNew("")
	if err != nil {
		panic(err)
	}

	ctx, _ := l.GetStyleContext()
	ctx.AddClass("labelbig")
	l.SetMarkup(fmt.Sprintf(label, args...))
	return l
}

// LabelWithImage represents a gtk.Label with a image to the right.
type LabelWithImage struct {
	Label *gtk.Label
	*gtk.Box
}

// LabelImageSize default width and height of the image for a LabelWithImage
const LabelImageSize = 20

// MustLabelWithImage returns a new LabelWithImage based on a gtk.Box containing
// a gtk.Label with a gtk.Image, the image is scaled at LabelImageSize.
func MustLabelWithImage(img, label string, args ...interface{}) *LabelWithImage {
	l := MustLabel(label, args...)
	b := MustBox(gtk.ORIENTATION_HORIZONTAL, 5)
	if(img != ""){
		b.Add(MustImageFromFileWithSize(img, LabelImageSize, LabelImageSize))
	}
	b.Add(l)

	return &LabelWithImage{Label: l, Box: b}
}

// MustButtonImage returns a new gtk.Button with the given label, image and
// clicked callback. If error panics.
func MustButtonImage(label, img string, clicked func()) *gtk.Button {
	return MustButtonImageFromImage(label, MustImageFromFile(img), clicked)
}

func MustButton(img *gtk.Image, clicked func()) *gtk.Button {
	b, err := gtk.ButtonNew()
	if err != nil {
		panic(err)
	}

	b.SetImage(img)
	b.SetImagePosition(gtk.POS_TOP)

	if clicked != nil {
		b.Connect("clicked", clicked)
	}

	return b
}

// MustButtonImageFromImage returns a new gtk.Button with the given label, image
// and clicked callback. If error panics.
func MustButtonImageFromImage(label string, img *gtk.Image, clicked func()) *gtk.Button {
	b, err := gtk.ButtonNewWithLabel(label)
	if err != nil {
		panic(err)
	}

	b.SetImage(img)
	b.SetAlwaysShowImage(true)
	b.SetImagePosition(gtk.POS_TOP)
	b.SetVExpand(true)
	b.SetHExpand(true)
	ctx, _ := b.GetStyleContext()
	ctx.AddClass("labelbig")

	if clicked != nil {
		b.Connect("clicked", clicked)
	}

	return b
}

// MustImageFromFileWithSize returns a new gtk.Image based on rescaled version
// of the given file.
func MustImageFromFileWithSize(img string, w, h int) *gtk.Image {
	p, err := gdk.PixbufNewFromFileAtScale(imagePath(img), w, h, true)
	if err != nil {
		Logger.Error(err)
	}

	i, err := gtk.ImageNewFromPixbuf(p)
	if err != nil {
		panic(err)
	}

	return i
}

// MustImageFromFile returns a new gtk.Image based on the given file, If error
// panics.
func MustImageFromFile(img string) *gtk.Image {
	i, err := gtk.ImageNewFromFile(imagePath(img))
	if err != nil {
		panic(err)
	}

	return i
}

// MustCSSProviderFromFile returns a new gtk.CssProvider for a given css file,
// If error panics.
func MustCSSProviderFromFile(css string) *gtk.CssProvider {
	p, err := gtk.CssProviderNew()
	if err != nil {
		panic(err)
	}

	if err := p.LoadFromPath(filepath.Join(StylePath, css)); err != nil {
		panic(err)
	}

	return p
}

func imagePath(img string) string {
	return filepath.Join(StylePath, ImageFolder, img)
}

// MustOverlay returns a new gtk.Overlay, if error panics.
func MustOverlay() *gtk.Overlay {
	o, err := gtk.OverlayNew()
	
	ctx, _ := o.GetStyleContext()
	ctx.AddClass("labelbig")
	if err != nil {
		panic(err)
	}

	return o
}

//Создаем билдера, есали не создали - паникуем
func MustBuilder() *gtk.Builder {
	bldr, err := gtk.BuilderNew()
	if err != nil {
		panic(err)
	}

	return bldr
}

func GladePath(glade string) string {
	return filepath.Join(StylePath, GladeFolder, glade)
}

//Создаем Выпадающее меню, если не создали - паникуем
func MustComboBoxText() *gtk.ComboBoxText {
	cmb, err := gtk.ComboBoxTextNew()
	
	ctx, _ := cmb.GetStyleContext()
	ctx.AddClass("combobox")
	if err != nil {
		panic(err)
	}

	return cmb
}

// //Создаем , если не создали - паникуем
// func MustRadioButton(groupLang *glib.SList, label string) *gtk.RadioButton {
	// rb, err := gtk.RadioButtonNewWithLabel(groupLang, label)
	

	// ctx, _ := rb.GetStyleContext()
	// ctx.AddClass("combobox")
	// if err != nil {
		// panic(err)
	// }

	// return rb
// }



// MustImageFromFile returns a new gtk.Image based on the given file, If error
// panics.
// func (ui *UI) MustFormFromFile(glade string) *gtk.Builder {
	// b, err := gtk.AddFromFile(gladePath(glade))
	// if err != nil {
		// panic(err)
	// }

	// return b
// }