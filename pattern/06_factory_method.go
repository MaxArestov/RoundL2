package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
Паттерн "Фабричный метод" (Factory Method) является порождающим паттерном проектирования,
который предоставляет интерфейс для создания объектов, но позволяет подклассам выбирать класс создаваемого объекта.
Это позволяет делегировать создание объектов подклассам, обеспечивая гибкость и возможность расширения системы.

Плюсы:
-Упрощение кода: паттерн позволяет разделить код создания объектов от остального кода,
	что упрощает его понимание и поддержку.
-Гибкость: паттерн позволяет подклассам определять тип создаваемых объектов, что делает код более гибким.
-Производительность: паттерн может улучшить производительность,
	поскольку он позволяет избежать повторного создания объектов одного и того же типа.

Минусы:
-Усложнение кода(несмотря на упрощение в плюсах)): паттерн может усложнить код,
	если в приложении используется большое количество подклассов Фабрики.
-Непонятность кода: паттерн может сделать код менее понятным, если он используется в сложных ситуациях.

Пример: Библиотеки для доступа к базам данных могут использовать фабричные методы
для создания конкретных соединений с разными типами баз данных
(MySQL, PostgreSQL, MongoDB и т. д.) без изменения клиентского кода.
*/

// TransportFactory Интерфейс фабрики
type TransportFactory interface {
	CreateTransport() Transport
}

// Transport Интерфейс для созданных продуктов (транспортных средств)
type Transport interface {
	Drive()
}

// CarFactory Конкретная фабрика для создания автомобилей
type CarFactory struct{}

func (f *CarFactory) CreateTransport() Transport {
	return &Car{}
}

// MotorcycleFactory Конкретная фабрика для создания мотоциклов
type MotorcycleFactory struct{}

func (f *MotorcycleFactory) CreateTransport() Transport {
	return &Motorcycle{}
}

// Car Конкретный продукт (автомобиль)
type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Driving a car")
}

// Motorcycle Конкретный продукт (мотоцикл)
type Motorcycle struct{}

func (m *Motorcycle) Drive() {
	fmt.Println("Riding a motorcycle")
}

func main() {
	// Создание фабрик для автомобилей и мотоциклов
	carFactory := &CarFactory{}
	motorcycleFactory := &MotorcycleFactory{}

	// Создание и использование транспортных средств
	car := carFactory.CreateTransport()
	motorcycle := motorcycleFactory.CreateTransport()

	car.Drive()        // Вывод: "Driving a car"
	motorcycle.Drive() // Вывод: "Riding a motorcycle"
}
