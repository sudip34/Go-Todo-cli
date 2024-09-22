package main

import "fmt"

func main() {
	fmt.Println("Inside main in todo app")
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	// load the file
	storage.Load(&todos)
	todos.add("Buy milk")
	todos.add("Buy bread")
	todos.add("Buy benana")
	todos.toggle(0)
	todos.print()
	storage.save(todos)
}
