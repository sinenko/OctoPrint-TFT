package ui

import (
	"fmt"
	"sort"
	"time"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/gotk3/gotk3/gtk"
	"github.com/mcuadros/go-octoprint"
	"github.com/mcuadros/OctoPrint-TFT/ui_lang"
	exe "github.com/mcuadros/OctoPrint-TFT/ui_exec"
)

var filesPanelInstance *filesPanel

// const (
	// LocalFlash  octoprint.Location = "local/flash"
// )

type filesPanel struct {
	CommonPanel
	selectedFolder string
	list *gtk.Box
	scroll *gtk.ScrolledWindow
}

func FilesPanel(ui *UI, parent Panel) Panel {
	if filesPanelInstance == nil {
		m := &filesPanel{CommonPanel: NewCommonPanel(ui, parent)}
		m.initialize()
		filesPanelInstance = m
	} else {
		filesPanelInstance.doLoadFiles()
	}
	return filesPanelInstance
}

func (m *filesPanel) initialize() {
	m.list = MustBox(gtk.ORIENTATION_VERTICAL, 0)
	m.list.SetVExpand(true)

	m.scroll, _ = gtk.ScrolledWindowNew(nil, nil)
	m.scroll.Add(m.list)

	box := MustBox(gtk.ORIENTATION_VERTICAL, 0)
	box.Add(m.scroll)
	box.Add(m.createActionBar())
	m.Grid().Add(box)

	exe.MountFlash()
	m.doLoadFiles()
}

func (m *filesPanel) createActionBar() gtk.IWidget {
	bar := MustBox(gtk.ORIENTATION_HORIZONTAL, 5)
	bar.SetHAlign(gtk.ALIGN_END)
	bar.SetHExpand(true)
	bar.SetMarginTop(5)
	bar.SetMarginBottom(5)
	bar.SetMarginEnd(5)

	bar.Add(m.createDeleteButton())
	bar.Add(m.createSpace())
	bar.Add(m.createSpace())
	bar.Add(m.createSpace())
	bar.Add(m.createSpace())
	bar.Add(m.createRefreshButton())
	bar.Add(m.createScrollUp())
	bar.Add(m.createScrollDown())
	//bar.Add(m.createInitReleaseSDButton())
	bar.Add(MustButton(MustImageFromFileWithSize("back.svg", 80, 80), m.doGoBackWithScrollReset))

	return bar
}



func (m *filesPanel) createRefreshButton() gtk.IWidget {
	return MustButton(MustImageFromFileWithSize("refresh.svg", 80, 80), m.doLoadFiles)
}

func (m *filesPanel) createSpace() gtk.IWidget {
	return MustImageFromFileWithSize("space.svg", 77, 77)
}

func (m *filesPanel) createScrollDown() gtk.IWidget {
	return MustButton(MustImageFromFileWithSize("navdown.svg", 80, 80), m.doScrollDown)
}

func (m *filesPanel) createScrollUp() gtk.IWidget {
	return MustButton(MustImageFromFileWithSize("navup.svg", 80, 80), m.doScrollUp)
}

func (m *filesPanel) createDeleteButton() gtk.IWidget {
	return MustButton(MustImageFromFileWithSize("deleteall.svg", 80, 80), m.doDeleteFiles)
}



func getFilesFromFolder(s []*octoprint.FileInformation, name string) ([]*octoprint.FileInformation) {
	folders := strings.Split(name, "/")

	for i := 0; i < len(folders); i++ {
		for _, f := range s {
			if (f.Name != folders[i]) {
				continue
			}
			
			if (f.IsFolder()) {
				if ((i+1) < len(folders)){
					s = f.Children
				} else {
					return f.Children
				}
			}
		}
	}
	
	return []*octoprint.FileInformation{}
}

func (m *filesPanel) doLoadFiles() {
	Logger.Info(ui_lang.Translate("Refreshing list of files"))
	// m.doRefreshSD()
	exe.MountFlash()

	local := m.doLoadFilesFromLocation(octoprint.Local)
	sdcard := m.doLoadFilesFromLocation(octoprint.SDCard)

	s := byDate(local)
	s = append(s, sdcard...)
	sort.Sort(s)

	EmptyContainer(&m.list.Container)
	for _, f := range s {
		if (!f.IsFolder()) {
			continue
		}

		m.addFile(m.list, f)
	}
	for _, f := range s {
		if f.IsFolder() {
			continue
		}

		m.addFile(m.list, f)
	}

	m.list.ShowAll()
}

