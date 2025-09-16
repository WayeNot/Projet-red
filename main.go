package main

import (
	"fmt"
	"red/game"
)

func main(){
	fmt.Println("-------------------------------------")
	fmt.Println("	  Bienvenue dans ")
	fmt.Println("	 CHÔMAGE & DRAGONS")
	fmt.Println("La quête du loose, du PMU et du RSA")
	fmt.Println("-------------------------------------")
	player := red.InitCharacter("")
	player.DisplayPlayer()

	menu := red.Menu{
		Name: "Menu Principal - Chômage & Dragons",
		Choices: []red.Choice{
			{
				Label: "Afficher les stats du personnage",
				Action: func(p *red.Character) {
					p.DisplayPlayer()
				},
			},
			{
				Label: "Ouvrir l’inventaire",
				Action: func(p *red.Character) {
					p.AccessInventory()
				},
			},
			{
				Label: "Accéder au marchand",
				Action: func(p *red.Character) {
					p.MenuMerchant()
				},
			},
			{
				Label: "Quitter",
				Action: func(p *red.Character) {
					fmt.Println("jsp")
				},
			},
		},
	}

	menu.Display(&player)
}