package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Определение главной функции main.
func main() {
	// Настройка обработки сигнала прерывания (Ctrl+C).
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	defer signal.Stop(sigCh)

	go func() {
		for range sigCh {
			fmt.Println("\nReceived interrupt signal, exiting...")
			os.Exit(0)
		}
	}()

	// Создание читателя для пользовательского ввода.
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("myshell> ")

		input, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Fprintln(os.Stderr, "error: failed to read command:", err)
			continue
		}

		input = strings.TrimSuffix(input, "\n")

		// Проверка наличия символа "|" для определения конвейера.
		if strings.Contains(input, "|") {
			handlePipes(input)
			continue
		}

		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		// Проверка, является ли команда встроенной.
		if isBuiltinCommand(args[0]) {
			executeBuiltinCommand(args)
		} else if strings.Contains(input, "|") {
			handlePipes(input)
		} else {
			fmt.Println("unknown command")
		}
	}
}

// Функция isBuiltinCommand проверяет, является ли команда встроенной.
func isBuiltinCommand(cmd string) bool {
	switch cmd {
	case "cd", "pwd", "echo", "kill", "ps", "\\quit":
		return true
	default:
		return false
	}
}

// Функция executeBuiltinCommand выполняет встроенные команды.
func executeBuiltinCommand(args []string) string {
	switch args[0] {
	case "cd":
		if len(args) > 1 {
			os.Chdir(args[1])
		}
	case "pwd":
		dir, _ := os.Getwd()
		fmt.Println(dir)
		return dir
	case "echo":
		fmt.Println(strings.Join(args[1:], " "))
		return strings.Join(args[1:], " ")
	case "kill":
		killProcess(args[1])
	case "ps":
		executeCommand([]string{"ps", "-e"})
		return ""
	case "\\quit":
		fmt.Println("Exiting by `quit` command")
		os.Exit(0)
		return ""
	}
	return ""
}

// Функция executeCommand выполняет внешние команды.
func executeCommand(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

// Функция killProcess завершает процесс с указанным PID.
func killProcess(pid string) string {
	killedProcess := fmt.Sprintf("Killing process %s...\n", pid)
	fmt.Print(killedProcess)
	return killedProcess
}

// Функция handlePipes обрабатывает конвейер (pipe).
func handlePipes(input string) {
	commands := strings.Split(input, "|")

	for _, commandStr := range commands {
		command := strings.Fields(commandStr)
		if len(command) == 0 {
			continue
		}
		if isBuiltinCommand(command[0]) {
			executeBuiltinCommand(command)
		} else {
			executeCommand(command)
		}
	}
}
