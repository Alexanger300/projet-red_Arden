package inventory

import (
	"encoding/json"
	"fmt"
	"os"
)

const inventoryFile = "Inventory.json"
const maxPerItem = 10 // Limite max par type d’objet
var MaxTotal = 40     //  Limite globale exportée et modifiable

// Charger l’inventaire depuis le fichier
func loadInventory() (map[string]int, error) {
	inventory := make(map[string]int)

	if _, err := os.Stat(inventoryFile); os.IsNotExist(err) {
		return inventory, nil
	}

	data, err := os.ReadFile(inventoryFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &inventory)
	if err != nil {
		return nil, err
	}

	return inventory, nil
}

// Sauvegarder l’inventaire
func saveInventory(inventory map[string]int) error {
	data, err := json.MarshalIndent(inventory, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(inventoryFile, data, 0644)
}

// Compter le total d’objets
func totalItems(inventory map[string]int) int {
	total := 0
	for _, qty := range inventory {
		total += qty
	}
	return total
}

// Ajouter un objet
func AddItem(name string, qty int) {
	inventory, err := loadInventory()
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	// Limite par type
	if inventory[name]+qty > maxPerItem {
		fmt.Printf("❌ Impossible d'ajouter %d %s. Limite : %d (vous en avez déjà %d).\n",
			qty, name, maxPerItem, inventory[name])
		return
	}

	// Limite globale
	if totalItems(inventory)+qty > MaxTotal {
		fmt.Printf("❌ Votre sac est plein (%d/%d objets).\n", totalItems(inventory), MaxTotal)
		return
	}

	inventory[name] += qty

	if err := saveInventory(inventory); err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	fmt.Printf("✅ Ajouté : %d %s | Total : %d (Sac: %d/%d)\n",
		qty, name, inventory[name], totalItems(inventory), MaxTotal)
}

// Retirer un objet
func RemoveItem(name string, qty int) {
	inventory, err := loadInventory()
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	if inventory[name] < qty {
		fmt.Printf("❌ Pas assez de %s (vous en avez %d).\n", name, inventory[name])
		return
	}

	inventory[name] -= qty
	if inventory[name] == 0 {
		delete(inventory, name)
	}

	if err := saveInventory(inventory); err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	fmt.Printf("✅ Retiré : %d %s | Restant : %d (Sac: %d/%d)\n",
		qty, name, inventory[name], totalItems(inventory), MaxTotal)
}

// Vérifier si on possède un objet
func HasItem(name string, qty int) bool {
	inventory, err := loadInventory()
	if err != nil {
		fmt.Println("Erreur:", err)
		return false
	}
	return inventory[name] >= qty
}

// Afficher inventaire
func ShowInventory() {
	inventory, err := loadInventory()
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	if len(inventory) == 0 {
		fmt.Println("📦 Inventaire vide.")
		return
	}

	fmt.Printf("=== 📦 Inventaire (%d/%d objets) ===\n", totalItems(inventory), MaxTotal)
	for item, qty := range inventory {
		fmt.Printf("- %s : %d/%d\n", item, qty, maxPerItem)
	}
}

// ✅ Améliorer la capacité du sac
func UpgradeBag(slots int) {
	MaxTotal += slots
	fmt.Printf("👜 Votre sac peut maintenant contenir %d objets au total !\n", MaxTotal)
}
