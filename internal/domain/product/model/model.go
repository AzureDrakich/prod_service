package model

type Product struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Image_id      string `json:"image_id"`
	Price         string `json:"price"`
	Currency_id   int    `json:"currency_id"`
	Rating        string `json:"rating"`
	Category_id   string `json:"category_id"`
	Specification string `json:"specification"`
	Created_at    string `json:"created_at"`
	Updated_at    string `json:"updated_at"`
}
