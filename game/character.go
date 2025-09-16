package red

import (
	"fmt"
	"slices"
	"strings"
)

type Character struct {
	Name       string
	Pv         int
	PvMax     int
	Shield     int
	Shield_max int
	Level      int
	Xp         int
	Money      int
	IsDead    bool
	MaxInventory int
	Inventory  []Inventory
}

type Inventory struct {
	Id       int
	Quantity int
}

type Item struct {
	Id    				int
	Name  				string
	Price 				int
	Icon  				string
	addHealth 			int
	removeEnemyHealth 	int
	giveInventory		int
	IsForgeron 			bool
	IsUsable 			bool
}

func New(Name string, Pv, PvMax, Shield, Shield_max, Level, Xp int, Money int, IsDead bool, MaxInventory int, Inventory []Inventory) Character {
	return Character{Name, Pv, PvMax, Shield, Shield_max, Level, Xp, Money, IsDead, MaxInventory, Inventory}
}

func InitCharacter(charName string) Character {
	name := ""
	if charName == "" {
		fmt.Println("\n")
		name = AskPlayerString("Comment vous appelez vous ?")
		if len(name) > 0 {
            finalName := strings.ToUpper(name[:1]) + strings.ToLower(name[1:])
            name = finalName
        }
	} else {
        name = strings.ToUpper(charName[:1]) + strings.ToLower(charName[1:])
	}
	fmt.Println("\n")

	char := New(name, 100, 100, 0, 100, 1, 0, 100, false, 10, []Inventory{})
	char.Pv = char.PvMax / 2
	return char
}

func (p Character) DisplayPlayer() {
	fmt.Println("--------------------------------------------------")
	fmt.Println("Nom du personnage : ", p.Name)
	fmt.Println("Vie / Bouclier du personnage : [", p.Pv, "|", p.Shield, "]")
	fmt.Println("Level / Xp : [", p.Level, " | ", p.Xp, "]")
	fmt.Println("Argent : ", p.Money, "₣")
	fmt.Println("--------------------------------------------------")
}

func (c *Character) HandleDeath() {
	if c.Pv <= 0 {
		c.IsDead = true
		fmt.Println(c.Name, "est mort... mais ressuscite à 50% de sa vie !")
		c.Pv = c.PvMax / 2
		c.IsDead = false
	}
}

func (c Character) GetItemNumber() int {
	result := 0
	for _, v := range c.Inventory {
		result += v.Quantity
	}
	return result
}

func (c Character) AccessInventory() {
	inv := c.Inventory

	fmt.Println("-------------------------------------")
	fmt.Println(" Votre Inventaire : (", c.GetItemNumber(), " / ", c.MaxInventory, " ) ")
	fmt.Println("-------------------------------------")
	if len(inv) > 0 {
		for k, v := range inv {
			item := allItems[v.Id]
			fmt.Println(k + 1,") Item : ", item.Icon, " | ", item.Name, "x", v.Quantity, "/ Prix :", item.Price, "₣")
		}
	} else {
		fmt.Println("Aucun item pour le moment !")
	}
	fmt.Println("---------------------------")
}

func (c *Character) AddItem(itemId, itemQuantity int) {
	if c.GetItemNumber() >= c.MaxInventory {
		fmt.Println("Impossible d'ajouter cet item, vous n'avez plus de places !")
		return
	}
	for i, v := range c.Inventory {
		if v.Id == itemId {
			c.Inventory[i].Quantity += itemQuantity
			return
		}
	}
	c.Inventory = append(c.Inventory, Inventory{
		Id:       itemId,
		Quantity: itemQuantity,
	})
}

func (c *Character) RemoveItem(itemId, itemQuantity int) {
	for k, v := range c.Inventory {
		if v.Id == itemId {
			if itemQuantity >= v.Quantity {
				c.Inventory = slices.Delete(c.Inventory, k, k+1)
			} else {
				c.Inventory[k].Quantity -= itemQuantity
			}
			return
		}
	}
	fmt.Println("Vous n’avez pas cet objet dans votre inventaire.")
}

func GetItemIdExist(itemId int) bool {
	_, exist := allItems[itemId]
	return exist
}

// 

func (c *Character) UseItem(itemId, q int) {
	if !GetItemIdExist(itemId) {
		fmt.Println("Erreur dans l'item sélectionné !")
		return
	}

	item := allItems[itemId]

	if !item.IsUsable {
		fmt.Println("Cet item n'est pas utilisable !")
		return
	}

	if item.addHealth > 0 {
		if c.Pv + item.addHealth <= c.PvMax {
			c.Pv += item.addHealth
		} else {
			c.Pv = c.PvMax
		}
	}

	if item.giveInventory > 0 {
		c.MaxInventory += item.giveInventory
	}

	if item.removeEnemyHealth > 0 {
		fmt.Println("Vous attaqué l'ennemi !")
	}

	for _, n := range c.Inventory {
		if n.Id == itemId && n.Quantity >= q {
			c.RemoveItem(itemId, q)
			return
		}
	}
	fmt.Println("Vous n'avez pas assez de cet item !")
}

func (c *Character) UpdateMoney(q int, s string) {
	switch s {
		case "-":
			if (c.Money - q) >= 0 {
				c.Money -= q
			}
		case "+":
			c.Money += q
	}
}
