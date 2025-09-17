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
	Id                   int
	Name                 string
	Price                int
	Icon                 string
	AddHealth            int
	removeEnemyHealth    int
	TimeUsing            int
	giveInventory        int
	AddPvMax             int
	PtsAttack            int
	IsForgeron           bool
	IsUsableInGame       bool
	CancelPoison         bool
	RessourceForForgeron bool
	itemNeeded           int
	itemNeededQuantity   int
	RewardIG             bool
	canLooseItem         int
}

var allItems = map[int]Item{

	// ------------------- Ressources Marchand ↓

	1: {Id: 1, Name: "Pain sec", Price: 2, Icon: "🍞", IsForgeron: false, AddHealth: 5, IsUsableInGame: true},
	2: {Id: 2, Name: "Kebab du Destin", Price: 15, Icon: "🍔", IsForgeron: false, AddHealth: 15, IsUsableInGame: true},

	3: {Id: 3, Name: "8.6 tiède", Price: 4, Icon: "🍺", IsForgeron: false, AddHealth: 5, IsUsableInGame: true},
	4: {Id: 4, Name: "Canette de Kronenbourg", Price: 10, Icon: "🥤", IsForgeron: false, AddHealth: 10, IsUsableInGame: true},

	5: {Id: 5, Name: "Stylo BIC", Price: 3, Icon: "🖊", IsForgeron: false, IsUsableInGame: true},
	6: {Id: 6, Name: "Formulaire Cerfa 666-B", Price: 6, Icon: "📄", IsForgeron: false, RessourceForForgeron: true},
	7: {Id: 7, Name: "Carte Navigo périmée", Price: 1, Icon: "🎫", IsForgeron: false, RessourceForForgeron: true},
	8: {Id: 8, Name: "CV Légendaire", Price: 20, Icon: "📃", IsForgeron: true, PtsAttack: 10, itemNeeded: 4, itemNeededQuantity: 1}, // stylo BIC

	9:  {Id: 9, Name: "Sac à dos troué", Price: 20, Icon: "🎒", IsForgeron: false, giveInventory: 5, IsUsableInGame: false, canLooseItem: 20},
	10: {Id: 10, Name: "Sac à dos Décathlon", Price: 35, Icon: "🎒", IsForgeron: false, giveInventory: 10, IsUsableInGame: false},

	// Ressources de craft pour le Forgeron ↓

	11: {Id: 11, Name: "Laine de Chèvre", Price: 5, Icon: "🐐", IsForgeron: false, RessourceForForgeron: true, RewardIG: true},
	12: {Id: 12, Name: "Attestation Pôle-Emploi", Price: 7, Icon: "📑", IsForgeron: false, RessourceForForgeron: true, RewardIG: true},
	13: {Id: 13, Name: "Ticket resto usagé", Price: 2, Icon: "🎟", IsForgeron: false, RessourceForForgeron: true, RewardIG: true},
	14: {Id: 14, Name: "Badge CGT", Price: 10, Icon: "👑", IsForgeron: false, RessourceForForgeron: true, RewardIG: true},
	15: {Id: 15, Name: "Lingot de SMIC", Price: 12, Icon: "💰", IsForgeron: false, RessourceForForgeron: true, RewardIG: true},

	// ------------------- Ressources Forgeron ↓

	16: {Id: 16, Name: "Casquette du Chômeur", Price: 15, Icon: "🧢", IsForgeron: true, AddPvMax: 5, itemNeeded: 9, itemNeededQuantity: 1},             // laine de chèvre
	17: {Id: 17, Name: "Costume d’Entretien Froissé", Price: 20, Icon: "🤵", IsForgeron: true, AddPvMax: 10, itemNeeded: 10, itemNeededQuantity: 1},    // attestation
	18: {Id: 18, Name: "Chaussures de Sécurité Abîmées", Price: 30, Icon: "🥾", IsForgeron: true, AddPvMax: 15, itemNeeded: 11, itemNeededQuantity: 1}, // ticket resto

	19: {Id: 19, Name: "Épée en SMIC", Price: 25, Icon: "🗡️", IsForgeron: true, PtsAttack: 15, itemNeeded: 13, itemNeededQuantity: 1},       // lingot de SMIC
	20: {Id: 20, Name: "Arc de Syndicaliste", Price: 18, Icon: "🏹", IsForgeron: true, PtsAttack: 12, itemNeeded: 12, itemNeededQuantity: 1}, // badge CGT
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
	if c.GetItemNumber()+itemQuantity > c.MaxInventory {
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
	for {
		if len(c.Inventory) > 0 {
			fmt.Println("-------------------------------------")
			fmt.Println(" Votre Inventaire : (", c.GetItemNumber(), " / ", c.MaxInventory, " ) ")
			fmt.Println("-------------------------------------")
			for k, v := range c.Inventory {
				item := allItems[v.Id]
				fmt.Println(k+1, ") Item : ", item.Icon, " | ", item.Name, "x", v.Quantity, "/ Prix :", item.Price, "€")
			}
		} else {
			fmt.Println("---------------------------")
			fmt.Println("Aucun item pour le moment !")
			fmt.Println("---------------------------")
		}
		ChoiceUser := AskPlayerInt("Tapez le numéro d'un item pour intéragir avec (0 pour annuler) : ")

		if ChoiceUser == 0 {
			return
		}
	}
}

func (c *Character) UseItem(itemId, q int) {
	if !GetItemIdExist(itemId) {
		fmt.Println("Erreur dans l'item sélectionné !")
		return
	}

	item := allItems[itemId]

	if item.TimeUsing > 0 {
		time.Sleep(time.Duration(item.TimeUsing))
	}

	if item.AddHealth > 0 {

		if c.Pv+item.AddHealth <= c.PvMax {
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

	if item.canLooseItem > 0 {
		fmt.Println("Vous avez perdu un item de votre sac à dos troué")
	}

	if item.giveInventory > 0 {
		c.MaxInventory += item.giveInventory
		fmt.Printf("Vous avez équipé un sac à dos, votre inventaire est passé à %v slots\n", c.MaxInventory)
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
