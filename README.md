# todo-cli

A simple command-line interface (CLI) task manager written in Go. Stores data locally in `tasks.yaml`.

## Binaries Download

Download the executable for your OS:

- [Windows (64-bit)](https://github.com/tori-sst/todo-cli/releases/latest/download/todo.exe)
- [Linux (64-bit)](https://github.com/tori-sst/todo-cli/releases/latest/download/todo_linux)


## Usage

Add a task:
```bash
./todo -add "Task title"
```
List all tasks:
```bash
./todo -list "all"
```
Filter tasks by status (todo, doing, done):
```bash
./todo -list "status"
```
Update task status:
```bash
./todo -status ID -to "status"
```
Delete a task:
```bash
./todo -delete ID
```
