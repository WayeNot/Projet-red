package red

import (
	"fmt"
	"bufio"
	"os"
)

// func main() {
// 	badam := askPlayer("Comment vous appelez-vous ?")
// 	fmt.Println("answer : ", badam)
// 	age := askPlayer("age ?")
// 	fmt.Println(age)
// }

func askPlayer(question string) interface{} {
	reader := bufio.NewReader(os.Stdin)
    fmt.Print(question, " : ")
    text, _ := reader.ReadString('\n')
	return text
}