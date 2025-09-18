package red

import "fmt"

type Combat struct {
	boss Boss
	player Character
	playerTour bool
	winner bool
}

func InitCombat(boss Boss, player Character) Combat {
	return Combat{boss, player, true, false}
}

func (c *Combat) Start() {
	fmt.Println("Le combat commence contre", c.boss.GetName(), "!")
	for !c.boss.IsDead() && c.player.Pv > 0 {
		c.Tour()
	}
	if c.boss.IsDead() {
		fmt.Println("Victoire ! Tu as battu ", c.boss.GetName())
		c.winner = true
	} else {
		fmt.Println("Défaite... le chômage gagne encore.")
		c.winner = false
	}
}

func (c *Combat) Tour() {
	if c.playerTour {
		menu := Menu{
			Name: "Ton tour " + c.player.Name,
			Choices: []Choice{
				{
					Label: "Attaquer",
					Action: func(p *Character) {
						damage := 5 // dégats de base
						for _, inv := range c.player.Inventory {
							item := allItems[inv.Id]
							if item.PtsAttack > 0 {
								damage += item.PtsAttack
							}
						}
						fmt.Println("Tu frappes le boss et infliges ", damage, " dégâts !")
						c.boss.RemovePV(damage)
						c.playerTour = false
					},
				},
				{
					Label: "Utiliser un item",
					Action: func(p *Character) {
						c.player.AccessInventory()
						itemId := AskPlayerInt("Choisis l’ID d’un item à utiliser :")
						c.player.UseItem(itemId, 1)
						c.playerTour = false
					},
				},
				{
					Label: "Passer le tour",
					Action: func(p *Character) {
						fmt.Println("Tu ne fais rien ce tour...")
						c.playerTour = false
					},
				},
			},
		}
		menu.Display(&c.player)
	} else {
		fmt.Println("Tour du Boss", c.boss.GetName())
		attack := RandomNbr(len(c.boss.GetAttacks()))
		c.boss.GetAttacks()[attack](&c.player)
		c.playerTour = true
	}
}

func (c *Combat) IsPlayerTour() bool {
	return c.playerTour
}

func (c *Combat) GetBoss() Boss {
	return c.boss
}

func (c *Combat) GetPlayer() Character {
	return c.player
}

func (c *Combat) IsWinner() bool {
	return c.winner
}