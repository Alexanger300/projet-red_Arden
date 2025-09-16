package main

import (
	"fmt"
	"time"
) //Importer TOUT le fichier CombatTourParTour (jpense)

func trainingFight() {
	if character.SPD < monster.SPD {
		fmt.Println("C'est au tour de", monster.Name)
		for {
			CharacterTurn(&character)
			time.Sleep(2 * time.Second)
			goblinpattern(&monster)
			if monster.HP <= 0 || character.HP <= 0 {
				break
			}
		}
	} else {
		fmt.Println("C'est au tour de", character.Name)
		for {
			goblinpattern(&monster)
			time.Sleep(2 * time.Second)
			CharacterTurn(&character)
			if monster.HP <= 0 || character.HP <= 0 {
				break
			}
		}
	}
	if character.HP <= 0 {
		fmt.Println(character.Name, "a été vaincu...")
		fmt.Println("GAME OVER")
	} else if monster.HP <= 0 {
		fmt.Println(monster.Name, "a été vaincu !")
		fmt.Println("Victoire !")
	}
}
