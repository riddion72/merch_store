package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"avito-shop-service/internal/config"
	"avito-shop-service/internal/postgress"
	"avito-shop-service/internal/token"
	"avito-shop-service/internal/transport/server"
	"avito-shop-service/pkg/logger"
)

const (
	shutdownTimeout = 5 * time.Second
)

func main() {
	configPath := flag.String("config", "config/config.yaml", "config file path")
	flag.Parse()

	cfg := config.ParseConfig(*configPath)

	logger.MustInitlogger(cfg.Logger.Level)

	db, err := postgress.NewConnection(context.Background(), cfg.DB)
	if err != nil {
		logger.Error("error connecting to database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	repo := postgress.New(db)

	jwtToken := token.New(cfg.Secret)

	//

	serv := server.New(hnd.Route(), cfg.Server)

	logger.Info("starting server", slog.String("address", cfg.Server.Address))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit

	logger.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err = serv.Shutdown(ctx); err != nil {
		logger.Error("shutdown server error", slog.String("error", err.Error()))
	}

}
