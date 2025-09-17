package red

import (
	"fmt"
	"time"
)

type Moderation struct {
	Logs		[]Logs
}

type Logs struct {
	id 		int
	msg 	string
	date 	string
}

func (c *Character) AdminMenu() {
	fmt.Println("\033[31m")
	awaitingDev := "Cette fonctionnalité est en cours de développement !"
	ClearTerminal()
	fmt.Println()
	fmt.Println("Si ce n'est pas fait exprès, gg ta trouver un easter-egg ^^")
	for {
		commandAdmin := AskPlayerString("Insérer votre commande (💀 A toi de trouver ! 💀) - (QuitAdmin pour revenir en arrière)")

		if commandAdmin == "QuitAdmin" {
			fmt.Println("\033[0m")
			return
		}

		if commandAdmin == "logs" {
			// viewLogs()
			fmt.Println(awaitingDev)
		}

		if commandAdmin == "ClearInventory" {
			for _, v := range c.Inventory {
				c.RemoveItem(v.Id, v.Quantity)
			}
			fmt.Println("Tous les items viennent d'être remove !")
		} else if commandAdmin == "Heal" {
			c.Pv = c.PvMax
		} else if commandAdmin == "SetHealth" {
			fmt.Println("Votre vie actuelle : ", c.Pv)
			c.Pv = AskPlayerInt("Combien de vie voulez-vous vous mettre ?")
		} else if commandAdmin == "SetMoney" {
			fmt.Println("Montant de vos poches : ", c.Money)
			c.Money = AskPlayerInt("Combien voulez-vous d'argent ?")
		} else if commandAdmin == "SetXp" {
			fmt.Println("Nombre de Xp actuelle : ", c.Xp)
			c.Xp = AskPlayerInt("Combien voulez-vous de Xp ?")
		} else if commandAdmin == "AddItem" {
			c.AddItem(AskPlayerInt("Id de l'item souhaitée"), AskPlayerInt("Quantité d'item souhaité ?"))
		} else if commandAdmin == "RemoveItem" {
			c.AccessInventory()
			c.RemoveItem(AskPlayerInt("Id de l'item souhaitée"), AskPlayerInt("Quantité d'item souhaité ?"))
		}

		// Commandes RP

		if commandAdmin == "InflationON" {
			for _, v := range allItems {
				v.Price *= 2
			}
		} else if commandAdmin == "InflationOff" {
			for _, v := range allItems {
				v.Price /= 2
			}
		} else if commandAdmin == "GodMode" {
			c.MaxInventory += 999999999999999
			c.Money += 999999999999999
			c.PvMax += 999999999999999
			c.Pv += 999999999999999
			c.Xp += 999999999999999
		} else if commandAdmin == "RmGodMode" {
			c.MaxInventory -= 999999999999999
			c.Money -= 999999999999999
			c.PvMax -= 999999999999999
			c.Pv -= 999999999999999
			c.Xp -= 999999999999999
		}

		// Commandes GamePlay

		if commandAdmin == "SpawnMerchant" {
			c.MenuMerchant()
		} else if commandAdmin == "SpawnForgeron" {
			c.MenuForgeron()
		}
	}
}

func (m *Moderation) SetLogs(msg string) {
	m.Logs = append(m.Logs, Logs{
		id: len(m.Logs),
		msg: msg,
		date: time.Now().Format("02/01/2006 15:04:05"),
	})
}

func (m *Moderation) viewLogs() {

}