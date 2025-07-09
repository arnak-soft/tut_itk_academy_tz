package models

import "github.com/google/uuid"

type Wallet struct {
	ID      uuid.UUID `json:"walletId"`
	Balance float64   `json:"balance"`
}