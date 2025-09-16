package menu

import (
	"fmt"
)

func ShowMenu() {
	fmt.Println("	    Menu Principal	   ")
	fmt.Println("1.     Afficher les informations du personnage ")
	fmt.Println("2.     Afficher l'inventaire   ")
	fmt.Println("3.     Qui sont-ils ?")
	fmt.Println("4.     Quitter le jeu   ")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println("DisplayFull()")
	case 2:
		fmt.Println("ShowInventory()")
	case 3:
		fmt.Println("ABBA/Steven Spielberg")
	case 4:
		fmt.Println("Quitter le jeu...")
	default:
		fmt.Println("Choix invalide, veuillez reessayer.")
		ShowMenu()
	}
}
