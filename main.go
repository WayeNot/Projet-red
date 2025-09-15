package main

import (
	"fmt"
	"red/game"
)

func main(){
	fmt.Println("_______________________________")
	fmt.Println("	  Bienvenue dans ")
	fmt.Println("	 CHÔMAGE & DRAGONS")
	fmt.Println("La quête du loose, du PMU et du RSA")
	fmt.Println("_______________________________")
	name := red.AskPlayerString("Comment vous appelez vous ?")
	fmt.Println(name)
	player := red.InitCharacter(name)
	player.DisplayPlayer()
}