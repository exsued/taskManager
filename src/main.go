package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	//"taskmanager/routes"
)

var (
	Tasks = []string{"gimba", "banga", "bonga", "gimba", "banga", "bonga", "gimba", "banga", "bonga", "gimba", "banga", "bonga", "gimba", "banga", "bonga"}
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

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	argStr := r.URL.Path[len("/remove-task/"):]
	arg, err := strconv.Atoi(argStr)
	if err != nil {
		http.Error(w, "Такого индекса нет", http.StatusBadRequest)
		return
	}
	fmt.Println("Фишки работают!")
	Tasks = append(Tasks[:arg], Tasks[arg+1:]...)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", neuter(fs)))
	http.HandleFunc("/", Index)
	http.HandleFunc("/remove-task/", DeleteTask)
	log.Fatalln(http.ListenAndServe(":8080", nil))
	fmt.Println(os.Getwd())
}
