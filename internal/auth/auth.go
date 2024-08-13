package auth

import (
    "time"

    "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

func GenerateJWT(username string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func HashPassword(password string) (string, error) {
    return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CheckPasswordHash(password, hash string) bool {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
