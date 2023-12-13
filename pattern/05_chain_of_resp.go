package pattern

import (
	"fmt"
	"net/http"
)

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/
/*
Паттерн "Цепочка вызовов" (Chain of Responsibility) является поведенческим паттерном проектирования,
который позволяет передавать запросы последовательно по цепочке обработчиков.
Каждый обработчик решает, может ли он обработать запрос, и передает его дальше по цепи, если не может.
Это позволяет создавать гибкие и расширяемые системы обработки запросов.

Плюсы:
-Уменьшение связанности: Паттерн позволяет избежать прямой зависимости между отправителем и получателем запроса.
-Гибкость и расширяемость: Цепочка обработчиков легко расширяется
	путем добавления новых обработчиков или изменения порядка их вызова.
-Простота добавления новых обработчиков: Новый обработчик можно добавить, не меняя существующий код.

Минусы:
-Гарантия обработки: Не всегда гарантируется обработка запроса, если цепь не настроена правильно.
-Увеличение сложности: Если цепь слишком длинная или сложная, это может усложнить понимание и отладку.

Пример: паттерн "Цепочка вызовов" может использоваться для обработки запроса HTTP.
Каждый обработчик выполняет свою специфическую проверку (аутентификация, авторизация и основная обработка)
и передает запрос дальше по цепочке. Это позволяет легко добавлять и изменять обработчики в цепи обработки запросов
HTTP без изменения существующего кода.
*/

// Handler Интерфейс обработчика запроса
type Handler interface {
	HandleRequest(request *http.Request)
	SetNext(handler Handler)
}

// AuthHandler Конкретный обработчик - проверка аутентификации
type AuthHandler struct {
	nextHandler Handler
}

func (h *AuthHandler) HandleRequest(request *http.Request) {
	// Проверка аутентификации пользователя
	fmt.Println("AuthHandler: Checking authentication...")
	// Передача запроса следующему обработчику в цепочке
	if h.nextHandler != nil {
		h.nextHandler.HandleRequest(request)
	}
}

func (h *AuthHandler) SetNext(handler Handler) {
	h.nextHandler = handler
}

// AuthorizationHandler Конкретный обработчик - проверка авторизации
type AuthorizationHandler struct {
	nextHandler Handler
}

func (h *AuthorizationHandler) HandleRequest(request *http.Request) {
	// Проверка авторизации пользователя
	fmt.Println("AuthorizationHandler: Checking authorization...")
	// Передача запроса следующему обработчику в цепочке
	if h.nextHandler != nil {
		h.nextHandler.HandleRequest(request)
	}
}

func (h *AuthorizationHandler) SetNext(handler Handler) {
	h.nextHandler = handler
}

// MainHandler Конкретный обработчик - основная обработка
type MainHandler struct{}

func (h *MainHandler) HandleRequest(request *http.Request) {
	// Обработка основной логики запроса
	fmt.Println("MainHandler: Handling request...")
}

func (h *MainHandler) SetNext(handler Handler) {
	// Опциональная реализация SetNext для MainHandler
	// Этот метод может быть пустым, так как MainHandler является последним в цепочке
}

func main() {
	// Создаем цепочку обработчиков
	mainHandler := &MainHandler{}
	authHandler := &AuthHandler{}
	authorizationHandler := &AuthorizationHandler{}

	// Настраиваем цепь вызовов
	authHandler.SetNext(authorizationHandler)
	authorizationHandler.SetNext(mainHandler)

	// Пример запроса
	request, _ := http.NewRequest("GET", "/", nil)

	// Обработка запроса начинается с первого обработчика в цепочке
	authHandler.HandleRequest(request)
}
