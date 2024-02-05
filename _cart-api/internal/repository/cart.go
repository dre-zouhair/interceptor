package repository

type Item struct {
	ID string `json:"id" validate:"required,min=5,max=20"`
}

type CartRepository struct {
	carts map[string][]Item
}

func NewCartRepository() *CartRepository {
	return &CartRepository{
		carts: make(map[string][]Item, 0),
	}
}

type ICartTRepository interface {
	Add(userID string, item Item) error
}

func (r CartRepository) Add(userID string, item Item) error {
	if _, ok := r.carts[userID]; !ok {
		r.carts[userID] = make([]Item, 0)
	}
	r.carts[userID] = append(r.carts[userID], item)
	return nil
}
