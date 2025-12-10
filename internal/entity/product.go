package entity

import (
	"github.com/rs/xid"
)

const (
	New         = "new"
	Used        = "used"
	Refurbished = "refurbished"
)

type Product struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Condition     string `json:"condition"`
	Price         Price  `json:"price"`
	OriginalPrice *Price `json:"original_price,omitempty"` // for discounts
	Stock         int    `json:"stock"`
	SoldQuantity  int    `json:"sold_quantity"`
}

func NewProduct(title, condition string, price Price, stock int, soldQuantity int) (*Product, error) {
	newId := xid.New().String()

	return &Product{
		ID:            newId,
		Title:         title,
		Condition:     condition,
		Price:         price,
		OriginalPrice: &price,
		Stock:         stock,
		SoldQuantity:  soldQuantity,
	}, nil
}

type Price struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type Image struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	SecureURL string `json:"secure_url"`
	IsMain    bool   `json:"is_main"`
}

type Attribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Seller struct {
	ID              string  `json:"id"`
	Nickname        string  `json:"nickname"`
	ReputationLevel string  `json:"reputation_level"`
	Transactions    int     `json:"transactions"`
	PositiveRating  float64 `json:"positive_rating"`
}

type ShippingInfo struct {
	FreeShipping bool             `json:"free_shipping"`
	Mode         string           `json:"mode"`
	Methods      []ShippingMethod `json:"methods"`
}

type ShippingMethod struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Cost          float64 `json:"cost"`
	EstimatedDays string  `json:"estimated_days"`
}

type Location struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type ReviewsSummary struct {
	Average      float64 `json:"average"`
	TotalReviews int     `json:"total_reviews"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
