package red

import "fmt"

// Slot_max 	int -- Pour l'inventaire

type Character struct {
	Name       	string
	Pv         	int
	Pv_max     	int
	Shield     	int
	Shield_max 	int
	Level      	int
	Xp         	int
	Money      	int
	Is_dead    	bool
	Inventory  	[]Inventory
}

type Inventory struct {
	Id			int
	Quantity 	int
}

type Item struct {
	Id 			int
	Name     	string
	Price    	int
}

var allItems = map[int]Item{
	1: {Id: 1, Name: "Pomme", Price: 10},
	2: {Id: 2, Name: "Épée", Price: 350},
}

func strTest(placeHolder string) string {
	var response string
	fmt.Println(placeHolder)
	fmt.Scanln(&response)
	return response
}

func New(Name string, Pv, Pv_max, Shield, Shield_max, Level, Xp int, Money int, Is_dead bool, Inventory []Inventory) Character {
	return Character{Name, Pv, Pv_max, Shield, Shield_max, Level, Xp, Money, Is_dead, Inventory}
}

func initCharacter() {
	name := strTest("Quel est le nom de votre personnage ?")

	char := New(name, 100, 100, 100, 100, 1, 0, 1000, false, []Inventory{})
	char.displayPlayer()
	char.accessInventory()
}

func (p Character) displayPlayer() {
	fmt.Println("--------------------------------------------------")
	fmt.Println("Nom du personnage : ", p.Name)
	fmt.Println("Vie / Bouclier du personnage : [", p.Pv, "|", p.Shield, "]")
	fmt.Println("Level / Xp : [", p.Level, " | ", p.Xp, "]")
	fmt.Println("Argent : ", p.Money)
}

func (p Character) accessInventory() {
	inv := p.Inventory
	if len(inv) > 0 {
		print("Yes")
	} else {
		print("Aucun item pour le moment !")
	}
}

// func red() {
// 	initCharacter()
// 	Character.accessInventory(Character{})
// }
