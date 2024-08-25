package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []Product

var productList = []Product{
	{
		ID:          1,
		Name:        "Hot Chocolate",
		Description: "Classic HotChocolate served with biscuits",
		Price:       17.50,
		SKU:         "chocolate",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	}, {
		ID:          2,
		Name:        "Tea",
		Description: "Classic Ginger Tea",
		Price:       10.50,
		SKU:         "tea",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
	{
		ID:          3,
		Name:        "Coffee",
		Description: "Classic Black Coffee",
		Price:       10.50,
		SKU:         "coffee",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
}

func GetProducts() Products {
	return productList
}

// converts product List to JSON
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}