package main

import (
	"errors"
	"fmt"
	"strconv"
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
		return nil, err
	}

	allWeaponsData = data
	return data, nil
}

func fetchWeaponsByQuery(query string) ([]Weapon, error) {
	if len(query) == 0 || query[0] != '$' {
		_, err := strconv.Atoi(query)
		if err != nil {
			return fetchWeaponsByField("name", query)
		}
		return fetchWeaponsByField("id", query)
	}

	query = query[1:]
	parts := strings.Split(query, ":")
	if len(parts) != 2 {
		return nil, errors.New("invalid query $" + query)
	}

	fmt.Println(parts[0] + "\t" + parts[1])

	return fetchWeaponsByField(parts[0], parts[1])
}

func fetchWeaponsByField(field, value string) ([]Weapon, error) {
	allWeapons, err := fetchAllWeaponsData()
	if err != nil {
		return nil, err
	}

	var filtered []Weapon
	switch strings.ToLower(field) {
	case "name":
		filtered = filter(allWeapons, func(w Weapon) bool { return searchSubstring(w.Name, value) })
	case "id":
		i, err := strconv.Atoi(value)
		if err != nil {
			return nil, errors.New("expected integer in ID query, encountered " + value)
		}
		filtered = filter(allWeapons, func(w Weapon) bool { return w.ID == i })
	case "slug":
		filtered = filter(allWeapons, func(w Weapon) bool { return searchSubstring(w.Slug, value) })
	case "type":
		filtered = filter(allWeapons, func(w Weapon) bool { return searchSubstring(string(w.Type), value) })
	case "damagetype":
		filtered = filter(allWeapons, func(w Weapon) bool { return searchSubstring(string(w.DamageType), value) })
	case "eldersealtype":
		filtered = filter(allWeapons, func(w Weapon) bool { return searchSubstring(string(w.Elderseal), value) })
	case "rarity":
		i, err := strconv.Atoi(value)
		if err != nil {
			return nil, errors.New("expected integer in rarity query, encountered " + value)
		}
		filtered = filter(allWeapons, func(w Weapon) bool { return w.Rarity == i })
	}

	return filtered, nil
}
