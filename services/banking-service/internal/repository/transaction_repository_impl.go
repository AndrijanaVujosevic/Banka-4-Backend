package repository

import (
	"banking-service/internal/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(ctx context.Context, transaction *model.Transaction) error {
	return r.db.WithContext(ctx).Create(transaction).Error
}

func (r *transactionRepository) GetByID(ctx context.Context, id uint) (*model.Transaction, error) {
	var transaction model.Transaction
	result := r.db.WithContext(ctx).First(&transaction, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &transaction, result.Error
}

func (r *transactionRepository) GetByPayerAccountNumber(ctx context.Context, accountNumber string) ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	err := r.db.WithContext(ctx).Where("payer_account_number = ?", accountNumber).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *transactionRepository) GetByRecipientAccountNumber(ctx context.Context, accountNumber string) ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	err := r.db.WithContext(ctx).Where("recipient_account_number = ?", accountNumber).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *transactionRepository) Update(ctx context.Context, transaction *model.Transaction) error {
	return r.db.WithContext(ctx).Save(transaction).Error
}
