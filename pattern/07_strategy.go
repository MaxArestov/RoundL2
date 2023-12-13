package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
Паттерн "Стратегия" является поведенческим паттерном проектирования,
который определяет семейство алгоритмов, инкапсулирует каждый из них
и делает их взаимозаменяемыми. Паттерн "Стратегия" позволяет клиентскому
коду выбирать подходящий алгоритм из семейства и использовать его независимо от самого алгоритма.

Плюсы:
-Гибкость: паттерн позволяет клиенту выбирать алгоритм в зависимости от конкретных потребностей.
-Расширяемость: паттерн позволяет легко добавлять новые алгоритмы.
-Упрощение кода: паттерн позволяет упростить код клиента, поскольку он не зависит от конкретного алгоритма.

Минусы:
-Усложнение структуры программы: Введение дополнительных классов и интерфейсов может усложнить структуру программы.

Пример: При разработке приложения, которое работает с данными, можно использовать паттерн "Стратегия"
для выбора алгоритма кэширования данных в зависимости от требований к производительности и доступности данных.
*/

// CacheStrategy Интерфейс для стратегии кэширования
type CacheStrategy interface {
	Get(key string) string
	Set(key, value string)
}

// InMemoryCache Конкретная стратегия - кэширование в памяти
type InMemoryCache struct {
	cache map[string]string
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		cache: make(map[string]string),
	}
}

func (c *InMemoryCache) Get(key string) string {
	return c.cache[key]
}

func (c *InMemoryCache) Set(key, value string) {
	c.cache[key] = value
}

// DatabaseCache Конкретная стратегия - кэширование в базе данных
type DatabaseCache struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func NewDatabaseCache() *DatabaseCache {
	return &DatabaseCache{}
}

func (c *DatabaseCache) Get(key string) string {
	// Здесь реализуется логика получения данных из базы данных
	// В данном примере просто возвращается пустая строка
	return ""
}

func (c *DatabaseCache) Set(key, value string) {
	// Здесь реализуется логика сохранения данных в базе данных
}

// Клиентский код, использующий стратегию кэширования
func main() {
	// Выбор стратегии кэширования в зависимости от требований приложения
	inMemoryCache := NewInMemoryCache()
	databaseCache := NewDatabaseCache()

	// Использование стратегии кэширования
	key := "some_key"
	value := "some_value"

	// Кэширование в памяти
	inMemoryCache.Set(key, value)
	fmt.Println("In-Memory Cache:", inMemoryCache.Get(key)) // Вывод: "In-Memory Cache: some_value"

	// Кэширование в базе данных
	databaseCache.Set(key, value)
	fmt.Println("Database Cache:", databaseCache.Get(key)) // Вывод: "Database Cache: some_value"
}
