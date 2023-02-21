package get_user_orders

import (
	"encoding/json"
	"fmt"
	"net/http"

	"applicationDesignTest/internal/common"
	"applicationDesignTest/internal/domain/entity"
)

type GetOrdersRequest struct {
	UserID string `json:"user_id,omitempty"`
}

type GetOrdersResponse struct {
	Orders []GetOrdersResponseOrder `json:"orders"`
}

type GetOrdersResponseOrder struct {
	ID   string `json:"id,omitempty"`
	Room string `json:"room,omitempty"`
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

func parseRequest(r *http.Request) (GetOrdersRequest, error) {
	var req GetOrdersRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return GetOrdersRequest{}, fmt.Errorf("invalid JSON: %w", err)
	}

	if req.UserID == "" {
		return GetOrdersRequest{}, fmt.Errorf("user ID can not be empty")
	}

	return req, nil
}

func composeResponse(orders []entity.Order, reservations []entity.RoomReservation) (GetOrdersResponse, error) {
	resp := GetOrdersResponse{
		Orders: make([]GetOrdersResponseOrder, 0, len(orders)),
	}

	reservationsMap := make(map[string]entity.RoomReservation)
	for _, v := range reservations {
		reservationsMap[v.ID()] = v
	}

	for _, v := range orders {
		r, ok := reservationsMap[v.ReservationID()]
		if !ok {
			err := fmt.Errorf("reservation `%s` not found fro order `%s`", v.ReservationID(), v.ID())
			return GetOrdersResponse{}, err
		}

		resp.Orders = append(resp.Orders, GetOrdersResponseOrder{
			ID:   v.ID(),
			Room: r.Room().ID(),
			From: r.From().Format(common.DateLayout),
			To:   r.To().Format(common.DateLayout),
		})
	}

	return resp, nil
}
