package equipment

import "fmt"

// Définition d’un équipement
type Equipment struct {
	Name  string
	HP    int
	Mana  int
	Atk   int
	Def   int
	Spd   int
	Crit  int
	Slot  string // "Head", "Body", "Legs"
	Class string // Classe qui peut l’équiper
}

// Ensemble équipé par un personnage
type EquipmentSet struct {
	Head Equipment
	Body Equipment
	Legs Equipment
}

// Base des équipements disponibles par classe
var EquipmentPools = map[string][]Equipment{
	"Paladin": {
		{Name: "Casque de paladin", HP: 15, Def: 10, Slot: "Head", Class: "Paladin"},
		{Name: "Armure bénie", HP: 40, Def: 20, Atk: 5, Slot: "Body", Class: "Paladin"},
		{Name: "Jambières lourdes", HP: 20, Def: 15, Slot: "Legs", Class: "Paladin"},
	},
	"Géant": {
		{Name: "Heaume massif", HP: 25, Def: 15, Slot: "Head", Class: "Géant"},
		{Name: "Plastron de colosse", HP: 60, Def: 30, Atk: 10, Slot: "Body", Class: "Géant"},
		{Name: "Jambières de pierre", HP: 30, Def: 20, Slot: "Legs", Class: "Géant"},
	},
	"Mage": {
		{Name: "Chapeau mystique", Mana: 20, Crit: 5, Slot: "Head", Class: "Mage"},
		{Name: "Robe enchantée", Mana: 40, Spd: 10, Slot: "Body", Class: "Mage"},
		{Name: "Bottes de lévitation", Spd: 15, Mana: 10, Slot: "Legs", Class: "Mage"},
	},
	"Guérisseur": {
		{Name: "Capuche de prêtre", Mana: 15, Def: 5, Slot: "Head", Class: "Guérisseur"},
		{Name: "Robe de lumière", Mana: 35, HP: 20, Slot: "Body", Class: "Guérisseur"},
		{Name: "Sandales bénies", Mana: 10, Spd: 5, Slot: "Legs", Class: "Guérisseur"},
	},
}

// Équiper un objet
func (set *EquipmentSet) Equip(item Equipment, class string) {
	if item.Class != class {
		fmt.Printf("❌ %s ne peut pas être équipé par un %s.\n", item.Name, class)
		return
	}

	switch item.Slot {
	case "Head":
		set.Head = item
	case "Body":
		set.Body = item
	case "Legs":
		set.Legs = item
	default:
		fmt.Println("❌ Emplacement invalide :", item.Slot)
		return
	}
	fmt.Printf("✅ %s équipé sur %s\n", item.Name, item.Slot)
}

// Retirer un équipement
func (set *EquipmentSet) Unequip(slot string) {
	switch slot {
	case "Head":
		fmt.Printf("❌ %s retiré de la tête\n", set.Head.Name)
		set.Head = Equipment{}
	case "Body":
		fmt.Printf("❌ %s retiré du corps\n", set.Body.Name)
		set.Body = Equipment{}
	case "Legs":
		fmt.Printf("❌ %s retiré des jambes\n", set.Legs.Name)
		set.Legs = Equipment{}
	default:
		fmt.Println("❌ Emplacement invalide :", slot)
	}
}

// Calcul des bonus totaux
func (set EquipmentSet) TotalStats() (hp, mana, atk, def, spd, crit int) {
	hp = set.Head.HP + set.Body.HP + set.Legs.HP
	mana = set.Head.Mana + set.Body.Mana + set.Legs.Mana
	atk = set.Head.Atk + set.Body.Atk + set.Legs.Atk
	def = set.Head.Def + set.Body.Def + set.Legs.Def
	spd = set.Head.Spd + set.Body.Spd + set.Legs.Spd
	crit = set.Head.Crit + set.Body.Crit + set.Legs.Crit
	return
}

// Afficher l’équipement
func (set EquipmentSet) Display() {
	fmt.Println("=== 🛡️ Équipement Actuel ===")
	if set.Head.Name != "" {
		fmt.Println("Tête :", set.Head.Name)
	} else {
		fmt.Println("Tête : Aucun")
	}
	if set.Body.Name != "" {
		fmt.Println("Corps :", set.Body.Name)
	} else {
		fmt.Println("Corps : Aucun")
	}
	if set.Legs.Name != "" {
		fmt.Println("Jambes :", set.Legs.Name)
	} else {
		fmt.Println("Jambes : Aucun")
	}
}
