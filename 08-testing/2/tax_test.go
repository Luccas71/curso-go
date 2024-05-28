package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax := CalculateTax(100.0)
	assert.Equal(t, 5.0, tax)

	tax = CalculateTax(15000.25)
	assert.Equal(t, 10.0, tax)

	tax = CalculateTax(25000.25)
	assert.Equal(t, 20.0, tax)

	tax = CalculateTax(0.0)
	assert.Equal(t, 0.0, tax)
}

// teste com mock
func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "error saving tax")

	repository.AssertExpectations(t)
}
