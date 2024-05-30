package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProdu(t *testing.T) {
	p, err := NewProduct("Produto 1", 10.0)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Produto 1", p.Name)
	assert.Equal(t, 10.0, p.Price)
}

func TestNewProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 10.0)
	assert.Nil(t, p)
	assert.Equal(t, ErrorNameIsRequired, err)
}
func TestNewProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Produto 1", 0)
	assert.Nil(t, p)
	assert.Equal(t, ErrorPriceIsRequired, err)
}
func TestNewProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Produto 1", -10.0)
	assert.Nil(t, p)
	assert.Equal(t, ErrorInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("Produto 1", 10.0)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
