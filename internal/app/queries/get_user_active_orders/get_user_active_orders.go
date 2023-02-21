package get_user_active_orders

import (
	"context"
	"fmt"

	"applicationDesignTest/internal/domain/entity"
)

type OrdersRepository interface {
	ByUserID(ctx context.Context, userID string, activeOnly bool) ([]entity.Order, error)
}

type GetUserOrdersHandler struct {
	orderRepository OrdersRepository
}

func NewGetUserOrdersHandler(orderRepository OrdersRepository) *GetUserOrdersHandler {
	return &GetUserOrdersHandler{
		orderRepository: orderRepository,
	}
}

func (h *GetUserOrdersHandler) Do(ctx context.Context, q GetUserActiveOrdersQuery) (GetUserActiveOrdersResponse, error) {
	orders, err := h.orderRepository.ByUserID(ctx, q.UserID, true)
	if err != nil {
		return GetUserActiveOrdersResponse{}, fmt.Errorf("error on orderRepository.ByUserID: %w", err)
	}

	return GetUserActiveOrdersResponse{
		Orders: orders,
	}, nil
}
