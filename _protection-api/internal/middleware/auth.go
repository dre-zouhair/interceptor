package middleware

import (
	"net/http"
	"slices"
	"strings"

	"github.com/uptrace/bunrouter"
)

type authMiddleware struct {
	apiTokens []string
}

func NewAuthMiddleware(tokens []string) IAuthMiddleware {
	// api tokens should be retrieved from a DB
	return &authMiddleware{
		apiTokens: tokens,
	}
}

type IAuthMiddleware interface {
	AuthMiddleware(next bunrouter.HandlerFunc) bunrouter.HandlerFunc
}

func (m authMiddleware) AuthMiddleware(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		authorization := req.Header.Get("Authorization")
		if !strings.HasPrefix(authorization, "Bearer") {
			http.Error(w, "unauthorized access", http.StatusUnauthorized)
			return nil
		}

		authorizationParts := strings.Split(authorization, "Bearer")

		token := authorizationParts[0]
		if len(authorizationParts) > 1 {
			token = strings.TrimSpace(authorizationParts[1])
		}

		if !slices.Contains(m.apiTokens, token) {
			http.Error(w, "unauthorized access", http.StatusUnauthorized)
			return nil
		}
		return next(w, req)
	}
}
