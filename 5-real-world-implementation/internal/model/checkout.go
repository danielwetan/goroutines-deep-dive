package model

type CheckoutRequest struct {
	StoreName string   `json:"store_name"`
	Products  []string `json:"products"`
}

type CheckoutResponse struct {
	StoreName string `json:"store_name"`
	InvoiceID string `json:"invoice_id"`
}
