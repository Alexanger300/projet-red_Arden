package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func ShowInventory() {
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
		return
	}
}

func main() {
	ShowInventory()
}
