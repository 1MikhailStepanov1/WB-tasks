package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) Hello() {
	fmt.Printf("Hello, my name is %s!\nI'm %d old.\n", h.Name, h.Age)
}

type Action struct {
	Human
	ActionType string
}

func (a *Action) SetActionType(s string) {
	a.ActionType = s
}

func (a *Action) Perform() {
	fmt.Printf("%s performing action: %s\n", a.Name, a.ActionType)
}

func main() {
	bob := Action{Human: Human{Name: "Bob", Age: 22}, ActionType: "write code"}
	bob.Hello()
	bob.Perform()
}
