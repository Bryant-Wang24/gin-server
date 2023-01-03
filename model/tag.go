package model

import "time"

type Tag struct {
	ID         int       `json:"id" gorm:"primary_key"`
	Name       string    `json:"name"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	ArticleNum int       `json:"articleNum"`
	Status     int       `json:"status"`
}

type UpdateTag struct {
	Name       string    `json:"name"`
	ID         int       `json:"id"`
	CreateTime time.Time `json:"createTime"`
}
