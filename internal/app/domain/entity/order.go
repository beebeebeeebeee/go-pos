package entity

type OrderID string

type Order struct {
	ID         OrderID
	OrderItems []OrderItem
	TotalPrice float64
}
