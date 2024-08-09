package model

import (
	"database/sql"
	"todo-app-go/internal/db"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func GetAllTodos() ([]Todo, error) {
	rows, err := db.DB.Query("SELECT id, title, completed FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func GetTodoByID(id int) (Todo, error) {
	row := db.DB.QueryRow("SELECT id, title, completed FROM todos WHERE id = ?", id)
	var todo Todo
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
		if err == sql.ErrNoRows {
			return todo, nil
		}
		return todo, err
	}
	return todo, nil
}

func CreateTodo(title string) (int64, error) {
	result, err := db.DB.Exec("INSERT INTO todos (title) VALUES (?)", title)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func UpdateTodo(id int, title string, completed bool) error {
	_, err := db.DB.Exec("UPDATE todos SET title = ?, completed = ? WHERE id = ?", title, completed, id)
	return err
}

func DeleteTodo(id int) error {
	_, err := db.DB.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}
