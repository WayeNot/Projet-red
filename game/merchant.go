package red

import (
	"fmt"
	"math/rand"
	"time"
)

var allItems = map[int]Item{
	// Potions ↓

	1: {Id: 1, Name: "Potion de vie", Price: 3, Icon: "➕", IsForgeron: false, addHealth: 15},
	2: {Id: 2, Name: "Potion de poison", Price: 6, Icon: "🤢", IsForgeron: false, removeEnemyHealth: 10},

	4: {Id: 4, Name: "Sac à dos", Price: 30, Icon: "🎒", IsForgeron: false, giveInventory: 10},
}

func RandomNbr(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max) + 1 // +1 pour éviter l’index 0
}

func (c *Character) MenuMerchant() {
	var playerChoice int

	fmt.Println("-------------------------------------")
	fmt.Println("     Bienvenue chez le Marchand      ")
	fmt.Println("-------------------------------------")
	fmt.Println("Voici les options disponibles :")

	fmt.Println(" - Appuyez sur 1 pour acheter")
	fmt.Println(" - Appuyez sur 2 pour vendre")
	fmt.Println(" - Appuyez sur 0 pour quitter")

	for {
		fmt.Print("Votre choix : ")
		fmt.Scanln(&playerChoice)

		switch playerChoice {
		case 1:
			c.BuyMerchantItem()
		case 2:
			c.sellMerchantItem()
		case 0:
			fmt.Println("Vous quittez le marchand.")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func (c *Character) BuyMerchantItem() {
	allChoice := []int{}

	for len(allChoice) < 3 {
		randomNbr := RandomNbr(len(allItems))
		
		if allItems[randomNbr].IsForgeron {
			continue
		}

		exists := false

		for _, v := range allChoice {
			if v == randomNbr {
				exists = true
			}
		}
		if !exists {
			allChoice = append(allChoice, randomNbr)
		}
	}

	fmt.Println("Articles disponibles à l’achat :")
	for k, v := range allChoice {
		item := allItems[v]
		fmt.Println(k+1, ")", item.Icon, "|", item.Name, "- Prix :", item.Price, "₣")
	}

	var nbrChoice int
	fmt.Print("Quel article choisissez-vous (0 pour annuler) ? ")
	fmt.Scan(&nbrChoice)

	if nbrChoice == 0 {
		fmt.Println("Achat annulé.")
		return
	}
	if nbrChoice < 1 || nbrChoice > len(allChoice) {
		fmt.Println("Choix invalide.")
		return
	}

	chosenId := allChoice[nbrChoice-1]
	chosen := allItems[chosenId]

	if c.Money < chosen.Price {
		fmt.Println("₣ - Vous n’avez pas assez d’argent !")
		return
	}
	if c.GetItemNumber() >= c.MaxInventory {
		fmt.Println("Vous n’avez plus de place dans votre inventaire !")
		return
	}

	c.AddItem(chosen.Id, 1)
	c.UpdateMoney(chosen.Price, "-")
	fmt.Println("Vous avez acheté :", chosen.Icon, chosen.Name)
}

func (c *Character) sellMerchantItem() {
	inv := c.Inventory
	if len(inv) == 0 {
		fmt.Println("Aucun item à vendre.")
		return
	}

	fmt.Println("Votre inventaire :")
	for k, v := range inv {
		item := allItems[v.Id]
		fmt.Println(k+1, ")", item.Icon, "|", item.Name, "x", v.Quantity, "- Prix de vente :", item.Price/2, "₣")
	}

	var userChoice int
	fmt.Print("Tapez le numéro de l’objet à vendre (0 pour annuler) : ")
	fmt.Scanln(&userChoice)

	if userChoice == 0 {
		fmt.Println("Vente annulée.")
		return
	}
	if userChoice < 1 || userChoice > len(inv) {
		fmt.Println("Choix invalide.")
		return
	}

	chosenInv := inv[userChoice-1]
	item := allItems[chosenInv.Id]
	price := item.Price / 2

	c.RemoveItem(item.Id, 1)
	c.UpdateMoney(price, "+")
	fmt.Printf("Vous avez vendu %v %v et gagné %d₣\n", item.Icon, item.Name, price)
}
