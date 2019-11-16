package product

type Product struct {
	ID           int64     `json:"id"`
	PriceInCents int64     `json:"price_in_cents"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Discount     *Discount `json:"discount"`
}

type Discount struct {
	Percent      float32 `json:"pct"`
	ValueInCents int64   `json:"value_in_cents"`
}
