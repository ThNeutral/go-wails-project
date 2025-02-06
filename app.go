package main

import (
	"context"
)

type PayloadWeaponArray struct {
	Data  []Weapon
	Error error
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

func (a *App) GetAllWeapons() PayloadWeaponArray {
	weapons, err := fetchAllWeaponsData()
	return PayloadWeaponArray{
		Data:  weapons,
		Error: err,
	}
}

func (a *App) GetWeaponsByQuery(query string) PayloadWeaponArray {
	var weapons []Weapon
	var err error

	if query == "" {
		weapons, err = fetchAllWeaponsData()
	} else {
		weapons, err = fetchWeaponsByQuery(query)
	}

	return PayloadWeaponArray{
		Data:  weapons,
		Error: err,
	}
}
