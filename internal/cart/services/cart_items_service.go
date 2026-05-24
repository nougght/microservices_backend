package services

import (
	"context"
	"store-server/internal/cart/models"
	"store-server/internal/cart/repositories"

	"github.com/google/uuid"
)

type CartItemsService struct {
	repo *repositories.CartItemsRepository
}

func NewCartItemsService(repo *repositories.CartItemsRepository) *CartItemsService {
	return &CartItemsService{repo: repo}
}

func (s *CartItemsService) AddToCart(ctx context.Context, item *models.CartItem) error {
	return s.repo.AddToCart(ctx, item)
}

func (s *CartItemsService) GetCartItemsByCartID(ctx context.Context, id uuid.UUID) ([]models.CartItem, error) {
	return s.repo.GetCartItemsByCartID(ctx, id)
}

func (s *CartItemsService) UpdateCartItemQuantity(ctx context.Context, id uuid.UUID, quantity int) error {
	return s.repo.UpdateCartItemQuantity(ctx, id, quantity)
}

func (s *CartItemsService) DeleteFromCartById(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteFromCartById(ctx, id)
}

func (s *CartItemsService) DeleteItemsByIDs(ctx context.Context, ids []uuid.UUID) error {
	return s.repo.DeleteItemsByIDs(ctx, ids)
}
