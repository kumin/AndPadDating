package entities

import (
	"time"

	"github.com/shopspring/decimal"
)

type TransactionType int

const (
	Topup TransactionType = iota
	Spent
)

type WalletTransaction struct {
	Id              int64           `json:"id,omitempty"`
	UserId          int64           `json:"user_id,omitempty"`
	TransactionType TransactionType `json:"transaction_type,omitempty"`
	CreatedAt       time.Time       `json:"created_at,omitempty"`
	Amount          decimal.Decimal `json:"amount,omitempty"`
	Metadata        string          `json:"metadata,omitempty"`
}

func (WalletTransaction) TableName() string {
	return "wallet_transaction"
}
