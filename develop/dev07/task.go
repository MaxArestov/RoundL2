package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

// Определение переменной or как функции, принимающей переменное количество каналов
// и возвращающей один канал.
var or func(channels ...<-chan interface{}) <-chan interface{}

// Инициализация функции or в блоке init.
func init() {
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		out := make(chan interface{}) // Создание выходного канала
		done := make(chan struct{})   // Создание канала для сигнала завершения

		// Запуск горутин для каждого входного канала
		for _, ch := range channels {
			go func(c <-chan interface{}) {
				select {
				case <-c:
					close(done) // Если канал закрылся, закрываем done
				case <-done: // Если done закрылся, завершаем горутину
				}
			}(ch)
		}

		// Горутина, ожидающая сигнала закрытия канала done
		go func() {
			defer close(out) // По завершении закрываем выходной канал
			<-done           // Ожидание закрытия канала done
		}()

		return out // Возвращение выходного канала
	}
}

// Функция sig создает канал, который закрывается через заданное время.
func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)    // Закрытие канала по завершении таймера
		time.Sleep(after) // Ожидание заданного времени
	}()
	return c
}

func main() {
	sigIntChan := make(chan interface{})  // Канал для сигнала прерывания
	sigInt := make(chan os.Signal, 1)     // Канал для системных сигналов
	signal.Notify(sigInt, syscall.SIGINT) // Перехват сигнала SIGINT (Ctrl+C)

	// Горутина, закрывающая sigIntChan при получении сигнала SIGINT
	go func() {
		<-sigInt
		close(sigIntChan)
	}()

	fmt.Println("You can close program immediately by pressing Ctrl+C!")

	start := time.Now() // Запоминаем время начала выполнения

	// Ожидание закрытия одного из каналов
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(15*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
		sigIntChan,
	)

	// Вывод времени, прошедшего с начала выполнения
	fmt.Printf("done after %v\n", time.Since(start))
}
