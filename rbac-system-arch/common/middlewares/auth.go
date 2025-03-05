package middlewares

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

var Seckey = "my_secret_key"

type Claims struct {
	UserId string   `json:"userId"`
	Roles  []string `json:"roles"`
	jwt.RegisteredClaims
}

func CreateToken(id string, roles []string) (token string, err error) {
	tokenStr := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{UserId: id, Roles: roles})
	token, err = tokenStr.SignedString([]byte(Seckey))
	if err != nil {
		return "", err
	}
	return token, nil
}

type contextKey string

const (
	UserIDKey contextKey = "userId"
	RolesKey  contextKey = "roles"
)

func AuthMiddleware(next http.Handler, permission []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		tokenStr := authHeader[len("Bearer "):]
		token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(Seckey), nil
		})
		if err != nil {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		if !token.Valid {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		claims, ok := token.Claims.(*Claims)
		if !ok {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserIDKey, claims.UserId)
		ctx = context.WithValue(ctx, RolesKey, claims.Roles)
		r = r.WithContext(ctx)
		r = r.WithContext(context.WithValue(r.Context(), "isApiKey", false))
		roles := make(map[string]bool)
		for i := range claims.Roles {
			roles[claims.Roles[i]] = true
		}
		for i := range permission {
			if _, ok := roles[permission[i]]; !ok {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
