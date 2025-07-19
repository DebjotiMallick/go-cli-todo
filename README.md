# 📝 CLI To-Do List App (Built with Go)

A simple, fast, and persistent command-line To-Do application written in Go.  
Manage your tasks right from the terminal — add, complete, delete, and list your tasks with ease.

---

## ⚙️ Features

- ✅ Add new tasks
- 📋 List all tasks
- 🗑️ Delete tasks by number
- ✔️ Mark tasks as done
- 🧹 Clear completed or all tasks
- 📊 Task summary via `status`
- 💾 Data stored in `tasks.json` (local file-based persistence)

---

## 🚀 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/yourusername/todo-cli.git
cd todo-cli
```


### 2. Build the binary
```bash
go build -o todo
```
### 3. Run commands
```bash
./todo add "Buy milk"
./todo list
./todo done 1
./todo delete 2
./todo clear --done
./todo status
```

### 💡 Command Reference
| Command              | Description                        |
| -------------------- | ---------------------------------- |
| todo add "task"      | Add a new task                     |
| todo list            | List all tasks                     |
| todo done [number]   | Mark a task as completed           |
| todo delete [number] | Delete a task by number            |
| todo clear --done    | Clear only completed tasks         |
| todo clear --all     | Clear all tasks explicitly         |
| todo status          | Show total, pending, and completed |

### 📁 Data Storage
All tasks are saved to a local tasks.json file in the project folder, so your task list persists between sessions.

### 🛠 Built With
Go — lightning-fast CLI development

### 📜 License
MIT License

### 🙌 Author
Debjoti Mallick

### Give it a ⭐ if you liked it!