package main

type Slot struct {
	Rank int `json:"rank"`
}

type ElementType string

const (
	Fire      ElementType = "fire"
	Water     ElementType = "water"
	Ice       ElementType = "ice"
	Thunder   ElementType = "thunder"
	Dragon    ElementType = "dragon"
	Blast     ElementType = "blast"
	Poison    ElementType = "poison"
	Sleep     ElementType = "sleep"
	Paralysis ElementType = "paralysis"
	Stun      ElementType = "stun"
)

type CraftingCost struct {
	Quantity int  `json:"quantity"`
	Item     Item `json:"item"`
}

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rarity      int    `json:"rarity"`
	CarryLimit  int    `json:"carryLimit"`
	SellPrice   int    `json:"sellPrice"`
	BuyPrice    int    `json:"buyPrice"`
}
