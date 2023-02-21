package di

import (
	createOrder "applicationDesignTest/internal/app/commands/create_order"
	reserveRoom "applicationDesignTest/internal/app/commands/reserve_room"
	getReservations "applicationDesignTest/internal/app/queries/get_reservations"
	getUserActiveOrders "applicationDesignTest/internal/app/queries/get_user_active_orders"
	reserve_room "applicationDesignTest/internal/app/service/booking"
	log "applicationDesignTest/internal/common/logger"
	transactionManager "applicationDesignTest/internal/common/transaction_manager"
	"applicationDesignTest/internal/infra/databus"
	"applicationDesignTest/internal/infra/repositories"
)

type App struct {
	Log      *log.Log
	Commands *Commands
	Queries  *Queries
}

type Commands struct {
	ReserveRoom *reserveRoom.ReserveRoomCommandHandler
	CreateOrder *createOrder.CreateOrderCommandHandler
}

type Queries struct {
	GetUserActiveOrders *getUserActiveOrders.GetUserOrdersHandler
	GetReservations     *getReservations.GetReservationsByIDsHandler
}

func NewApp(log *log.Log) *App {
	app := &App{
		Log: log,
	}

	orderRepository := repositories.NewOrderRepository()
	hotelRoomRepository := repositories.NewHotelRoomRepository()
	trxManager := transactionManager.NewTrxManager()
	dbus := databus.NewDatabus()

	bookingService := reserve_room.NewBookingService(hotelRoomRepository, trxManager)

	app.Commands = &Commands{
		ReserveRoom: reserveRoom.NewReserveRoomCommandHandler(bookingService),
		CreateOrder: createOrder.NewCreateOrderCommandHandler(orderRepository, bookingService, trxManager, dbus),
	}

	app.Queries = &Queries{
		GetUserActiveOrders: getUserActiveOrders.NewGetUserOrdersHandler(orderRepository),
		GetReservations:     getReservations.NewGetReservationsByIDsHandler(hotelRoomRepository),
	}

	return app
}
