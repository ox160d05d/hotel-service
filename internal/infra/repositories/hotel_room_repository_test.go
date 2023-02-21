package repositories

import (
	"context"
	"fmt"
	"testing"
	"time"

	"applicationDesignTest/internal/common"
	"applicationDesignTest/internal/domain/entity"
)

func TestHotelRoomRepository_GetHotelRoom(t *testing.T) {
	testCases := []struct {
		name          string
		inHotelID     string
		inRoomID      string
		inForUpdate   bool
		expectedOut   *entity.Room
		expectedError error
	}{
		{
			name:          "ok, room found",
			inHotelID:     common.DefaultHotelID,
			inRoomID:      "room1",
			inForUpdate:   false,
			expectedOut:   entity.NewRoom("room1", common.DefaultHotelID),
			expectedError: nil,
		},
		{
			name:          "room not found",
			inHotelID:     common.DefaultHotelID,
			inRoomID:      "room123456789",
			inForUpdate:   false,
			expectedOut:   nil,
			expectedError: nil,
		},
		{
			name:          "hotel not found, error",
			inHotelID:     "hotel123456789",
			inRoomID:      "room1",
			inForUpdate:   false,
			expectedOut:   nil,
			expectedError: fmt.Errorf("hotel `hotel123456789` not found"),
		},
	}

	r := NewHotelRoomRepository()

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			room, err := r.GetHotelRoom(context.Background(), tt.inHotelID, tt.inRoomID, tt.inForUpdate)

			if tt.expectedError != nil {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}

				if tt.expectedError.Error() != err.Error() {
					t.Fatalf("expected error `%s`, got `%s`", tt.expectedError.Error(), err.Error())
				}
			} else if err != nil {
				t.Fatalf("expected no error, got `%s`", err.Error())
			}

			if tt.expectedOut == nil {
				if room != nil {
					t.Fatalf("expected nil, got room `%s`", room.ID())
				}
			} else {
				if room == nil {
					t.Fatalf("expected room `%s`, got room nil", tt.expectedOut.ID())
				}

				if tt.expectedOut.ID() != room.ID() {
					t.Fatalf("expected room `%s`, got room `%s`", tt.expectedOut.ID(), room.ID())
				}
			}
		})
	}
}

func TestHotelRepository_Save(t *testing.T) {
	r := NewHotelRoomRepository()
	ctx := context.Background()
	newRoomID := "new-room"

	newRoom := entity.NewRoom(newRoomID, common.DefaultHotelID)

	err := r.Save(ctx, *newRoom)
	if err != nil {
		t.Fatalf("expected no error, got `%s`", err)
	}

	room, err := r.GetHotelRoom(ctx, common.DefaultHotelID, newRoomID, false)
	if err != nil {
		t.Fatalf("expected no error, got `%s`", err)
	} else if room == nil {
		t.Fatalf("expected not nil")
	} else if room.ID() != newRoomID {
		t.Fatalf("expected `%s`, got `%s`", newRoomID, room.ID())
	}

	newReservationID := "new-reservation"
	userID := "user123456"
	from := time.Date(2058, 1, 1, 12, 0, 0, 0, time.UTC)
	to := time.Date(2058, 1, 2, 12, 0, 0, 0, time.UTC)

	_ = newRoom.AddReservation(newReservationID, userID, from, to)
	_ = r.Save(ctx, *newRoom)

	reservations, err := r.GetReservationsByIDs(ctx, []string{newReservationID})
	if err != nil {
		t.Fatalf("expected no error, got `%s`", err)
	} else if len(reservations) != 1 {
		t.Fatalf("expected 1 found reservation, got %d", len(reservations))
	} else if reservations[0].ID() != newReservationID {
		t.Fatalf("expected reservation `%s`, got `%s`", newReservationID, reservations[0].ID())
	}
}

func TestHotelRoomRepository_ConfirmReservation(t *testing.T) {
	r := NewHotelRoomRepository()
	ctx := context.Background()
	newRoomID := "new-room"

	newRoom := entity.NewRoom(newRoomID, common.DefaultHotelID)
	_ = r.Save(ctx, *newRoom)

	newReservationID := "new-reservation"
	userID := "user123456"
	from := time.Date(2058, 1, 1, 12, 0, 0, 0, time.UTC)
	to := time.Date(2058, 1, 2, 12, 0, 0, 0, time.UTC)

	_ = newRoom.AddReservation(newReservationID, userID, from, to)
	_ = r.Save(ctx, *newRoom)

	err := r.ConfirmReservation(ctx, newReservationID)
	if err != nil {
		t.Fatalf("expected no erorr, got `%s`", err)
	}

	reservations, err := r.GetReservationsByIDs(ctx, []string{newReservationID})
	if err != nil {
		t.Fatalf("expected no erorr, got `%s`", err)
	} else if len(reservations) != 1 {
		t.Fatalf("expected %d reservations, got %d", 1, len(reservations))
	} else if reservations[0].Status() != entity.ReservationStatusConfirmed {
		t.Fatalf("expected status `%s`, got `%s`", entity.ReservationStatusConfirmed, reservations[0].Status())
	}

}
