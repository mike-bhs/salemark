package search

import (
	"fmt"
	xid "github.com/rs/xid"
)

type Product struct {
	Model string `json:"model"`
}

func (p Product) String() string {
	return fmt.Sprintf("\n\tProduct: %s", p.Model)
}

func (p Product) Id() string {
	return xid.New().String()
}

func (p Product) Type() string {
	return "products"
}

func SeedProducts() []Product {
	var products []Product

	for i := 0; i < 5; i++ {
		p := Product{Model: fmt.Sprintf("iPhone %s", xid.New().String())}
		products = append(products, p)
	}

	for i := 0; i < 5; i++ {
		p := Product{Model: fmt.Sprintf("Google Pixel %s", xid.New().String())}
		products = append(products, p)
	}

	return products
}
