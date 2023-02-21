package entity

type OrderStatus string

const OrderStatusNew OrderStatus = "new"
const OrderStatusConfirmed OrderStatus = "confiremed"
const OrderStatusCanceled OrderStatus = "cancelled"

func (s OrderStatus) IsActive() bool {
	if s == OrderStatusConfirmed || s == OrderStatusNew {
		return true
	}

	return false
}

func NewOrder(id string, userID string, reservationID string) *Order {
	return &Order{
		id:            id,
		userID:        userID,
		reservationID: reservationID,
		status:        OrderStatusNew,
	}
}

type Order struct {
	id     string
	status OrderStatus

	// В исходном примере был email, но пользователь может его поменять, нужен суррогатный ID
	userID string

	// атрибуты from\to\room хранятся в сущности reservations, пока что не нужна денормализация и дублирование их в order
	reservationID string
}

func (o *Order) ID() string {
	return o.id
}

func (o *Order) Status() OrderStatus {
	return o.status
}

func (o *Order) UserID() string {
	return o.userID
}

func (o *Order) ReservationID() string {
	return o.reservationID
}
