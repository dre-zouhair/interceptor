package service

import "dre-zouhair/modules/cart-api/internal/repository"

type cartService struct {
	cartRepository repository.ICartTRepository
}

func NewCartService(cartRepository repository.ICartTRepository) ICartService {
	return &cartService{
		cartRepository: cartRepository,
	}
}

type ICartService interface {
	Add(userID string, item repository.Item) error
	Get(userID string) ([]repository.Item, error)
}

func (s cartService) Add(userID string, item repository.Item) error {
	return s.cartRepository.Add(userID, item)
}

func (s cartService) Get(userID string) ([]repository.Item, error) {
	return s.cartRepository.Get(userID)
}
