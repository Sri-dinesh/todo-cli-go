package main

func main() {
	todos := Todos{}

	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)

	todos.add("Complete the project")
	// todos.add("Review code")
	todos.add("Write documentation")

	todos.toggle(1)

	todos.print()

	storage.save(todos)

	// todos.validateIndex(2)

}