func (m *filesPanel) doGoBackWithScrollReset() {
	m.doScrollHome()
	m.UI.GoHistory()
}

func (m *filesPanel) doRefreshSD() {
	if err := (&octoprint.SDRefreshRequest{}).Do(m.UI.Printer); err != nil {
		Logger.Error(err)
	}
}

func (m *filesPanel) doDeleteFiles() {
	m.doDeleteFilesFromLocation(octoprint.Local)
}


func (m *filesPanel) doScrollDown() {
	m.scroll.GetVAdjustment().SetValue(m.scroll.GetVAdjustment().GetValue() + 376)
}

func (m *filesPanel) doScrollUp() {
	m.scroll.GetVAdjustment().SetValue(m.scroll.GetVAdjustment().GetValue() - 376)
}

func (m *filesPanel) doScrollHome() {
	m.scroll.GetVAdjustment().SetValue(m.scroll.GetVAdjustment().GetLower())
}

func (m *filesPanel) doLoadFilesFromLocation(l octoprint.Location) []*octoprint.FileInformation {
	r := &octoprint.FilesRequest{Location: l, Recursive: true}
	files, err := r.Do(m.UI.Printer)
	if err != nil {
		Logger.Error(err)
		return []*octoprint.FileInformation{}
	}

	return files.Files
}

func (m *filesPanel) doDeleteFilesFromLocation(l octoprint.Location) {
	local := m.doLoadFilesFromLocation(l)
	for _, f := range local {
		if f.IsFolder() {
			continue
		}

		r := &octoprint.DeleteFileRequest{Location: l, Path: f.Path}
		err := r.Do(m.UI.Printer)
		if err != nil {
			Logger.Error(err)
		}
	}
	m.doLoadFiles()
	return
}

func (m *filesPanel) addFile(b *gtk.Box, f *octoprint.FileInformation) {
	frame, _ := gtk.FrameNew("")

	name := MustLabel(f.Name)
	name.SetMarkup(fmt.Sprintf("<big>%s</big>", filenameEllipsis(f.Name)))
	name.SetHExpand(true)

	info := MustLabel("")
	info.SetMarkup(fmt.Sprintf(ui_lang.Translate("<small>Uploaded: <b>%s</b> - Size: <b>%s</b></small>"),
		ui_lang.FindAndTranslate("week ago", ui_lang.FindAndTranslate("weeks ago", ui_lang.FindAndTranslate("minutes ago", ui_lang.FindAndTranslate("minute ago", ui_lang.FindAndTranslate("hour ago", ui_lang.FindAndTranslate("hours ago", ui_lang.FindAndTranslate("hours from now", ui_lang.FindAndTranslate("day ago", ui_lang.FindAndTranslate("days ago", ui_lang.FindAndTranslate("month ago", ui_lang.FindAndTranslate("months ago", ui_lang.FindAndTranslate("year ago", ui_lang.FindAndTranslate("years ago", ui_lang.FindAndTranslate("a long while ago", humanize.Time(f.Date.Time))))))))))))))), 
		ui_lang.FindAndTranslate("B",ui_lang.FindAndTranslate("mB", ui_lang.FindAndTranslate("kB", humanize.Bytes(uint64(f.Size))))),
	))

	labels := MustBox(gtk.ORIENTATION_VERTICAL, 5)
	labels.Add(name)
	labels.Add(info)

	actions := MustBox(gtk.ORIENTATION_HORIZONTAL, 5)
	if(f.IsFolder()){
		actions.Add(m.createOpenFolderButton("open_folder.svg", f))
	} else {
		// actions.Add(m.createLoadAndPrintButton("load.svg", f, false))
		actions.Add(m.createLoadAndPrintButton("status.svg", f, true))
	}

	file := MustBox(gtk.ORIENTATION_HORIZONTAL, 5)
	file.SetMarginTop(5)
	file.SetMarginEnd(5)
	file.SetMarginStart(5)
	file.SetMarginBottom(5)
	file.SetHExpand(true)

	if(f.IsFolder()){
		file.Add(MustImageFromFileWithSize("folder.svg", 82, 82)) //////62 62
	} else {
		file.Add(MustImageFromFileWithSize("file.svg", 82, 82))  //////62 62
	}

	file.Add(labels)
	file.Add(actions)

	frame.Add(file)
	b.Add(frame)
}

