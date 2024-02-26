package mysql

import (
	"context"

	"github.com/kumin/BityDating/entities"
	"github.com/kumin/BityDating/infras"
	"github.com/kumin/BityDating/repos"
	"github.com/shopspring/decimal"
)

var _ repos.WalletRepo = &WalletMysqlRepo{}

type WalletMysqlRepo struct {
	db *infras.MysqlConnector
}

func NewWalletMysqlRepo(
	db *infras.MysqlConnector,
) *WalletMysqlRepo {
	return &WalletMysqlRepo{
		db: db,
	}
}

func (w *WalletMysqlRepo) CreateOne(
	ctx context.Context,
	transaction *entities.WalletTransaction,
) (*entities.WalletTransaction, error) {
	if err := w.db.Client.WithContext(ctx).Create(transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (w *WalletMysqlRepo) GetTotalAmount(
	ctx context.Context,
	userId int64,
) (*decimal.Decimal, error) {
	var totalAmount decimal.Decimal
	rawQuery := `
  SELECT SUM(
    CASE
      WHEN transaction_type = 1 THEN amount
      ELSE -amount
    END) total_amount
  FROM wallet_transaction
  WHERE user_id = ?
  `
	if err := w.db.Client.WithContext(ctx).
		Raw(rawQuery, userId).Scan(&totalAmount).Error; err != nil {
		return &decimal.Zero, err
	}
	return &totalAmount, nil
}

func (w *WalletMysqlRepo) ListTransactions(
	ctx context.Context,
	userId int64,
	page, limit int,
) ([]*entities.WalletTransaction, error) {
	var transactions []*entities.WalletTransaction
	if err := w.db.Client.WithContext(ctx).Where("user_id=?", userId).Limit(limit).Offset(page * limit).Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}
