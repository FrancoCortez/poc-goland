package main

import (
	cencopim "cenco-pim"
	"cenco-pim/adapter/db"
	"cenco-pim/adapter/dependecy"
	"cenco-pim/config"
	lr "cenco-pim/util/logger"
	vr "cenco-pim/util/validator"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	appConf := config.AppConfig()
	logger := lr.New(appConf.Debug)
	validator := vr.New()
	initDb := db.InitDb()
	if err != nil {
		logger.Fatal().Err(err).Msg("")
		return
	}
	container := cencopim.InitializeServer(initDb, logger)
	appRouter := dependecy.GetRouter(&container, logger, validator, initDb)
	address := fmt.Sprintf(":%d", appConf.Server.Port)
	srv := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	//go func() {
	//
	//	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
	//		log.Printf("%s %s", method, route)
	//		return nil
	//	}
	//	if err := chi.Walk(appRouter, walkFunc); err != nil {
	//		log.Error().Str("Logging err: %s\n", err.Error())
	//	}
	//
	//	if err := http.ListenAndServe(address, appRouter); err != nil {
	//		log.Error().Str("method", "main").Msgf("%v", err)
	//		os.Exit(1)
	//	}
	//}()
	//
	//os.Exit(0)

	closed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		logger.Info().Msgf("Shutting down server %v", address)

		ctx, cancel := context.WithTimeout(context.Background(), appConf.Server.TimeoutIdle)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			logger.Warn().Err(err).Msg("Server shutdown failure")
		}

		sqlDB, err := initDb.DB()
		if err == nil {
			if err = sqlDB.Close(); err != nil {
				logger.Warn().Err(err).Msg("Db connection closing failure")
			}
		}

		close(closed)
	}()

	logger.Info().Msgf("Starting server %v", address)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server startup failure")
	}

	<-closed
}
