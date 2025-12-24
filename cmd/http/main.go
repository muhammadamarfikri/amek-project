package main

import (
	"context"
	"fmt"
	"hotel/config"
	"hotel/database"
	"hotel/delivery/rest/handler"
	"hotel/impl/repository"
	"hotel/router"
	"hotel/user"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {

	appConfig := config.AppServer()
	postgresConfig := config.GetPostgresConfig()
	postgresSync := config.GetPostgresSync()
	log := config.NewLogger()
	e := echo.New()

	fmt.Println("config", postgresConfig)

	postgresDb := database.InitPostgresDB(postgresConfig, postgresSync)

	repository := repository.NewRepository(postgresDb, log)

	userUsecase := user.NewUserUsecase(repository, log)

	userHandler := handler.NewUserHandler(userUsecase, log)

	router.UserRouter(e, userHandler)

	e.HideBanner = true
	e.HidePort = true

	// ===== GRACEFUL SHUTDOWN CONTEXT =====

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	// ===== START SERVER (NON-BLOCKING) =====

	go func() {
		log.Info("starting http server",
			zap.String("port", appConfig.Port),
		)

		if err := e.Start(":" + appConfig.Port); err != nil && err != http.ErrServerClosed {
			log.Fatal("failed to start server", zap.Error(err))
		}
	}()

	// ===== WAIT FOR SIGTERM / CTRL+C =====
	<-ctx.Done()
	log.Info("shutdown signal received")

	// ===== GRACEFUL SHUTDOWN WITH TIMEOUT =====
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(shutdownCtx); err != nil {
		log.Error("server shutdown failed", zap.Error(err))
	}

	log.Info("server exited properly")

}
