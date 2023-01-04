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
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateTagStatus struct {
	ID     int `json:"id"`
	Status int `json:"status"`
}
