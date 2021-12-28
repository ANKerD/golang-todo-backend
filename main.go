package main

import "fmt"

type Todo struct {
	title string
	content string
}

var todos []Todo

func list() []Todo {
	todos = append(todos, todos[0], todos[0])
	return todos
}

func main() { 
	todos = []Todo{
		{title: "oi", content:  "hello warudo"},
	}
	fmt.Println(list())
}