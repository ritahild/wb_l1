package main

import (
	"errors"
	"fmt"
	"strings"
)

type Human struct {
	Name string
	Age  int
}

func (h *Human) String() string {
	return fmt.Sprintf("%s %d years old", h.Name, h.Age)
}

func (h *Human) Rename(name string) error {
	if len(name) < 3 {
		return errors.New("name too short")
	}
	h.Name = strings.Title(strings.ToLower(name))
	return nil
}

type Action struct {
	Human
}

func NewAction(name string, age int) *Action {
	return &Action{
		Human{
			Name: name,
			Age:  age,
		},
	}
}

func main() {
	action := NewAction("Van", 300)

	fmt.Println(action)

	action.Rename("Bill")

	fmt.Println(action)
}
