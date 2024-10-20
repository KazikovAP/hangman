[![Go](https://img.shields.io/badge/-Go-464646?style=flat-square&logo=Go)](https://go.dev/)

# hangman
# Консольная игра "Виселица"

---
## Описание проекта
В игре игрок пытается угадать загаданное слово, вводя буквы по одной за раз. Слово выбирается случайно из предварительно заданного списка слов. Количество попыток ограничено, и за каждую неверную догадку визуализируется часть виселицы и фигурки висельника. Если игрок не может угадать букву, выводится подсказка к загадонному слову.

---
## Технологии
* Go 1.23.0
* DDD (Domain Driven Design)

---
## Запуск игры

**1. Клонировать репозиторий:**
```
git clone https://github.com/KazikovAP/hangman.git
```

**2. Запустить игру:**
```
go run cmd/hangman/main.go
```

## Пример игры
```
Не забудьте переключиться на русский язык!

Загаданное слово: _ _ _ _ _ _ 
          _______
          |/
          |
          |
          |
          |
          |
          |
          |
        __|________
        |         |


Допущено ошибок: 0: Пока ошибок нет
Осталось ошибок: 7

Введите следующую букву:
```
```
Загаданное слово: _ А _ А _ _
          _______
          |/
          |
          |
          |
          |
          |
          |
          |
        __|________
        |         |


Допущено ошибок: 0: Пока ошибок нет
Осталось ошибок: 7
```
```
Загаданное слово: _ А _ А _ О
          _______
          |/
          |     ( )
          |
          |
          |
          |
          |
          |
        __|________
        |         |


Допущено ошибок: 1: К
Осталось ошибок: 6
```
```
Загаданное слово: _ А _ А Л О
          _______
          |/
          |     ( )
          |     _|_
          |    /   \\
          |
          |
          |
          |
        __|________
        |         |


Допущено ошибок: 4: К, П, Г, В
Осталось ошибок: 3

Подсказка: старт чего-то
```
```
Загаданное слово: Н А Ч А Л О
          _______
          |/
          |     ( )
          |     _|_
          |    /   \\
          |
          |
          |
          |
        __|________
        |         |


Допущено ошибок: 4: К, П, Г, В
Осталось ошибок: 3

Поздравляем, вы победили!
Отгаданное слово: НАЧАЛО
Игра завершена!
```

---
## Разработал:
[Aleksey Kazikov](https://github.com/KazikovAP)

---
## Лицензия:
[MIT](https://opensource.org/licenses/MIT)
