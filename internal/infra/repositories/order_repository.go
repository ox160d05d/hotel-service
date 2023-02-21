package repositories

import (
	"context"
	"fmt"

	"applicationDesignTest/internal/domain/entity"
)

type OrderRepository struct {
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

func (r *OrderRepository) Add(_ context.Context, order entity.Order) error {
	ordersMutex.Lock()
	defer ordersMutex.Unlock()

	if _, ok := orders[order.ID()]; !ok {
		orders[order.ID()] = order
		return nil
	}

	return fmt.Errorf("order with ID `%s` already exists", order.ID())
}

func (r *OrderRepository) ByUserID(_ context.Context, userID string, activeOnly bool) ([]entity.Order, error) {
	ordersMutex.RLock()
	defer ordersMutex.RUnlock()

	var res []entity.Order
	for _, v := range orders {
		if v.UserID() != userID {
			continue
		}

		if !v.Status().IsActive() {
			continue
		}

		res = append(res, v)
	}

	return res, nil
}
