package main

import (
	"fmt"
)

func ShowMenu() {
	fmt.Println("	    Menu Principal	   ")
	fmt.Println("1.     Afficher les informations du personnage ")
	fmt.Println("2.     Afficher l'inventaire   ")
	fmt.Println("3.     Quitter le jeu   ")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println("DisplayInfo()")
	case 2:
		fmt.Println("ShowInventory()")
	case 3:
		fmt.Println("Au revoir!")
	default:
		fmt.Println("Choix invalide, veuillez reessayer.")
		ShowMenu()
	}
}
func main() {
	ShowMenu()
}
