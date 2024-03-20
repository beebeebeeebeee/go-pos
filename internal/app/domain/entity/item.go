package entity

type ItemID string

type Item struct {
	ID    ItemID  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func NewItem(
	id ItemID,
	name string,
	price float64,
) Item {

	return Item{
		ID:    id,
		Name:  name,
		Price: price,
	}
}
