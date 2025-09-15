package red

import (
	"fmt"
	"bufio"
	"os"
	"unicode"
	"strings"
	"strconv"
)

func AskPlayer(question string) interface{} {
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
func AskPlayerString(question string) string {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print(question, " : ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	
	return text
}

func AskPlayerInt(question string) int {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print(question, " : ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	if isInt(text) {
		if number, err := strconv.Atoi(text); err == nil {
			return number
		}
	}
	return 0
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