package main

import (
	"fmt"
	"time"
)

func poisonPot() { //Importer bibliothèque Character et/ou inventaire
	for i := 0; i < 3; i++ {
		fmt.Println("Vous subissez des dégâts de poison !")
		HP -= 5
		time.Sleep(1 * time.Second)
	}
}

func main() {
	poisonPot()
}
