package save

import (
	"encoding/json"
	"fmt"
	"os"
)

type GameState struct {
	Name      string
	Class     string
	Money     int
	Progress  string
	Inventory map[string]int
}

// Sauvegarder la partie
func SaveGame(state GameState) {
	file, err := os.Create("save.json")
	if err != nil {
		fmt.Println("❌ Erreur lors de la sauvegarde :", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(state)
	if err != nil {
		fmt.Println("❌ Erreur lors de l'encodage :", err)
		return
	}

	fmt.Println("💾 Sauvegarde réussie.")
}

// Charger la partie
func LoadGame() GameState {
	file, err := os.Open("save.json")
	if err != nil {
		fmt.Println("⚠️ Aucune sauvegarde trouvée.")
		return GameState{}
	}
	defer file.Close()

	var state GameState
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&state)
	if err != nil {
		fmt.Println("❌ Erreur lors du chargement :", err)
		return GameState{}
	}

	fmt.Println("✅ Sauvegarde chargée avec succès.")
	return state
}
