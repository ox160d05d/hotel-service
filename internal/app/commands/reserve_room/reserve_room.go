package reserve_room

import (
	"context"
	"fmt"
	"time"

	"applicationDesignTest/internal/domain/entity"
)

type BookingService interface {
	Reserve(
		ctx context.Context,
		hotelID string,
		userID string,
		roomID string,
		from time.Time,
		to time.Time,
	) (reservationID string, err error)
}

type HotelRepository interface {
	GetHotelRoom(ctx context.Context, hotelID string, roomID string, forUpdate bool) (*entity.Room, error)
	Save(ctx context.Context, room entity.Room) error
}

type TxManager interface {
	WithTransaction(ctx context.Context, txFunc func(ctx context.Context) error) (err error)
}

type ReserveRoomCommandHandler struct {
	bookingService BookingService
}

func NewReserveRoomCommandHandler(
	bookingService BookingService,
) *ReserveRoomCommandHandler {
	return &ReserveRoomCommandHandler{
		bookingService: bookingService,
	}
}

func (h *ReserveRoomCommandHandler) Do(ctx context.Context, c ReserveRoomCommand) (ReserveRoomResult, error) {
	reservationID, err := h.bookingService.Reserve(ctx, c.HotelID, c.UserID, c.RoomID, c.From, c.To)
	if err != nil {
		return ReserveRoomResult{}, fmt.Errorf("errro on bookingService.Reserve: %w", err)
	}

	return ReserveRoomResult{ReservationID: reservationID}, nil
}
