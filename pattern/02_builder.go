package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
Паттерн "Строитель" - это порождающий паттерн проектирования, который позволяет создавать сложные объекты с множеством
параметров, предоставляя шаг за шагом процесс конструирования. Он используется, чтобы упростить создание объектов с
различными конфигурациями, обеспечивая гибкость и читаемость кода.

Плюсы:
-Гибкость в создании объектов: Позволяет создавать сложные объекты с разными конфигурациями,
	не затрудняя интерфейс класса.
-Читаемость кода: Улучшает читаемость кода, так как параметры строителя можно именовать и явно указывать, что создается.
-Соблюдение принципа единственной ответственности Liskov: Разделяет процесс конструирования объекта от самого объекта,
	что соответствует принципу единственной ответственности.

Минусы:
-Дополнительный код: Требует создания отдельного строителя для каждого типа объекта, что может увеличить объем кода.
-Может препятствовать/усложнять внедрение зависимостей.

Пример на практике - создание HTML-документа.
*/

type Computer struct {
	Motherboard string
	CPU         string
	GPU         string
	RAM         string
	Storage     string
}

// ComputerBuilder Интерфейс строителя
type ComputerBuilder interface {
	SetMotherboard() ComputerBuilder
	SetCPU() ComputerBuilder
	SetGPU() ComputerBuilder
	SetRAM() ComputerBuilder
	SetStorage() ComputerBuilder
	Build() Computer
}

// GamingComputerBuilder Конкретный строитель
type GamingComputerBuilder struct {
	computer Computer
}

func NewGamingComputerBuilder() *GamingComputerBuilder {
	return &GamingComputerBuilder{}
}

func (b *GamingComputerBuilder) SetMotherboard() ComputerBuilder {
	b.computer.Motherboard = "A4TECH"
	return b
}

func (b *GamingComputerBuilder) SetCPU() ComputerBuilder {
	b.computer.CPU = "AMD Ryzen 5 6600"
	return b
}

func (b *GamingComputerBuilder) SetGPU() ComputerBuilder {
	b.computer.GPU = "Nvidia RTX 4060"
	return b
}

func (b *GamingComputerBuilder) SetRAM() ComputerBuilder {
	b.computer.RAM = "32GB RAM"
	return b
}

func (b *GamingComputerBuilder) SetStorage() ComputerBuilder {
	b.computer.Storage = "1TB SSD"
	return b
}

func (b *GamingComputerBuilder) Build() Computer {
	return b.computer
}

func main() {
	// Создаем строителя
	builder := NewGamingComputerBuilder()

	// Собираем компьютер
	gamingComputer := builder.
		SetMotherboard().
		SetCPU().
		SetGPU().
		SetRAM().
		SetStorage().
		Build()

	fmt.Printf("Gaming Computer Configuration:\nMotherboard: %s\nCPU: %s\nGPU: %s\nRAM: %s\nStorage: %s\n",
		gamingComputer.Motherboard, gamingComputer.CPU, gamingComputer.GPU, gamingComputer.RAM, gamingComputer.Storage)
}
