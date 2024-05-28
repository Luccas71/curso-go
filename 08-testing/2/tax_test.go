package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax := CalculateTax(100.0)
	assert.Equal(t, 5.0, tax)
}
