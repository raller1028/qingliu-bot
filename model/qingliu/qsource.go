package qingliu

type QSource struct {
	Applicant string    `json:"applicant"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Country   string    `json:"country"`
	Province  string    `json:"province"`
	City      string    `json:"city"`
	ZipCode   string    `json:"zip_code"`
	Address   string    `json:"address"`
	Product   []Product `json:"product"`
	Order     string    `json:"order"`
	Server    string    `json:"server"`
	Company   string    `json:"company"`
	Trial     float64   `json:"trial"`
	Weight    float64   `json:"weight"`
}
type Product struct {
	Name     string  `json:"name"`
	SKU      string  `json:"sku"`
	Price    float64 `json:"price"`
	Num      int     `json:"num"`
	From     string  `json:"from"`
	Currency string  `json:"currency"`
}
