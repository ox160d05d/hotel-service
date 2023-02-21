package reserve_room

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	reserveRoom "applicationDesignTest/internal/app/commands/reserve_room"
	"applicationDesignTest/internal/common"
	"applicationDesignTest/internal/domain/entity"
)

type ReserveRoomCommandHandler interface {
	Do(ctx context.Context, c reserveRoom.ReserveRoomCommand) (reserveRoom.ReserveRoomResult, error)
}

func NewReserveRoomHandler(reserveRoomCommandHandler ReserveRoomCommandHandler) http.HandlerFunc {
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

		// dates already validated in parseRequest
		from, _ := time.Parse(common.DateLayout, req.From)
		to, _ := time.Parse(common.DateLayout, req.To)

		result, err := reserveRoomCommandHandler.Do(r.Context(), reserveRoom.ReserveRoomCommand{
			HotelID: common.DefaultHotelID,
			UserID:  req.UserID,
			RoomID:  req.RoomID,
			From:    from,
			To:      to,
		})

		if err != nil {
			if errors.Is(entity.AlreadyReservedErr, err) {
				// сделано по образцу с исходным приложением; на реальных проектах нужно обсуждать с фронтендерами -
				// они могут попросить возвращать status 200 + {error_code: ...}, или что-то ещё (и это ок)
				http.Error(w, err.Error(), http.StatusConflict)
				return
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(ReserveRoomResponse{ReservationID: result.ReservationID})
	}
}
