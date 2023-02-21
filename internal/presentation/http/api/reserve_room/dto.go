package reserve_room

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"applicationDesignTest/internal/common"
)

type ReserveRoomRequest struct {
	UserID string `json:"user_id,omitempty"`
	RoomID string `json:"room_id,omitempty"`
	From   string `json:"from,omitempty"`
	To     string `json:"to,omitempty"`
}

type ReserveRoomResponse struct {
	ReservationID string `json:"reservation_id"`
}

func parseRequest(r *http.Request) (ReserveRoomRequest, error) {
	var req ReserveRoomRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return ReserveRoomRequest{}, fmt.Errorf("invalid JSON: %w", err)
	}

	if req.UserID == "" {
		return ReserveRoomRequest{}, fmt.Errorf("user ID can not be empty")
	}

	if req.RoomID == "" {
		return ReserveRoomRequest{}, fmt.Errorf("room ID can not be empty")
	}

	var (
		from, to time.Time
		err      error
	)

	if _, err = time.Parse(common.DateLayout, req.From); err != nil {
		return ReserveRoomRequest{}, fmt.Errorf("invalid from date `%s`: %w", req.From, err)
	}

	if _, err = time.Parse(common.DateLayout, req.To); err != nil {
		return ReserveRoomRequest{}, fmt.Errorf("invalid to date `%s`: %w", req.To, err)
	}

	if from.Before(time.Now()) {
		return ReserveRoomRequest{}, fmt.Errorf("`from` date must be greater than now")
	}

	if from.After(to) || from.Equal(to) {
		return ReserveRoomRequest{}, fmt.Errorf("`from` date must be after `to` date")
	}

	return req, nil
}
