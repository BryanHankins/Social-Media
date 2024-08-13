package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/yourusername/socialmedia/internal/db"
    "github.com/yourusername/socialmedia/internal/models"
    "github.com/yourusername/socialmedia/internal/utils"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
    var post models.Post
    json.NewDecoder(r.Body).Decode(&post)
    db.DB.Create(&post)
    utils.RespondJSON(w, http.StatusCreated, post)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var post models.Post
    if err := db.DB.Preload("Comments").Preload("Likes").First(&post, id).Error; err != nil {
        http.Error(w, "Post not found", http.StatusNotFound)
        return
    }

    utils.RespondJSON(w, http.StatusOK, post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var post models.Post
    if err := db.DB.First(&post, id).Error; err != nil {
        http.Error(w, "Post not found", http.StatusNotFound)
        return
    }

    json.NewDecoder(r.Body).Decode(&post)
    db.DB.Save(&post)
    utils.RespondJSON(w, http.StatusOK, post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    if err := db.DB.Delete(&models.Post{}, id).Error; err != nil {
        http.Error(w, "Post not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
