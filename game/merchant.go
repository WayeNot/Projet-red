package red

import (
	"fmt"
	"math/rand"
	"time"
)

func (c *Character) MenuMerchant() {
	canSell := false
	canBuy := false

	if len(c.Inventory) > 0 {
		canSell = true
	}
	
	if c.Money > 0 {
		canBuy = true
	}

	fmt.Print("\033[H\033[2J")

	fmt.Println("-------------------------------------")
	fmt.Println("     Bienvenue chez le marchand      ")
	fmt.Println(" Tapper quit pour revenir en arrière ")
	fmt.Println("-------------------------------------")

	fmt.Println(" Voici les options disponibles : ")

	if canSell && canBuy {
		fmt.Println("")
		fmt.Println(" Appuyez sur 1 pour acheter un item ")
		fmt.Println(" Appuyez sur 2 pour vendre un item ")
		fmt.Println("---------------------------------")
	} else if canSell && !canBuy {
		fmt.Println("")
		fmt.Println(" Appuyez sur 1 pour vendre un item ")
		fmt.Println("-----------------------------------")
	} else if !canSell && canBuy {
		fmt.Println("")
		fmt.Println(" Appuyez sur 1 pour acheter un item ")
		fmt.Println("-----------------------------------")
	} else {
		fmt.Println("---------------------------------------")
		fmt.Println(" Vous ne pouvez ni vendre ni acheter ! ")
		fmt.Println("---------------------------------------")
	}
}

func RandomNbr(max int) int {
	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(max)
	return randomNumber
}

func (c *Character) BuyMerchantItem() {
	allChoice := []int{}
	countChoices := 0
	var nbrChoice int
	count := 0

	for countChoices <= 2 {
		countChoices += 1
		allChoice = append(allChoice, RandomNbr(len(allItems) - 1))
	}

	for k,v := range allChoice {
		count += 1
		fmt.Println(k + 1,") Item : ", allItems[v].Icon, " | ", allItems[v].Name, "/ Prix :", allItems[v].Price, "₣")
	}

	fmt.Println("Quelle item choisissez-vous ?", 1, " - ", count, "/ (quit) pour sortir")
	fmt.Scan(&nbrChoice)
	
	for _, v := range allChoice {
		if nbrChoice == v {
			c.AddItem(nbrChoice, 1)
			c.UpdateMoney(allItems[v].Price, "-")
		}
	}
}

func (c *Character) sellMerchantItem() {
	inv := c.Inventory
	MaxInventoryItems := 10
	var userChoice int

	fmt.Print("\033[H\033[2J")
	fmt.Println("-------------------------------------")
	fmt.Println("    Votre Inventaire : (", c.GetItemNumber(), " / ", MaxInventoryItems, " )      ")
	fmt.Println(" Tapper quit pour revenir en arrière ")
	fmt.Println("-------------------------------------")

	if len(inv) > 0 {
		for k, v := range inv {
			item := allItems[v.Id]
			fmt.Println(k + 1,") Item : ", item.Icon, " | ", item.Name, "/ Prix :", item.Price, "₣")
		}
		for {
			fmt.Println("Tapez le nombre de gauche pour vendre cet objet !")
			fmt.Scanln(&userChoice)

			item, e := allItems[userChoice]

			if e {
				c.RemoveItem(item.Id, 1)
				c.UpdateMoney(item.Price / (15 * 100), "+")
				fmt.Printf("[ Item choisi | Récompense de vente ] - [ %v | %v ]", item.Icon, item.Price)
			} else {
				fmt.Println("Cet item est inexistant !")
			}
		}
	} else {
		fmt.Println("Aucun item pour le moment !")
	}
	fmt.Println("---------------------------")
}
