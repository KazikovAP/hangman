package main

import (
	"hangman/config"
	"hangman/internal/application"
	"hangman/internal/domain"
	"hangman/internal/infrastructure"
	"log/slog"
	"os"
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
