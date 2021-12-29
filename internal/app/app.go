package app

import (
	"articles/internal/config"
	handler "articles/internal/handlers"
	"articles/internal/repository"
	"articles/internal/server"
	"articles/internal/service"
	"articles/pkg/db"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	cfg := config.Init()

	mongoClient, err := db.NewClient(cfg.Mongo.URI, cfg.Mongo.User, cfg.Mongo.Password)
	if err != nil {
		log.Println(err)
	}

	db := mongoClient.Database(cfg.Mongo.DbName)

	repos := repository.NewRepositories(db)
	services := service.NewServices(service.Deps{
		Repos: repos,
	})
	handlers := handler.NewHandler(services)

	srv := server.NewServer(cfg, handlers.Init(cfg))

	go func() {
		srv.Run()
	}()

	log.Println("server successfully started!")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		log.Println("failed to stop server!")
	}

	if err := mongoClient.Disconnect(context.Background()); err != nil {
		log.Println("failed to disconnect db")
	}
}
