package red

import (
	"fmt"
	"bufio"
	"os"
	"unicode"
	"strings"
	"strconv"
)

// func main() {
//  	badam := askPlayer("Comment vous appelez-vous ?")
//  	fmt.Println("answer : ", badam)
// 	fmt.Printf("%T\n", badam)
//  	age := askPlayer("age ?")
// }

func askPlayer(question string) interface{} {
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