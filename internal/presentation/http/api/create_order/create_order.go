package create_order

import (
	"context"
	"encoding/json"
	"net/http"

	createOrder "applicationDesignTest/internal/app/commands/create_order"
)

type CreateOrderCommandHandler interface {
	Do(ctx context.Context, c createOrder.CreateOrderCommand) (createOrder.CreateOrderResult, error)
}

func NewReserveRoomHandler(createOrderCommandHandler CreateOrderCommandHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		req, err := parseRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := createOrderCommandHandler.Do(r.Context(), createOrder.CreateOrderCommand{
			UserID:        req.UserID,
			ReservationID: req.ReservationID,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(CreateOrderResponse{OrderID: result.OrderID})
	}
}
