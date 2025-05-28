package main

import (
	"context"
	"database/sql"
	"github.com/joho/godotenv"
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

	err := godotenv.Load(".env.local")
	if err != nil {
		logger.Error("Error loading .env file", slog.String("err", err.Error()))
	}

	sqlDSN := os.Getenv("SQL_DSN")
	TG_KEY := os.Getenv("TG_KEY")

	db, err := repositories.NewSQLDB(sqlDSN)
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

	bot, err := tgbotapi.NewBotAPI(TG_KEY)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	logger.Info("Authorized on account $s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 25

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		go middleware.MessageMiddleware(ctx, ClientDataService, bot, update, logger)
	}
}
