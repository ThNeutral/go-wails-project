package main

type WeaponType string

const (
	GreatSword     WeaponType = "great-sword"
	LongSword      WeaponType = "long-sword"
	SwordAndShield WeaponType = "sword-and-shield"
	DualBlades     WeaponType = "dual-blades"
	Hammer         WeaponType = "hammer"
	HuntingHorn    WeaponType = "hunting-horn"
	Lance          WeaponType = "lance"
	Gunlance       WeaponType = "gunlance"
	SwitchAxe      WeaponType = "switch-axe"
	ChargeBlade    WeaponType = "charge-blade"
	InsectGlaive   WeaponType = "insect-glaive"
	LightBowgun    WeaponType = "light-bowgun"
	HeavyBowgun    WeaponType = "heavy-bowgun"
	Bow            WeaponType = "bow"
)

type Attack struct {
	Display int `json:"display"`
	Raw     int `json:"raw"`
}

type WeaponElement struct {
	Type   ElementType `json:"type"`
	Damage int         `json:"damage"`
	Hidden bool        `json:"hidden"`
}

type WeaponCraftingInfo struct {
	Craftable         bool           `json:"craftable"`
	Previous          int            `json:"previous"`
	Branches          []int          `json:"branches"`
	CraftingMaterials []CraftingCost `json:"craftingMaterials"`
	UpgradeMaterials  []CraftingCost `json:"upgradeMaterials"`
}

type WeaponAssets struct {
	Icon  string `json:"icon"`
	Image string `json:"image"`
}

type WeaponSharpness struct {
	Red    int `json:"red"`
	Orange int `json:"orange"`
	Yellow int `json:"yellow"`
	Green  int `json:"green"`
	Blue   int `json:"blue"`
	White  int `json:"white"`
	Purple int `json:"purple"`
}

type EldersealType string

const (
	Low     EldersealType = "low"
	Average EldersealType = "average"
	High    EldersealType = "high"
)

type DamageType string

const (
	Blunt    DamageType = "blunt"
	Piercing DamageType = "piercing"
	Slashing DamageType = "slashing"
)

type WeaponAttributes struct {
	Affinity int `json:"affinity"`
	Defense  int `json:"defense"`
}
