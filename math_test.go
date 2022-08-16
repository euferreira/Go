package main

func TestSoma(t *testing.T) {
	total := Somar(1, 2)

	if total != 3 {
		t.Errorf("Soma de 1 + 2 deve ser 3, mas foi %d", total)
	}
}