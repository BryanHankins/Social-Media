package middleware

import (
    "net/http"
    "strings"

    "github.com/dgrijalva/jwt-go"
    "github.com/yourusername/socialmedia/internal/auth"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenStr := r.Header.Get("Authorization")
        if tokenStr == "" {
            http.Error(w, "Missing token", http.StatusUnauthorized)
            return
        }

        tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

        claims := &auth.Claims{}
        token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
            return auth.JwtKey, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    })
}
