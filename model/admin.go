package model

import (
	"time"
)

type Auth struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type Admin struct {
	Auth
	ID        int       `json:"id" gorm:"primary_key"`
	CreateAt  time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
