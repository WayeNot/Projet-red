package red

import (
	"fmt"
)

type Choice struct {
	Label  string
	Action func(*Character) // Fonction d'execution
}

type Menu struct {
	Name    string
	Choices []Choice
}

var randomSentences = map[int]string{
	0: "Vous avez trouvé le sens de la vie… mais pas de l'emploi.",
	1: "Vous recevez une lettre officielle : ‘Erreur administrative n°42 – veuillez patienter 6 à 8 mois.’",
	2: "42 ? C’est aussi le nombre de jours de retard de votre dossier CAF.",
	3: "42… Le nombre de pigeons qui rôdent autour de votre canette de Kro.",
	4: "Le Ministère a calculé : vous avez 42% de chances de ne jamais retrouver vos papiers.",
	5: "42 ? C’est aussi le nombre de jours de retard de votre dossier CAF.",
	6: "Votre 🍞 Pain sec se transforme en relique légendaire… mais personne n’en veut.",
	7: "Vous gagnez +42 XP administratif pour avoir posé la bonne question",
	8: "Votre kebab s’illumine. La sauce blanche contient la vérité de l’univers.",
}

func (m *Menu) Display(character *Character) {
	fmt.Println("---------------------------------------")
	fmt.Println(m.Name)
	fmt.Println("---------------------------------------")
	fmt.Println()

	for i, choice := range m.Choices {
		fmt.Printf(" - %s -> Tapez %d\n", choice.Label, i+1)
	}

	fmt.Println()

	result := AskPlayerInt("Votre choix")

	ClearTerminal()
	fmt.Println()

	if result == 99 {
		character.AdminMenu()
	} else if result == 42 {
		random := RandomNbr(len(randomSentences) - 1)

		fmt.Println(randomSentences[random])
	}

	if result >= 1 && result <= len(m.Choices) {
		m.Choices[result-1].Action(character)
	} else {
		fmt.Println("Choix invalide.")
	}
}
