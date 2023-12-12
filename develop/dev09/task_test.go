package main

import (
	"os"
	"testing"
)

func TestDownloadWebsite(t *testing.T) {
	url := "https://example.com"
	output := "test.html"

	err := DownloadWebsite(url, output)
	if err != nil {
		t.Errorf("Ошибка при скачивании сайта: %v", err)
	}

	// Проверяем, что файл существует.
	_, err = os.Stat(output)
	if err != nil {
		t.Errorf("Файл не был создан: %v", err)
	}

	// Очищаем файл после выполнения теста.
	defer os.Remove(output)
}

func TestDownloadResource(t *testing.T) {
	resourceURL := "https://example.com/image.png"
	filename := "image.png"

	err := DownloadResource(resourceURL)
	if err != nil {
		t.Errorf("Ошибка при загрузке ресурса: %v", err)
	}

	// Проверяем, что файл существует.
	_, err = os.Stat(filename)
	if err != nil {
		t.Errorf("Файл не был создан: %v", err)
	}

	// Очищаем файл после выполнения теста.
	defer os.Remove(filename)
}
