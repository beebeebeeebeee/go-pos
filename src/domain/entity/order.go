package entity

type OrderItem struct {
	Item       Item
	Quantity   int
	TotalPrice float64
}

type Order struct {
	OrderID    string
	OrderItems []OrderItem
	TotalPrice float64
}
