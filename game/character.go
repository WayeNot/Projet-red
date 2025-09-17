package red

import (
	"fmt"
	"strings"
)

type Character struct {
	Name         string
	Pv           int
	PvMax        int
	Xp           int
	Money        int
	IsDead       bool
	MaxInventory int
	LoseRate	 int
	Inventory    []Inventory
}

func New(Name string, Pv, PvMax, Xp int, Money int, IsDead bool, MaxInventory, LoseRate int, Inventory []Inventory) Character {
	return Character{Name, Pv, PvMax, Xp, Money, IsDead, MaxInventory, LoseRate, Inventory}
}

func InitCharacter(charName string) Character {
	name := ""
	if charName == "" {
		fmt.Println("\n")
		name = AskPlayerString("Comment vous appelez vous ?")
		if len(name) > 0 {
			finalName := strings.ToUpper(name[:1]) + strings.ToLower(name[1:])
			name = finalName
		}
	} else {
		name = strings.ToUpper(charName[:1]) + strings.ToLower(charName[1:])
	}
	char := New(name, 100, 100, 0, 100, false, 10, 0, []Inventory{})
	char.Pv = char.PvMax / 2
	return char
}

func (c Character) DisplayPlayer() {
	fmt.Println("--------------------------------------------------")
	fmt.Println("Nom du personnage : ", c.Name)
	fmt.Println("Vie du personnage : [", c.Pv, "]")
	fmt.Println("XP administratif : [", c.Xp, "]")
	fmt.Println("Argent : ", c.Money, "€")
	fmt.Println("Item dans l'inventaire :", c.GetItemNumber(), "/", c.MaxInventory)
	fmt.Println("--------------------------------------------------")
}

func (c *Character) HandleDeath() {
	if c.Pv <= 0 {
		c.IsDead = true
		fmt.Println(c.Name, "est mort... mais ressuscite à 50% de sa vie !")
		c.Pv = c.PvMax / 2
		c.IsDead = false
	}
}

func (c *Character) UpdateMoney(q int, s string) {
	switch s {
	case "-":
		if (c.Money - q) >= 0 {
			c.Money -= q
		}
	case "+":
		c.Money += q
	}
}

func (c *Character) AddPV(pv int) {
	if c.GetPV()+pv > c.GetMaxPV() {
		fmt.Println("Action impossicle, PV limités à ", c.GetMaxPV())
	} else {
		c.Pv += pv
	}
}

func (c *Character) RemovePV(pv int) {
	if c.GetPV()-pv <= 0 {
		c.SetPV(0)
		c.IsDead = true
	} else {
		c.SetPV(c.GetPV() - pv)
	}
}

func (c *Character) SetPV(pv int) {
	c.Pv = pv
}

func (c *Character) GetPV() int {
	return c.Pv
}

func (c *Character) GetMaxPV() int {
	return c.PvMax
}

func (c *Character) UpdateXp(q int, s string) {
	switch s {
		case "-":
			c.Xp -= q
		case "+":
			c.Xp += q
		}
}
