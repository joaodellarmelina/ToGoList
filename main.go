package main

import (
	"net/http"

	"github.com/RamsesMartins/ToGoList/pkg/routes"
)

func main() {
	http.HandleFunc("/", routes.AllTasks)
	http.HandleFunc("/task", routes.Task)
	http.HandleFunc("/task/add", routes.AddTask)

	println("Teste")
	http.ListenAndServe(":8030", nil)
}