func (m *filesPanel) addUpDirectory(b *gtk.Box, pathDiretory string) {
	frame, _ := gtk.FrameNew("")

	name := MustLabel("Up")
	name.SetMarkup(fmt.Sprintf("<big>%s</big>", filenameEllipsis(ui_lang.Translate("UP DIR"))))
	name.SetHExpand(true)

	info := MustLabel(ui_lang.Translate("Go to prev directory"))

	labels := MustBox(gtk.ORIENTATION_VERTICAL, 5)
	labels.Add(name)
	labels.Add(info)

	actions := MustBox(gtk.ORIENTATION_HORIZONTAL, 5)
	actions.Add(MustButton(MustImageFromFileWithSize("folder_up.svg", 80, 80),   //////60 60
	func() {
		
			// m.doRefreshSD()
			
			local := m.doLoadFilesFromLocation(octoprint.Local)

			s := byDate(local)
			sort.Sort(s)

			EmptyContainer(&m.list.Container)
			
			folders := strings.Split(pathDiretory, "/")
			//Если уровень вложенности меньше двух, значит прудыдущая папка - корень
			if(len(folders)<2) {
				m.doLoadFiles()
			} else {
			
				folders = folders[:len(folders)-1]
				upFolderString := strings.Join(folders,"/")
				filesFromFolder := getFilesFromFolder(s, upFolderString)
				
				m.addUpDirectory(m.list, upFolderString) ////////////
				for _, f := range filesFromFolder {
					m.addFile(m.list, f)
				}

				m.list.ShowAll()
			}
		
	}))

	file := MustBox(gtk.ORIENTATION_HORIZONTAL, 5)
	file.SetMarginTop(5)
	file.SetMarginEnd(5)
	file.SetMarginStart(5)
	file.SetMarginBottom(5)
	file.SetHExpand(true)

	file.Add(MustImageFromFileWithSize("dotdot.svg", 82, 82))  //////62 62

	file.Add(labels)
	file.Add(actions)

	frame.Add(file)
	b.Add(frame)
}

func (m *filesPanel) createOpenFolderButton(img string, f *octoprint.FileInformation) gtk.IWidget {

	return MustButton(MustImageFromFileWithSize(img, 60, 60), //////60 60 
	func() {
		
			// m.doRefreshSD()
			
			local := m.doLoadFilesFromLocation(octoprint.Local)

			s := byDate(local)
			sort.Sort(s)

			EmptyContainer(&m.list.Container)
			filesFromFolder := getFilesFromFolder(s, f.Path)
			
			m.addUpDirectory(m.list, f.Path)
			for _, f := range filesFromFolder {
				m.addFile(m.list, f)
			}

			m.list.ShowAll()
		
	})
}

func (m *filesPanel) createLoadAndPrintButton(img string, f *octoprint.FileInformation, print bool) gtk.IWidget {
	return MustButton(
		MustImageFromFileWithSize(img, 60, 60),
		MustConfirmDialog(m.UI.w, ui_lang.Translate("Are you sure you want to proceed?"), func() {
			r := &octoprint.SelectFileRequest{}
			r.Location = octoprint.Local
			r.Path = f.Path
			r.Print = print

			m.doScrollHome()
			Logger.Infof(ui_lang.Translate("Loading file %q, printing: %v"), f.Name, print)
			if err := r.Do(m.UI.Printer); err != nil {
				Logger.Error(err)
				return
			}
		}),
	)
}

func (m *filesPanel) createInitReleaseSDButton() gtk.IWidget {
	release := MustImageFromFileWithSize("sd_eject.svg", 80, 80)
	init := MustImageFromFileWithSize("sd.svg", 80, 80)
	b := MustButton(release, nil)

	state := func() {
		time.Sleep(50 * time.Millisecond)
		switch m.isReady() {
		case true:
			b.SetImage(release)
		case false:
			b.SetImage(init)
		}
	}

	b.Connect("clicked", func() {
		var err error
		if !m.isReady() {
			err = (&octoprint.SDInitRequest{}).Do(m.UI.Printer)
		} else {
			err = (&octoprint.SDReleaseRequest{}).Do(m.UI.Printer)
		}

		if err != nil {
			Logger.Error(err)
		}

		state()
	})

	return b
}

func (m *filesPanel) isReady() bool {
	state, err := (&octoprint.SDStateRequest{}).Do(m.UI.Printer)
	if err != nil {
		Logger.Error(err)
		return false
	}

	return state.Ready
}

type byDate []*octoprint.FileInformation

func (s byDate) Len() int           { return len(s) }
func (s byDate) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s byDate) Less(i, j int) bool { return s[j].Date.Time.Before(s[i].Date.Time) }
