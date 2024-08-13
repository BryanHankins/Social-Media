package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/yourusername/socialmedia/internal/db"
    "github.com/yourusername/socialmedia/internal/models"
    "github.com/yourusername/socialmedia/internal/utils"
)

func FollowUser(w http.ResponseWriter, r *http.Request) {
    followerID := utils.ExtractUserIDFromRequest(r)
    vars := mux.Vars(r)
    followingID := vars["id"]

    follow := models.Follow{FollowerID: followerID, FollowingID: utils.StrToUint(followingID)}
    db.DB.Create(&follow)

    w.WriteHeader(http.StatusCreated)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
    followerID := utils.ExtractUserIDFromRequest(r)
    vars := mux.Vars(r)
    followingID := vars["id"]

    db.DB.Where("follower_id = ? AND following_id = ?", followerID, followingID).Delete(&models.Follow{})

    w.WriteHeader(http.StatusOK)
}

func GetFollowers(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userID := vars["id"]

    var followers []models.Follow
    db.DB.Where("following_id = ?", userID).Find(&followers)

    utils.RespondJSON(w, http.StatusOK, followers)
}

func GetFollowing(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userID := vars["id"]

    var following []models.Follow
    db.DB.Where("follower_id = ?", userID).Find(&following)

    utils.RespondJSON(w, http.StatusOK, following)
}
