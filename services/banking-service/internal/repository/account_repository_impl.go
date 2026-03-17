package repository

import (
	"banking-service/internal/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) Create(ctx context.Context, account *model.Account) error {
	return r.db.WithContext(ctx).Create(account).Error
}

func (r *accountRepository) AccountNumberExists(ctx context.Context, accountNumber string) (bool, error) {
	var count int64

	err := r.db.WithContext(ctx).
		Model(&model.Account{}).
		Where("account_number = ?", accountNumber).
		Count(&count).
		Error

	return count > 0, err
}

func (r *accountRepository) GetByAccountNumber(ctx context.Context, accountNumber string) (*model.Account, error) {
	var account model.Account
	result := r.db.WithContext(ctx).Where("account_number = ?", accountNumber).First(&account)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &account, result.Error
}

func (r *accountRepository) Update(ctx context.Context, account *model.Account) error {
	return r.db.WithContext(ctx).Save(account).Error
}
