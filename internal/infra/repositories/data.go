package repositories

import (
	"sync"
	"time"

	"applicationDesignTest/internal/common"
	"applicationDesignTest/internal/domain/entity"
)

var roomsMutex sync.RWMutex
var ordersMutex sync.RWMutex

var orders map[string]entity.Order

// map[hotelID]map[roomId]{...}
var rooms map[string]map[string]entity.Room

func init() {
	orders = make(map[string]entity.Order, 0)
	rooms = make(map[string]map[string]entity.Room, 0)

	from := time.Date(2076, 11, 20, 15, 0, 0, 0, time.UTC)
	to := time.Date(2076, 11, 25, 15, 0, 0, 0, time.UTC)

	room1 := entity.NewRoom("room1", common.DefaultHotelID)
	_ = room1.AddReservation("reservation1", "user1", from, to)

	rooms[common.DefaultHotelID] = make(map[string]entity.Room)
	rooms[common.DefaultHotelID][room1.ID()] = *room1

	order1 := entity.NewOrder("order1", "user1", "reservation1")
	orders["order1"] = *order1
}
