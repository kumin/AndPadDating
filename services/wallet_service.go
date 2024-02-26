package services

import (
	"context"

	"github.com/kumin/BityDating/entities"
	"github.com/kumin/BityDating/repos"
	"github.com/shopspring/decimal"
)

type WalletService struct {
	walletRepo repos.WalletRepo
}

func NewWalletService(
	walletRepo repos.WalletRepo,
) *WalletService {
	return &WalletService{
		walletRepo: walletRepo,
	}
}

func (w *WalletService) CreateTransaction(
	ctx context.Context,
	transaction *entities.WalletTransaction,
) (*entities.WalletTransaction, error) {
	return w.walletRepo.CreateOne(ctx, transaction)
}

func (w *WalletService) GetTotal(
	ctx context.Context,
	userId int64,
) (*decimal.Decimal, error) {
	return w.walletRepo.GetTotalAmount(ctx, userId)
}

func (w *WalletService) ListTransactions(
	ctx context.Context,
	userId int64,
	page, limit int,
) ([]*entities.WalletTransaction, error) {
	return w.walletRepo.ListTransactions(ctx, userId, page, limit)
}
