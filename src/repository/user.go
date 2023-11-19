package repository

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `json:"userName"`
	Email    string `gorm:"not null" gorm:"unique" json:"email"`
	Password string `json:"password"`
}

func (receiver *User) Save() error {
	result := db.Create(receiver)
	return result.Error
}

func (receiver *User) FindByEmail(email string) ([]User, error) {
	var users []User
	result := db.Where("email = ?", email).Find(&users)
	return users, result.Error
}

func (receiver *User) FindById(userId uint) error {
	result := db.First(receiver, userId)
	return result.Error
}

func (receiver *User) Update() error {
	result := db.Save(receiver)
	return result.Error
}
