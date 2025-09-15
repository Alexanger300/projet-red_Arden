package class

type ClassStats struct {
	HPMax       int
	HP          int
	Atk         int
	Def         int
	ManaMax     int
	Mana        int
	Spd         int
	Crit        int
	Weapon      string
	Description string
}

var Classes = map[string]ClassStats{
	"Paladin": {
		HPMax: 100, HP: 100, Atk: 60, Def: 60, ManaMax: 40, Mana: 40, Spd: 40, Crit: 5, Weapon: "Épée sacrée",
		Description: "Nourri par la foi et la justice, le Paladin protège les faibles et lutte contre les hérétiques. Sa force n’est pas seulement dans ses bras, mais dans sa croyance inébranlable.",
	},
	"Géant": {
		HPMax: 150, HP: 150, Atk: 70, Def: 80, ManaMax: 20, Mana: 20, Spd: 25, Crit: 3, Weapon: "Poings",
		Description: "Né des montagnes, le Géant n’a jamais connu la peur. Ses poings sont des armes, son corps un mur. Il n’obéit à personne, mais quand il choisit un camp, il le défend jusqu’à la mort.",
	},
	"Mage": {
		HPMax: 75, HP: 75, Atk: 40, Def: 35, ManaMax: 100, Mana: 100, Spd: 50, Crit: 10, Weapon: "Grimoire",
		Description: "Le Mage a délaissé les lames et les boucliers. Dans son vieux grimoire sommeillent des flammes, des éclairs et des ombres interdites. Le monde le craint, car là où il passe, la réalité elle-même se plie.",
	},
	"Guérisseur": {
		HPMax: 75, HP: 75, Atk: 35, Def: 40, ManaMax: 85, Mana: 85, Spd: 45, Crit: 5, Weapon: "Bâton de vie",
		Description: "Le Guérisseur n’est pas un guerrier, mais un gardien de vie. Son bâton n’apporte pas la mort mais la lumière. Beaucoup se moquent de lui, jusqu’au jour où ses soins sauvent tout un bataillon.",
	},
}
