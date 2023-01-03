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
