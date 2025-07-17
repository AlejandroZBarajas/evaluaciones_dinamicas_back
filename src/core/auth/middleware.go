package auth

import (
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "no autorizado", http.StatusUnauthorized)
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		tok, err := ParseToken(tokenStr)
		if err != nil || !tok.Valid {
			http.Error(w, "token inválido", http.StatusUnauthorized)
			return
		}

		// Aquí puedes extraer claims si quieres pasar datos al contexto

		next.ServeHTTP(w, r)
	})
}
