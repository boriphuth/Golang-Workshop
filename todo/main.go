package main

import (
	"net/http"

	"github.com/labstack/echo"
)

var todoList = []todo{}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/todo", func(c echo.Context) error {
		var t todo
		error := c.Bind(t)
		if error != nil {
			return error
		}
		NewTodo(t.Topic)
		return c.JSON(http.StatusOK, nil)
	})
    e.GET("/todos", func(c echo.Context) error {
		list :=  ListTodo()
		return c.JSON(http.StatusOK, list)
	})
	e.Logger.Fatal(e.Start(":1323"))
}

type todo struct {
	Topic string
}

func NewTodo(s string) {
	todoList = append(todoList, todo{Topic: s})
}

func ListTodo() []todo{
	return todoList
}
