package introduction

import (
	"fmt"

	"github.com/Alexanger300/projet-red_Forge/character"
	"github.com/Alexanger300/projet-red_Forge/inn"
)

func Introduction() {
	fmt.Println(`
Au cœur d'un royaume déchiré par la guerre, les terres d'Arden sombrent dans les flammes.
Les dragons ravagent les villages, les armées s'entre-déchirent, et la peste ronge les survivants.
Dans ce chaos, quatre figures se lèvent.
Chacune guidée par un destin différent, mais toutes appelées par la même prophétie :
"Quand les ténèbres couvriront le ciel, un héros renaîtra de la poussière."
À toi de choisir ton rôle dans cette épopée.`)

	// 🚀 Création du personnage
	player := character.InitCharacter()
	player.DisplaySummary()

	// Passe directement à l’auberge (hub du village)
	inn.Inn(&player)
}
