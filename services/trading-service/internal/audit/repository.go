package audit

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Save(ctx context.Context, entry *AuditLog) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Save(ctx context.Context, entry *AuditLog) error {
	return r.db.WithContext(ctx).Create(entry).Error
}
