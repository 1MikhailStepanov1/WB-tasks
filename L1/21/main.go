package main

import "fmt"

type RussianSocket interface {
	InsertRUPlug() string
}

type AmericanPlug interface {
	InsertUSPlug() string
}

type USPlug struct{}

func (us *USPlug) InsertUSPlug() string {
	return "Подключена американская вилка"
}

type RUSocket struct {
}

func (ru *RUSocket) InsertRUPlug() string {
	return "Подключение успешно"
}

type USToRUAdapter struct {
	usPlug USPlug
}

func (a *USToRUAdapter) InsertRUPlug() string {
	return "Адаптер: " + a.usPlug.InsertUSPlug() + " через переходник"
}

func main() {
	ruSocket := RUSocket{}
	fmt.Println(ruSocket.InsertRUPlug())

	usPlug := USPlug{}

	adapter := USToRUAdapter{
		usPlug: usPlug,
	}
	fmt.Println(adapter.InsertRUPlug())
}
