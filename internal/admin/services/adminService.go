package services

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/admin"
)

type AdminService struct {
	adminRepo admin.Repository
}

func NewAdminService(adminRepo admin.Repository) admin.Service {
	return &AdminService{adminRepo: adminRepo}
}

