package main

import (
	"context"
)

type ResponseWeaponArray struct {
	Data  []Weapon
	Error error
}

type ResponseWeaponTree struct {
	Data  *WeaponTreeNode
	Error error
}

type RequestQuery struct {
	Field string
	Query string
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetAllWeapons() ResponseWeaponArray {
	weapons, err := fetchAllWeaponsData()
	return ResponseWeaponArray{
		Data:  weapons,
		Error: err,
	}
}

func (a *App) GetWeaponsByQuery(query RequestQuery) ResponseWeaponArray {
	weapons, err := fetchWeaponsByQuery(query)
	return ResponseWeaponArray{
		Data:  weapons,
		Error: err,
	}
}

func (a *App) GetWeaponTree(ID int) ResponseWeaponTree {
	tree, err := generateWeaponTree(ID)
	return ResponseWeaponTree{
		Data:  tree,
		Error: err,
	}
}
