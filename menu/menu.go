package red

import (
	"fmt"
)

type Choice struct {
	Label string          	
	Action func(Character) // Fonction d'execution
}

type Menu struct {
	Name    string
	Choices []Choice
}
*
func (m *Menu) Display(character Character) {
	fmt.Println("_______________________________________")
	fmt.Println(m.Name)
	fmt.Println("_______________________________________")

	for i, choice := range m.Choices {
		fmt.Printf(" - %s -> Tapez %d\n", choice.Label, i+1)
	}

	fmt.Println("_______________________________________")

	result := AskPlayerType("Votre choix :", 0).(int)

	if result >= 1 && result <= len(m.Choices) {
		m.Choices[result-1].Action(character)
	} else {
		fmt.Println("Choix invalide.")
	}
}
