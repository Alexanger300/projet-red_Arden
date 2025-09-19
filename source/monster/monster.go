package monster

import "fmt"

// Statut appliqué au monstre
type Status struct {
	Name     string // ex: "Poison"
	Duration int    // nb de tours restants
	Damage   int    // dégâts par tour (si applicable)
}

// Structure d’un monstre
type Monster struct {
	Name       string
	HPMax      int
	HP         int
	Atk        int
	Def        int
	Spd        int
	BasicAtk   string
	SpecialAtk string
	ExpReward  int
	GoldReward int
	Loot       string

	Statuses []Status
}

// Est-ce que le monstre est vivant ?
func (m *Monster) IsAlive() bool { return m.HP > 0 }

// Appliquer un statut si non présent
func (m *Monster) ApplyStatus(name string, duration, damage int) {
	for _, s := range m.Statuses {
		if s.Name == name {
			fmt.Printf("%s est déjà affecté par %s.\n", m.Name, name)
			return
		}
	}
	m.Statuses = append(m.Statuses, Status{Name: name, Duration: duration, Damage: damage})
	fmt.Printf("%s est maintenant affecté par %s (%d tours).\n", m.Name, name, duration)
}

// À appeler à chaque tour pour appliquer les effets des statuts
func (m *Monster) UpdateStatuses() {
	if len(m.Statuses) == 0 {
		return
	}
	var remaining []Status
	for _, s := range m.Statuses {
		// Poison (ou tout statut avec dégâts)
		if s.Damage > 0 {
			m.HP -= s.Damage
			if m.HP < 0 {
				m.HP = 0
			}
			fmt.Printf("☠️  %s subit %d dégâts (%s). HP: %d/%d\n",
				m.Name, s.Damage, s.Name, m.HP, m.HPMax)

			if m.HP == 0 {
				fmt.Printf("💀 %s succombe à %s.\n", m.Name, s.Name)
				return
			}
		}

		s.Duration--
		if s.Duration > 0 {
			remaining = append(remaining, s)
		} else {
			fmt.Printf("%s n’est plus affecté par %s.\n", m.Name, s.Name)
		}
	}
	m.Statuses = remaining
}

// Création de différents types de monstres

// Gobelin — rapide, peu défendu
func NewGoblin() *Monster {
	return &Monster{
		Name:       "Gobelin",
		HPMax:      20,
		HP:         20,
		Atk:        4,
		Def:        2,
		Spd:        6,
		BasicAtk:   "Coup de Dague",
		SpecialAtk: "Lancer de Dague",
		ExpReward:  5,
		GoldReward: 8,
		Loot:       "Peau de gobelin",
	}
}

// Sanglier — costaud, bonne charge
func NewBoar() *Monster {
	return &Monster{
		Name:       "Sanglier",
		HPMax:      30,
		HP:         30,
		Atk:        5,
		Def:        2,
		Spd:        5,
		BasicAtk:   "Coup de Tête",
		SpecialAtk: "Charge",
		ExpReward:  7,
		GoldReward: 12,
		Loot:       "Peau de sanglier",
	}
}

// Loup — très rapide, attaque mordante
func NewWolf() *Monster {
	return &Monster{
		Name:       "Loup",
		HPMax:      25,
		HP:         25,
		Atk:        7,
		Def:        3,
		Spd:        7,
		BasicAtk:   "Morsure",
		SpecialAtk: "Crocs Sauvages",
		ExpReward:  6,
		GoldReward: 10,
		Loot:       "Fourrure de loup",
	}
}
