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

// testando varios valores de um vez
func TestCalcTexBatch(t *testing.T) {
	type calcTex struct {
		amount, expect float64
	}

	table := []calcTex{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1200.0, 10.0},
		{120.0, 5.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expect {
			t.Errorf("Expected %f but got %f", item.expect, result)
		}
	}
}
