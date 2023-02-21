package repositories

import (
	"context"
	"fmt"
	"sync"

	"applicationDesignTest/internal/domain/entity"
)

type HotelRoomRepository struct {
	sync.Mutex
}

func NewHotelRoomRepository() *HotelRoomRepository {
	return &HotelRoomRepository{}
}

func (r *HotelRoomRepository) GetHotelRoom(
	_ context.Context,
	hotelID string,
	roomID string,
	_ bool,
) (*entity.Room, error) {
	roomsMutex.RLock()
	defer roomsMutex.RUnlock()

	if _, ok := rooms[hotelID]; !ok {
		return nil, fmt.Errorf("hotel `%s` not found", hotelID)
	}

	if r, ok := rooms[hotelID][roomID]; !ok {
		return nil, nil
	} else {
		copyVal := r
		return &copyVal, nil
	}
}

func (r *HotelRoomRepository) Save(_ context.Context, room entity.Room) error {
	roomsMutex.Lock()
	defer roomsMutex.Unlock()

	if _, ok := rooms[room.HotelID()]; !ok {
		rooms[room.HotelID()] = make(map[string]entity.Room)
	}

	rooms[room.HotelID()][room.ID()] = room

	return nil
}

func (r *HotelRoomRepository) GetReservationsByIDs(
	_ context.Context,
	reservationsIDs []string,
) ([]entity.RoomReservation, error) {
	roomsMutex.RLock()
	defer roomsMutex.RUnlock()

	var res []entity.RoomReservation
	for _, id := range reservationsIDs {
		for _, hotelRooms := range rooms {
			for _, room := range hotelRooms {
				for _, reservation := range room.Reservations() {
					if reservation.ID() == id {
						res = append(res, reservation)
					}
				}
			}
		}
	}

	return res, nil
}

func (r *HotelRoomRepository) ConfirmReservation(_ context.Context, reservationID string) error {
	roomsMutex.Lock()
	defer roomsMutex.Unlock()

	for _, hotelRooms := range rooms {
		for _, room := range hotelRooms {
			for i, reservation := range room.Reservations() {
				if reservation.ID() == reservationID {
					w := &reservation
					w.SetStatus(entity.ReservationStatusConfirmed)
					room.Reservations()[i] = *w
					return nil
				}
			}
		}
	}

	return fmt.Errorf("reservation `%s` not found", reservationID)
}
