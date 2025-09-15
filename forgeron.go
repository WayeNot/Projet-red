package main

import (
	"fmt"
)

type Item struct {
	Id    int
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
	player := Player{Name: "Héros du PMU", Gold: 100, PV: 50, Shield: 20, Items: []Item{
		{Id: 1, Name: "Épée cassée", Price: 20},
	}}

	shop := []Item{
		{Id: 2, Name: "Grand bouclier", Price: 40},
		{Id: 3, Name: "Épée du CDD (requis éppée cassée)", Price: 100},
		{Id: 4, Name: "Épée courte", Price: 50},
	}

	accessInventory(&player, shop)
}

func accessInventory(player *Player, shop []Item) {
	for {
		fmt.Println("\n------ Forge ------")
		fmt.Println("1. Fabriquer un équipement")
		fmt.Println("2. Échanger un objet")
		fmt.Println("3. Quitter")
		fmt.Print("Ton choix : ")

		var action int
		fmt.Scan(&action)

		switch action {
		case 1:
			buyItem(player, shop)
		case 2:
			exchangeItem(player, shop)
		case 87:
			fmt.Println("Go c'est moins bien que le Python !")
		case 3:
			fmt.Println("Tu quittes la Forge !")
			return
		default:
			fmt.Println("Choix invalide !")
		}
	}
}

func buyItem(player *Player, shop []Item) {
	fmt.Println("\nVoici ce que tu peux fabriquer :")
	for i, item := range shop {
		fmt.Printf("%d. %s (Prix : %d)\n", i+1, item.Name, item.Price)
	}

	fmt.Print("Quel objet veux-tu acheter ? (0 pour annuler) : ")
	var choice int
	fmt.Scan(&choice)

	if choice <= 0 || choice > len(shop) {
		fmt.Println("Achat annulé.")
		return
	}

	item := shop[choice-1]
	if player.Gold >= item.Price {
		player.Gold -= item.Price
		player.Items = append(player.Items, item)
		fmt.Printf("Tu as fabriqué : %s ! (Or restant : %d)\n", item.Name, player.Gold)
	} else {
		fmt.Println("Tu n'as pas assez d'or !")
	}
}

func exchangeItem(player *Player, shop []Item) {
	if len(player.Items) == 0 {
		fmt.Println("Tu n'as aucun objet à échanger !")
		return
	}

	fmt.Println("\nTes objets :")
	for i, item := range player.Items {
		fmt.Printf("%d. %s\n", i+1, item.Name)
	}

	fmt.Print("Quel objet veux-tu échanger ? (0 pour annuler) : ")
	var playerChoice int
	fmt.Scan(&playerChoice)

	if playerChoice <= 0 || playerChoice > len(player.Items) {
		fmt.Println("Échange annulé.")
		return
	}

	playerItem := player.Items[playerChoice-1]

	fmt.Println("\nObjets disponibles chez le forgeron :")
	for i, item := range shop {
		fmt.Printf("%d. %s (Prix : %d)\n", i+1, item.Name, item.Price)
	}

	fmt.Print("Contre quel objet veux-tu échanger ? (0 pour annuler) : ")
	var shopChoice int
	fmt.Scan(&shopChoice)

	if shopChoice <= 0 || shopChoice > len(shop) {
		fmt.Println("Échange annulé.")
		return
	}

	shopItem := shop[shopChoice-1]

	player.Items[playerChoice-1] = shopItem

	fmt.Printf("Tu as échangé ton %s contre un %s !\n", playerItem.Name, shopItem.Name)
}
