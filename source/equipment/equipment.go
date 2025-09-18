package equipment

import "fmt"

// === Définition d’un équipement ===
type Equipment struct {
	Name  string
	HP    int
	Mana  int
	Atk   int
	Def   int
	Spd   int
	Crit  int
	Slot  string // "Weapon", "Head", "Body", "Legs"
	Class string // Classe qui peut l’équiper
}

// === Ensemble équipé par un personnage ===
type EquipmentSet struct {
	Weapon Equipment
	Head   Equipment
	Body   Equipment
	Legs   Equipment
}

// Calcul des bonus totaux (⚠️ sans l’arme !)
func (set EquipmentSet) TotalStats() (hp, mana, atk, def, spd, crit int) {
	items := []Equipment{set.Head, set.Body, set.Legs}
	for _, item := range items {
		hp += item.HP
		mana += item.Mana
		atk += item.Atk
		def += item.Def
		spd += item.Spd
		crit += item.Crit
	}
	return
}

// Afficher l’équipement
func (set EquipmentSet) Display() {
	fmt.Println("=== 🛡️ Équipement Actuel ===")
	if set.Weapon.Name != "" {
		fmt.Println("Arme :", set.Weapon.Name)
	} else {
		fmt.Println("Arme : Aucune")
	}
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

// === Pool d’équipements disponibles par classe ===
var EquipmentPools = map[string][]Equipment{
	"Paladin": {
		{Name: "Épée sacrée", Atk: 15, Crit: 5, Slot: "Weapon", Class: "Paladin"},
		{Name: "Casque de paladin", HP: 15, Def: 10, Slot: "Head", Class: "Paladin"},
		{Name: "Armure bénie", HP: 40, Def: 20, Atk: 5, Slot: "Body", Class: "Paladin"},
		{Name: "Jambières lourdes", HP: 20, Def: 15, Slot: "Legs", Class: "Paladin"},
	},
	"Géant": {
		{Name: "Gantelets colossaux", Atk: 20, Slot: "Weapon", Class: "Géant"},
		{Name: "Heaume massif", HP: 25, Def: 15, Slot: "Head", Class: "Géant"},
		{Name: "Plastron de colosse", HP: 60, Def: 30, Atk: 10, Slot: "Body", Class: "Géant"},
		{Name: "Jambières de pierre", HP: 30, Def: 20, Slot: "Legs", Class: "Géant"},
	},
	"Mage": {
		{Name: "Grimoire ancien", Mana: 40, Atk: 10, Slot: "Weapon", Class: "Mage"},
		{Name: "Chapeau mystique", Mana: 20, Crit: 5, Slot: "Head", Class: "Mage"},
		{Name: "Robe enchantée", Mana: 40, Spd: 10, Slot: "Body", Class: "Mage"},
		{Name: "Bottes de lévitation", Spd: 15, Mana: 10, Slot: "Legs", Class: "Mage"},
	},
	"Guérisseur": {
		{Name: "Bâton de vie", Mana: 20, Atk: 8, Slot: "Weapon", Class: "Guérisseur"},
		{Name: "Capuche de prêtre", Mana: 15, Def: 5, Slot: "Head", Class: "Guérisseur"},
		{Name: "Robe de lumière", Mana: 35, HP: 20, Slot: "Body", Class: "Guérisseur"},
		{Name: "Sandales bénies", Mana: 10, Spd: 5, Slot: "Legs", Class: "Guérisseur"},
	},
}

// === Équiper un objet ===
func (set *EquipmentSet) Equip(item Equipment, class string) {
	if item.Class != class {
		fmt.Printf("❌ %s ne peut pas être équipé par un %s.\n", item.Name, class)
		return
	}

	switch item.Slot {
	case "Weapon":
		set.Weapon = item
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
	case "Weapon":
		fmt.Printf("❌ %s retiré (arme)\n", set.Weapon.Name)
		set.Weapon = Equipment{}
	case "Head":
		fmt.Printf("❌ %s retiré (tête)\n", set.Head.Name)
		set.Head = Equipment{}
	case "Body":
		fmt.Printf("❌ %s retiré (corps)\n", set.Body.Name)
		set.Body = Equipment{}
	case "Legs":
		fmt.Printf("❌ %s retiré (jambes)\n", set.Legs.Name)
		set.Legs = Equipment{}
	default:
		fmt.Println("❌ Emplacement invalide :", slot)
	}
}
