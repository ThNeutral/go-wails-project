package main

import (
	"strings"
)

type Weapon struct {
	ID         int                `json:"id"`
	Slug       string             `json:"slug"`
	Name       string             `json:"name"`
	Type       WeaponType         `json:"type"`
	Rarity     int                `json:"rarity"`
	Attack     Attack             `json:"attack"`
	Slots      []Slot             `json:"slots"`
	Elements   []WeaponElement    `json:"elements"`
	Crafting   WeaponCraftingInfo `json:"crafting"`
	Assets     WeaponAssets       `json:"assets"`
	Durability []WeaponSharpness  `json:"durability"`
	Elderseal  EldersealType      `json:"elderseal"`
	DamageType DamageType         `json:"damageType"`
	Attributes WeaponAttributes   `json:"attributes"`
}

const WeaponsEndpoint Endpoint = "weapons"

var allWeaponsData []Weapon

func fetchAllWeaponsData() ([]Weapon, error) {
	if len(allWeaponsData) != 0 {
		return allWeaponsData, nil
	}

	data, err := fetchData[[]Weapon](WeaponsEndpoint)
	if err != nil {
		return []Weapon{}, err
	}

	allWeaponsData = data

	return data, nil
}

func fetchWeaponsByQuery(query string) ([]Weapon, error) {
	// TODO
	return []Weapon{}, nil
}

func fetchWeaponsByName(name string) ([]Weapon, error) {
	allWeapons, err := fetchAllWeaponsData()
	if err != nil {
		return []Weapon{}, err
	}

	var filtered []Weapon
	tree := generateRadixTree(name)
	for _, weapon := range allWeapons {
		if tree.search(strings.ToLower(weapon.Name)) {
			filtered = append(filtered, weapon)
		}
	}

	return filtered, nil
}

func fetchWeaponsBySlug(slug string) ([]Weapon, error) {
	allWeapons, err := fetchAllWeaponsData()
	if err != nil {
		return []Weapon{}, err
	}

	var filtered []Weapon
	tree := generateRadixTree(slug)
	for _, weapon := range allWeapons {
		if tree.search(strings.ToLower(weapon.Slug)) {
			filtered = append(filtered, weapon)
		}
	}

	return filtered, nil
}

func fetchWeaponsByType(t WeaponType) ([]Weapon, error) {
	allWeapons, err := fetchAllWeaponsData()
	if err != nil {
		return []Weapon{}, err
	}

	var filtered []Weapon
	tree := generateRadixTree(string(t))
	for _, weapon := range allWeapons {
		if tree.search(strings.ToLower(string(weapon.Type))) {
			filtered = append(filtered, weapon)
		}
	}

	return filtered, nil
}

func fetchWeaponsByDamageType(t DamageType) ([]Weapon, error) {
	allWeapons, err := fetchAllWeaponsData()
	if err != nil {
		return []Weapon{}, err
	}

	var filtered []Weapon
	tree := generateRadixTree(string(t))
	for _, weapon := range allWeapons {
		if tree.search(strings.ToLower(string(weapon.DamageType))) {
			filtered = append(filtered, weapon)
		}
	}

	return filtered, nil
}

func fetchWeaponsByEldersealType(t EldersealType) ([]Weapon, error) {
	allWeapons, err := fetchAllWeaponsData()
	if err != nil {
		return []Weapon{}, err
	}

	var filtered []Weapon
	tree := generateRadixTree(string(t))
	for _, weapon := range allWeapons {
		if tree.search(strings.ToLower(string(weapon.Elderseal))) {
			filtered = append(filtered, weapon)
		}
	}

	return filtered, nil
}

func fetchWeaponByID(id int) (Weapon, error) {
	allWeapons, err := fetchAllWeaponsData()
	if err != nil {
		return Weapon{}, err
	}

	for _, weapon := range allWeapons {
		if weapon.ID == id {
			return weapon, nil
		}
	}

	return Weapon{}, nil
}

func fetchWeaponByRarity(rarity int) (Weapon, error) {
	allWeapons, err := fetchAllWeaponsData()
	if err != nil {
		return Weapon{}, err
	}

	for _, weapon := range allWeapons {
		if weapon.Rarity == rarity {
			return weapon, nil
		}
	}

	return Weapon{}, nil
}
