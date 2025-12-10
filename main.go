package main

func main() {
	todos := Todos{}

	todos.add("Complete the project")
	todos.add("Review code")
	todos.add("Write documentation")


    todos.toggle(1)

    todos.print()
    
    


    todos.validateIndex(2)
}