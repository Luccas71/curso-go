package database

import (
	"github.com/Luccas71/curso-go/09-API/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	//conexao com db
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	//retornar um db
	return &User{DB: db}
}

func (u *User) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
