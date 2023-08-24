package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	//"taskmanager/routes"
)

type Task struct {
	ID               int
	Title            string
	Description      string
	Priority         int
	Background       bool
	NotificationTime string
	Parent           *Task
	Children         []*Task // Список дочерних задач
}

var (
	Tasks = []Task{
		{
			ID:               0,
			Title:            "Задача 1",
			Description:      "Описание задачи 1",
			Priority:         3,
			Background:       false,
			NotificationTime: "10:00",
		},
		{
			ID:               1,
			Title:            "Подзадача 1.2",
			Description:      "Описание подзадачи 1.2",
			Priority:         3,
			Background:       false,
			NotificationTime: "10:00",
		},
	}
)

func neuter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") || r.URL.Path == "" {
			http.NotFound(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	Tasks[0].Children = append(Tasks[0].Children, &Tasks[1])
	Tasks[1].Parent = &Tasks[0]

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", neuter(fs)))
	http.HandleFunc("/", Index)
	http.HandleFunc("/remove-task/", DeleteTask)
	http.HandleFunc("/update-task/", UpdateTask)
	http.HandleFunc("/complete-task/", CompleteTask)
	log.Fatalln(http.ListenAndServe(":8080", nil))
	fmt.Println(os.Getwd())
}
