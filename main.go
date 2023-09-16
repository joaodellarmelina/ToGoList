package main

import (
	"net/http"

	"github.com/RamsesMartins/ToGoList/pkg/routes"
)

func main() {
	http.HandleFunc("/", routes.AllTasks)
	http.HandleFunc("/task", routes.Task)
	http.HandleFunc("/task/add", routes.AddTask)
	http.HandleFunc("/api/todo", routes.Tasks)

	println("Só rodar baby")
	http.ListenAndServe(":8030", nil)
}
