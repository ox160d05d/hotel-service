package get_user_orders

import (
	"context"
	"encoding/json"
	"net/http"

	getUserActiveOrders "applicationDesignTest/internal/app/queries/get_reservations"
	getUserOrders "applicationDesignTest/internal/app/queries/get_user_active_orders"
)

type GetUserOrdersQueryHandler interface {
	Do(ctx context.Context, q getUserOrders.GetUserActiveOrdersQuery) (getUserOrders.GetUserActiveOrdersResponse, error)
}

type GetReservationsQueryHandler interface {
	Do(
		ctx context.Context,
		q getUserActiveOrders.GetReservationsByIDsQuery,
	) (getUserActiveOrders.GetReservationsByIDsResponse, error)
}

func NewGetUserOrdersHandler(
	getUserOrdersQueryHandler GetUserOrdersQueryHandler,
	getReservationsQueryHandler GetReservationsQueryHandler,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		req, err := parseRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		ordersQuery := getUserOrders.GetUserActiveOrdersQuery{UserID: req.UserID}
		orders, err := getUserOrdersQueryHandler.Do(ctx, ordersQuery)
		if err != nil {
			http.Error(w, "something went wrong, try later", http.StatusInternalServerError)
			return
		}

		reservationsIDs := make([]string, 0, len(orders.Orders))
		for _, v := range orders.Orders {
			reservationsIDs = append(reservationsIDs, v.ReservationID())
		}

		reservationsQuery := getUserActiveOrders.GetReservationsByIDsQuery{ReservationIDs: reservationsIDs}
		reservations, err := getReservationsQueryHandler.Do(ctx, reservationsQuery)
		if err != nil {
			http.Error(w, "something went wrong, try later", http.StatusInternalServerError)
			return
		}

		resp, err := composeResponse(orders.Orders, reservations.Reservations)
		if err != nil {
			http.Error(w, "something went wrong, try later", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(resp)
	}
}
