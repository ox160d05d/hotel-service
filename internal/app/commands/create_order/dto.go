package create_order

type CreateOrderCommand struct {
	UserID        string
	ReservationID string
}

type CreateOrderResult struct {
	OrderID string
}
