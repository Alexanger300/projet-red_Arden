package money

import (
	"fmt"
)

type Money struct {
	// Définition des attributs
	amount   int
	currency string
}

func (m Money) String() string {
	// Affichage du montant et de la devise
	return fmt.Sprintf("%d %s", m.amount, m.currency)
}

func (m *Money) Add(amount int) {
	// Ajout d'un montant
	m.amount += amount
}

func (m *Money) Spend(amount int) bool {
	// Dépense d'un montant si suffisant
	if amount > m.amount {
		return false
	}
	m.amount -= amount
	return true
}
