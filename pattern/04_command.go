package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
Паттерн "Команда" (Command) является поведенческим паттерном проектирования,
который позволяет инкапсулировать запросы или операции в отдельные объекты.
Эти объекты могут быть переданы, сохранены и выполнены в разное время,
что обеспечивает более гибкую и расширяемую систему.

Плюсы:
-Отделение отправителя и получателя: Команда инкапсулирует запрос
	и его параметры, что разделяет отправителя команды от получателя.
-Отмена и повтор операций: Паттерн позволяет реализовать отмену и повтор операций (Undo/Redo).
-Поддержка отложенных операций: Команды могут быть выполнены в будущем,
	что полезно для планирования и управления операциями.

Минусы:
-Использование паттерна может привести к
	увеличению количества страктов, особенно если есть множество команд.
-Усложнение кода: Для мелких операций паттерн "Команда" может быть излишним и усложнять код.

Пример: В умных домах команды могут представлять действия,
такие как включение/выключение света, регулирование температуры
и управление устройствами безопасности. Паттерн "Команда" позволяет
управлять всеми устройствами через единый интерфейс.
*/

// Command Интерфейс команды
type Command interface {
	Execute()
}

// InsertCommand Конкретная команда - вставка текста
type InsertCommand struct {
	Text   string
	Editor *TextEditor
}

func (c *InsertCommand) Execute() {
	c.Editor.InsertText(c.Text)
}

// DeleteCommand Конкретная команда - удаление текста
type DeleteCommand struct {
	Position int
	Editor   *TextEditor
}

func (c *DeleteCommand) Execute() {
	c.Editor.DeleteText(c.Position)
}

// TextEditor Класс редактора
type TextEditor struct {
	Text string
}

func (e *TextEditor) InsertText(text string) {
	e.Text += text
}

func (e *TextEditor) DeleteText(position int) {
	if position >= 0 && position < len(e.Text) {
		e.Text = e.Text[:position] + e.Text[position+1:]
	}
}

func main() {
	editor := &TextEditor{}

	// Создаем команды
	insertCommand := &InsertCommand{Text: "Hello, ", Editor: editor}
	deleteCommand := &DeleteCommand{Position: 0, Editor: editor}

	// Выполняем команды
	insertCommand.Execute()
	deleteCommand.Execute()

	fmt.Println(editor.Text) // Вывод: "ello, "
}
