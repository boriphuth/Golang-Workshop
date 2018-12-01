package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

var todoList = []todo{}

func main1() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/todo", func(c echo.Context) error {

		t := &todo{}
		if err := c.Bind(t); err != nil {
			return err
		}
		NewTodo(t.Topic)
		return c.JSON(http.StatusOK, nil)
	})

	e.GET("/todos", func(c echo.Context) error {
		list := ListTodo()
		return c.JSON(http.StatusOK, list)
	})

	e.PUT("/todo", func(c echo.Context) error {

		t := todo{}
		if err := c.Bind(&t); err != nil {
			return err
		}
		UpdateTodo(t)
		return c.JSON(http.StatusOK, t)
	})

	// Routes
	// e.POST("/users", createUser)
	// e.GET("/users/:id", getUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}

type todo struct {
	ID    string `json:"id"`
	Topic string `json:"topic"`
	Done  bool   `json:"done"`
}

func NewTodo(s string) {
	todoList = append(todoList, todo{ID: uuid.New().String(), Topic: s})
}

func ListTodo() []todo {
	return todoList
}

func UpdateTodo(t todo) []todo {
	for i := range todoList {
		if todoList[i].ID == t.ID {
			// Found!
			todoList[i] = t
		}
	}
	return todoList
}
