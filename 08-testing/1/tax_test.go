package tax

import "testing"

func TestCalculatetax(t *testing.T) {
	amount := 400.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}
