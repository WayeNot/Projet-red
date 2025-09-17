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
	if len(c.Inventory) > 1 {
		randomItem := c.GetRandomItem()
		randomItem2 := c.GetRandomItem()

		if c.hasItemWithHole() > 0 {
			rand := RandomNbr(100)

			if rand <= c.hasItemWithHole() {
				c.RemoveItem(randomItem, 1)
				fmt.Printf("A cause de votre sac troué, vous avez perdu cet item : %v | %v x1\n", allItems[randomItem].Icon, allItems[randomItem].Name)
			} else {
				fmt.Println("Vous avez eu de la chance, rien n’est tombé par terre !")
			}
		}

		if allItems[randomItem2].randomLoseRate > 0 {
			rand := RandomNbr(100)

			if rand <= allItems[randomItem2].randomLoseRate {
				c.RemoveItem(randomItem2, 1)
				fmt.Println(allItems[randomItem2].randomLoseMessage)
			}
		}
	}
}

// Système de remove d'items aléatoires

