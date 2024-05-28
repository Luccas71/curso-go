package tax

import "testing"

// func TestCalculatetax(t *testing.T) {
// 	amount := 400.0
// 	expected := 5.0

// 	result := CalculateTax(amount)

// 	if result != expected {
// 		t.Errorf("Expected %f but got %f", expected, result)
// 	}
// }

// testando varios valores de um vez
// func TestCalcTexBatch(t *testing.T) {
// 	type calcTex struct {
// 		amount, expect float64
// 	}

// 	table := []calcTex{
// 		{500.0, 5.0},
// 		{1000.0, 10.0},
// 		{1200.0, 10.0},
// 		{120.0, 5.0},
// 	}

// 	for _, item := range table {
// 		result := CalculateTax(item.amount)
// 		if result != item.expect {
// 			t.Errorf("Expected %f but got %f", item.expect, result)
// 		}
// 	}
// }

// func BenchmarkCalculatetax(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		CalculateTax(500.0)
// 	}
// }
// func BenchmarkCalculatetax2(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		CalculateTax2(500.0)
// 	}
// }

// coverage => determina o quanto do meu codigo foi testado
// go test -coverprofile=coverage.out
// mostra a cobertura do codigo em html
// go tool cover -html=coverage.out

//rodando benchmark
// go test -bench=.

//rodando apenas o bench
// go test -bench=. -run=^#

// -count=4   => quantas vezes vai rodar
// -time=3s   => vai rodar por 3 second cada count

// fuzz
func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -1.25, 0, 1580.0, 200.89}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Received %f but expected 0", result)
		}
		if amount >20000 && result != 20 {
			t.Errorf("Received %f but expected 20", result)
		}
	})
}
