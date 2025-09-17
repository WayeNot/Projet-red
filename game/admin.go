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
			fmt.Println(awaitingDev)
		}

		switch commandAdmin {
		case "ClearInventory":
			for _, v := range c.Inventory {
				c.RemoveItem(v.Id, v.Quantity)
			}
			fmt.Println("Tous les items viennent d'être remove !")
		case "Heal":
			c.Pv = c.PvMax
		case "SetHealth":
			fmt.Println("Votre vie actuelle : ", c.Pv)
			c.Pv = AskPlayerInt("Combien de vie voulez-vous vous mettre ?")
		case "SetMoney":
			fmt.Println("Montant de vos poches : ", c.Money)
			c.Money = AskPlayerInt("Combien voulez-vous d'argent ?")
		case "SetXp":
			fmt.Println("Nombre de Xp actuelle : ", c.Xp)
			c.Xp = AskPlayerInt("Combien voulez-vous de Xp ?")
		case "AddItem":
			c.AddItem(AskPlayerInt("Id de l'item souhaitée"), AskPlayerInt("Quantité d'item souhaité ?"))
		case "RemoveItem":
			c.AccessInventory()
			c.RemoveItem(AskPlayerInt("Id de l'item souhaitée"), AskPlayerInt("Quantité d'item souhaité ?"))
		}

		// Commandes RP

		switch commandAdmin {
		case "InflationON":
			for _, v := range allItems {
				v.Price *= 2
			}
		case "InflationOff":
			for _, v := range allItems {
				v.Price /= 2
			}
		case "GodMode":
			c.MaxInventory += 999999999999999
			c.Money += 999999999999999
			c.PvMax += 999999999999999
			c.Pv += 999999999999999
			c.Xp += 999999999999999
		case "RmGodMode":
			c.MaxInventory -= 999999999999999
			c.Money -= 999999999999999
			c.PvMax -= 999999999999999
			c.Pv -= 999999999999999
			c.Xp -= 999999999999999
		}

		// Commandes GamePlay

		switch commandAdmin {
		case "SpawnMerchant":
			c.MenuMerchant()
		case "SpawnForgeron":
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