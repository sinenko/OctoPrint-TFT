package ui

import (
	"fmt"
	"log"
	"math"
	"strings"
	"sync"
	"time"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/mcuadros/OctoPrint-TFT/ui_lang"
)

// Set at compilation time.
var Version = "0.1.x"
var Build = "no-set"

const panelW = 5
const panelH = 4

type Panel interface {
	Grid() *gtk.Grid
	Show()
	Hide()
	Parent() Panel
}

type CommonPanel struct {
	UI *UI
	g  *gtk.Grid
	b  *BackgroundTask
	p  Panel

	buttons []gtk.IWidget
}

func NewCommonPanel(ui *UI, parent Panel) CommonPanel {
	g := MustGrid()
	g.SetRowHomogeneous(true)
	g.SetColumnHomogeneous(true)

	return CommonPanel{UI: ui, g: g, p: parent}
}

func (p *CommonPanel) Initialize() {
	last := panelW * panelH
	if last < len(p.buttons) {
		cols := math.Ceil(float64(len(p.buttons)) / float64(panelW))
		last = int(cols) * panelW
	}

	for i := len(p.buttons) + 1; i < last; i++ {
		p.AddButton(MustBox(gtk.ORIENTATION_HORIZONTAL, 0))
	}

	p.AddBackButton(MustButtonImage(ui_lang.Translate("Back"), "back.svg", p.UI.GoHistory))
}

func (p *CommonPanel) Parent() Panel {
	return p.p
}

func (p *CommonPanel) AddButton(b gtk.IWidget) {
	x := len(p.buttons) % panelW
	y := len(p.buttons) / panelW
	p.g.Attach(b, x+1, y, 1, 1)
	p.buttons = append(p.buttons, b)
}

func (p *CommonPanel) AddBackButton(b gtk.IWidget) {
	x := len(p.buttons) % panelW
	p.g.Attach(b, x+1, 3, 1, 1)
	p.buttons = append(p.buttons, b)
}

func (p *CommonPanel) Show() {
	if p.b != nil {
		p.b.Start()
	}
}

func (p *CommonPanel) Hide() {
	if p.b != nil {
		p.b.Close()
	}
}

func (p *CommonPanel) Grid() *gtk.Grid {
	return p.g
}

type BackgroundTask struct {
	close chan bool

	d       time.Duration
	task    func()
	running bool
	sync.Mutex
}

func NewBackgroundTask(d time.Duration, task func()) *BackgroundTask {
	return &BackgroundTask{
		task: task,
		d:    d,

		close: make(chan bool, 1),
	}
}

func (t *BackgroundTask) Start() {
	t.Lock()
	defer t.Unlock()

	Logger.Debug(ui_lang.Translate("New background task started"))
	go t.loop()

	t.running = true
}

func (t *BackgroundTask) Close() {
	t.Lock()
	defer t.Unlock()
	if !t.running {
		return
	}

	t.close <- true
	t.running = false
}

func (t *BackgroundTask) loop() {
	t.execute()

	ticker := time.NewTicker(t.d)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			t.execute()
		case <-t.close:
			Logger.Debug(ui_lang.Translate("Background task closed"))
			return
		}
	}
}

func (t *BackgroundTask) execute() {
	_, err := glib.IdleAdd(t.task)
	if err != nil {
		log.Fatal(ui_lang.Translate("IdleAdd() failed:"), err)
	}
}

type StepButton struct {
	Current  int
	Steps    []Step
	Callback func()

	*gtk.Button
	sync.RWMutex
}

type Step struct {
	Label string
	Value interface{}
}

func MustStepButton(image string, s ...Step) *StepButton {
	var l string
	if len(s) != 0 {
		l = s[0].Label
	}

	b := &StepButton{
		Button: MustButtonImage(l, image, nil),
		Steps:  s,
	}

	b.Connect("clicked", b.clicked)
	return b
}

func (b *StepButton) Label() string {
	b.RLock()
	defer b.RUnlock()

	return b.Steps[b.Current].Label
}

func (b *StepButton) Value() interface{} {
	b.RLock()
	defer b.RUnlock()

	return b.Steps[b.Current].Value
}

func (b *StepButton) AddStep(s Step) {
	b.Lock()
	defer b.Unlock()

	if len(b.Steps) == 0 {
		b.SetLabel(s.Label)
	}

	b.Steps = append(b.Steps, s)
}

func (b *StepButton) clicked() {
	b.RLock()
	defer b.RUnlock()

	b.Current++
	if b.Current >= len(b.Steps) {
		b.Current = 0
	}

	b.SetLabel(b.Steps[b.Current].Label)

	if b.Callback != nil {
		b.Callback()
	}
}

func MustConfirmDialog(parent *gtk.Window, msg string, cb func()) func() {
	return func() {
		win := gtk.MessageDialogNewWithMarkup(
			parent,
			gtk.DIALOG_MODAL,
			gtk.MESSAGE_INFO,
			gtk.BUTTONS_OK_CANCEL,
			"",
		)

		win.SetMarkup(CleanHTML(msg))
		defer win.Destroy()

		box, _ := win.GetContentArea()
		box.SetMarginStart(15)
		box.SetMarginEnd(15)
		box.SetMarginTop(15)
		box.SetMarginBottom(15)

		ctx, _ := win.GetStyleContext()
		ctx.AddClass("dialog")

		if win.Run() == int(gtk.RESPONSE_OK) {
			cb()
		}
	}
}

func EmptyContainer(c *gtk.Container) {
	ch := c.GetChildren()
	defer ch.Free()

	ch.Foreach(func(i interface{}) {
		c.Remove(i.(gtk.IWidget))
	})
}

var translatedTags = [][2]string{{"strong", "b"}}
var disallowedTags = []string{"p"}

func CleanHTML(html string) string {
	for _, tag := range translatedTags {
		html = replaceHTMLTag(html, tag[0], tag[1])
	}

	for _, tag := range disallowedTags {
		html = replaceHTMLTag(html, tag, " ")
	}

	return html
}

func replaceHTMLTag(html, from, to string) string {
	for _, pattern := range []string{"<%s>", "</%s>", "<%s/>"} {
		to := to
		if to != "" && to != " " {
			to = fmt.Sprintf(pattern, to)
		}

		html = strings.Replace(html, fmt.Sprintf(pattern, from), to, -1)
	}

	return html
}