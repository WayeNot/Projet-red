package red

import (
	"fmt"

	"math/rand/v2"
)

type Combat struct {
	boss Boss
	player Character
	playerTour bool
	winner bool
}

func InitCombat(boss Boss, player Character) Combat {
	return Combat{boss, player, false, false}
}

func (c *Combat) Tour() {
	if c.IsPlayerTour() {
		fmt.Println("A votre tour ", c.player.Name, " !")
	} else {
		fmt.Println("Au tour de ", c.boss.GetName(), " !")
		attack := rand.IntN(len(c.boss.GetAttacks()))
		c.boss.GetAttacks()[attack](&c.player)
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