package service

import "github.com/dre-zouhair/modules/cart-api/internal/repository"

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
}

func (s cartService) Add(userID string, item repository.Item) error {
	return s.cartRepository.Add(userID, item)
}
