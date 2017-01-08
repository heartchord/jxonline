// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/lxn/walk"

	dcl "github.com/lxn/walk/declarative"
)

var isSpecialMode = walk.NewMutableCondition()

// MyMainWindow :
type MyMainWindow struct {
	*walk.MainWindow
	prevFilePath string
}

type Directory struct {
	name     string
	parent   *Directory
	children []*Directory
}

func NewDirectory(name string, parent *Directory) *Directory {
	return &Directory{name: name, parent: parent}
}

var _ walk.TreeItem = new(Directory)

func (d *Directory) Text() string {
	return d.name
}

func (d *Directory) Parent() walk.TreeItem {
	if d.parent == nil {
		// We can't simply return d.parent in this case, because the interface
		// value then would not be nil.
		return nil
	}

	return d.parent
}

func (d *Directory) ChildCount() int {
	if d.children == nil {
		// It seems this is the first time our child count is checked, so we
		// use the opportunity to populate our direct children.
		if err := d.ResetChildren(); err != nil {
			log.Print(err)
		}
	}

	return len(d.children)
}

func (d *Directory) ChildAt(index int) walk.TreeItem {
	return d.children[index]
}

func (d *Directory) Image() interface{} {
	return d.Path()
}

func (d *Directory) ResetChildren() error {
	d.children = nil

	dirPath := d.Path()

	if err := filepath.Walk(d.Path(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			if info == nil {
				return filepath.SkipDir
			}
		}

		name := info.Name()

		if !info.IsDir() || path == dirPath || shouldExclude(name) {
			return nil
		}

		d.children = append(d.children, NewDirectory(name, d))

		return filepath.SkipDir
	}); err != nil {
		return err
	}

	return nil
}

func (d *Directory) Path() string {
	elems := []string{d.name}

	dir, _ := d.Parent().(*Directory)

	for dir != nil {
		elems = append([]string{dir.name}, elems...)
		dir, _ = dir.Parent().(*Directory)
	}

	return filepath.Join(elems...)
}

type DirectoryTreeModel struct {
	walk.TreeModelBase
	roots []*Directory
}

var _ walk.TreeModel = new(DirectoryTreeModel)

func NewDirectoryTreeModel() (*DirectoryTreeModel, error) {
	model := new(DirectoryTreeModel)

	drives, err := walk.DriveNames()
	if err != nil {
		return nil, err
	}

	for _, drive := range drives {
		switch drive {
		case "A:\\", "B:\\":
			continue
		}

		model.roots = append(model.roots, NewDirectory(drive, nil))
	}

	return model, nil
}

func (*DirectoryTreeModel) LazyPopulation() bool {
	// We don't want to eagerly populate our tree view with the whole file system.
	return true
}

func (m *DirectoryTreeModel) RootCount() int {
	return len(m.roots)
}

func (m *DirectoryTreeModel) RootAt(index int) walk.TreeItem {
	return m.roots[index]
}

type FileInfo struct {
	Name     string
	Size     int64
	Modified time.Time
}

type FileInfoModel struct {
	walk.SortedReflectTableModelBase
	dirPath string
	items   []*FileInfo
}

var _ walk.ReflectTableModel = new(FileInfoModel)

func NewFileInfoModel() *FileInfoModel {
	return new(FileInfoModel)
}

func (m *FileInfoModel) Items() interface{} {
	return m.items
}

func (m *FileInfoModel) SetDirPath(dirPath string) error {
	m.dirPath = dirPath
	m.items = nil

	if err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			if info == nil {
				return filepath.SkipDir
			}
		}

		name := info.Name()

		if path == dirPath || shouldExclude(name) {
			return nil
		}

		item := &FileInfo{
			Name:     name,
			Size:     info.Size(),
			Modified: info.ModTime(),
		}

		m.items = append(m.items, item)

		if info.IsDir() {
			return filepath.SkipDir
		}

		return nil
	}); err != nil {
		return err
	}

	m.PublishRowsReset()

	return nil
}

func (m *FileInfoModel) Image(row int) interface{} {
	return filepath.Join(m.dirPath, m.items[row].Name)
}

func shouldExclude(name string) bool {
	switch name {
	case "System Volume Information", "pagefile.sys", "swapfile.sys":
		return true
	}

	return false
}

