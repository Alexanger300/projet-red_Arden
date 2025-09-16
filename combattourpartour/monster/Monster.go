package main

type Monster struct {
	Name        string
	HPMax       int
	HP          int
	ATK         int
	DEF         int
	SPD         int
	Basic_ATK   string
	Special_ATK string
}

func initGoblin() Monster {
	return Monster{
		Name:        "Gobelin",
		HPMax:       30,
		HP:          30,
		ATK:         5,
		DEF:         2,
		SPD:         3,
		Basic_ATK:   "Coup de Dague",
		Special_ATK: "Lancer de Dague",
	}
}
func initSanglier() Monster {
	return Monster{
		Name:        "Sanglier",
		HPMax:       40,
		HP:          40,
		ATK:         6,
		DEF:         3,
		SPD:         2,
		Basic_ATK:   "Coup de Tête",
		Special_ATK: "Charge",
	}
}
