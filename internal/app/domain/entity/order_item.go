package entity

type OrderItemID string

type OrderItem struct {
	ID         OrderItemID `json:"id"`
	Item       Item        `json:"item"`
	Quantity   int         `json:"quantity"`
	TotalPrice float64     `json:"totalPrice"`
}
