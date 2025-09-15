package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func addItems(itemName string, quantity int) {
	file, err := os.ReadFile("Inventory.json")
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	var inventory map[string]int
	err = json.Unmarshal(file, &inventory)
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	inventory[itemName] -= quantity

	updatedData, err := json.MarshalIndent(inventory, "", "  ")
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	err = os.WriteFile("Inventory.json", updatedData, 0644)
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	fmt.Println("Inventaire mis à jour :", inventory)
}

func main() {
	addItems("Potion de soin", 3)
	addItems("Épée en fer", 1)
}
