package main

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {
	for i := 0; i < 6; i++ {
		currentTime := time.Now() // Создаем переменную текущего времени.

		ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org") // Создаем переменную точного времени через NTP.
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка при получении времени: %s\n", err) // Обрабатываем ошибки.
			os.Exit(1)                                                        // В случае ошибки завершаем работу не с нулевым кодом
		}

		differenceTime := ntpTime.Sub(currentTime).Round(time.Millisecond) // Вычисляем разницу во времени.
		if differenceTime < 0 {                                            // Вычисляем именно модуль разницы во времени.
			differenceTime *= -1
		}

		timeFormat := "02.01.2006 15:04:05.000" // Приводим время к удобочитаемому формату

		//Выводим время в Stdout
		fmt.Printf("Точное время: %s\n", ntpTime.Format(timeFormat))
		fmt.Printf("Текущее время: %s\n", currentTime.Format(timeFormat))
		fmt.Printf("Разница во времени с сервером NTP: %s\n\n", differenceTime)

		time.Sleep(2 * time.Second)
	}
}
