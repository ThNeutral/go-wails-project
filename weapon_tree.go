package main

import (
	"errors"
	"strconv"
)

type WeaponTreeNode struct {
	Prev *WeaponTreeNode   `json:"prev"`
	Next []*WeaponTreeNode `json:"next"`
	ID   int               `json:"id"`
	Name string            `json:"name"`
}

var visitedNodes map[int]bool

func generateWeaponTree(ID int) (*WeaponTreeNode, error) {
	headWeaponByID, err := fetchWeaponsByQuery(RequestQuery{Field: "id", Query: strconv.Itoa(ID)})
	if err != nil {
		return nil, err
	}
	if len(headWeaponByID) != 1 {
		return nil, errors.New("invalid ID: " + strconv.Itoa(ID))
	}

	headNode := &WeaponTreeNode{
		Prev: nil,
		Next: []*WeaponTreeNode{},
		ID:   headWeaponByID[0].ID,
		Name: headWeaponByID[0].Name,
	}

	visitedNodes = make(map[int]bool)

	err = buildPreviousTree(headNode, headWeaponByID[0])
	if err != nil {
		return nil, err
	}

	err = buildNextTree(headNode, headWeaponByID[0])
	if err != nil {
		return nil, err
	}

	return headNode, nil
}

func buildPreviousTree(parentNode *WeaponTreeNode, weapon Weapon) error {
	currentWeapon := weapon
	currentNode := parentNode

	for currentWeapon.Crafting.Previous.String() != "" {
		nextID, err := currentWeapon.Crafting.Previous.Int64()
		if err != nil {
			return err
		}

		if visitedNodes[int(nextID)] {
			break
		}
		visitedNodes[int(nextID)] = true

		nextWeaponByID, err := fetchWeaponsByQuery(RequestQuery{Field: "id", Query: strconv.Itoa(int(nextID))})
		if err != nil || len(nextWeaponByID) == 0 {
			return errors.New("invalid previous weapon ID: " + strconv.Itoa(int(nextID)))
		}

		newNode := &WeaponTreeNode{
			Prev: nil,
			Next: nil,
			ID:   nextWeaponByID[0].ID,
			Name: nextWeaponByID[0].Name,
		}
		currentNode.Prev = newNode
		currentNode = newNode
		currentWeapon = nextWeaponByID[0]
		buildNextTree(currentNode, currentWeapon)
	}

	return nil
}

func buildNextTree(parentNode *WeaponTreeNode, weapon Weapon) error {
	for _, nextID := range weapon.Crafting.Branches {
		if visitedNodes[nextID] {
			return nil
		}
		visitedNodes[nextID] = true

		nextWeaponByID, err := fetchWeaponsByQuery(RequestQuery{Field: "id", Query: strconv.Itoa(nextID)})
		if err != nil || len(nextWeaponByID) == 0 {
			return errors.New("invalid next weapon ID: " + strconv.Itoa(nextID))
		}

		nextWeapon := nextWeaponByID[0]
		newNode := &WeaponTreeNode{
			Prev: nil,
			Next: []*WeaponTreeNode{},
			ID:   nextWeapon.ID,
			Name: nextWeapon.Name,
		}

		parentNode.Next = append(parentNode.Next, newNode)
		err = buildNextTree(newNode, nextWeapon)
		if err != nil {
			return err
		}
	}

	return nil
}
