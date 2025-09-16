package main

import (
	"fmt"
	"math/rand"
	"time"
) //Importer Biblioteque Monstre+Character(jpense)

func goblinpattern(monster *Monster) {
	{
		gobelin := initGoblin()
		fmt.Println("C'est au tour de", gobelin.Name)
		time.Sleep(1 * time.Second)
		random_atk := rand.Intn(3)
		if random_atk < 2 {
			fmt.Println(gobelin.Name, "utilise son attaque basique :", gobelin.Basic_ATK)
			character.HP -= gobelin.ATK - character.DEF
			fmt.Println("Il inflige", gobelin.ATK-character.DEF, "dégâts à", character.Name)
		} else {
			fmt.Println(gobelin.Name, "utilise son attaque spéciale :", gobelin.Special_ATK)
			character.HP -= (gobelin.ATK * 2) - character.DEF
			fmt.Println("Il inflige", (gobelin.ATK*2)-character.DEF, "dégâts à", character.Name)
		}
	}
}
