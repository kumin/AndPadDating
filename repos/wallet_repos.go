package repos

import (
	"context"

	"github.com/kumin/BityDating/entities"
	"github.com/shopspring/decimal"
)

type WalletRepo interface {
	CreateOne(ctx context.Context, transaction *entities.WalletTransaction) (*entities.WalletTransaction, error)
	GetTotalAmount(ctx context.Context, userId int64) (decimal.Decimal, error)
	ListTransactions(ctx context.Context, userId int64, page, limit int) ([]*entities.WalletTransaction, error)
}
