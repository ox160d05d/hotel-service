package api

import (
	"net/http"

	"applicationDesignTest/internal/common/di"
	createOrder "applicationDesignTest/internal/presentation/http/api/create_order"
	getUserOrders "applicationDesignTest/internal/presentation/http/api/get_user_orders"
	reserveRoom "applicationDesignTest/internal/presentation/http/api/reserve_room"
)

type HttpHandlers struct {
	GetUserOrdersHandler http.HandlerFunc
	ReserveRoomHandler   http.HandlerFunc
	CreateOrderHandler   http.HandlerFunc
}

func NewHttpHandlers(commands *di.Commands, queries *di.Queries) *HttpHandlers {
	return &HttpHandlers{
		GetUserOrdersHandler: getUserOrders.NewGetUserOrdersHandler(
			queries.GetUserActiveOrders,
			queries.GetReservations,
		),
		ReserveRoomHandler: reserveRoom.NewReserveRoomHandler(commands.ReserveRoom),
		CreateOrderHandler: createOrder.NewReserveRoomHandler(commands.CreateOrder),
	}
}
