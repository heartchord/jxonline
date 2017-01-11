package main

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"fmt"

	"github.com/heartchord/goblazer"
	"github.com/lxn/walk"
)

// DirectoryNode is a node of a directory tree
type DirectoryNode struct {
	name     string
	alias    string
	parent   *DirectoryNode
	children []*DirectoryNode
}

// NewDirectoryNode creates a new node of the directory tree
func NewDirectoryNode(name string, parent *DirectoryNode) *DirectoryNode {
	return &DirectoryNode{name: name, parent: parent}
}

// SetAlias returns the text of the directory node text
func (d *DirectoryNode) SetAlias(alias string) {
	d.alias = alias
}

// Text returns the text of the directory node text
func (d *DirectoryNode) Text() string {
	if len(d.alias) != 0 {
		return d.alias
	}
	return d.name
}

// Parent returns the parent of the directory node
func (d *DirectoryNode) Parent() walk.TreeItem {
	if d.parent == nil { // We can't simply return d.parent in this case, because the interface value then would not be nil
		return nil
	}
	return d.parent
}

// ChildCount returns the number of children of the directory node
func (d *DirectoryNode) ChildCount() int {
	if d.children == nil {
		// 孩子如果为空，重新生成所有孩子
		if err := d.ResetChildren(); err != nil {
			log.Print(err)
		}
	}
	return len(d.children)
}

// ChildAt returns the child node at the specified index
func (d *DirectoryNode) ChildAt(index int) walk.TreeItem {
	return d.children[index]
}

// FindChild returns the child index by child.name
func (d *DirectoryNode) FindChild(name string) int {
	for i, v := range d.children {
		if v.name == name {
			return i
		}
	}

	return -1
}

// Image :
func (d *DirectoryNode) Image() interface{} {
	return d.Path()
}

// ResetChildren rebuilds all children nodes of current directory node
func (d *DirectoryNode) ResetChildren() error {
	d.children = nil

	dirPath := d.Path()
	if err := filepath.Walk(
		d.Path(), // 目录路径
		func(path string, info os.FileInfo, err error) error { // 目录所有子节点处理函数
			if err != nil {
				if info == nil {
					return filepath.SkipDir
				}
			}

			name := info.Name()
			if !info.IsDir() || path == dirPath || shouldExclude(name) {
				return nil
			}

			d.children = append(d.children, NewDirectoryNode(name, d))

			return filepath.SkipDir
		}); err != nil {
		return err
	}

	return nil
}

// Path returns the path of the directory node
func (d *DirectoryNode) Path() string {
	elems := []string{d.name}

	dir, _ := d.Parent().(*DirectoryNode)
	for dir != nil {
		elems = append([]string{dir.name}, elems...)
		dir, _ = dir.Parent().(*DirectoryNode)
	}

	return filepath.Join(elems...)
}

// DirectoryTreeModel provides a directory tree model
type DirectoryTreeModel struct {
	walk.TreeModelBase
	roots []*DirectoryNode
}

// NewDirectoryTreeModel creates a new directory tree model
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
		dnode := NewDirectoryNode(drive, nil)
		dnode.SetAlias(fmt.Sprintf("本地磁盘(%s)", goblazer.GetPathName(drive)))
		model.roots = append(model.roots, dnode)
	}

	// 加入桌面路径
	u, err := user.Current()
	if err == nil {
		drive := u.HomeDir + "\\Desktop"
		dnode := NewDirectoryNode(drive, nil)
		dnode.SetAlias("桌面")
		model.roots = append(model.roots, dnode)
	}

	return model, nil
}

// LazyPopulation returns if the model prefers on-demand population.
func (*DirectoryTreeModel) LazyPopulation() bool {
	// 返回false将展示整个文件系统
	return true
}

// RootCount returns
func (m *DirectoryTreeModel) RootCount() int {
	return len(m.roots)
}

// RootAt creates root node count based on the specific path
func (m *DirectoryTreeModel) RootAt(index int) walk.TreeItem {
	return m.roots[index]
}

// FileInfo is a struct contains some infomation of a file
type FileInfo struct {
	Name     string
	Size     int64
	Modified time.Time
}

// FileInfoModel is a file info model
type FileInfoModel struct {
	walk.SortedReflectTableModelBase
	dirPath string
	items   []*FileInfo
}

// NewFileInfoModel creates a new file info model
func NewFileInfoModel() *FileInfoModel {
	return new(FileInfoModel)
}

// Items returns
func (m *FileInfoModel) Items() interface{} {
	return m.items
}

// SetDirPath creates
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

// Image :
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
