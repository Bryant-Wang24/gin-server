package model

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Admin struct {
	Auth
	ID int `json:"id" gorm:"primary_key"`
	// CreateAt time.Time `json:"create_at"`
	// UpdateAt time.Time `json:"update_at"`
	// DeletedAt time.Time `json:"deleted_at"`
}
