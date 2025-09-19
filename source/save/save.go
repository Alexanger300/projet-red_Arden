package save

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Alexanger300/projet-red_Forge/asset/css"
)

// Structure sérialisable pour l'équipement
type EquipmentState struct {
	Head   string `json:"head,omitempty"`
	Body   string `json:"body,omitempty"`
	Legs   string `json:"legs,omitempty"`
	Weapon string `json:"weapon,omitempty"`
}

// Structure de sauvegarde complète
type GameState struct {
	SlotID    int            `json:"slot_id"`   // ID du slot
	Name      string         `json:"name"`      // Nom du personnage
	Class     string         `json:"class"`     // Classe
	Money     int            `json:"money"`     // Argent
	Progress  string         `json:"progress"`  // Étape actuelle
	Inventory map[string]int `json:"inventory"` // Inventaire (objet → quantité)
	Equip     EquipmentState `json:"equip"`     // Équipement (struct claire)
	Level     int            `json:"level"`     // Niveau
	Exp       int            `json:"exp"`
	ExpNext   int            `json:"exp_next"` // Exp pour passer au lvl suivant
}

// Sauvegarder la partie
func SaveGame(state GameState) {
	fileName := fmt.Sprintf("save_slot_%d.json", state.SlotID)

	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		fmt.Println("❌ Erreur lors de l'encodage :", err)
		return
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Println("❌ Erreur lors de l'écriture :", err)
		return
	}

	fmt.Printf("💾 Partie sauvegardée dans le slot %d.\n", state.SlotID)
}

// Charger une partie depuis un slot
func LoadGame(slotID int) (GameState, error) {
	fileName := fmt.Sprintf("save_slot_%d.json", slotID)

	data, err := os.ReadFile(fileName)
	if err != nil {
		return GameState{}, fmt.Errorf("⚠️ Aucun fichier trouvé pour le slot %d", slotID)
	}

	var state GameState
	err = json.Unmarshal(data, &state)
	if err != nil {
		return GameState{}, fmt.Errorf("❌ Erreur lors du chargement du slot %d", slotID)
	}
	css.Clear()
	fmt.Printf("✅ Sauvegarde du slot %d chargée avec succès.\n", slotID)
	return state, nil
}

// Vérifie si un slot existe
func SlotExists(slotID int) bool {
	fileName := fmt.Sprintf("save_slot_%d.json", slotID)
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

// Supprimer un slot
func DeleteSlot(slotID int) {
	fileName := fmt.Sprintf("save_slot_%d.json", slotID)
	err := os.Remove(fileName)
	if err != nil {
		fmt.Printf("❌ Impossible de supprimer le slot %d.\n", slotID)
		return
	}
	fmt.Printf("🗑️ Slot %d supprimé.\n", slotID)
}

// Nouvelle fonction : choisir un slot
func ChooseSlot() int {
	var slot int
	fmt.Println("\n=== Sélection du slot ===")
	fmt.Println("1. Slot 1")
	fmt.Println("2. Slot 2")
	fmt.Println("3. Slot 3")
	fmt.Print("Votre choix : ")
	fmt.Scan(&slot)

	if slot < 1 || slot > 3 {
		fmt.Println("❌ Choix invalide, slot 1 sélectionné par défaut.")
		return 1
	}
	return slot
}
