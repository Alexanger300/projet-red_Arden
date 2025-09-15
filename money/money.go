package money

import "fmt"

type Money struct {
	Amount   int
	Currency string
}

func (m Money) String() string {
	return fmt.Sprintf("%d %s", m.Amount, m.Currency)
}

// Ajoute un montant
func (m *Money) Add(amount int) {
	m.Amount += amount
}

// Dépense un montant si suffisant, retourne vrai/faux
func (m *Money) Spend(amount int) bool {
	if amount > m.Amount {
		return false
	}
	m.Amount -= amount
	return true
}

// Constructeur pour initialiser de l’or
func NewGold(amount int) Money {
	return Money{
		Amount:   amount,
		Currency: "or",
	}
}
