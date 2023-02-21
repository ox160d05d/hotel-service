package reserve_room

import "time"

type ReserveRoomCommand struct {
	HotelID string
	UserID  string
	RoomID  string
	From    time.Time
	To      time.Time
}

type ReserveRoomResult struct {
	ReservationID string
}
