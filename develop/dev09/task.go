package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// Определение флагов командной строки.
	url := flag.String("url", "", "URL сайта для скачивания")
	output := flag.String("output", "index.html", "Имя файла для сохранения")
	flag.Parse()

	// Проверка обязательного флага "url".
	if *url == "" {
		fmt.Println("Использование: wget -url <URL> [-output <имя_файла>]")
		return
	}

	err := DownloadWebsite(*url, *output)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}

// DownloadWebsite скачивает веб-сайт по указанному URL и сохраняет его в файл.
func DownloadWebsite(url, output string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	// Если файл является HTML, загружаем связанные ресурсы.
	if strings.HasSuffix(output, ".html") {
		err := DownloadResources(output)
		if err != nil {
			return err
		}
	}

	fmt.Println("Завершено. Сайт успешно скачан.")
	return nil
}

// DownloadResources загружает связанные ресурсы (ссылки на изображения, скрипты и другие) из HTML-файла.
func DownloadResources(htmlFile string) error {
	file, err := os.Open(htmlFile)
	if err != nil {
		return err
	}
	defer file.Close()

	tokenizer := html.NewTokenizer(file)

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return nil
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data == "a" || token.Data == "link" || token.Data == "script" || token.Data == "img" {
				for _, attr := range token.Attr {
					if attr.Key == "href" || attr.Key == "src" {
						resourceURL := attr.Val
						fmt.Println("Загрузка ресурса:", resourceURL)
						err := DownloadResource(resourceURL)
						if err != nil {
							fmt.Println("Ошибка при загрузке ресурса:", err)
						}
					}
				}
			}
		}
	}
}

// DownloadResource загружает отдельный ресурс по его URL и сохраняет в текущей директории.
func DownloadResource(resourceURL string) error {
	resp, err := http.Get(resourceURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Извлечение имени файла из URL.
	tokens := strings.Split(resourceURL, "/")
	filename := tokens[len(tokens)-1]

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Ресурс успешно загружен:", filename)
	return nil
}
