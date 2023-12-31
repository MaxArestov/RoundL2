package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
Паттерн "Состояние" является поведенческим паттерном проектирования,
который позволяет объекту изменять свое поведение в зависимости от внутреннего состояния.
Он позволяет объекту иметь различное поведение в разных состояниях и управлять переходами между состояниями.

Плюсы:
-Разделение ответственностей: Каждое состояние инкапсулируется в отдельном стракте,
	что делает код более читаемым и поддерживаемым.
-Гибкость и расширяемость: Легко добавлять новые состояния
	и изменять поведение объекта без изменения его основного кода.
-Соблюдение принципа единственной ответственности:
	Каждый стракт состояния имеет только одну ответственность - определение поведения в данном состоянии.

Минусы:
-Увеличение количества классов: Применение паттерна может привести к увеличению числа страктов в программе,
	что может усложнить структуру.

Пример: В веб-приложениях у пользовательских сессий может быть несколько состояний,
таких как "вход", "выход", "активная сессия".
Паттерн "Состояние" позволяет управлять переходами между этими состояниями и определять,
какие операции доступны пользователю в каждом состоянии.
*/

// SessionState Интерфейс для состояния сессии
type SessionState interface {
	Login()
	Logout()
	ViewContent()
}

// LoggedInState Конкретное состояние - сессия в состоянии "вход"
type LoggedInState struct{}

func (ls *LoggedInState) Login() {
	fmt.Println("Вы уже вошли в систему.")
}

func (ls *LoggedInState) Logout() {
	fmt.Println("Выход из системы.")
}

func (ls *LoggedInState) ViewContent() {
	fmt.Println("Просмотр контента.")
}

// LoggedOutState Конкретное состояние - сессия в состоянии "выход"
type LoggedOutState struct{}

func (los *LoggedOutState) Login() {
	fmt.Println("Вход в систему.")
}

func (los *LoggedOutState) Logout() {
	fmt.Println("Вы уже вышли из системы.")
}

func (los *LoggedOutState) ViewContent() {
	fmt.Println("Для просмотра контента войдите в систему.")
}

// UserSession Сессия пользователя
type UserSession struct {
	state SessionState
}

func (us *UserSession) SetState(state SessionState) {
	us.state = state
}

func main() {
	// Создание сессии пользователя и начальное состояние - "выход"
	session := &UserSession{state: &LoggedOutState{}}

	// Попытка просмотра контента в состоянии "выход"
	session.state.ViewContent()

	// Вход в систему
	session.state.Login()

	// Попытка входа в систему в состоянии "вход"
	session.state.Login()

	// Просмотр контента после входа
	session.state.ViewContent()

	// Выход из системы
	session.state.Logout()

	// Попытка выхода из системы в состоянии "выход"
	session.state.Logout()
}
