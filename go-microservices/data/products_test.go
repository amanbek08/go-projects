package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Aman",
		Price: 5.00,
		SKU:   "abs-asd-dfs",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
