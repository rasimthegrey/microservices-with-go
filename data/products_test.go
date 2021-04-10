package data

import "testing"

func TestChechsValidation(t *testing.T) {
	p := &Product{
		Name:  "rasim",
		Price: 2.2,
		SKU:   "abcd-abcd-abcd",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
