package middleware

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/je117er/ocg-final-project/internal/config"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/utils"
	"net/http"
	"strings"
)

var logger = utils.SugarLog()

func JWTVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("x-access-token")

		header = strings.TrimSpace(header)

		if header == "" {
			utils.ResponseWithJson(w, http.StatusForbidden, utils.ResponseError{Message: "Missing auth token"})
			return
		}

		tk := &models.Token{}

		if _, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Config("JWT_SECRET_KEY")), nil
		}); err != nil {
			logger.Error(err.Error())
			utils.ResponseWithJson(w, http.StatusForbidden, utils.ResponseError{Message: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
