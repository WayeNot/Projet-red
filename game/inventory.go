package red

import (
	"fmt"
	"slices"
	"time"
)

type Inventory struct {
	Id       int
	Quantity int
}

type Item struct {
	Id    					int
	Name  					string
	Price 					int
	Icon  					string
	AddHealth 				int
	removeEnemyHealth 		int
	TimeUsing				int
	giveInventory			int
	AddPvMax				int
	PtsAttack				int
	IsForgeron 				bool
	IsUsable 				bool
	CancelPoison			bool
	RessourceForForgeron 	bool
}

var allItems = map[int]Item{
	1: {Id: 1, Name: "Potion de soin", Price: 3, Icon: "➕", IsForgeron: false, AddHealth: 15, IsUsable: true},
	2: {Id: 2, Name: "Grande Potion de soin", Price: 7, Icon: "💖", IsForgeron: false, AddHealth: 35, IsUsable: true},
	3: {Id: 3, Name: "Potion de soin suprême", Price: 20, Icon: "✨", IsForgeron: false, AddHealth: 50, IsUsable: true},

	4: {Id: 4, Name: "Potion de poison", Price: 6, Icon: "🤢", IsForgeron: false, removeEnemyHealth: 10, TimeUsing: 3, IsUsable: true},
	5: {Id: 5, Name: "Potion d’antidote", Price: 6, Icon: "🤢", IsForgeron: false, CancelPoison: true, IsUsable: true},

	6: {Id: 6, Name: "Pomme", Price: 2, Icon: "🍎", IsForgeron: false, AddHealth: 5, IsUsable: true},
	7: {Id: 7, Name: "Pain", Price: 4, Icon: "🍞", IsForgeron: false, AddHealth: 10, IsUsable: true},
	8: {Id: 8, Name: "Viande séchée", Price: 6, Icon: "🍖", IsForgeron: false, AddHealth: 20, IsUsable: true},

	9: {Id: 9, Name: "Sac à dos", Price: 30, Icon: "🎒", IsForgeron: false, giveInventory: 10, IsUsable: true},

	// Ressources de craft pour le Forgeron ↓

	10: {Id: 10, Name: "Plume de corbeau", Price: 1, Icon: "🪶", IsForgeron: false, RessourceForForgeron: true},
	11: {Id: 11, Name: "Cuir de sanglier", Price: 3, Icon: "🐗", IsForgeron: false, RessourceForForgeron: true},
	12: {Id: 12, Name: "Fourrure de loup", Price: 4, Icon: "🐺", IsForgeron: false, RessourceForForgeron: true},
	13: {Id: 13, Name: "Peau de troll", Price: 7, Icon: "👹", IsForgeron: false, RessourceForForgeron: true},
	14: {Id: 14, Name: "Lingot de fer", Price: 6, Icon: "⛓️", IsForgeron: false, RessourceForForgeron: true},
	15: {Id: 15, Name: "Bois solide", Price: 3, Icon: "🪵", IsForgeron: false, RessourceForForgeron: true},
	16: {Id: 16, Name: "Bois flexible", Price: 2, Icon: "🌿", IsForgeron: false, RessourceForForgeron: true},
	17: {Id: 17, Name: "Corde", Price: 2, Icon: "🧵", IsForgeron: false, RessourceForForgeron: true},

	// ------------------- Ressources Forgeron ↓

	18: {Id: 18, Name: "Chapeau de l’aventurier", Price: 5, Icon: "🎩", IsForgeron: true, AddPvMax: 10},
	19: {Id: 19, Name: "Tunique de l’aventurier", Price: 5, Icon: "🥋", IsForgeron: true, AddPvMax: 25},
	20: {Id: 20, Name: "Bottes de l’aventurier", Price: 5, Icon: "🥾", IsForgeron: true, AddPvMax: 15},

	21: {Id: 21, Name: "Épée en fer", Price: 20, Icon: "🗡️", IsForgeron: true, PtsAttack: 15},
	22: {Id: 22, Name: "Arc de chasseur", Price: 15, Icon: "🏹", IsForgeron: true, PtsAttack: 10},

}

func (c Character) GetItemNumber() int {
	result := 0
	for _, v := range c.Inventory {
		result += v.Quantity
	}
	return result
}

func GetItemIdExist(itemId int) bool {
	_, exist := allItems[itemId]
	return exist
}

func (c *Character) AddItem(itemId, itemQuantity int) {
	if c.GetItemNumber() + itemQuantity > c.MaxInventory {
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

	if item.TimeUsing > 0 {
		time.Sleep(time.Duration(item.TimeUsing))
	}

	if item.AddHealth > 0 {
		
		if c.Pv + item.AddHealth <= c.PvMax {
			c.Pv += item.AddHealth
		} else {
			c.Pv = c.PvMax
		}
		fmt.Printf("Vous avez récupéré %v PV", item.AddHealth)
	}

	if item.AddPvMax > 0 {
		c.PvMax += item.AddPvMax
		fmt.Printf("Vous avez %v points de vies supplémentaires !", item.AddPvMax)
	}

	if item.giveInventory > 0 {
		c.MaxInventory += item.giveInventory
		fmt.Printf("Votre inventaire augmente de %v slots", item.giveInventory)
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