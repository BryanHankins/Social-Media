package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/yourusername/socialmedia/internal/auth"
    "github.com/yourusername/socialmedia/internal/db"
    "github.com/yourusername/socialmedia/internal/models"
    "github.com/yourusername/socialmedia/internal/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)
    hashedPassword, _ := auth.HashPassword(user.Password)
    user.Password = hashedPassword
    db.DB.Create(&user)
    utils.RespondJSON(w, http.StatusCreated, user)
}

func Login(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)

    var dbUser models.User
    db.DB.Where("username = ?", user.Username).First(&dbUser)

    if !auth.CheckPasswordHash(user.Password, dbUser.Password) {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    token, _ := auth.GenerateJWT(user.Username)
    utils.RespondJSON(w, http.StatusOK, map[string]string{"token": token})
}
