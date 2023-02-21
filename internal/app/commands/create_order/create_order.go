package create_order

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"applicationDesignTest/internal/domain/entity"
)

type BookingService interface {
	ConfirmReservation(ctx context.Context, reservationID string) error
}

type OrderRepository interface {
	Add(ctx context.Context, order entity.Order) error
}

type Databus interface {
	PublishOrderCreatedEvent(ctx context.Context, order entity.Order) error
}

type TxManager interface {
	WithTransaction(ctx context.Context, txFunc func(ctx context.Context) error) (err error)
}

type CreateOrderCommandHandler struct {
	orderRepository OrderRepository
	bookingService  BookingService
	databus         Databus
	txManager       TxManager
	idGenerator     func() string
}

func NewCreateOrderCommandHandler(
	orderRepository OrderRepository,
	bookingService BookingService,
	txManager TxManager,
	databus Databus,
) *CreateOrderCommandHandler {
	gen := func() string {
		// обычно uuid, но по условиям задачи используем только std библиотеку
		return strconv.FormatInt(time.Now().UnixNano(), 10)
	}

	return &CreateOrderCommandHandler{
		orderRepository: orderRepository,
		databus:         databus,
		bookingService:  bookingService,
		txManager:       txManager,
		idGenerator:     gen,
	}
}

func (h *CreateOrderCommandHandler) Do(ctx context.Context, c CreateOrderCommand) (CreateOrderResult, error) {
	newOrderID := h.idGenerator()
	newOrder := entity.NewOrder(newOrderID, c.UserID, c.ReservationID)

	err := h.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		err := h.orderRepository.Add(ctx, *newOrder)
		if err != nil {
			return fmt.Errorf("error on orderRepository.Add: %w", err)
		}

		if err = h.bookingService.ConfirmReservation(ctx, c.ReservationID); err != nil {
			return fmt.Errorf("error on bookingService.ConfirmReservation: %w", err)
		}

		// По эвенту окончательно подверждаем резервацию, отправляем email, выставляем счёт и т.п.
		// В идеале - записываем event в БД, потом в фоне пушим в databus (transactional outbox pattern)
		if err := h.databus.PublishOrderCreatedEvent(ctx, *newOrder); err != nil {
			return fmt.Errorf("error on databus.PublishOrderCreatedEvent: %w", err)
		}

		return nil
	})

	if err != nil {
		return CreateOrderResult{}, err
	}

	return CreateOrderResult{OrderID: newOrderID}, nil
}
