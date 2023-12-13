package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/* Паттерн "Фасад" - это структурный паттерн проектирования, который предоставляет упрощенный интерфейс
для взаимодействия с сложной системой, скрывая ее детали реализации.

Плюсы:
-Упрощает работу с комплексными системами, предоставляя простой интерфейс.
-Облегчает поддержку и изменение системы, так как изменения внутри системы не затрагивают реализацию фасада.

Минусы:
-Может создать дополнительный слой абстракции, что может быть излишним в простых системах.
-Не всегда подходит, если клиенту нужно более сложное управление подсистемой.

Пример использования:
Веб-приложение может использовать фасад для управления всеми авторизацией и аутентификацией пользователей.
Фасад будет предоставлять методы для регистрации, входа и управления пользователями, скрывая сложности,
связанные с хранением и проверкой учетных данных.
*/

import "fmt"

// HomeFacade для управления домашними устройствами
type HomeFacade struct {
	light  *Light
	airCon *AirConditioner
}

func NewHomeFacade() *HomeFacade {
	return &HomeFacade{
		light:  &Light{},
		airCon: &AirConditioner{},
	}
}

func (hf *HomeFacade) TurnOnEverything() {
	fmt.Println("Turning on all devices")
	hf.light.TurnOn()
	hf.airCon.TurnOn()
}

func (hf *HomeFacade) TurnOffEverything() {
	fmt.Println("Turning off all devices")
	hf.light.TurnOff()
	hf.airCon.TurnOff()
}

// Light: светильник
type Light struct{}

func (l *Light) TurnOn() {
	fmt.Println("Light is on")
}

func (l *Light) TurnOff() {
	fmt.Println("Light is off")
}

// AirConditioner: кондиционер
type AirConditioner struct{}

func (ac *AirConditioner) TurnOn() {
	fmt.Println("Air conditioner is on")
}

func (ac *AirConditioner) TurnOff() {
	fmt.Println("Air conditioner is off")
}

func main() {
	home := NewHomeFacade()

	// Включаем все устройства с помощью фасада
	home.TurnOnEverything()

	// Выключаем все устройства с помощью фасада
	home.TurnOffEverything()
}
