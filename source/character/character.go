package character

import (
	"fmt"
	"time"

	"github.com/Alexanger300/projet-red_Forge/asset/css"
	"github.com/Alexanger300/projet-red_Forge/source/class"
	"github.com/Alexanger300/projet-red_Forge/source/equipment"
	"github.com/Alexanger300/projet-red_Forge/source/money"
	"github.com/Alexanger300/projet-red_Forge/source/monster"
	"github.com/Alexanger300/projet-red_Forge/source/save"
	"github.com/Alexanger300/projet-red_Forge/source/skills"
)

// Gestion des statuts
type Status struct {
	Name     string
	Duration int
	Damage   int
}

// Structure du personnage
type Character struct {
	Name    string
	Gender  string
	Class   string
	Level   int
	Exp     int
	ExpNext int

	HP      int
	MaxHP   int
	Mana    int
	MaxMana int
	Atk     int
	Def     int
	Spd     int
	Crit    int
	Weapon  string
	Wallet  money.Money
	Skills  []skills.Skill

	// Stats de base (recalcul à chaque update)
	BaseHP   int
	BaseMana int
	BaseAtk  int
	BaseDef  int
	BaseSpd  int
	BaseCrit int

	Statuses  []Status
	Inventory map[string]int
	Weapons   map[string]equipment.Equipment
	Equip     equipment.EquipmentSet
}

// --- Création du personnage ---

