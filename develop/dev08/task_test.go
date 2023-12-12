package main

import (
	"os"
	"strings"
	"testing"
)

func TestCdCommand(t *testing.T) {
	os.Chdir("/") // Переход в корневую директорию для теста
	executeBuiltinCommand([]string{"cd", "tmp"})
	dir, _ := os.Getwd()
	if !strings.Contains(dir, "tmp") {
		t.Errorf("cd command failed, current directory: %s", dir)
	}
}

func TestPwdCommand(t *testing.T) {
	dirBefore, _ := os.Getwd()
	output := executeBuiltinCommand([]string{"pwd"})
	if output != dirBefore {
		t.Errorf("pwd command failed, expected: %s, got: %s", dirBefore, output)
	}
}

func TestEchoCommand(t *testing.T) {
	output := executeBuiltinCommand([]string{"echo", "hello"})
	if output != "hello" {
		t.Errorf("echo command failed, expected: hello, got: %s", output)
	}
}
