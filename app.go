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

func (a *App) GetWeaponByID(id int) (Weapon, error) {
	weapon, err := fetchWeaponByID(id)

	if err != nil {
		fmt.Println(err)
		return Weapon{}, err
	}

	return weapon, nil
}

func (a *App) GetWeaponsByQuery(query string) ([]Weapon, error) {
	weapons, err := fetchWeaponsByQuery(query)

	if err != nil {
		fmt.Println(err)
		return []Weapon{}, err
	}

	return weapons, nil
}
