// Responsável por gerenciar as rotas da aplicação
package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RamsesMartins/ToGoList/pkg/database"
	"github.com/RamsesMartins/ToGoList/pkg/model"
)

var db = database.GetDB()

// Contém as rotas para CRUD da aplicação
func Tasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tasks_get(w, r)
	case "POST":
		tasks_post(w, r)
	case "DELETE":
		tasks_delete(w, r)
	}
}

func tasks_get(w http.ResponseWriter, r *http.Request) {
	id_str := r.URL.Query().Get("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		println(err)
		return
	} else {
		var task model.Task
		db.First(&task, id)
		w.Header().Set("Content-Type", "application/json")
		res, err := json.Marshal(&task)
		if err != nil {
			println(err)
			return
		} else {
			w.Write(res)
		}
		return
	}
}

func tasks_post(w http.ResponseWriter, r *http.Request) {
	var task model.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		println(err)
		return
	}
	db.Create(&task)

}

func tasks_delete(w http.ResponseWriter, r *http.Request) {
	id_str := r.URL.Query().Get("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		println(err)
		return
	}
	var task model.Task
	db.Delete(&task, id)
}
