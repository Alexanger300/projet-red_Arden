package save

import (
	"encoding/json"
	"fmt"
	"os"
)

type GameState struct {
	Name      string
	Class     string
	Inventory map[string]int
	Money     int
	Progress  string // ex: "foret", "montagne"
}

func SaveGame(state GameState) {
	data, _ := json.MarshalIndent(state, "", "  ")
	os.WriteFile("save.json", data, 0644)
	fmt.Println("💾 Partie sauvegardée !")
}

func LoadGame() GameState {
	var state GameState
	data, err := os.ReadFile("save.json")
	if err != nil {
		fmt.Println("⚠️ Pas de sauvegarde trouvée.")
		return state
	}
	json.Unmarshal(data, &state)
	fmt.Println("📂 Partie chargée.")
	return state
}
