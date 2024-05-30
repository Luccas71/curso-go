package database

import "github.com/Luccas71/curso-go/09-API/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email entity.User) (entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindById(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
