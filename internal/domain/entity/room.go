package entity

import (
	"fmt"
	"time"
)

var AlreadyReservedErr = fmt.Errorf("room already reserved")
var AlreadyExistsErr = fmt.Errorf("reservation already exists")

type ReservationStatus string

const ReservationStatusNew ReservationStatus = "new"
const ReservationStatusConfirmed ReservationStatus = "confirmed"
const ReservationStatusCancelled ReservationStatus = "cancelled"

type Room struct {
	id           string
	hotelID      string
	reservations []RoomReservation
}

func NewRoom(
	id string,
	hotelID string,
) *Room {
	return &Room{
		id:           id,
		hotelID:      hotelID,
		reservations: make([]RoomReservation, 0),
	}
}

func (r *Room) Reservations() []RoomReservation {
	return r.reservations
}

func (r *Room) ID() string {
	return r.id
}

func (r *Room) HotelID() string {
	return r.hotelID
}

func (r *Room) AddReservation(id string, userID string, from time.Time, to time.Time) error {
	if !from.Before(to) {
		return fmt.Errorf("`from` date must be earlier than `to` date")
	}

	for _, existingReservation := range r.reservations {
		if existingReservation.status == ReservationStatusCancelled {
			continue
		}

		if id == existingReservation.id {
			return AlreadyExistsErr
		}

		if (!from.Before(existingReservation.from) && !from.After(existingReservation.to)) ||
			(!to.Before(existingReservation.from) && !to.After(existingReservation.to)) {
			return AlreadyReservedErr
		}

		if (!existingReservation.from.Before(from) && !existingReservation.from.After(from)) ||
			(!existingReservation.to.Before(from) && !existingReservation.to.After(to)) {
			return AlreadyReservedErr
		}
	}

	reservation := RoomReservation{
		id:     id,
		userID: userID,
		room:   r,
		status: ReservationStatusNew,
		from:   from,
		to:     to,
	}
	r.reservations = append(r.reservations, reservation)

	return nil
}

type RoomReservation struct {
	id     string
	userID string
	room   *Room
	status ReservationStatus
	from   time.Time
	to     time.Time
}

func (r *RoomReservation) SetStatus(status ReservationStatus) {
	r.status = status
}

func (r *RoomReservation) Room() *Room {
	return r.room
}

func (r *RoomReservation) ID() string {
	return r.id
}

func (r *RoomReservation) UserID() string {
	return r.userID
}

func (r *RoomReservation) Status() ReservationStatus {
	return r.status
}

func (r *RoomReservation) From() time.Time {
	return r.from
}

func (r *RoomReservation) To() time.Time {
	return r.to
}
