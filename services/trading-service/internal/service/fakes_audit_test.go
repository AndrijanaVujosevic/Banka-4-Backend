package service

import (
	"context"

	"github.com/RAF-SI-2025/Banka-4-Backend/services/trading-service/internal/audit"
)

type fakeAuditRepo struct {
	saveErr error
}

func (f *fakeAuditRepo) Save(_ context.Context, _ *audit.AuditLog) error {
	return f.saveErr
}
