package search

import (
	"fmt"
	xid "github.com/rs/xid"
)

type Brand struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (b Brand) String() string {
	return fmt.Sprintf("\n\tBrand: %s, Description: %s", b.Name, b.Description)
}

func (b Brand) Type() string {
	return "brands"
}

func (b Brand) Id() string {
	return xid.New().String()
}

func SeedBrands() []Brand {
	var brands []Brand

	for i := 0; i < 5; i++ {
		b := Brand{Name: fmt.Sprintf("Apple %s", xid.New().String()), Description: "apppllllee"}
		brands = append(brands, b)
	}

	for i := 0; i < 5; i++ {
		b := Brand{Name: fmt.Sprintf("Google %s", xid.New().String()), Description: "gogogogogle"}
		brands = append(brands, b)
	}

	return brands
}
