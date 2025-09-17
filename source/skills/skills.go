package skills

// Skill représente un sort ou une capacité
type Skill struct {
	Name        string
	Description string
	ManaCost    int
	Power       int  // Puissance brute (base pour calculer dégâts/soins)
	IsHeal      bool // true = soin, false = dégâts
	IsMagic     bool // différencie attaques physiques et magiques
}

// Liste des sorts de base par classe
var ClassSkills = map[string][]Skill{
	"Paladin": {
		{
			Name:        "Coup d'Épée sacrée",
			Description: "Frappe l'ennemi avec l'épée sacrée, infligeant des dégâts bénis.",
			ManaCost:    0,
			Power:       20,
			IsHeal:      false,
			IsMagic:     false,
		},
		{
			Name:        "Lumière sacrée",
			Description: "Soigne légèrement un allié.",
			ManaCost:    10,
			Power:       15,
			IsHeal:      true,
			IsMagic:     true,
		},
	},
	"Géant": {
		{
			Name:        "Coup de poing colossal",
			Description: "Une attaque brutale avec ses poings massifs.",
			ManaCost:    0,
			Power:       25,
			IsHeal:      false,
			IsMagic:     false,
		},
		{
			Name:        "Écrasement",
			Description: "Écrase le sol ou l'ennemi, infligeant de lourds dégâts.",
			ManaCost:    15,
			Power:       40,
			IsHeal:      false,
			IsMagic:     false,
		},
	},
	"Mage": {
		{
			Name:        "Projectiles magiques",
			Description: "Tire une salve d'éclats d'énergie magique.",
			ManaCost:    5,
			Power:       15,
			IsHeal:      false,
			IsMagic:     true,
		},
		{
			Name:        "Éclair magique",
			Description: "Projette un éclair destructeur d'énergie pure.",
			ManaCost:    10,
			Power:       30,
			IsHeal:      false,
			IsMagic:     true,
		},
	},
	"Guérisseur": {
		{
			Name:        "Coup de Bâton de vie",
			Description: "Frappe l'ennemi avec le bâton sacré, faible dégâts mais précis.",
			ManaCost:    0,
			Power:       10,
			IsHeal:      false,
			IsMagic:     false,
		},
		{
			Name:        "Soin mineur",
			Description: "Restaure un peu de PV à un allié.",
			ManaCost:    8,
			Power:       20,
			IsHeal:      true,
			IsMagic:     true,
		},
	},
}

// Sort spécial débloqué via un livre (uniquement pour Mage)
var Fireball = Skill{
	Name:        "Boule de feu",
	Description: "Lance une boule de feu infligeant de lourds dégâts magiques.",
	ManaCost:    20,
	Power:       50,
	IsHeal:      false,
	IsMagic:     true,
}
