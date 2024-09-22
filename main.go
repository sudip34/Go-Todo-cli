package main

import "fmt"

func main() {
	fmt.Println("Inside main in todo app")
	todos := Todos{}
	todos.add("Buy milk")
	todos.add("Buy bread")
	todos.add("Buy benana")

	fmt.Printf("%+v\n\n", todos)

	todos.delete(0)
	fmt.Printf("%+v", todos)
}
