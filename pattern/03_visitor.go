package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
Паттерн "Посетитель" (Visitor) является поведенческим паттерном проектирования,
который позволяет добавлять новые операции к классам объектов без изменения их кода.
Он достигается путем создания внешнего объекта (посетителя), который содержит
новые методы для выполнения операций над объектами.

Плюсы:
-Разделение ответственности: Паттерн "Посетитель" разделяет операции над объектами
	и сами объекты, что способствует соблюдению принципа единственной ответственности.
-Гибкость: Добавление новых операций происходит через создание новых посетителей,
	что делает систему более гибкой и расширяемой.
-Читаемость кода: Операции над объектами выносятся в отдельные структуры (посетители), что улучшает читаемость кода.

Минусы:
-Усложнение структуры: Внедрение паттерна "Посетитель" может усложнить структуру кода и добавить дополнительные классы.
-Нарушение инкапсуляции: Посетитель получает доступ к внутреннему состоянию объектов, что может нарушить инкапсуляцию.

На практике в приложениях для работы с базами данных, паттерн "Посетитель" может использоваться для анализа
и обработки данных, например, для вычисления статистики, фильтрации или преобразования данных.
*/

type Visitable interface {
	Accept(visitor Visitor)
}

// Product Конкретный элемент - продукт
type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func (p *Product) Accept(visitor Visitor) {
	visitor.VisitProduct(p)
}

// Visitor Интерфейс посетителя
type Visitor interface {
	VisitProduct(product *Product)
}

// TotalPriceVisitor Конкретный посетитель - вычисление общей стоимости заказа
type TotalPriceVisitor struct {
	TotalPrice float64
}

func (v *TotalPriceVisitor) VisitProduct(product *Product) {
	v.TotalPrice += float64(product.Quantity) * product.Price
}

func main() {
	products := []Visitable{
		&Product{Name: "Laptop", Price: 1000, Quantity: 2},
		&Product{Name: "Mouse", Price: 20, Quantity: 5},
		&Product{Name: "Keyboard", Price: 50, Quantity: 3},
	}

	totalPriceVisitor := &TotalPriceVisitor{}

	// Вычисление общей стоимости заказа с помощью посетителя
	for _, product := range products {
		product.Accept(totalPriceVisitor)
	}

	fmt.Printf("Total Price: $%.2f\n", totalPriceVisitor.TotalPrice)
}
