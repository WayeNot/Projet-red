package red

import (
	"fmt"
)

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
			return
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