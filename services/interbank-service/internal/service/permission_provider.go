package service

import (
	"context"

	"github.com/RAF-SI-2025/Banka-4-Backend/common/pkg/auth"
	"github.com/RAF-SI-2025/Banka-4-Backend/common/pkg/jwt"
	"github.com/RAF-SI-2025/Banka-4-Backend/common/pkg/permission"
)

// NoopPermissionProvider returns an empty permission set for every token.
// interbank-service guards its frontend routes by identity type alone
// (auth.RequireIdentityType) — granular permission checks belong to the
// services that own the underlying resources (banking, trading, user).
type NoopPermissionProvider struct{}

func NewNoopPermissionProvider() auth.PermissionProvider {
	return &NoopPermissionProvider{}
}

func (NoopPermissionProvider) GetPermissions(_ context.Context, _ *jwt.Claims) ([]permission.Permission, error) {
	return nil, nil
}
