package get_user_active_orders

import "applicationDesignTest/internal/domain/entity"

type GetUserActiveOrdersQuery struct {
	UserID string
}

type GetUserActiveOrdersResponse struct {
	Orders []entity.Order
}
