package domain

import (
	"fmt"
	"strings"

	"github.com/KazikovAP/hangman/internal/errors"
)

type Game struct {
	Attempts     int
	MaxAttempts  int
	Letters      map[rune]struct{}
	WrongLetters map[rune]struct{}
	UniqueErrors int
	Word         Word
	HangStages   []string
}

func NewGame(word Word, maxAttempts int) *Game {
	return &Game{
		MaxAttempts:  maxAttempts,
		Letters:      make(map[rune]struct{}),
		WrongLetters: make(map[rune]struct{}),
		UniqueErrors: 0,
		Word:         word,
	}
}

func (g *Game) CheckLetter(r rune) bool {
	if _, guessed := g.Letters[r]; guessed {
		return true
	}

	if strings.ContainsRune(g.Word.Word, r) {
		g.Letters[r] = struct{}{}
		return true
	}

	if !strings.ContainsRune(g.Word.Word, r) {
		if _, exists := g.WrongLetters[r]; !exists {
			g.WrongLetters[r] = struct{}{}
			g.UniqueErrors++
			g.Attempts++
		}
	}

	return false
}

func (g *Game) IsGameOver() bool {
	return g.Attempts >= g.MaxAttempts || g.IsUserWon()
}

func (g *Game) IsUserWon() bool {
	for _, letter := range g.Word.Word {
		if _, guessed := g.Letters[letter]; !guessed {
			return false
		}
	}

	return true
}

func (g *Game) WrongLettersToString() []string {
	wrongLetters := make([]string, 0, len(g.WrongLetters))
	for letter := range g.WrongLetters {
		wrongLetters = append(wrongLetters, string(letter))
	}

	return wrongLetters
}

func (g *Game) IsUserLost() bool {
	return g.Attempts >= g.MaxAttempts
}

func (g *Game) SetWord() {
	if err := g.Word.SetRandomWord(); err != nil {
		fmt.Println(errors.SetWordError{Err: err})
		return
	}
}
