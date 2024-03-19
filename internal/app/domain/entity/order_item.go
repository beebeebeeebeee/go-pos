package entity

type OrderItemID string

type OrderItem struct {
	ID         OrderItemID
	Item       Item
	Quantity   int
	TotalPrice float64
}
