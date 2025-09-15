package red

import (
	"fmt"
	"slices"
	"strings"
)

type Character struct {
	Name       string
	Pv         int
	Pv_max     int
	Shield     int
	Shield_max int
	Level      int
	Xp         int
	Money      int
	Is_dead    bool
	Inventory  []Inventory
}

type Inventory struct {
	Id       int
	Quantity int
}

type Item struct {
	Id    int
	Name  string
	Price int
	Icon  string
	addHealth int
}

type CharacterType struct {
	Id 		   int
	Name       string
	Pv         int
	Pv_max     int
	Shield     int
	Shield_max int
	Money      int
}

var chooseCharacter = map[int]CharacterType {
	1: {Id: 1, Name: ""},
}

var allItems = map[int]Item {
	1: {Id: 1, Name: "Pomme", Price: 10, Icon: "🍎", addHealth: 20},
	2: {Id: 2, Name: "Épée", Price: 350, Icon: "🗡️"},
}

func New(Name string, Pv, Pv_max, Shield, Shield_max, Level, Xp int, Money int, Is_dead bool, Inventory []Inventory) Character {
	return Character{Name, Pv, Pv_max, Shield, Shield_max, Level, Xp, Money, Is_dead, Inventory}
}

func InitCharacter(charName string) Character {
	var name string
	if charName == "" {
		name = AskPlayerString("Quel est le nom de votre joueur ? ")
		if len(name) > 0 {
            finalName := strings.ToUpper(name[:1]) + strings.ToLower(name[1:])
            name = finalName
        }
	} else {
        name = strings.ToUpper(charName[:1]) + strings.ToLower(charName[1:])
	}

	char := New(name, 100, 100, 0, 100, 1, 0, 0, false, []Inventory{})
	char.Pv = char.Pv_max / 2
	char.DisplayPlayer()
	return char
}

func (p Character) DisplayPlayer() {
	fmt.Println("--------------------------------------------------")
	fmt.Println("Nom du personnage : ", p.Name)
	fmt.Println("Vie / Bouclier du personnage : [", p.Pv, "|", p.Shield, "]")
	fmt.Println("Level / Xp : [", p.Level, " | ", p.Xp, "]")
	fmt.Println("Argent : ", p.Money)
	fmt.Println("--------------------------------------------------")
}

func (c *Character) SetPlayerDead() {
	c.Is_dead = true
	c.Pv = 0
	c.Shield = 0
}

func (c *Character) IsDead() {
	if c.Pv <= 0 {
		c.Pv = c.Pv_max / 2
	}
}

func (p Character) AccessInventory() {
	inv := p.Inventory
	if len(inv) > 0 {
		for _, v := range inv {
			item := allItems[v.Id]
			fmt.Println(item.Icon, " (",item.Name,") ", " - ", "x", v.Quantity)
		}
	} else {
		print("Aucun item pour le moment !")
	}
}

func (c *Character) AddItem(itemId, itemQuantity int) {
	c.Inventory = append(c.Inventory, Inventory{
		Id:       itemId,
		Quantity: itemQuantity,
	})
}

func (c *Character) RemoveItem(itemId, itemQuantity int) {
	for k, v := range c.Inventory {
		if itemId == v.Id {
			if v.Quantity == itemQuantity {
				c.Inventory = slices.Delete(c.Inventory, k, k+1)
			} else {
				c.Inventory[k].Quantity = v.Quantity - itemQuantity
			}
		}
	}
}

func getItemIdExist(itemId int) bool {
	_, exist := allItems[itemId]
	return exist
}

// 

func (c *Character) useItem(itemId int) {
	if !getItemIdExist(itemId) {
		fmt.Println("Erreur dans l'item sélectionné !")
		return
	}

	if allItems[itemId].addHealth != 0 {
		fmt.Println("Cet item n'est pas utilisable !")
		return
	}

	c.RemoveItem(itemId, 1)
	if c.Pv + allItems[itemId].addHealth <= c.Pv_max {
		c.Pv += allItems[itemId].addHealth
	} else {
		c.Pv = c.Pv_max
	}
}