package main

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

const WeaponsDataEndpoint Endpoint = "weapons"

func fetchAllWeaponsData() ([]Weapon, error) {
	return fetchData[[]Weapon](WeaponsDataEndpoint)
}

const WeaponDataByIDEndpoint Endpoint = "weapons/"

func fetchAllWeaponByIDData(id string) (Weapon, error) {
	return fetchData[Weapon](WeaponDataByIDEndpoint + Endpoint(id))
}
