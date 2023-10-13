package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		log.Println("Error al abrir la base de datos:", err)
		return err
	}

	// Comprobar si la tabla "tasks" existe, y si no, crearla
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		done BOOLEAN
	)`)
	if err != nil {
		log.Println("Error al crear la tabla 'tasks':", err)
		return err
	}

	DB = db
	log.Println("Conexión a la base de datos establecida correctamente.")
	return nil
}

type Task struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"status"`
}

func GetTasks() ([]Task, error) {

	rows, err := DB.Query("SELECT * FROM tasks")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := make([]Task, 0)

	for rows.Next() {
		singleTask := Task{}
		err = rows.Scan(&singleTask.Id, &singleTask.Title, &singleTask.Done)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, singleTask)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTaskByID(taskID int) (Task, error) {
	var task Task

	err := DB.QueryRow("SELECT id, title, done FROM tasks WHERE id = ?", taskID).
		Scan(&task.Id, &task.Title, &task.Done)

	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func CreateTask(title string) (int, error) {
	stmt, err := DB.Prepare("INSERT INTO tasks(title, done) VALUES(?, ?)")
	if err != nil {
		log.Println("Error al preparar la declaración SQL:", err)
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(title, false)
	if err != nil {
		log.Println("Error al ejecutar la declaración SQL:", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener el último ID insertado:", err)
		return 0, err
	}

	log.Printf("Registro creado con éxito. ID: %d, Título: %s\n", id, title)
	return int(id), nil
}

func UpdateTask(taskID int, done bool) error {
	stmt, err := DB.Prepare("UPDATE tasks SET done = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(done, taskID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTask(taskID int) error {
	stmt, err := DB.Prepare("DELETE FROM tasks WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(taskID)
	if err != nil {
		return err
	}

	return nil
}
