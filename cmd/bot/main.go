package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"mkk_notification_bot/internal/middleware"
	"mkk_notification_bot/internal/processors"
	"mkk_notification_bot/internal/repositories"
	"mkk_notification_bot/services"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	loggerJSONHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(loggerJSONHandler)
	slog.SetDefault(logger)

	//SQLDSN := os.Getenv("SQL_DSN")

	SQLDSN := ""
	db, err := repositories.NewSQLDB(SQLDSN)
	if err != nil {
		logger.Error(
			"error connecting to database",
			slog.String("err", err.Error()),
		)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.Error(
				"error closing db",
				slog.String("err", err.Error()),
			)
		}
	}(db)

	dbRepository := repositories.NewSQLRepository(db, logger)
	clientDataProcessor := processors.NewClientDataProcessor(dbRepository, logger)
	ClientDataService := services.NewClientDataService(clientDataProcessor, logger)

	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	fmt.Printf("Authorized on account $s\n", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 25

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		go middleware.MessageMiddleware(ctx, ClientDataService, bot, update, logger)
	}
}
