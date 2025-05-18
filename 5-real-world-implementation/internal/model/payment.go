package model

type PaymentDetail struct {
	Product Product `json:"product"`
	Payment Payment `json:"payment"`
}

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Payment struct {
	Status string `json:"status"`
}
