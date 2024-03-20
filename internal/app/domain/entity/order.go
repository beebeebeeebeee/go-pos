package entity

type OrderID string

type Order struct {
	ID         OrderID     `json:"id"`
	OrderItems []OrderItem `json:"orderItems"`
	TotalPrice float64     `json:"totalPrice"`
}
