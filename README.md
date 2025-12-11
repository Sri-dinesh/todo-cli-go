# Todo CLI

A simple and efficient command-line todo application built with Go.

## Features

- ✅ Add new todo items
- ✅ List all todos in a formatted table
- ✅ Edit existing todos
- ✅ Toggle todo completion status
- ✅ Delete todos
- ✅ Persistent storage using JSON

## Installation

### Build from Source

```bash
# Clone the repository
git clone https://github.com/Sri-dinesh/todo-cli-go.git
cd todo-cli-go

# Install dependencies
go mod download

# Build the application
go build -o todo
```

## Usage

### Add a Todo

```bash
./todo -add "Buy groceries"
```

### List All Todos

```bash
./todo -list
```

This displays a formatted table with:
| Column | Description |
|--------|-------------|
| # | Todo index |
| Title | Todo description |
| Completed | ✅ or ❌ status |
| Created At | Creation timestamp |
| Completed At | Completion timestamp (if completed) |

### Edit a Todo

```bash
./todo -edit "0:Updated todo title"
```

Format: `index:new_title`

### Toggle Todo Status

```bash
./todo -toggle 0
```

Marks a todo as completed or uncompleted.

### Delete a Todo

```bash
./todo -del 0
```

Removes the todo at the specified index.

## Command Reference

| Flag      | Type   | Description                          |
| --------- | ------ | ------------------------------------ |
| `-add`    | string | Add a new todo item                  |
| `-list`   | bool   | List all todo items                  |
| `-edit`   | string | Edit a todo (format: `id:new_title`) |
| `-toggle` | int    | Toggle completion status by index    |
| `-del`    | int    | Delete a todo by index               |

## Architecture

### Components

- **main.go**: Initializes storage, loads todos, executes commands, and saves changes.
- **command.go**: Handles command-line argument parsing using Go's `flag` package.
- **todo.go**: Defines the `Todo` struct and `Todos` slice type with all CRUD operations.
- **storage.go**: Generic JSON storage implementation using Go generics.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
