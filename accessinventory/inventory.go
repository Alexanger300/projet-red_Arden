package inventory

import (
	"encoding/json"
	"fmt"
	"os"
)

const inventoryFile = "Inventory.json"

// Charge l’inventaire depuis le fichier JSON
func loadInventory() (map[string]int, error) {
	inventory := make(map[string]int)

	// Si le fichier n’existe pas, retourne un inventaire vide
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

// Sauvegarde l’inventaire dans le fichier JSON
func saveInventory(inventory map[string]int) error {
	data, err := json.MarshalIndent(inventory, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(inventoryFile, data, 0644)
}

// Ajoute un objet
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
	fmt.Println("Ajout :", qty, name, "| Inventaire :", inventory)
}

// Retire un objet
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
		delete(inventory, name) // supprime les objets à 0
	}

	if err := saveInventory(inventory); err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	fmt.Println("Suppression :", qty, name, "| Inventaire :", inventory)
}

// Affiche l’inventaire complet
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
