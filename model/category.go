package model

import "time"

type Category struct {
	ID         int       `json:"id" gorm:"primary_key"`
	Name       string    `json:"name"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	ArticleNum int       `json:"articleNum"`
}

type UpdateCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
