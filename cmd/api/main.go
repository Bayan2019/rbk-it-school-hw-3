package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/Bayan2019/rbk-it-school-hw-3/internal/client"
	"github.com/Bayan2019/rbk-it-school-hw-3/internal/config"
	"github.com/Bayan2019/rbk-it-school-hw-3/internal/repository/postgres"
	"github.com/Bayan2019/rbk-it-school-hw-3/internal/server"
	"github.com/Bayan2019/rbk-it-school-hw-3/internal/service"
)

func main() {
	cfg := config.MustLoad()

	db, err := postgres.NewDB(cfg.Database)
	if err != nil {
		log.Fatalf("connect database: %v", err)
	}
	defer db.Close()

	userRepo := postgres.NewUserRepository(db)
	cityRepo := postgres.NewCityRepository(db)

	userService := service.NewUserService(userRepo)
	cityService := service.NewCityService(cityRepo)

	osmClient := client.NewOsmClient(cfg.Api)

	handler := server.NewHandler(userService, cityService, osmClient)
	router := server.NewRouter(handler)

	srv := &http.Server{
		Addr:         ":" + cfg.App.Port,
		Handler:      router,
		ReadTimeout:  cfg.App.ReadTimeout,
		WriteTimeout: cfg.App.WriteTimeout,
		IdleTimeout:  cfg.App.IdleTimeout,
	}

	go func() {
		log.Printf("server started on :%s", cfg.App.Port)

		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen server: %v", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = srv.Shutdown(shutdownCtx)
	if err != nil {
		log.Printf("shutdown server: %v", err)
	}

	log.Println("server stopped")
}
