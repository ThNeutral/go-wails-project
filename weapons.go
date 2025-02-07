package main

import (
	"errors"
	"strconv"
	"strings"
)

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

func fetchWeaponsByQuery(request RequestQuery) ([]Weapon, error) {
	if len(request.Query) == 0 {
		return []Weapon{}, errors.New("empty query is not allowed")
	}

	return fetchWeaponsByField(request.Field, request.Query)
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
	default:
		return []Weapon{}, errors.New("unknown query " + field)
	}

	return filtered, nil
}
