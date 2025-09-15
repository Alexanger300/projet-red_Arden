package main

import "fmt" //Essayer d'importer inventaire+Character

func takePot() {
	fmt.Println("Vous avez pris une potion de soin !")
	HP += 20
	if HP > HPMax {
		HP = HPMax
		removeitem("Potion de soin", 1) //Fonction à créer dans inventaire.go
	}
}
