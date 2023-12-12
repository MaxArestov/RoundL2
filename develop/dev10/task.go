package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	// Определение флагов командной строки.
	timeoutFlag := flag.Duration("timeout", 10*time.Second, "Таймаут подключения")
	flag.Parse()

	// Проверка наличия аргументов host и port.
	if flag.NArg() != 2 {
		fmt.Println("Использование: go-telnet [--timeout=<таймаут>] host port")
		return
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	// Установка соединения с сервером.
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), *timeoutFlag)
	if err != nil {
		fmt.Println("Ошибка при подключении к серверу:", err)
		return
	}
	defer conn.Close()

	// Настройка обработки сигнала прерывания (Ctrl+C или Ctrl+D).
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	go func() {
		<-sigCh
		fmt.Println("\nЗавершение программы.")
		conn.Close()
		os.Exit(0)
	}()

	// Копирование данных между STDIN и сокетом.
	go func() {
		_, err := io.Copy(conn, os.Stdin)
		if err != nil {
			fmt.Println("Ошибка при отправке данных:", err)
		}
	}()

	// Копирование данных между сокетом и STDOUT.
	_, err = io.Copy(os.Stdout, conn)
	if err != nil {
		fmt.Println("Ошибка при чтении данных:", err)
	}
}