func main() {
	dcl.MustRegisterCondition("isSpecialMode", isSpecialMode)

	mw := new(MyMainWindow)
	bp := new(RoleBakPage)
	dp := new(RoleDbPage)

	var openAction, showAboutBoxAction *walk.Action
	var recentMenu *walk.Menu
	var treeView *walk.TreeView
	var tableView *walk.TableView
	var webView *walk.WebView

	treeModel, err := NewDirectoryTreeModel()
	if err != nil {
		log.Fatal(err)
	}
	tableModel := NewFileInfoModel()

	if err := (dcl.MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "角色存档信息查看工具",
		MenuItems: []dcl.MenuItem{
			dcl.Menu{
				// &符号：Alt快捷键，快捷键即&后面字符
				Text: "文件(&F)",
				Items: []dcl.MenuItem{
					dcl.Action{
						AssignTo: &openAction,
						Text:     "Open",
						Image:    "../../gameresource/img/open.png",
						// Shortcut字段：设置Action的快捷键
						Shortcut:    dcl.Shortcut{Modifiers: walk.ModControl, Key: walk.KeyO},
						OnTriggered: mw.openActionTriggered,
					},
					dcl.Menu{
						AssignTo: &recentMenu,
						Text:     "Recent",
					},
					dcl.Separator{},
					dcl.Action{
						Text:        "Exit",
						OnTriggered: func() { mw.Close() },
					},
				},
			},
			dcl.Menu{
				Text: "帮助(&H)",
				Items: []dcl.MenuItem{
					dcl.Action{
						AssignTo:    &showAboutBoxAction,
						Text:        "About",
						OnTriggered: mw.showAboutBoxActionTriggered,
					},
				},
			},
		},

		ToolBar: dcl.ToolBar{
			ButtonStyle: dcl.ToolBarButtonImageBeforeText,
			Items: []dcl.MenuItem{
				//dcl.ActionRef{Action: &openAction},
				dcl.Menu{
					Text:  "New A",
					Image: "../../gameresource/img/document-new.png",
					Items: []dcl.MenuItem{
						dcl.Action{
							Text:        "A",
							OnTriggered: mw.newActionTriggered,
						},
						dcl.Action{
							Text:        "B",
							OnTriggered: mw.newActionTriggered,
						},
						dcl.Action{
							Text:        "C",
							OnTriggered: mw.newActionTriggered,
						},
					},
					OnTriggered: mw.newActionTriggered,
				},
				dcl.Separator{},
				dcl.Menu{
					Text:  "View",
					Image: "../../gameresource/img/document-properties.png",
					Items: []dcl.MenuItem{
						dcl.Action{
							Text:        "X",
							OnTriggered: mw.changeViewActionTriggered,
						},
						dcl.Action{
							Text:        "Y",
							OnTriggered: mw.changeViewActionTriggered,
						},
						dcl.Action{
							Text:        "Z",
							OnTriggered: mw.changeViewActionTriggered,
						},
					},
				},
				dcl.Separator{},
				dcl.Action{
					Text:        "Special",
					Image:       "../../gameresource/img/system-shutdown.png",
					OnTriggered: mw.specialActionTriggered,
				},
			},
		},

		ContextMenuItems: []dcl.MenuItem{
			dcl.ActionRef{Action: &showAboutBoxAction},
		},
		MinSize: dcl.Size{Width: 800, Height: 600},
		Layout:  dcl.VBox{},

		Children: []dcl.Widget{
			dcl.TabWidget{
				Pages: []dcl.TabPage{
					dcl.TabPage{
						AssignTo: &bp.self,
						Title:    "Role Bak",
						Layout:   dcl.HBox{},
						Children: []dcl.Widget{
							dcl.HSplitter{
								Children: []dcl.Widget{
									dcl.TreeView{
										AssignTo: &treeView,
										Model:    treeModel,
										OnCurrentItemChanged: func() {
											dir := treeView.CurrentItem().(*Directory)
											if err := tableModel.SetDirPath(dir.Path()); err != nil {
												walk.MsgBox(
													mw.MainWindow,
													"Error",
													err.Error(),
													walk.MsgBoxOK|walk.MsgBoxIconError)
											}
										},
									},

									dcl.TableView{
										AssignTo:      &tableView,
										StretchFactor: 2,
										Columns: []dcl.TableViewColumn{
											dcl.TableViewColumn{
												DataMember: "Name",
												Width:      192,
											},
											dcl.TableViewColumn{
												DataMember: "Size",
												Format:     "%d",
												Alignment:  dcl.AlignFar,
												Width:      64,
											},
											dcl.TableViewColumn{
												DataMember: "Modified",
												Format:     "2006-01-02 15:04:05",
												Width:      120,
											},
										},
										Model: tableModel,
										OnCurrentIndexChanged: func() {
											var url string
											if index := tableView.CurrentIndex(); index > -1 {
												name := tableModel.items[index].Name
												dir := treeView.CurrentItem().(*Directory)
												url = filepath.Join(dir.Path(), name)
											}

											webView.SetURL(url)
										},
									},
									dcl.WebView{
										AssignTo:      &webView,
										StretchFactor: 2,
									},
								},
							},
						},
					},

					dcl.TabPage{
						AssignTo: &dp.self,
						Title:    "Role DB",
						Layout:   dcl.VBox{},
					},
				},
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	addRecentFileActions := func(texts ...string) {
		for _, text := range texts {
			a := walk.NewAction()
			a.SetText(text)
			a.Triggered().Attach(mw.openActionTriggered)
			recentMenu.Actions().Add(a)
		}
	}

	addRecentFileActions("Foo", "Bar", "Baz")

	mw.Run()
}

func (mw *MyMainWindow) openImage() error {
	dlg := new(walk.FileDialog)

	dlg.FilePath = mw.prevFilePath
	dlg.Filter = "Image Files (*.emf;*.bmp;*.exif;*.gif;*.jpeg;*.jpg;*.png;*.tiff)|*.emf;*.bmp;*.exif;*.gif;*.jpeg;*.jpg;*.png;*.tiff"
	dlg.Title = "Select an Image"

	if ok, err := dlg.ShowOpen(mw); err != nil {
		return err
	} else if !ok {
		return nil
	}

	mw.prevFilePath = dlg.FilePath

	return nil
}

func (mw *MyMainWindow) openActionTriggered() {
	if err := mw.openImage(); err != nil {
		log.Print(err)
	}
}

func (mw *MyMainWindow) newActionTriggered() {
	walk.MsgBox(mw, "New", "Newing something up... or not.", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) changeViewActionTriggered() {
	walk.MsgBox(mw, "Change View", "By now you may have guessed it. Nothing changed.", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) showAboutBoxActionTriggered() {
	walk.MsgBox(mw, "About", "Walk Actions Example", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) specialActionTriggered() {
	walk.MsgBox(mw, "Special", "Nothing to see here.", walk.MsgBoxIconInformation)
}
