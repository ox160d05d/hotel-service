package entity

import (
	"testing"
	"time"
)

func TestRoom_AddReservation(t *testing.T) {
	userID := "user1"
	hotelID := "hotel1"
	roomID := "room1"
	reservationID := "reservation1"
	from := time.Date(2075, 12, 17, 10, 0, 0, 0, time.UTC)
	to := time.Date(2075, 12, 19, 10, 0, 0, 0, time.UTC)

	room := NewRoom(roomID, hotelID)
	_ = room.AddReservation(reservationID, userID, from, to)

	err := room.AddReservation(reservationID, userID, from, to)
	if err == nil || err != AlreadyExistsErr {
		t.Fatalf("expected AlreadyExistsErr error, got %v", err)
	}

	// same [start|end] dates, ok
	newReservationID := "new-reservation"
	err = room.AddReservation(newReservationID, userID, from, to)
	if err == nil || err != AlreadyReservedErr {
		t.Fatalf("expected AlreadyReservedErr error, got %v", err)
	}

	// [start1|start2 - end1 - end2] case, error
	from = time.Date(2075, 12, 17, 10, 0, 0, 0, time.UTC)
	to = time.Date(2075, 12, 25, 10, 0, 0, 0, time.UTC)
	err = room.AddReservation(newReservationID, userID, from, to)
	if err == nil || err != AlreadyReservedErr {
		t.Fatalf("expected AlreadyReservedErr error, got %v", err)
	}

	// invalid dates
	from = time.Date(2140, 1, 1, 0, 0, 0, 0, time.UTC)
	to = time.Date(2130, 1, 1, 0, 0, 0, 0, time.UTC)
	err = room.AddReservation(newReservationID, userID, from, to)
	expectedErrStr := "`from` date must be earlier than `to` date"
	if err == nil || err.Error() != expectedErrStr {
		t.Fatalf("expected `%s` error, got %v", expectedErrStr, err)
	}

	// [start1 - start2 - end1 - end2] case, error
	from = time.Date(2075, 12, 18, 10, 0, 0, 0, time.UTC)
	to = time.Date(2075, 12, 25, 10, 0, 0, 0, time.UTC)
	err = room.AddReservation(newReservationID, userID, from, to)
	if err == nil || err != AlreadyReservedErr {
		t.Fatalf("expected AlreadyReservedErr error, got %v", err)
	}

	// [start2 - start1 - end2 - end1] case, error
	from = time.Date(2075, 12, 15, 10, 0, 0, 0, time.UTC)
	to = time.Date(2075, 12, 18, 10, 0, 0, 0, time.UTC)
	err = room.AddReservation(newReservationID, userID, from, to)
	if err == nil || err != AlreadyReservedErr {
		t.Fatalf("expected AlreadyReservedErr error, got %v", err)
	}

	// [start2 - start1 - end1 - end2] case, error
	from = time.Date(2075, 12, 10, 10, 0, 0, 0, time.UTC)
	to = time.Date(2075, 12, 29, 10, 0, 0, 0, time.UTC)
	err = room.AddReservation(newReservationID, userID, from, to)
	if err == nil || err != AlreadyReservedErr {
		t.Fatalf("expected AlreadyReservedErr error, got %v", err)
	}

	// [start1 - start2 - end2 - end1] case, error
	from = time.Date(2075, 12, 18, 10, 0, 0, 0, time.UTC)
	to = time.Date(2075, 12, 18, 19, 0, 0, 0, time.UTC)
	err = room.AddReservation(newReservationID, userID, from, to)
	if err == nil || err != AlreadyReservedErr {
		t.Fatalf("expected AlreadyReservedErr error, got %v", err)
	}

	// [start1 - end1, start2 - end2] case, ok
	from = time.Date(2075, 12, 20, 10, 0, 0, 0, time.UTC)
	to = time.Date(2075, 12, 21, 10, 0, 0, 0, time.UTC)
	err = room.AddReservation(newReservationID, userID, from, to)
	if err != nil {
		t.Fatalf("expected no error, got `%s`", err)
	}

	if len(room.Reservations()) != 2 {
		t.Fatalf("expected 2 reservations, got `%d`", len(room.Reservations()))
	}
}
