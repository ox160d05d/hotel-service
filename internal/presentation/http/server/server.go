package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"applicationDesignTest/internal/common/di"
	"applicationDesignTest/internal/presentation/http/api"
)

const gracefulStopTimeout = time.Second * 5

func newMux(commands *di.Commands, queries *di.Queries) *http.ServeMux {
	h := api.NewHttpHandlers(commands, queries)

	mux := http.NewServeMux()
	mux.HandleFunc("/order", h.CreateOrderHandler)
	mux.HandleFunc("/reservation", h.ReserveRoomHandler)
	mux.HandleFunc("/orders", h.GetUserOrdersHandler)

	return mux
}

func Serve(ctx context.Context, commands *di.Commands, queries *di.Queries) error {
	mux := newMux(commands, queries)

	httpPort := "8080"
	if p := os.Getenv("HTTP_PORT"); p != "" {
		httpPort = p
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%s", httpPort),
		Handler: mux,
	}

	errChan := make(chan error, 2)
	defer close(errChan)

	ctx, cancelFn := context.WithCancel(ctx)
	defer cancelFn()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		<-ctx.Done()
		ctx, cancelFn := context.WithTimeout(ctx, gracefulStopTimeout)
		defer cancelFn()

		errChan <- httpServer.Shutdown(ctx)
	}()

	go func() {
		defer wg.Done()
		defer cancelFn()

		if err := httpServer.ListenAndServe(); err != nil {
			if errors.Is(http.ErrServerClosed, err) {
				errChan <- nil
				return
			}

			errChan <- err
		}
	}()

	err := <-errChan
	wg.Wait()

	return err
}
