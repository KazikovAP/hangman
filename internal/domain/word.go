package domain

import (
	"crypto/rand"
	"hangman/config"
	"math/big"
)

type Word struct {
	Word       string
	Hint       string
	Dictionary []wordWithHint
}

type wordWithHint struct {
	Word string
	Hint string
}

func NewWord() Word {
	return Word{}
}

func (w *Word) Build(cfg *config.Config) error {
	for _, word := range cfg.Words {
		w.Dictionary = append(w.Dictionary, wordWithHint{Word: word.Word, Hint: word.Hint})
	}

	return nil
}

func (w *Word) SetRandomWord() error {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(w.Dictionary))))
	if err != nil {
		return err
	}

	word := w.Dictionary[int(nBig.Int64())]

	w.Word = word.Word
	w.Hint = word.Hint

	return nil
}
