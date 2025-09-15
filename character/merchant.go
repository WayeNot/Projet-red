package main

import (
	"fmt"
)

type Item struct {
	Name  string
	Price int
}

type Player struct {
	Name   string
	Gold   int
	PV     int
	Shield int
	Items  []Item
}

func main() {
	player := Player{Name: "Héros du PMU", Gold: 100, PV: 50, Shield: 20, Items: []Item{}}

	shop := []Item{
		{Name: "Potion de soin", Price: 0 },
		{Name: "Grande potion de soin", Price: 80 },
		{Name: "Bouclier", Price: 50 },
	}

	accessInventory(&player, shop)
}

func accessInventory(player *Player, shop []Item) {
	for {
		fmt.Println("\n--- Marchand ---")
		fmt.Println("1. Acheter un objet")
		fmt.Println("2. Vendre un objet")
		fmt.Println("3. Quitter")
		fmt.Print("Ton choix : ")

		var action int
		fmt.Scan(&action)

		switch action {
		case 1:
			buyItem(player, shop)
		case 2:
			sellItem(player)
		case 3:
			fmt.Println("À bientôt, mec !")
			return
		default:
			fmt.Println("Choix invalide !")
		}
	}
}

func buyItem(player *Player, shop []Item) {
	fmt.Println("\nVoici ce que j'ai à vendre :")
	for i, item := range shop {
		fmt.Printf("%d. %s (Prix : %d)\n", i+1, item.Name, item.Price)
	}

	fmt.Print("Que veux-tu acheter ? (numéro, 0 pour annuler) : ")
	var choix int
	fmt.Scan(&choix)

	if choix == 0 {
		return
	}
	if choix < 1 || choix > len(shop) {
		fmt.Println("Choix invalide !")
		return
	}

	itemChoisi := shop[choix-1]

	if player.Gold < itemChoisi.Price {
		fmt.Println("Tu n'as pas assez d'or !")
		return
	}

	player.Gold -= itemChoisi.Price
	player.Items = append(player.Items, itemChoisi)

	fmt.Printf("Tu as acheté : %s\n", itemChoisi.Name)
	fmt.Printf("Or restant : %d\n", player.Gold)
}

func sellItem(player *Player) {
	if len(player.Items) == 0 {
		fmt.Println("Tu n'as rien à vendre.")
		return
	}

	fmt.Println("\nVoici ton inventaire :")
	

	fmt.Print("Quel objet veux-tu vendre ? (numéro, 0 pour annuler) : ")
	var choix int
	fmt.Scan(&choix)

	if choix == 0 {
		return
	}
	if choix < 1 || choix > len(player.Items) {
		fmt.Println("Choix invalide !")
		return
	}

	itemVendu := player.Items[choix-1]
	player.Gold += itemVendu.Price / 2

	player.Items = append(player.Items[:choix-1], player.Items[choix:]...)

	fmt.Printf("Tu as vendu : %s\n", itemVendu.Name)
	fmt.Printf("Or actuel : %d\n", player.Gold)
}
