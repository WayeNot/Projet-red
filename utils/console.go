package red

import (
	"fmt"
	"bufio"
	"os"
	"unicode"
	"strings"
	"strconv"
)

func AskPlayer(question string) any {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question, " : ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	if isInt(text) {
		if number, err := strconv.Atoi(text); err == nil {
			return number
		}
	}
	return text
}

/*
	Si on veut forcer le type de la réponse du joueur.
*/
func AskPlayerType(question string, forcedtType any) any {
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Print(question, " : ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		switch forcedtType.(type) {
		case int:
			if isInt(text) {
				num, _ := strconv.Atoi(text)
				return num
			}
		case string:
			if text != "" {
				return text
			}
		default:
			return text
		}

		fmt.Println("Mauvaise réponse ! Veuillez rentrer le bon type.")
	}
}

func isInt(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}