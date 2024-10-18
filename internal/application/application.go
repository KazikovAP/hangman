package application

import (
	"hangman/internal/domain"
	"hangman/internal/infrastructure"
	"strings"
	"unicode"
)

const (
	errorsAllowed  int = 7
	showHintErrors int = 3
)

type App struct {
	game *domain.Game
	io   *infrastructure.IOAdapter
}

func New(game *domain.Game, io *infrastructure.IOAdapter) *App {
	return &App{game: game, io: io}
}

func (a *App) Start() error {
	a.game.SetWord()

	for !a.game.IsGameOver() {
		hint := ""
		if a.game.Attempts > showHintErrors {
			hint = a.game.Word.Hint
		}

		a.PrintGameState(hint)

		letter := a.io.UserInput()
		letter = strings.TrimSpace(letter)
		runes := []rune(letter)

		if len(runes) > 0 {
			runeLetter := runes[0]
			runeLetter = unicode.ToUpper(runeLetter)
			a.game.CheckLetter(runeLetter)
		}
	}

	a.PrintGameState("")

	return nil
}

func (a *App) PrintHangman() string {
	index := a.game.Attempts
	if index < 0 || index >= len(domain.HangmanStages) {
		index = len(domain.HangmanStages) - 1
	}

	return domain.HangmanStages[index]
}

func (a *App) HiddenWord() string {
	result := ""

	for _, letter := range a.game.Word.Word {
		if _, guessed := a.game.Letters[letter]; guessed {
			result += string(letter) + " "
		} else {
			result += "_ "
		}
	}

	return result
}

func (a *App) GetWrongLetters() string {
	return strings.Join(a.game.WrongLettersToString(), ", ")
}

func (a *App) PrintGameState(hint string) {
	a.io.PrintOut(
		a.HiddenWord(),
		a.PrintHangman(),
		a.game.Attempts,
		errorsAllowed-a.game.UniqueErrors,
		a.game.IsUserWon(),
		a.game.IsUserLost(),
		a.game.Word.Word,
		a.GetWrongLetters(),
		hint,
	)
}
