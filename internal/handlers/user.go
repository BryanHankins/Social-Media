package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/yourusername/socialmedia/internal/db"
    "github.com/yourusername/socialmedia/internal/models"
    "github.com/yourusername/socialmedia/internal/utils"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var user models.User
    if err := db.DB.First(&user, id).Error; err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    utils.RespondJSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var user models.User
    if err := db.DB.First(&user, id).Error; err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    json.NewDecoder(r.Body).Decode(&user)
    db.DB.Save(&user)
    utils.RespondJSON(w, http.StatusOK, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    if err := db.DB.Delete(&models.User{}, id).Error; err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
