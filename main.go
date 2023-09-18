package main

import (
	"net/http"

	"github.com/RamsesMartins/ToGoList/pkg/routes"
)

func main() {
	http.HandleFunc("/", routes.AllTasks)
	http.HandleFunc("/task", routes.Task)
	http.HandleFunc("/task/add", routes.AddTask)

	println("Run in: http://localhost:8030")
	http.ListenAndServe(":8030", nil)s
}
