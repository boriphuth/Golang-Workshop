package main

import (
	"reflect"
	"testing"
)

func TestNewTodo(t *testing.T) {
	todoList = []todo{}
	NewTodo("Learn golang")

	if todoList[0].Topic != "Learn golang" {
		t.Error("it shoud store Learn golang but get", todoList[0].Topic)
	}
}

func TestListTodo(t *testing.T) {
	todoList = []todo{}
	NewTodo("Learn golang")
	NewTodo("Learn git")
	NewTodo("Learn VueJS")

	listTodo := ListTodo()

	expected := []todo{
		{Topic: "Learn golang"},
		{Topic: "Learn git"},
		{Topic: "Learn VueJS"},
	}

	if !reflect.DeepEqual(listTodo, expected) {
		t.Errorf("%v\nis expected but get\n%v\n", expected, listTodo)
	}

}
