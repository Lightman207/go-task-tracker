# Go Task Tracker

A simple command-line task tracker built with **Go**.

This project was created as a practical exercise to learn Go fundamentals, work with JSON persistence, handle files, and build a small CLI application with a clean and understandable structure.

**Project:** [https://roadmap.sh/projects/task-tracker](https://roadmap.sh/projects/task-tracker/solutions?u=6708f8e8fb4be684db21ad1e)

---

## ✨ Features

- ✅ Create tasks
- 📋 View all tasks
- ✏️ Update tasks
- 🗑️ Delete tasks
- 🔎 Search tasks by title
- 🏷️ Filter tasks by status
- 🚦 Set task priority
- 🕐 Track creation and update timestamps
- 💾 Persist tasks using JSON
- 🔢 Maintain unique task IDs after application restart

---

## 🛠️ Tech Stack

- **Go**
- Go Standard Library
- `bufio`
- `encoding/json`
- `os`
- `strconv`
- `strings`
- `time`

No external dependencies are required.

---

## 📦 Project Structure

    go-task-tracker/
    │
    ├── main.go
    ├── task.go
    ├── task_service.go
    ├── input.go
    ├── storage.go
    ├── tasks.json
    ├── go.mod
    ├── .gitignore
    └── README.md

### Files

| File | Purpose |
|---|---|
| `main.go` | CLI menu and application entry point |
| `task.go` | Task model |
| `task_service.go` | Task creation, update, deletion, search and filtering |
| `input.go` | User input and input parsing |
| `storage.go` | JSON file storage |
| `tasks.json` | Persistent task data |

---

## 🚀 Getting Started

### Requirements

- Go 1.22 or newer

Check your installed version:

    go version

### Clone the repository

    git clone https://github.com/YOUR_USERNAME/go-task-tracker.git
    cd go-task-tracker

### Run

    go run .

---

## 🖥️ CLI

The application provides a simple interactive menu:

    __ Roadmap __
    1. Добавить задачу
    2. Показать задачи
    3. Изменить задачу
    4. Удалить задачу
    5. Поиск задачи
    6. Фильтр задач
    7. Выйти

### Task example

    ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
    ID: 1
    Название: Изучить Go
    Описание: Разобраться с JSON
    Статус: IN PROGRESS
    Приоритет: HIGH
    Создана: 27.09.2026 11:40
    Изменена: 27.09.2026 12:15
    ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

---

## 💾 Data Persistence

Tasks are stored in a local `tasks.json` file.

When the application starts:

    tasks.json
        ↓
    ReadFile
        ↓
    JSON
        ↓
    []Task

When a task is created, updated, or deleted:

    []Task
        ↓
    JSON
        ↓
    WriteFile
        ↓
    tasks.json

Example stored data:

    [
      {
        "ID": 1,
        "Title": "Learn Go",
        "Description": "Study JSON and file storage",
        "Status": "IN PROGRESS",
        "Priority": "HIGH",
        "CreatedAt": "2026-09-27T11:40:00+03:00",
        "UpdatedAt": "2026-09-27T12:15:00+03:00"
      }
    ]

---

## 🧩 Task Model

Each task contains:

| Field | Description |
|---|---|
| `ID` | Unique task identifier |
| `Title` | Task title |
| `Description` | Detailed task description |
| `Status` | `TODO`, `IN PROGRESS`, or `DONE` |
| `Priority` | `LOW`, `MEDIUM`, or `HIGH` |
| `CreatedAt` | Creation timestamp |
| `UpdatedAt` | Last update timestamp |

---

## 🏗️ Architecture

The project intentionally keeps the architecture simple.

                    ┌─────────────┐
                    │   main.go   │
                    │   CLI/Menu  │
                    └──────┬──────┘
                           │
                           ▼
                  ┌─────────────────┐
                  │ task_service.go │
                  │                 │
                  │ Add / Update    │
                  │ Delete / Search │
                  │ Filter          │
                  └───────┬─────────┘
                          │
              ┌───────────┴───────────┐
              ▼                       ▼
       ┌─────────────┐        ┌─────────────┐
       │  input.go   │        │ storage.go  │
       │             │        │             │
       │ User Input  │        │ JSON / File │
       └─────────────┘        └──────┬──────┘
                                     │
                                     ▼
                              ┌─────────────┐
                              │ tasks.json  │
                              └─────────────┘

The application avoids unnecessary abstractions and keeps responsibilities separated between CLI logic, task operations, input handling, and persistence.

---

## 📚 What I Practiced

This project helped me practice:

- Go modules
- Structs
- Slices
- Functions
- Multiple return values
- Error handling
- `bufio.Scanner`
- `strconv.Atoi`
- String operations
- Working with `time.Time`
- JSON serialization with `encoding/json`
- JSON deserialization
- Reading and writing files
- Searching and filtering slices
- Separating code across multiple files
- Basic application architecture
- Persistent local storage

---

## 🗺️ Roadmap

- [x] Create tasks
- [x] View tasks
- [x] Update tasks
- [x] Delete tasks
- [x] Search tasks
- [x] Filter tasks by status
- [x] Add task priorities
- [x] Add creation/update timestamps
- [x] Add JSON persistence
- [ ] Sort tasks by priority
- [ ] Sort tasks by date
- [ ] Filter tasks by priority
- [ ] Combine multiple filters
- [ ] Add automated tests
- [ ] Improve CLI formatting
- [ ] Improve error handling

---

## 🎯 Project Goal

The main goal of this project is to gain practical experience with Go by building a small application from scratch and gradually improving its structure and functionality.

The project follows the **Task Tracker** challenge from roadmap.sh:

[https://roadmap.sh/projects/task-tracker](https://roadmap.sh/projects/task-tracker/solutions?u=6708f8e8fb4be684db21ad1e)

---
