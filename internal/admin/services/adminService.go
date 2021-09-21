package services

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/je117er/ocg-final-project/internal/admin"
	"github.com/je117er/ocg-final-project/internal/config"
	"github.com/je117er/ocg-final-project/internal/models"
	"golang.org/x/crypto/bcrypt"
	"time"
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

	secretKey := config.Config("JWT_SECRET_KEY")
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
