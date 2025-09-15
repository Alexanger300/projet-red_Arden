package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func CreateInventory() { // Creer un inventaire vide
	Inventory := make(map[string]interface{})

	jsonData, err := json.MarshalIndent(Inventory, "", "  ")
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	file, err := os.Create("Inventory.json")
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Erreur:", err)
	}
	return
}

func removeItems(itemName string, quantity int) { //Enlever un item de l'inventaire
	file, err := os.ReadFile("Inventory.json")
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	Inventory := map[string]int{}
	err = json.Unmarshal(file, &Inventory)
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	Inventory[itemName] -= quantity // Enlever la quantite specifiee de l'item

	updatedData, err := json.MarshalIndent(Inventory, "", "  ")
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	err = os.WriteFile("Inventory.json", updatedData, 0644)
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	for item, qty := range Inventory {
		fmt.Println(item, ":", qty)
	}
}

func addItems(itemName string, quantity int) { //Ajouter un item a l'inventaire
	file, err := os.ReadFile("Inventory.json")
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	Inventory := map[string]int{}
	err = json.Unmarshal(file, &Inventory)
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	Inventory[itemName] += quantity // Ajouter la quantite a l'item

	updatedData, err := json.MarshalIndent(Inventory, "", "  ")
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	err = os.WriteFile("Inventory.json", updatedData, 0644)
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	for item, qty := range Inventory {
		fmt.Println(item, ":", qty)
	}
}
func main() {
	removeItems("Potion de soin", 3)
}
