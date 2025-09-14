package class

type ClassStats struct {
	// Statistiques de base pour chaque classe
	Pv     int
	Atk    int
	Def    int
	Mana   int
	Spd    int
	Crit   int
	Weapon string
}

var Classes = map[string]ClassStats{
	// Définition des statistiques pour chaque classe
	"Paladin": {
		Pv: 100, Atk: 60, Def: 60, Mana: 40, Spd: 40, Crit: 5, Weapon: "Épée sacrée",
	},
	"Géant": {
		Pv: 150, Atk: 70, Def: 80, Mana: 20, Spd: 25, Crit: 3, Weapon: "Poings",
	},
	"Mage": {
		Pv: 75, Atk: 40, Def: 35, Mana: 100, Spd: 50, Crit: 10, Weapon: "Grimoire",
	},
	"Guérisseur": {
		Pv: 75, Atk: 35, Def: 40, Mana: 85, Spd: 45, Crit: 5, Weapon: "Bâton de vie",
	},
}
