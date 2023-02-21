package create_order

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateOrderRequest struct {
	UserID        string `json:"user_id,omitempty"`
	ReservationID string `json:"reservation_id"`
}

type CreateOrderResponse struct {
	OrderID string `json:"order_id"`
}

func parseRequest(r *http.Request) (CreateOrderRequest, error) {
	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return CreateOrderRequest{}, fmt.Errorf("invalid JSON: %w", err)
	}

	if req.UserID == "" {
		return CreateOrderRequest{}, fmt.Errorf("user ID can not be empty")
	}

	if req.ReservationID == "" {
		return CreateOrderRequest{}, fmt.Errorf("reservation ID can not be empty")
	}

	return req, nil
}
