package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/zhas-off/test-music-player/internal/app"
	config "github.com/zhas-off/test-music-player/internal/config"
	repo "github.com/zhas-off/test-music-player/internal/repository"
	playlist "github.com/zhas-off/test-music-player/internal/service"
)

func main() {
	// Инициализируем сигнальный канал для graceful shutdown
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// Считываем данные из конфиг файла
	cfgPath := flag.String("config", "./config.yaml", "Path to config file")
	flag.Parse()

	// Инициализируем логирование
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	logger.Info().Msg("Playlist service has started")

	// Загружаем данные конфиг файла
	conf := config.NewConfig()
	err := conf.LoadConfig(*cfgPath)
	if err != nil {
		logger.Fatal().Err(err).Msg("config loading error")
	}

	database, err := repo.InitDatabase(conf.DbUrl)
	if err != nil {
		logger.Fatal().Err(err).Msg("database init error")
	}
	defer database.Conn.Close(context.Background())

	pl := playlist.Init()
	pl.Logger = &logger
	err = database.LoadToDatabaseIfNotExistBaseSongsSet()
	if err != nil {
		logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("random songs set loading error")
	}
	databaseList, err := database.Load()
	if err != nil {
		logger.Fatal().Err(err).Msg("playlist initializing error")
	}
	pl.LoadListToPlaylistFromDatabase(databaseList)
	logger.Info().Msg("playlist initialized")

	app := app.New(pl, database)
	app.Logger = &logger

	wg := &sync.WaitGroup{}
	app.Config = conf
	go func() {
		app.Run()
	}()

	wg.Add(1)
	go func() {
		pl.Run()
		wg.Done()
	}()

	<-signalCh
	fmt.Printf("Stopping the server")
	pl.ExitChan <- struct{}{}
	list, err := pl.GetList()
	if err != nil {
		logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("error getting list to upload to database")
	} else {
		err = database.Update(list)
		if err != nil {
			logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("database upload error")
		}
		logger.Info().Msg("state uploaded")
	}

	wg.Wait()
	logger.Info().Msg("service stopping")
}
