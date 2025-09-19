package class

type ClassStats struct {
	HP          int    // Points de vie de base
	Atk         int    // Attaque
	Def         int    // Défense
	Mana        int    // Mana de base
	Spd         int    // Vitesse
	Crit        int    // Taux critique
	Weapon      string // Arme de départ
	Description string // Description narrative
}

var Classes = map[string]ClassStats{
	"Paladin": {
		HP: 100, Atk: 10, Def: 8, Mana: 30, Spd: 5, Crit: 10, Weapon: "",
		Description: "Nourri par la foi et la justice, le Paladin protège les faibles et lutte contre les hérétiques. Sa force n'est pas seulement dans ses bras, mais dans sa croyance inébranlable.",
	},
	"Géant": {
		HP: 150, Atk: 12, Def: 10, Mana: 10, Spd: 3, Crit: 5, Weapon: "",
		Description: "Né des montagnes, le Géant n'a jamais connu la peur. Ses poings sont des armes, son corps un mur. Il n'obéit à personne, mais quand il choisit un camp, il le défend jusqu’à la mort.",
	},
	"Mage": {
		HP: 60, Atk: 6, Def: 3, Mana: 60, Spd: 6, Crit: 5, Weapon: "",
		Description: "Le Mage a délaissé les lames et les boucliers. Dans son vieux grimoire sommeillent des flammes, des éclairs et des ombres interdites. Le monde le craint, car là où il passe, la réalité elle-même se plie.",
	},
	"Guérisseur": {
		HP: 80, Atk: 4, Def: 5, Mana: 50, Spd: 5, Crit: 5, Weapon: "",
		Description: "Le Guérisseur n'est pas un guerrier, mais un gardien de vie. Son bâton n'apporte pas la mort mais la lumière. Beaucoup se moquent de lui, jusqu’au jour où ses soins sauvent tout un bataillon.",
	},
}
