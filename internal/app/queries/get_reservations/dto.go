package get_reservations

import "applicationDesignTest/internal/domain/entity"

type GetReservationsByIDsQuery struct {
	ReservationIDs []string
}

type GetReservationsByIDsResponse struct {
	Reservations []entity.RoomReservation
}
