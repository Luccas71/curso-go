package tax

import (
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
