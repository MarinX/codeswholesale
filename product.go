// product
package codeswholesale

import (
	"time"
)

type Item struct {
	Items []Product `json:"items"`
}

type Product struct {
	ProductID   string     `json:"productId"`
	Identifier  string     `json:"identifier"`
	Name        string     `json:"name"`
	Platform    string     `json:"platform"`
	Quantity    uint       `json:"quantity"`
	Images      []Image    `json:"images"`
	Regions     []string   `json:"regions"`
	Languages   []string   `json:"languages"`
	Prices      []Price    `json:"prices"`
	Links       []Link     `json:"links"`
	ReleaseDate *time.Time `json:"releaseDate"`
}
