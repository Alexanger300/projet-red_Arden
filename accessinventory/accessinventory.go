package main

import "fmt"

var inventory [][]string

func ShowInventory() {
	fmt.Println("Voici votre inventaire :")
	if len(inventory) == 0 {
		return
	}
	for _, item := range inventory {
		fmt.Println(item)
	}
}
func AddInInventory(item string, quantity int) {
	inventory = append(inventory, []string{item})
	for i := 1; i < quantity; i++ {
		inventory = append(inventory, []string{item})
	}
}
func main() {
	AddInInventory("Potion de vie", 3)
	ShowInventory()
}
