package entity

type User struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	IsAdmin  bool   `json:"is_admin"`
}
