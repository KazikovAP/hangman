package main

import (
	"log/slog"
	"os"

	"github.com/KazikovAP/hangman/internal/application"
	"github.com/KazikovAP/hangman/internal/domain"
	"github.com/KazikovAP/hangman/internal/infrastructure"

	"github.com/KazikovAP/hangman/config"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cfg, err := config.Init(nil)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	words := domain.NewWord()

	if err = words.Build(cfg); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	game := domain.NewGame(words, 7)

	io := infrastructure.NewIOAdapter(os.Stdin, os.Stdout, logger)

	app := application.New(game, io)

	if err = app.Start(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
