package inventory

import (
	"encoding/json"
	"fmt"
	"os"
)

const inventoryFile = "Inventory.json"

// Charger l’inventaire
func loadInventory() (map[string]int, error) {
	inventory := make(map[string]int)

	// Si le fichier n’existe pas → inventaire vide
	if _, err := os.Stat(inventoryFile); os.IsNotExist(err) {
		return inventory, nil
	}

	// Lire le fichier
	data, err := os.ReadFile(inventoryFile)
	if err != nil {
		return nil, err
	}

	// Convertir JSON → map
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

// Ajouter un objet
func AddItem(name string, qty int) {
	inventory, err := loadInventory()
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	inventory[name] += qty

	if err := saveInventory(inventory); err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	fmt.Printf("Ajout : %d %s | Inventaire : %v\n", qty, name, inventory)
}

// Retirer un objet
func RemoveItem(name string, qty int) {
	inventory, err := loadInventory()
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	if inventory[name] < qty {
		fmt.Println("Pas assez de", name, "dans l'inventaire.")
		return
	}

	inventory[name] -= qty
	if inventory[name] == 0 {
		delete(inventory, name) // supprime si 0
	}

	if err := saveInventory(inventory); err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	fmt.Printf("Suppression : %d %s | Inventaire : %v\n", qty, name, inventory)
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

// Afficher l’inventaire
func ShowInventory() {
	inventory, err := loadInventory()
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	if len(inventory) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}

	fmt.Println("=== Inventaire ===")
	for item, qty := range inventory {
		fmt.Printf("- %s : %d\n", item, qty)
	}
}
