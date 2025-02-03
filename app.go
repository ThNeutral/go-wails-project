package main

import (
	"context"
	"fmt"
)

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

func (a *App) GetAllWeapons() []Weapon {
	weapons, err := fetchAllWeaponsData()

	if err != nil {
		fmt.Println(err)
		return []Weapon{}
	}

	return weapons
}

func (a *App) GetWeaponByID(id string) Weapon {
	weapon, err := fetchAllWeaponByIDData(id)

	if err != nil {
		fmt.Println(err)
		return Weapon{}
	}

	return weapon
}
