package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"applicationDesignTest/internal/common/di"
	log "applicationDesignTest/internal/common/logger"
	"applicationDesignTest/internal/presentation/http/server"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	logger := log.NewLog()
	app := di.NewApp(logger)

	logger.LogInfo("server started...")
	if err := server.Serve(ctx, app.Commands, app.Queries); err != nil {
		logger.LogErrorf("failed to serve: %s", err)
		os.Exit(1)
	}

	logger.LogInfo("bye...")
}