func InitCharacter() Character {
	css.Clear()
	var c Character
	var choiceNumber int
	var confirm string
	confirmed := false
	// Nom du personnage
	text1 := "Entrez votre prénom : "
	fmt.Print("\n")
	for _, char := range text1 {
		fmt.Printf(css.Bold+"%c"+css.Reset, char)
		time.Sleep(50 * time.Millisecond)
	}
	time.Sleep(1 * time.Second)
	_, err := fmt.Scan(&c.Name)
	if err != nil {
		fmt.Println("Erreur :", err)
	}
	css.Clear()
	time.Sleep(1 * time.Second)
	// Sexe du personnage
	text2 := "Choisissez le sexe (Homme/Femme/Autre) : "
	for _, char := range text2 {
		fmt.Printf(css.Bold+"%c"+css.Reset, char)
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Scan(&c.Gender)
	time.Sleep(1 * time.Second)
	text3 := fmt.Sprintf("Voici le nom de votre personnage : %s (%s)\n", c.Name, c.Gender)
	fmt.Printf("\n")
	for _, char := range text3 {
		fmt.Printf(css.Bold+"%c"+css.Reset, char)
		time.Sleep(30 * time.Millisecond)
	}
	time.Sleep(1 * time.Second)
	css.Clear()
	// Choix de la classe
	for !confirmed {
		text4 := "\nQuelle Classe voulez-vous ?"
		for _, char := range text4 {
			fmt.Printf(css.Bold+"%c"+css.Reset, char)
			time.Sleep(30 * time.Millisecond)
		}
		time.Sleep(1 * time.Second)
		fmt.Print("\n")
		fmt.Println("1: Paladin ⚔️")
		fmt.Println("2: Géant 🪓")
		fmt.Println("3: Mage 🔮")
		fmt.Println("4: Guérisseur ✨")
		fmt.Print("Votre choix : ")
		fmt.Scan(&choiceNumber)

		var className string
		switch choiceNumber {
		case 1:
			className = "Paladin"
		case 2:
			className = "Géant"
		case 3:
			className = "Mage"
		case 4:
			className = "Guérisseur"
		default:
			fmt.Println("Choix invalide.")
			continue
		}

		stats := class.Classes[className]

		fmt.Printf("\n%s → ", className)
		for _, char := range stats.Description {
			fmt.Printf(css.Bold+"%c"+css.Reset, char)
			time.Sleep(20 * time.Millisecond)
		}
		fmt.Print("\n")
		time.Sleep(1 * time.Second)
		fmt.Printf("PV: %d | ATK: %d | DEF: %d | Mana: %d | SPD: %d | CRIT: %d%% | Arme: %s\n",
			stats.HP, stats.Atk, stats.Def, stats.Mana, stats.Spd, stats.Crit, stats.Weapon)
		time.Sleep(1 * time.Second)
		fmt.Print("Confirmez-vous votre choix ? (Oui/Non) : ")
		fmt.Scan(&confirm)

		if confirm == "Oui" || confirm == "oui" {
			c.Class = className

			// Stats de base
			c.BaseHP = stats.HP
			c.BaseMana = stats.Mana
			c.BaseAtk = stats.Atk
			c.BaseDef = stats.Def
			c.BaseSpd = stats.Spd
			c.BaseCrit = stats.Crit

			// Stats courantes
			c.HP = stats.HP
			c.MaxHP = stats.HP
			c.Mana = stats.Mana
			c.MaxMana = stats.Mana
			c.Atk = stats.Atk
			c.Def = stats.Def
			c.Spd = stats.Spd
			c.Crit = stats.Crit
			c.Weapon = stats.Weapon

			c.Wallet = money.Money{Amount: 100, Currency: "Gold"}
			c.Skills = skills.ClassSkills[className]

			// Expérience
			c.Level = 1
			c.Exp = 0
			c.ExpNext = 10

			// Inventaire
			c.Inventory = map[string]int{
				"Potion de soin":   2,
				"Potion de poison": 1,
			}

			// Armes forgées
			c.Weapons = make(map[string]equipment.Equipment)

			// Équipement vide
			c.Equip = equipment.EquipmentSet{}

			confirmed = true
		}
	}

	return c
}

// Vérifie si le joueur est vivant
func (c *Character) IsAlive() bool {
	return c.HP > 0
}

// Gain d’expérience
func (c *Character) GainExp(amount int) {
	c.Exp += amount
	fmt.Printf("✨ %s gagne %d points d'expérience ! (%d/%d)\n",
		c.Name, amount, c.Exp, c.ExpNext)

	if c.Exp >= c.ExpNext {
		c.LevelUp()
	}
}

// Passage de niveau
func (c *Character) LevelUp() {
	c.Level++
	c.Exp -= c.ExpNext
	c.ExpNext += 10

	switch c.Class {
	case "Paladin":
		c.BaseHP += 15
		c.BaseMana += 5
		c.BaseAtk += 3
		c.BaseDef += 4
	case "Géant":
		c.BaseHP += 20
		c.BaseAtk += 5
		c.BaseDef += 5
	case "Mage":
		c.BaseHP += 8
		c.BaseMana += 15
		c.BaseAtk += 2
		c.BaseDef += 1
	case "Guérisseur":
		c.BaseHP += 10
		c.BaseMana += 12
		c.BaseAtk += 1
		c.BaseDef += 2
	}

	// Recalcul
	c.RecalculateStatsFromEquipment()
	c.HP = c.MaxHP
	c.Mana = c.MaxMana

	fmt.Printf("\n🎉 %s passe au niveau %d !\n", c.Name, c.Level)
	fmt.Printf("Stats : HP %d | Mana %d | ATK %d | DEF %d\n",
		c.MaxHP, c.MaxMana, c.Atk, c.Def)
}

// Gestion inventaire
func (c *Character) AddItem(item string, qty int) {
	c.Inventory[item] += qty
	fmt.Printf("🧳 Vous obtenez %d x %s\n", qty, item)
}

func (c *Character) RemoveItem(item string, qty int) bool {
	if c.Inventory[item] < qty {
		return false
	}
	c.Inventory[item] -= qty
	if c.Inventory[item] == 0 {
		delete(c.Inventory, item)
	}
	return true
}

// Utiliser un objet sur joueur
func (c *Character) UseItem(item string, target *Character) {
	switch item {
	case "Potion de soin":
		if c.RemoveItem(item, 1) {
			heal := 50
			target.HP += heal
			if target.HP > target.MaxHP {
				target.HP = target.MaxHP
			}
			fmt.Printf("🍷 %s utilise une potion de soin → %s récupère %d PV (HP: %d/%d)\n",
				c.Name, target.Name, heal, target.HP, target.MaxHP)
		} else {
			fmt.Println("❌ Aucune potion de soin disponible.")
		}
	case "Élixir de mana":
		if c.RemoveItem(item, 1) {
			manaRestore := 30
			target.Mana += manaRestore
			if target.Mana > target.MaxMana {
				target.Mana = target.MaxMana
			}
			fmt.Printf("💧 %s utilise un élixir de mana → %s récupère %d PM (Mana: %d/%d)\n",
				c.Name, target.Name, manaRestore, target.Mana, target.MaxMana)
		} else {
			fmt.Println("❌ Aucun élixir de mana disponible.")
		}
	}
}

// Utiliser un objet sur monstre
func (c *Character) UseItemOnMonster(item string, target *monster.Monster) {
	switch item {
	case "Potion de poison":
		if c.RemoveItem(item, 1) {
			target.ApplyStatus("Poison", 3, 5)
		} else {
			fmt.Println("❌ Aucune potion de poison disponible.")
		}
	default:
		fmt.Println("❌ Objet inconnu :", item)
	}
}

// Apprendre une compétence
func (c *Character) LearnSkill(newSkill skills.Skill) {
	for _, s := range c.Skills {
		if s.Name == newSkill.Name {
			fmt.Println("❌ Vous connaissez déjà ce sort :", newSkill.Name)
			return
		}
	}
	c.Skills = append(c.Skills, newSkill)
	fmt.Println("✨ Nouveau sort appris :", newSkill.Name)
}

// Utiliser une compétence
func (c *Character) UseSkillOnMonster(skillName string, target *monster.Monster) {
	var s *skills.Skill
	for i := range c.Skills {
		if c.Skills[i].Name == skillName {
			s = &c.Skills[i]
			break
		}
	}
	if s == nil {
		fmt.Println("❌ Sort inconnu :", skillName)
		return
	}

	if c.Mana < s.ManaCost {
		fmt.Println("❌ Pas assez de mana pour lancer", s.Name)
		return
	}
	c.Mana -= s.ManaCost

	if s.IsHeal {
		healAmount := s.Power + (c.Mana / 10)
		c.HP += healAmount
		if c.HP > c.MaxHP {
			c.HP = c.MaxHP
		}
		fmt.Printf("✨ %s utilise %s → récupère %d PV (HP: %d/%d)\n",
			c.Name, s.Name, healAmount, c.HP, c.MaxHP)
		return
	}

	rawDamage := s.Power + (c.Atk / 2)
	if s.IsMagic {
		rawDamage = s.Power + (c.Mana / 5)
	}
	finalDamage := rawDamage - target.Def
	if finalDamage < 0 {
		finalDamage = 0
	}

	target.HP -= finalDamage
	if target.HP < 0 {
		target.HP = 0
	}

	fmt.Printf("🔥 %s utilise %s sur %s → %d dégâts (HP restants : %d/%d)\n",
		c.Name, s.Name, target.Name, finalDamage, target.HP, target.HPMax)
}

// Statuts
func (c *Character) ApplyStatus(name string, duration int, damage int) {
	for _, s := range c.Statuses {
		if s.Name == name {
			fmt.Printf("%s est déjà affecté par %s.\n", c.Name, name)
			return
		}
	}
	c.Statuses = append(c.Statuses, Status{Name: name, Duration: duration, Damage: damage})
	fmt.Printf("%s est maintenant affecté par %s (%d tours).\n", c.Name, name, duration)
}

func (c *Character) UpdateStatuses() {
	var remaining []Status
	for _, s := range c.Statuses {
		if s.Name == "Poison" {
			c.HP -= s.Damage
			fmt.Printf("☠️ %s subit %d dégâts de poison (HP: %d/%d)\n", c.Name, s.Damage, c.HP, c.MaxHP)
			if c.HP <= 0 {
				c.HP = 0
				fmt.Printf("💀 %s est mort à cause du poison !\n", c.Name)
				return
			}
		}
		s.Duration--
		if s.Duration > 0 {
			remaining = append(remaining, s)
		} else {
			fmt.Printf("%s n’est plus affecté par %s.\n", c.Name, s.Name)
		}
	}
	c.Statuses = remaining
}

// --- Recalcul des stats ---
func (c *Character) RecalculateStatsFromEquipment() {
	// Repartir des stats de base
	c.MaxHP = c.BaseHP
	c.MaxMana = c.BaseMana
	c.Atk = c.BaseAtk
	c.Def = c.BaseDef
	c.Spd = c.BaseSpd
	c.Crit = c.BaseCrit

	// Bonus armures (tête, corps, jambes)
	hp, mana, atk, def, spd, crit := c.Equip.TotalStats()
	c.MaxHP += hp
	c.MaxMana += mana
	c.Atk += atk
	c.Def += def
	c.Spd += spd
	c.Crit += crit

	// Bonus arme (séparé !)
	if c.Equip.Weapon.Name != "" {
		c.MaxHP += c.Equip.Weapon.HP
		c.MaxMana += c.Equip.Weapon.Mana
		c.Atk += c.Equip.Weapon.Atk
		c.Def += c.Equip.Weapon.Def
		c.Spd += c.Equip.Weapon.Spd
		c.Crit += c.Equip.Weapon.Crit
	}

	// Clamp PV/mana
	if c.HP > c.MaxHP {
		c.HP = c.MaxHP
	}
	if c.Mana > c.MaxMana {
		c.Mana = c.MaxMana
	}
}

// --- Affichages ---
func (c Character) DisplaySummary() {
	fmt.Println("\n--- Résumé ---")
	fmt.Printf("Nom   : %s (%s)\n", c.Name, c.Gender)
	fmt.Println("Classe:", c.Class)
	fmt.Println("Niveau:", c.Level, "| XP:", c.Exp, "/", c.ExpNext)
}

func (c Character) DisplayFull() {
	fmt.Println("\n=== Informations du personnage ===")
	fmt.Printf("Nom    : %s (%s)\n", c.Name, c.Gender)
	fmt.Printf("Classe : %s\n", c.Class)
	fmt.Printf("Niveau : %d | XP : %d/%d\n", c.Level, c.Exp, c.ExpNext)
	fmt.Printf("HP     : %d/%d\n", c.HP, c.MaxHP)
	fmt.Printf("Mana   : %d/%d\n", c.Mana, c.MaxMana)
	fmt.Printf("ATK    : %d | DEF: %d | SPD: %d | CRIT: %d%%\n", c.Atk, c.Def, c.Spd, c.Crit)
	fmt.Printf("Arme   : %s\n", c.Weapon)
	fmt.Printf("Or     : %d %s\n", c.Wallet.Amount, c.Wallet.Currency)

	if len(c.Skills) > 0 {
		fmt.Println("\n--- Compétences ---")
		for _, s := range c.Skills {
			fmt.Printf("- %s (Mana: %d) → %s\n", s.Name, s.ManaCost, s.Description)
		}
	} else {
		fmt.Println("Aucune compétence connue.")
	}

	if len(c.Statuses) > 0 {
		fmt.Println("\n--- Statuts ---")
		for _, s := range c.Statuses {
			fmt.Printf("- %s (%d tours restants)\n", s.Name, s.Duration)
		}
	}

	if len(c.Inventory) > 0 {
		fmt.Println("\n--- Inventaire ---")
		for item, qty := range c.Inventory {
			fmt.Printf("- %s x%d\n", item, qty)
		}
	}

	if len(c.Weapons) > 0 {
		fmt.Println("\n--- Armes forgées ---")
		for _, w := range c.Weapons {
			fmt.Printf("- %s (ATK: %d, Classe: %s)\n", w.Name, w.Atk, w.Class)
		}
	} else {
		fmt.Println("\n--- Armes forgées ---\nAucune arme.")
	}

	c.Equip.Display()
}

func (c Character) DisplayStatsBar() {

	fmt.Println(css.Beige + "\n=== Statistiques ===" + css.Reset)
	fmt.Printf(css.LightGreen+"Nom: %s"+css.Reset+" | "+css.Gray+"Classe: %s "+" | "+css.Gold+"Niveau: %d"+css.Reset+" | "+css.LightBlue+" XP: %d/%d\n"+css.Reset+css.Yellow+"Gold: %d\n"+css.Reset,
		c.Name, c.Class, c.Level, c.Exp, c.ExpNext, c.Wallet.Amount)
	fmt.Printf(css.Red+"HP: %d/%d"+css.Reset+" | "+css.Violet+"Mana: %d/%d"+css.Reset+" | "+css.Orange+"ATK: %d"+css.Reset+" | "+css.SteelBlue+"DEF: %d"+css.Reset+" | "+css.Green+"SPD: %d"+css.Reset+" | "+css.Yellow+"CRIT: %d%%\n"+css.Reset,
		c.HP, c.MaxHP, c.Mana, c.MaxMana, c.Atk, c.Def, c.Spd, c.Crit)
}

func (c Character) DisplayInventoryAndEquipment() {
	fmt.Println("\n=== Inventaire ===")
	if len(c.Inventory) == 0 {
		fmt.Println("Inventaire vide.")
	} else {
		for item, qty := range c.Inventory {
			fmt.Printf("- %s x%d\n", item, qty)
		}
	}

	fmt.Println("\n=== Armes forgées ===")
	if len(c.Weapons) == 0 {
		fmt.Println("Aucune arme forgée.")
	} else {
		for _, w := range c.Weapons {
			fmt.Printf("- %s (ATK: %d, Classe: %s)\n", w.Name, w.Atk, w.Class)
		}
	}

	fmt.Println("\n=== Équipement ===")
	if c.Equip.Head.Name != "" {
		fmt.Printf("Tête : %s\n", c.Equip.Head.Name)
	} else {
		fmt.Println("Tête : Aucun")
	}
	if c.Equip.Body.Name != "" {
		fmt.Printf("Corps : %s\n", c.Equip.Body.Name)
	} else {
		fmt.Println("Corps : Aucun")
	}
	if c.Equip.Legs.Name != "" {
		fmt.Printf("Jambes : %s\n", c.Equip.Legs.Name)
	} else {
		fmt.Println("Jambes : Aucun")
	}
}

// Charger une sauvegarde
func LoadFromSave(state save.GameState) Character {
	c := Character{
		Name:      state.Name,
		Class:     state.Class,
		Gender:    "Inconnu",
		Level:     state.Level,
		Exp:       state.Exp,
		ExpNext:   state.ExpNext,
		Wallet:    money.Money{Amount: state.Money, Currency: "Gold"},
		Inventory: map[string]int{},
		Equip:     equipment.EquipmentSet{},
		Weapons:   make(map[string]equipment.Equipment),
	}

	if state.Inventory != nil {
		c.Inventory = state.Inventory
	} else {
		c.Inventory = map[string]int{
			"Potion de soin":   2,
			"Potion de poison": 1,
		}
	}

	stats, ok := class.Classes[state.Class]
	if ok {
		// Base
		c.BaseHP = stats.HP
		c.BaseMana = stats.Mana
		c.BaseAtk = stats.Atk
		c.BaseDef = stats.Def
		c.BaseSpd = stats.Spd
		c.BaseCrit = stats.Crit

		// Courantes
		c.MaxHP = stats.HP
		c.HP = c.MaxHP
		c.MaxMana = stats.Mana
		c.Mana = c.MaxMana
		c.Atk = stats.Atk
		c.Def = stats.Def
		c.Spd = stats.Spd
		c.Crit = stats.Crit
		c.Weapon = stats.Weapon
		c.Skills = skills.ClassSkills[state.Class]
	}

	c.RecalculateStatsFromEquipment()
	return c
}
