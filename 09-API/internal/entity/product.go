package entity

import (
	"errors"
	"time"

	"github.com/Luccas71/curso-go/09-API/pkg/entity"
)

// validações
var (
	ErrorIDIsRequired    = errors.New("id is required")
	ErrorInvalidID       = errors.New("invalid id")
	ErrorNameIsRequired  = errors.New("name is required")
	ErrorPriceIsRequired = errors.New("price is required")
	ErrorInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrorIDIsRequired
	}
	if _, error := entity.ParseID(p.ID.String()); error != nil {
		return ErrorInvalidID
	}
	if p.Name == "" {
		return ErrorNameIsRequired
	}
	if p.Price == 0 {
		return ErrorPriceIsRequired
	}
	if p.Price < 0 {
		return ErrorInvalidPrice
	}
	return nil
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}
