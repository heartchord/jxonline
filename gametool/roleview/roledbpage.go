package main

import (
	"github.com/lxn/walk"
	dcl "github.com/lxn/walk/declarative"
)

// RoleDbPage :
type RoleDbPage struct {
	parent *MyMainWindow
	self   *walk.TabPage
}

// Create is
func (pg *RoleDbPage) Create(parent *MyMainWindow) *dcl.TabPage {
	pg.parent = parent

	return &dcl.TabPage{
		AssignTo: &pg.self,
		Title:    "Role DB",
		Layout:   dcl.VBox{},
	}
}
