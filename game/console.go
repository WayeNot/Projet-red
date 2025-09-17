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
	} else {
		AskPlayerInt(question)
	}

	return 0
}

func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
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