Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
# Ответ:

Программа выведет "error", потому что `err` является интерфейсом, 
который хранит в себе `nil` значение типа `*customError`. 
В Go, интерфейсное значение не равно `nil`, если его типовая часть не равна `nil`, 
даже если само значение равно `nil`. В этом примере `err` имеет тип, 
который не является `nil` (`*customError`), поэтому `err != nil` возвращает `true`.
```