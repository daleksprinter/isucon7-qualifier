package main

import (
	"time"
)

type User struct {
	ID          int64     `json:"-" db:"id"`
	Name        string    `json:"name" db:"name"`
	Salt        string    `json:"-" db:"salt"`
	Password    string    `json:"-" db:"password"`
	DisplayName string    `json:"display_name" db:"display_name"`
	AvatarIcon  string    `json:"avatar_icon" db:"avatar_icon"`
	CreatedAt   time.Time `json:"-" db:"created_at"`
}

type Message struct {
	ID        int64     `db:"id"`
	ChannelID int64     `db:"channel_id"`
	UserID    int64     `db:"user_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}

type ChannelInfo struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	UpdatedAt   time.Time `db:"updated_at"`
	CreatedAt   time.Time `db:"created_at"`
}

type MessageAndUser struct {
	ID        int64     `db:"id"`
	ChannelID int64     `db:"channel_id"`
	UserID    int64     `db:"user_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`

	UserPK          int64     `json:"-" db:"user_pk"`
	UserName        string    `json:"name" db:"name"`
	UserSalt        string    `json:"-" db:"salt"`
	UserPassword    string    `json:"-" db:"password"`
	UserDisplayName string    `json:"display_name" db:"display_name"`
	UserAvatarIcon  string    `json:"avatar_icon" db:"avatar_icon"`
	UserCreatedAt   time.Time `json:"-" db:"user_created_at"`
}
