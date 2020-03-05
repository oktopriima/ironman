package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type UserSocial struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	SocialName string    `json:"social_name"`
	SocialID   string    `json:"social_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (u *UserSocial) TableName() string {
	return "user_socials"
}
