package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Status      string  `json:"status"`
}

func GetProduct(productID int) (*Product, error) {
	url := fmt.Sprintf(
		"http://product-service:3002/products/%d",
		productID,
	)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var product Product

	err = json.NewDecoder(resp.Body).Decode(&product)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
