package red

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomNbr(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max) + 1 // +1 pour éviter l’index 0
}

func (c *Character) MenuSellers() {
	var playerChoice int

	for {
		fmt.Print("\033[H\033[2J")
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
		case 2 :
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
		fmt.Print("\033[H\033[2J")
		fmt.Println("-------------------------------------")
		fmt.Println("     Bienvenue chez le Marchand      ")
		fmt.Printf("          Quantité de ₣ : %v          \n", c.Money)
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
		case 2 :
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
	allChoice := []int{}

	for {
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
		fmt.Printf("          Quantité de ₣ : %v          \n", c.Money)
		fmt.Println("   Articles disponibles à l’achat :  ")
		fmt.Println("-------------------------------------")

		fmt.Println("")
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
		fmt.Println("Vous avez acheté : x1", chosen.Icon, chosen.Name)
	}
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

func (c *Character) MenuForgeron() {
	var playerChoice int

	fmt.Println("-------------------------------------")
	fmt.Println("     Bienvenue chez le forgeron      ")
	fmt.Println("-------------------------------------")
	fmt.Println("Voici les options disponibles :")

	fmt.Println(" - Appuyez sur 1 pour acheter")
	fmt.Println(" - Appuyez sur 0 pour quitter")

	for {
		fmt.Print("Votre choix : ")
		fmt.Scanln(&playerChoice)

		switch playerChoice {
		case 1:
			c.BuyForgeronItem()
		case 0:
			fmt.Println("Vous quittez le forgeron.")
			break
			
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func (c *Character) BuyForgeronItem() {
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

	fmt.Println("Articles disponibles à l’achat :")
	for k, v := range allChoice {
		item := allItems[v]
		itemNeed := allItems[item.itemNeeded]
		fmt.Println(k+1, ")", item.Icon, "|", item.Name, "- Prix :", item.Price, "₣ | Item nécessaire : ", itemNeed.Icon, " | ", itemNeed.Name, "x", item.itemNeededQuantity)
	}

	var nbrChoice int

	for {
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

		itemNeeded := allItems[allItems[nbrChoice].itemNeeded].Id
		for _, v := range c.Inventory {
			if v.Id == itemNeeded {
				c.AddItem(chosen.Id, 1)
				c.UpdateMoney(chosen.Price, "-")		
				break
			}
		}

		fmt.Println("Vous n'avez pas les items requis !")
	}
}