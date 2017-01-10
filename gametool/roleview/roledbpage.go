package main

import (
	"github.com/lxn/walk"
	dcl "github.com/lxn/walk/declarative"
)

// RoleDbPage :
type RoleDbPage struct {
	*walk.TabPage
}

// Create is
func (pg *RoleDbPage) Create() *dcl.TabPage {
	return &dcl.TabPage{
		AssignTo: &pg.TabPage,
		Title:    "Role DB",
		Layout:   dcl.VBox{},
	}
}
