package red

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomNbr(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max) + 1
}

func (c *Character) MenuSellers() {
	var playerChoice int

	for {
		ClearTerminal()
		fmt.Println("-------------------------------------")
		fmt.Println("    Voici les magasins disponibles   ")
		fmt.Println("-------------------------------------")
		fmt.Println("Voici les options disponibles :")

		fmt.Println(" - Appuyez sur 1 pour se rendre au Marchand\n - Appuyez sur 2 pour se rendre au Forgeron\n - Appuyez sur 0 pour revenir en arrière")
		fmt.Println("")
		fmt.Print("Votre choix : ")
		fmt.Scanln(&playerChoice)

		switch playerChoice {
		case 1:
			c.MenuMerchant()
		case 2:
			c.MenuForgeron()
		case 0:
			fmt.Println("Vous quittez le marchand.")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func (c *Character) MenuMerchant() {
	var playerChoice int

	for {
		ClearTerminal()
		fmt.Println("-------------------------------------")
		fmt.Println("     Bienvenue chez le Marchand      ")
		fmt.Printf("          Quantité de € : %v          \n", c.Money)
		fmt.Println("-------------------------------------")
		fmt.Println("Voici les options disponibles :")

		fmt.Println(" - Appuyez sur 1 pour acheter")
		if len(c.Inventory) > 0 {
			fmt.Println(" - Appuyez sur 2 pour vendre")
		}
		fmt.Println(" - Appuyez sur 0 pour quitter")
		fmt.Println("")
		fmt.Print("Votre choix : ")
		fmt.Scanln(&playerChoice)

		switch playerChoice {
		case 1:
			c.BuyMerchantItem()
		case 2:
			if len(c.Inventory) > 0 {
				c.sellMerchantItem()
			}
		case 0:
			fmt.Println("Vous quittez le marchand.")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func (c *Character) BuyMerchantItem() {
	for {
		allChoice := []int{}
		for len(allChoice) < 3 {
			randomNbr := RandomNbr(len(allItems))
			itemRnd := allItems[randomNbr]

			if itemRnd.IsForgeron || itemRnd.RewardIG {
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

		fmt.Println("-------------------------------------")
		fmt.Printf("          Quantité de € : %v          \n", c.Money)
		fmt.Println("   Articles disponibles à l’achat :  ")
		fmt.Println("-------------------------------------")

		fmt.Println("")

		for k, v := range allChoice {
			item := allItems[v]
			fmt.Println(k+1, ")", item.Icon, "|", item.Name, "- Prix :", item.Price, "€")
		}

		fmt.Println("")

		nbrChoice := AskPlayerInt("Quel article choisissez-vous (0 pour annuler) ? ")

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
			fmt.Println("€ - Vous n’avez pas assez d’argent !")
			return
		}
		if c.GetItemNumber() >= c.MaxInventory {
			fmt.Println("Vous n’avez plus de place dans votre inventaire !")
			return
		}

		c.AddItem(chosen.Id, 1)
		c.UpdateMoney(chosen.Price, "-")
		fmt.Printf("Vous avez acheté : %v | %v x1\n", chosen.Name, chosen.Icon)

		if !chosen.IsUsableInGame {
			c.UseItem(chosen.Id, 1)
		}
	}
}

func (c *Character) sellMerchantItem() {
	ClearTerminal()

	if len(c.Inventory) == 0 {
		fmt.Println("Aucun item à vendre.")
		return
	}

	for len(c.Inventory) > 0 {
		if len(c.Inventory) <= 0 {
			return
		}

		ClearTerminal()

		fmt.Println("-------------------------------------")
		fmt.Println(" Votre Inventaire : (", c.GetItemNumber(), "/", c.MaxInventory, ") ")
		fmt.Println("-------------------------------------")

		for k, v := range c.Inventory {
			item := allItems[v.Id]
			fmt.Println(k+1, ")", item.Icon, "|", item.Name, "x", v.Quantity, "- Prix de vente :", item.Price/2, "€")
		}

		userChoice := AskPlayerInt("Tapez le numéro de l’objet à vendre (0 pour annuler) : ")
		var itemNbr int
		chosenInv := c.Inventory[userChoice-1]
		item := allItems[chosenInv.Id]

		if userChoice == 0 {
			fmt.Println("Vente annulée.")
			return
		}

		if userChoice < 1 || userChoice > len(c.Inventory) {
			fmt.Println("Choix invalide.")
			return
		}

		if chosenInv.Quantity > 1 {
			itemNbr = AskPlayerInt("Combien voulez-vous en vendre ?")
		} else {
			itemNbr = 1
		}

		price := (item.Price / 2) * itemNbr

		c.RemoveItem(item.Id, itemNbr)
		c.UpdateMoney(price, "+")
		fmt.Printf("Vous avez vendu %v %v x %v et gagné %d€\n", item.Icon, item.Name, itemNbr, price)
	}
}

func (c *Character) MenuForgeron() {
	var playerChoice int

	for {
		ClearTerminal()

		fmt.Println("-------------------------------------")
		fmt.Println("     Bienvenue chez le forgeron      ")
		fmt.Println("-------------------------------------")
		fmt.Println("Voici les options disponibles :")

		fmt.Println(" - Appuyez sur 1 pour acheter")
		fmt.Println(" - Appuyez sur 0 pour quitter")

		fmt.Print("Votre choix : ")
		fmt.Scanln(&playerChoice)

		switch playerChoice {
		case 1:
			c.BuyForgeronItem()
		case 0:
			fmt.Println("Vous quittez le forgeron.")
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func (c *Character) BuyForgeronItem() {
	for {
		allChoice := []int{}
		for len(allChoice) < 3 {
			randomNbr := RandomNbr(len(allItems))

			if !allItems[randomNbr].IsForgeron {
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

		fmt.Println("-------------------------------------")
		fmt.Printf("          Quantité de € : %v          \n", c.Money)
		fmt.Println("   Articles disponibles à l’achat :  ")
		fmt.Println("-------------------------------------")

		fmt.Println("")

		for k, v := range allChoice {
			item := allItems[v]
			itemNeed := allItems[item.itemNeeded]
			fmt.Println(k+1, ")", item.Icon, "|", item.Name, "- Prix :", item.Price, "€ | Item nécessaire : ", itemNeed.Icon, " | ", itemNeed.Name, "x", item.itemNeededQuantity)
		}

		fmt.Println("")

		nbrChoice := AskPlayerInt("Quel article choisissez-vous (0 pour annuler) ?")

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
			fmt.Println("€ - Vous n’avez pas assez d’argent !")
			return
		}

		if c.GetItemNumber() >= c.MaxInventory {
			fmt.Println("Vous n’avez plus de place dans votre inventaire !")
			return
		}

		itemNeeded := allItems[allItems[nbrChoice].itemNeeded].Id

		for _, v := range c.Inventory {
			if v.Id == itemNeeded {
				c.AddItem(chosen.Id, 1)
				c.UpdateMoney(chosen.Price, "-")
				fmt.Printf("Vous avez acheté : %v | %v x1\n", chosen.Name, chosen.Icon)
				return
			}
		}

		fmt.Println("Vous n'avez pas les items requis !")
	}
}
