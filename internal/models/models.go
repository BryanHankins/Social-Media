package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username  string `gorm:"unique"`
    Email     string `gorm:"unique"`
    Password  string
    Followers []Follow `gorm:"foreignKey:FollowingID"`
    Following []Follow `gorm:"foreignKey:FollowerID"`
}

type Post struct {
    gorm.Model
    UserID    uint
    Content   string
    Likes     []Like
    Comments  []Comment
}

type Comment struct {
    gorm.Model
    PostID    uint
    UserID    uint
    Content   string
}

type Like struct {
    gorm.Model
    PostID    uint
    UserID    uint
}

type Follow struct {
    gorm.Model
    FollowerID  uint
    FollowingID uint
}

type Notification struct {
    gorm.Model
    UserID  uint
    Message string
    Seen    bool
}
