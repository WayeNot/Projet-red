package main

import (
	"fmt"
	"red/game"
)

func main(){
	fmt.Println("_______________________________")
	fmt.Println("	  Bienvenue dans ")
	fmt.Println("	 CHÔMAGE & DRAGONS")
	fmt.Println("La quête du loose, du PMU et du RSA")
	fmt.Println("_______________________________")
	name := red.AskPlayerString("Comment vous appelez vous ?")
	fmt.Println(name)
	player := red.InitCharacter(name)
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
					fmt.Println("Bientot sac a merde")
				},
			},
			{
				Label: "NTM",
				Action: func(p *red.Character) {
					fmt.Println("Bientot sac a merde")
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