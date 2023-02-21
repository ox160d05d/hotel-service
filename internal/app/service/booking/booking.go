package reserve_room

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"applicationDesignTest/internal/domain/entity"
)

type HotelRepository interface {
	GetHotelRoom(ctx context.Context, hotelID string, roomID string, forUpdate bool) (*entity.Room, error)
	Save(ctx context.Context, room entity.Room) error
	ConfirmReservation(ctx context.Context, reservationID string) error
}

type TxManager interface {
	WithTransaction(ctx context.Context, txFunc func(ctx context.Context) error) (err error)
}

type Booking struct {
	hotelRepository HotelRepository
	txManager       TxManager
	idGenerator     func() string
}

func NewBookingService(
	hotelRepository HotelRepository,
	txManager TxManager,
) *Booking {
	gen := func() string {
		// обычно uuid, но по условиям задачи используем только std библиотеку
		return strconv.FormatInt(time.Now().UnixNano(), 10)
	}

	return &Booking{
		hotelRepository: hotelRepository,
		txManager:       txManager,
		idGenerator:     gen,
	}
}

func (h *Booking) ConfirmReservation(ctx context.Context, reservationID string) error {
	if err := h.hotelRepository.ConfirmReservation(ctx, reservationID); err != nil {
		return fmt.Errorf("error on hotelRepository.ConfirmReservation: %w", err)
	}

	return nil
}

func (h *Booking) Reserve(
	ctx context.Context,
	hotelID string,
	userID string,
	roomID string,
	from time.Time,
	to time.Time,
) (reservationID string, err error) {
	newReservationID := h.idGenerator()

	err = h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		room, err := h.hotelRepository.GetHotelRoom(ctx, hotelID, roomID, true)
		if err != nil {
			return fmt.Errorf("error on hotelRepository.GetHotelRoom: %w", err)
		} else if room == nil {
			return fmt.Errorf("room `%s` not found", roomID)
		}

		if err := room.AddReservation(newReservationID, userID, from, to); err != nil {
			return fmt.Errorf("error on room.AddReservation: %w", err)
		}

		if err := h.hotelRepository.Save(ctx, *room); err != nil {
			return fmt.Errorf("error on hotelRepository.Save: %w", err)
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return newReservationID, nil
}
