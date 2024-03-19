package entity

type ItemID string

type Item struct {
	ID       ItemID
	Name     string
	Quantity int
	Price    float64
}
