package main

import (
	"fmt"
	"os"
	"red/game"
)

func main(){
	Introduction()
	player := red.InitCharacter("")
	red.ClearTerminal()
	player.DisplayPlayer()
	for {
		menu := red.Menu{
		Name: "Menu Principal - Chômage & Dragons",
		Choices: []red.Choice{
			{
				Label: "Afficher les stats du personnage",
				Action: func(p *red.Character) {
					red.ClearTerminal()
					p.DisplayPlayer()
				},
			},
			{
				Label: "Ouvrir l’inventaire",
				Action: func(p *red.Character) {
					red.ClearTerminal()
					p.AccessInventory()
				},
			},
			{
				Label: "Accéder au marchand",
				Action: func(p *red.Character) {
					red.ClearTerminal()
					p.MenuSellers()
				},
			},
			{
				Label: "Quitter",
				Action: func(p *red.Character) {
					fmt.Println("Merci d'avoir joué !")
					os.Exit(0)
				},
			},
		},
	}

	menu.Display(&player)
	}
}

func Introduction() {
	fmt.Println("===================================")
	fmt.Println("      📉 BIENVENUE DANS 📉")
	fmt.Println("     💀 CHÔMAGE & DRAGONS 💀")
	fmt.Println("La quête du loose, du PMU et du RSA")
	fmt.Println("===================================")
	fmt.Println("")
	fmt.Println("Tu incarnes un héros malgré lui, fraîchement inscrit à Pôle Emploi.")
	fmt.Println("Ton objectif ? Survivre à la jungle administrative,")
	fmt.Println("farmer des XP bureaucratiques, et peut-être, un jour, décrocher…")
	fmt.Println("un VRAI CDI Légendaire (Contrat à Durée Infinie).")
	fmt.Println("")
	fmt.Println("⚔️  OBJECTIFS :")
	fmt.Println(" - Crée ton personnage loser.")
	fmt.Println(" - Gère ton inventaire de galérien (8.6 tiède, CV moisi, etc).")
	fmt.Println(" - Affronte les boss ultimes : le DRH, le Banquier, et même…")
	fmt.Println("   le terrible Contrôleur de Pôle Emploi.")
	fmt.Println(" - Utilise tes potions de vie, tes sorts de clodo et ton skill naturel :")
	fmt.Println("   survivre avec 2€ sur le compte le 15 du mois.")
	fmt.Println("")
	fmt.Println("Tu es prêt ? Pas grave, personne ne l’est vraiment.")
	fmt.Println("Bonne chance, et que le RSA soit avec toi.")
	fmt.Println("===================================")
	fmt.Println("")
}