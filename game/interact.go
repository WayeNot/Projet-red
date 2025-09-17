package red

import "fmt"

// Système de perte d'item aléatoire

func (c *Character) hasItemWithHole() int {
	percent := 0
	for _, v := range c.Inventory {
		if allItems[v.Id].loseRate > 0 {
			if percent + allItems[v.Id].loseRate <= 100 {
				percent += allItems[v.Id].loseRate
			} else {
				percent = 100
			}
		}
	}
	return percent
}

func (c *Character) LooseItem() {
	randomItem := c.GetRandomItem()

	if c.hasItemWithHole() > 0 {
		rand := RandomNbr(100)

		if rand <= c.hasItemWithHole() {
			c.RemoveItem(randomItem, 1)
			fmt.Printf("A cause de votre sac troué, vous avez perdu cet item : %v | %v x1", allItems[randomItem].Icon, allItems[randomItem].Name)
		} else {
			fmt.Println("Vous avez eu de la chance, rien n’est tombé par terre !")
		}
	}
}