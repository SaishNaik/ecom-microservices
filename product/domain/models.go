package domain

type Product struct {
	Sku   string  `json:"sku"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
}
