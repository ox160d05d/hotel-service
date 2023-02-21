package databus

import (
	"context"

	"applicationDesignTest/internal/domain/entity"
)

type Databus struct {
}

func NewDatabus() *Databus {
	return &Databus{}
}

func (bus *Databus) PublishOrderCreatedEvent(ctx context.Context, order entity.Order) error {
	// nothing for now

	return nil
}
