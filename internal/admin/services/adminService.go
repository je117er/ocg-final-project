package services

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/je117er/ocg-final-project/internal/admin"
	"github.com/je117er/ocg-final-project/internal/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var (
	secretKey = "2fcc261931b586ef50afbce104b83b18f69672b8da5487fbee56779edda1b90d"
)

type AdminService struct {
	adminRepo admin.Repository
}

func NewAdminService(adminRepo admin.Repository) admin.Service {
	return &AdminService{adminRepo: adminRepo}
}

func (as *AdminService) Authenticate(ctx context.Context, username string, hashedPassword string) *models.LoginResponse {

	user, err := as.adminRepo.ByUsername(ctx, username)
	if err != nil {
		return &models.LoginResponse{
			Status:   false,
			Username: username,
			Message:  admin.ErrInvalidUsername,
			Token:    "",
		}
	}

	expiresAt := time.Now().Add(time.Minute * 60).Unix()

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword.String), []byte(hashedPassword)); err != nil {
		return &models.LoginResponse{
			Status:   false,
			Username: username,
			Message:  admin.ErrInvalidPassword,
			Token:    "",
		}
	}

	tk := &models.Token{
		UserID:         user.ID,
		Username:       user.Username.String,
		StandardClaims: &jwt.StandardClaims{ExpiresAt: expiresAt},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return &models.LoginResponse{
			Status:   false,
			Username: username,
			Message:  admin.ErrGeneratingToken,
			Token:    "",
		}
	}
	return &models.LoginResponse{
		Status:   true,
		Username: username,
		Message:  "Login success",
		Token:    tokenString,
	}
}
