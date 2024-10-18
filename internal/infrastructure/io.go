package infrastructure

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"strings"
)

type IOAdapter struct {
	r      io.Reader
	w      io.Writer
	logger *slog.Logger
}

func NewIOAdapter(r io.Reader, w io.Writer, logger *slog.Logger) *IOAdapter {
	return &IOAdapter{
		r:      r,
		w:      w,
		logger: logger,
	}
}

func (a *IOAdapter) Input() (string, error) {
	reader := bufio.NewReader(a.r)
	return reader.ReadString('\n')
}

func (a *IOAdapter) Output(content ...string) {
	message := strings.Join(content, "")
	_, err := fmt.Fprintln(a.w, message)

	if err != nil {
		a.logger.Error(err.Error())
	}
}

func (a *IOAdapter) PrintOut(result, figure string, errorsMade, errorsBalance int,
	isUserWon, isUserLost bool, word, wrongLetters, hint string) {
	a.Output("\nНе забудьте переключиться на русский язык!")
	a.Output(fmt.Sprintf("\nЗагаданное слово: %v", result))
	a.Output(figure)

	if errorsMade > 0 {
		a.Output(fmt.Sprintf("\nДопущено ошибок: %v: %s", errorsMade, wrongLetters))
	} else {
		a.Output(fmt.Sprintf("\nДопущено ошибок: %v: Пока ошибок нет", errorsMade))
	}

	a.Output(fmt.Sprintf("Осталось ошибок: %v\n", errorsBalance))

	if hint != "" {
		a.Output(fmt.Sprintf("Подсказка: %s\n", hint))
	}

	if isUserWon {
		a.Output(fmt.Sprintf("Поздравляем, вы победили!)\nОтгаданное слово: %v", word))
		a.Output("Игра завершена!")
	} else if isUserLost {
		a.Output(fmt.Sprintf("Жаль, но вы проиграли:(\nЗагаданное слово: %v", word))
		a.Output("Игра завершена!")
	}
}

func (a *IOAdapter) UserInput() string {
	var letter string

	a.Output("Введите следующую букву: ")
	_, err := fmt.Fscanln(a.r, &letter)

	if err != nil {
		a.logger.Error(err.Error())
		return ""
	}

	return strings.ToUpper(letter)
}
