package errors

import "fmt"

type SetWordError struct {
	Err error
}

func (e SetWordError) Error() string {
	return fmt.Sprintf("Ошибка при установке случайного слова: %v", e.Err)
}
