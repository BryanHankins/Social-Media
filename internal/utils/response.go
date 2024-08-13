package utils

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/dgrijalva/jwt-go"
    "github.com/yourusername/socialmedia/internal/auth"
)

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(payload)
}

func ExtractUserIDFromRequest(r *http.Request) uint {
    tokenStr := r.Header.Get("Authorization")
    claims := &auth.Claims{}
    tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

    token, _ := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return auth.JwtKey, nil
    })

    return StrToUint(claims.Subject)
}

func StrToUint(s string) uint {
    u64, _ := strconv.ParseUint(s, 10, 64)
    return uint(u64)
}
