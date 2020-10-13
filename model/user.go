package model

type User struct {
	ID    uint64 `gorm:"primary_key" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
