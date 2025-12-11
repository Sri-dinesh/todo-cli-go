package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *cmdFlags {
	cf := cmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo item")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo item by index & a new Todo Title, id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Delete a todo item by index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a todo item by index")
	flag.BoolVar(&cf.List, "list", false, "List all todo items")


	flag.Parse()

	return &cf
}


func (cf *cmdFlags) Execute (todos *Todos) {

	switch {

	case cf.List:
		todos.print()
		
	case cf.Add != "":
		todos.add(cf.Add)

	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid format for edit. Use id:new_title")
			os.Exit(1)
		}
		
		index, err := strconv.Atoi(parts[0])
		
		if err != nil {
			fmt.Println("Invalid index for edit. Use id:new_title")
			os.Exit(1)
		}
		
		todos.edit(index, parts[1])
		
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
		
	case cf.Del != -1:
		todos.delete(cf.Del)
		
	default:
		fmt.Println("Invalid Command")
		
	}
}