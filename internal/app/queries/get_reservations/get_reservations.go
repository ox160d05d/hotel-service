package get_reservations

import (
	"context"
	"fmt"

	"applicationDesignTest/internal/domain/entity"
)

type ReservationsRepository interface {
	GetReservationsByIDs(ctx context.Context, reservationsIDs []string) ([]entity.RoomReservation, error)
}

type GetReservationsByIDsHandler struct {
	reservationRepository ReservationsRepository
}

func NewGetReservationsByIDsHandler(reservationsRepository ReservationsRepository) *GetReservationsByIDsHandler {
	return &GetReservationsByIDsHandler{
		reservationRepository: reservationsRepository,
	}
}

func (h *GetReservationsByIDsHandler) Do(
	ctx context.Context,
	q GetReservationsByIDsQuery,
) (GetReservationsByIDsResponse, error) {
	if len(q.ReservationIDs) == 0 {
		return GetReservationsByIDsResponse{}, nil
	}

	reservations, err := h.reservationRepository.GetReservationsByIDs(ctx, q.ReservationIDs)
	if err != nil {
		return GetReservationsByIDsResponse{}, fmt.Errorf("error on reservationRepository.GetReservationsByIDs: %w", err)
	}

	return GetReservationsByIDsResponse{
		Reservations: reservations,
	}, nil
}
