package introduction

import (
	"fmt"
	"time"

	"github.com/Alexanger300/projet-red_Forge/asset/css"
	"github.com/Alexanger300/projet-red_Forge/source/character"
	"github.com/Alexanger300/projet-red_Forge/source/inn"
)

func Introduction() {
	css.Clear()
	text0 := `Au cœur d'un royaume déchiré par la guerre, `
	texte1 := `les terres d'Arden`
	texte2 := ` sombrent dans les flammes.`
	texte3 := `Les dragons ravagent les villages, les armées s'entre-déchirent, et la peste ronge les survivants.
Dans ce chaos, quatre figures se lèvent.`
	text4 := `À toi de choisir ton rôle dans cette épopée.`

	for _, c := range text0 {
		fmt.Print(css.Bold + string(c) + css.Reset)
		time.Sleep(30 * time.Millisecond)
	}
	for _, c := range texte1 {
		fmt.Print(css.Red + css.Bold + string(c) + css.Reset)
		time.Sleep(30 * time.Millisecond)
	}
	for _, c := range texte2 {
		fmt.Print(css.Bold + string(c) + css.Reset)
		time.Sleep(30 * time.Millisecond)
	}
	time.Sleep(700 * time.Millisecond)
	for _, c := range texte3 {
		fmt.Print(css.Bold + string(c) + css.Reset)
		time.Sleep(30 * time.Millisecond)
	}
	time.Sleep(1 * time.Second)
	fmt.Print("\n")
	for _, c := range text4 {
		fmt.Print(css.Red + css.Bold + string(c) + css.Reset)
		time.Sleep(50 * time.Millisecond) // pause entre chaque caractère
	}
	time.Sleep(1 * time.Second)
	// Création du personnage
	player := character.InitCharacter()

	// Passe directement à l’auberge (hub du village)
	inn.Inn(&player)
}
