package model

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  int    `json:"is_admin"`
}

type Admin struct {
	Auth
	ID int `json:"id" gorm:"primary_key"`
	// CreateAt time.Time `json:"create_at"`
	// UpdateAt time.Time `json:"update_at"`
	// DeletedAt time.Time `json:"deleted_at"`
}

type Conf struct {
	ClientId     string
	ClientSecret string //GitHub里所获取
	RedirectUrl  string //重定向URL
}
