package main

import "fmt"

type Class struct { //Initialisation des Stats pour CHAQUE Class
	Pv   int
	Atk  int
	Def  int
	Mana int
	Spd  int
	Crit int
}

func main() {
	//Initialisation des classe

	var Paladin Class
	var Mage Class
	var Géant Class
	var Guerisseur Class

	//Stats de début du Paladin
	Paladin.Pv = 100
	Paladin.Atk = 50
	Paladin.Def = 30
	Paladin.Mana = 50
	Paladin.Spd = 50
	Paladin.Crit = 0

	//Stats de début du Géant

	Géant.Pv = 80
	Géant.Atk = 100
	Géant.Def = 50
	Géant.Mana = 50
	Géant.Spd = 20
	Géant.Crit = 0

	//Stats de début du Guérisseur

	Guerisseur.Pv = 60
	Guerisseur.Atk = 40
	Guerisseur.Def = 100
	Guerisseur.Mana = 100
	Guerisseur.Spd = 25
	Guerisseur.Crit = 0

	//Stats de début de Mage

	Mage.Pv = 60
	Mage.Atk = 100
	Mage.Def = 40
	Mage.Mana = 100
	Mage.Spd = 30
	Mage.Crit = 0

	Paladin.affichage()
}
func (c Class) affichage() {
	fmt.Println("----------", c Class, "----------")
	fmt.Println("Pv :", c.Pv)
	fmt.Println("Atk:", c.Atk)
	fmt.Println("Def :", c.Def)
	fmt.Println("Mana:", c.Mana)
	fmt.Println("Speed:", c.Spd)
	fmt.Println("Mana:", c.Crit)
	fmt.Println("-----------------------")
}
