package main

import (
	"fmt"
	"os"
	red "red/game"
)

func main() {
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
					Label: "Centre d'échange",
					Action: func(p *red.Character) {
						red.ClearTerminal()
						p.MenuSellers()
					},
				},
				{
					Label: "Continuer les quêtes",
					Action: func(p *red.Character) {
						red.ClearTerminal()
						Quest1(&player)
					},
				},
				{
					Label: "Changer son pseudo (Cout 30€)",
					Action: func(p *red.Character) {
						p.EditName()
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
	fmt.Println()
	fmt.Println("Tu incarnes un héros malgré lui, fraîchement inscrit à Pôle Emploi.")
	fmt.Println("Ton objectif ? Survivre à la jungle administrative,")
	fmt.Println("farmer des XP bureaucratiques, et peut-être, un jour, décrocher…")
	fmt.Println("un VRAI CDI Légendaire (Contrat à Durée Infinie).")
	fmt.Println()
	fmt.Println("⚔️  OBJECTIFS :")
	fmt.Println(" - Crée ton personnage loser.")
	fmt.Println(" - Gère ton inventaire de galérien (8.6 tiède, CV moisi, etc).")
	fmt.Println(" - Affronte les boss ultimes : le DRH, le Banquier, et même…")
	fmt.Println("   le terrible Contrôleur de Pôle Emploi.")
	fmt.Println(" - Utilise tes potions de vie, tes sorts de clodo et ton skill naturel :")
	fmt.Println("   survivre avec 2€ sur le compte le 15 du mois.")
	fmt.Println()
	fmt.Println("Tu es prêt ? Pas grave, personne ne l’est vraiment.")
	fmt.Println("Bonne chance, et que le RSA soit avec toi.")
	fmt.Println("===================================")
	fmt.Println()
}

func Quest1(player *red.Character) {
	player.LooseItem()
	fmt.Println("Quête 1 : Pointer à Pôle Emploi")
	fmt.Println("===================================")
	fmt.Println("Il est lundi, il pleut, vous arrivez à Pôle Emploi et devez pointer pour signaler votre présence.")
	menu := red.Menu{
		Name: "Quête 1 :",
		Choices: []red.Choice{
			{
				Label: "Pointer",
				Action: func(p *red.Character) {
					fmt.Println("Vous êtes présent !")
					fmt.Println("+10 XP administratif et +10 € !")
					player.UpdateXp(10, "+")
					player.UpdateMoney(10, "+")
				},
			},
			{
				Label: "Ne pas pointer",
				Action: func(p *red.Character) {
					fmt.Println("Le manager vous appelle.. Il n'est pas content de votre absence !")
					fmt.Println("-10 €. Pôle Emploi vient de vous débiter 10 euros.. Et -10 XP administratif !")
					fmt.Println("Soyez présent à l'avenir !")
					player.UpdateMoney(10, "-")
					player.UpdateXp(10, "-")
				},
			},
		},
	}
	quest := red.InitQuest("Pointer à Pôle Emploi", 5, 0, 1, menu)
	quest.PlayQuest(player)
}

func Quest2(player *red.Character) {
	fmt.Println("Quête 2 : Acheter un kebab")
	fmt.Println("===================================")
	fmt.Println("Tu passes devant le Kebab en bas de ton appartement du Crous. Ça sent la sauce blanche divine.")
	menu := red.Menu{
		Name: "Quête 2 :",
		Choices: []red.Choice{
			{
				Label: "Acheter le Kebab (-15€)",
				Action: func(p *red.Character) {
					if p.Money >= 15 {
						p.UpdateMoney(15, "-")
						p.AddPV(100)
						fmt.Println("Tu manges le Kebab ! +100 PV d'énergie retrouvée !")
					} else {
						fmt.Println("La honte pas assez d'argent pour un Kebab, il est temps de trouver un taff..")
					}
				},
			},
			{
				Label: "Ignorer le kebab",
				Action: func(p *red.Character) {
					p.RemovePV(10)
					fmt.Println("🤢 T'as faim… -10 PV. Mais tu gardes tes sous.")
				},
			},
		},
	}
	quest := red.InitQuest("Acheter un kebab", 0, 0, 0, menu)
	quest.PlayQuest(player)
}

func Quest3(player *red.Character) {
	fmt.Println("Quête 3 : Affronter le Contrôleur Naolib")
	fmt.Println("===================================")
	fmt.Println("Tu montes dans le tram avec ta Carte Naolib périmée… Le contrôleur arrive !")
	menu := red.Menu{
		Name: "Quête 3 :",
		Choices: []red.Choice{
			{
				Label: "Tenter de négocier",
				Action: func(p *red.Character) {
					if p.Xp >= 20 {
						fmt.Println("Tu bluffes comme un pro. Le contrôleur t'épargne ! +10 XP.")
						p.UpdateXp(10, "+")
					} else {
						fmt.Println("Raté, tu te fais cramer. Amende : -20€.")
						p.UpdateMoney(20, "-")
					}
				},
			},
			{
				Label: "Fuir",
				Action: func(p *red.Character) {
					p.RemovePV(15)
					fmt.Println("🏃 Tu cours comme un dératé ! Tu t'en sors mais perds -15 PV.")
				},
			},
			{
				Label: "Accepter l’amende (-20€)",
				Action: func(p *red.Character) {
					p.UpdateMoney(20, "-")
					p.UpdateXp(10, "+")
					fmt.Println("Tu prends l’amende. -20€, mais +10 XP administratif pour la paperasse.")
				},
			},
		},
	}
	quest := red.InitQuest("Affronter le Contrôleur Naolib", 0, 0, 0, menu)
	quest.PlayQuest(player)
}

func Quest4(player *red.Character) {
	fmt.Println("Quête 4 : Premier entretien foireux")
	fmt.Println("===================================")
	fmt.Println("Un RH louche t’invite pour un entretien d’embauche...")
	menu := red.Menu{
		Name: "Quête 4 :",
		Choices: []red.Choice{
			{
				Label: "Sortir ton CV Légendaire",
				Action: func(p *red.Character) {
					// check inventaire ici si tu veux
					fmt.Println("Le RH est impressionné ! +20 XP.")
					p.UpdateXp(20, "+")
				},
			},
			{
				Label: "Dire n'importe quoi",
				Action: func(p *red.Character) {
					fmt.Println("Tu fais le con. Le RH se marre. -10 XP mais +10€ en pot-de-vin.")
					p.UpdateXp(10, "-")
					p.UpdateMoney(10, "+")
				},
			},
			{
				Label: "Ne pas venir",
				Action: func(p *red.Character) {
					fmt.Println("Tu dors chez toi. -15 XP, zéro réputation.")
					p.UpdateXp(15, "-")
				},
			},
		},
	}
	quest := red.InitQuest("Premier entretien foireux", 0, 0, 0, menu)
	quest.PlayQuest(player)
}

func Quest5(player *red.Character) {
	fmt.Println("Quête 5 : Le Banquier")
	fmt.Println("===================================")
	fmt.Println("Ton banquier t’appelle : « Monsieur, votre compte est à découvert »")
	menu := red.Menu{
		Name: "Quête 5 :",
		Choices: []red.Choice{
			{
				Label: "Supplier",
				Action: func(p *red.Character) {
					if p.Xp > 30 {
						fmt.Println("Le banquier est ému. +50€ sur ton compte !")
						p.UpdateMoney(50, "+")
					} else {
						fmt.Println("Le banquier s'en fout et bien fait pour toi, trouves un taff. -50€.")
						p.UpdateMoney(50, "-")
					}
				},
			},
			{
				Label: "L’arnaquer avec un faux Cerfa",
				Action: func(p *red.Character) {
					fmt.Println("Tu balances ton Cerfa 666-B. Le banquier panique. +100€ !")
					p.UpdateMoney(100, "+")
				},
			},
			{
				Label: "Fuir",
				Action: func(p *red.Character) {
					fmt.Println("Tu quittes l’agence. -10 PV mais tu gardes tes sous.")
					p.RemovePV(10)
				},
			},
		},
	}
	quest := red.InitQuest("Le Banquier", 0, 0, 0, menu)
	quest.PlayQuest(player)
}

func Quest6(player *red.Character) {
	fmt.Println("Quête 6 (Boss final) : Le Super Contrôleur de Pôle Emploi")
	fmt.Println("===================================")
	fmt.Println("⚔️ Le Super Contrôleur veut te RADIER. Combat tour par tour !")
	menu := red.Menu{
		Name: "Quête 6 :",
		Choices: []red.Choice{
			{
				Label: "Attaque au Stylo BIC (5 dmg)",
				Action: func(p *red.Character) {
					fmt.Println("🖊️ Tu plantes le Stylo BIC ! -5 PV au boss.")
				},
			},
			{
				Label: "Attaque à l’Épée en SMIC (15 dmg)",
				Action: func(p *red.Character) {
					fmt.Println("🗡️ Tu frappes avec l’Épée en SMIC ! -15 PV au boss.")
				},
			},
			{
				Label: "Utiliser un item (Pain, Canette, Kebab)",
				Action: func(p *red.Character) {
					fmt.Println("🎒 Tu fouilles ton sac… (à implémenter avec ton inventaire).")
				},
			},
			{
				Label: "Sort spécial : Appel à la grève (si Badge CGT)",
				Action: func(p *red.Character) {
					fmt.Println("👑 Tu invoques la CGT ! Le boss est stun un tour.")
				},
			},
			{
				Label: "Fuir",
				Action: func(p *red.Character) {
					fmt.Println("💀 Impossible, le Contrôleur est partout.")
				},
			},
		},
	}
	quest := red.InitQuest("Le Super Contrôleur de Pôle Emploi", 0, 0, 0, menu)
	quest.PlayQuest(player)
}
