package repository

import "errors"

type Item struct {
	ID string `json:"id" validate:"required,min=5,max=20"`
}

type cartRepository struct {
	carts map[string][]Item
}

func NewCartRepository() ICartTRepository {
	return &cartRepository{
		carts: make(map[string][]Item, 0),
	}
}

type ICartTRepository interface {
	Add(userID string, item Item) error
	Get(userID string) ([]Item, error)
}

func (r cartRepository) Add(userID string, item Item) error {
	if _, ok := r.carts[userID]; !ok {
		r.carts[userID] = make([]Item, 0)
	}
	r.carts[userID] = append(r.carts[userID], item)
	return nil
}

func (r cartRepository) Get(userID string) ([]Item, error) {
	items, ok := r.carts[userID]
	if ok {
		return items, nil
	}

	return nil, errors.New("not found")
}
