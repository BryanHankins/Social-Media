package main

import (
    "log"
    "net/http"

    "github.com/yourusername/socialmedia/internal/db"
    "github.com/yourusername/socialmedia/internal/handlers"
    "github.com/yourusername/socialmedia/internal/middleware"

    "github.com/gorilla/mux"
)

func main() {
    // Initialize the database
    db.Connect()
    db.Migrate()

    // Initialize the router
    r := mux.NewRouter()

    // Public routes
    r.HandleFunc("/register", handlers.Register).Methods("POST")
    r.HandleFunc("/login", handlers.Login).Methods("POST")

    // Protected routes
    protected := r.PathPrefix("/api").Subrouter()
    protected.Use(middleware.AuthMiddleware)
    protected.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
    protected.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
    protected.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
    protected.HandleFunc("/posts", handlers.CreatePost).Methods("POST")
    protected.HandleFunc("/posts/{id}", handlers.GetPost).Methods("GET")
    protected.HandleFunc("/posts/{id}", handlers.UpdatePost).Methods("PUT")
    protected.HandleFunc("/posts/{id}", handlers.DeletePost).Methods("DELETE")
    protected.HandleFunc("/posts/{id}/comments", handlers.CreateComment).Methods("POST")
    protected.HandleFunc("/posts/{id}/comments/{commentId}", handlers.DeleteComment).Methods("DELETE")
    protected.HandleFunc("/posts/{id}/like", handlers.LikePost).Methods("POST")
    protected.HandleFunc("/posts/{id}/unlike", handlers.UnlikePost).Methods("POST")
    protected.HandleFunc("/users/{id}/follow", handlers.FollowUser).Methods("POST")
    protected.HandleFunc("/users/{id}/unfollow", handlers.UnfollowUser).Methods("POST")
    protected.HandleFunc("/users/{id}/followers", handlers.GetFollowers).Methods("GET")
    protected.HandleFunc("/users/{id}/following", handlers.GetFollowing).Methods("GET")

    // Start the server
    log.Fatal(http.ListenAndServe(":8080", r))
}
