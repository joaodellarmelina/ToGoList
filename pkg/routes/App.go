package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/RamsesMartins/ToGoList/pkg/controller"
	"github.com/RamsesMartins/ToGoList/pkg/model"
)

/*
Handler para retornar todas as Tasks/Tarefas -
Method: GET
*/
func AllTasks(w http.ResponseWriter, r *http.Request) {
	//Validando Método
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Carregando o Template
	tmpl, err := template.ParseFiles("./pkg/views/index.html")
	if err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Inserindo os dados
	t := new([]model.Task)
	tasks := controller.GetAllTasks(t)

	//Executando Template
	err = tmpl.Execute(w, tasks)
	if err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//SUCESS
	w.WriteHeader(http.StatusOK)

}

/*
Handler para retornar uma Task/Tarefa de acordo com o parâmetro ID da URL -
Method: GET
*/
func Task(w http.ResponseWriter, r *http.Request) {

	// w.WriteHeader(http.StatusBadRequest)
	// return
	//Validando Método
	if r.Method == http.MethodGet {

		//Capturando o valor do parâmetro ID da URL e realizando a conversão para INT
		i := r.URL.Query().Get("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			println(err.Error())
			//Erro
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		//Carregando o template
		tmpl, err := template.ParseFiles("./pkg/views/task.html")
		if err != nil {
			println(err.Error())
			//Erro
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var t = new(model.Task)
		task := controller.GetTask(t, id)

		//Executando Template
		err = tmpl.Execute(w, task)
		if err != nil {
			println(err.Error())
			//Erro interno
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		//Sucess
		w.WriteHeader(http.StatusOK)

	} else if r.Method == "PUT" {

		//Capturando o valor do parâmetro ID da URL e realizando a conversão para INT
		i := r.URL.Query().Get("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			println(err.Error())
			//Erro Interno
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		//Alterando o Status da Task
		var t = new(model.Task)
		controller.DoneTask(t, id)

		//Success
		w.WriteHeader(http.StatusOK)

	}
}

/*
Handler para criar uma Task/Tarefa de acordo com o form -
Method: GET
*/
func AddTask(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var task = model.Task{}

		fmt.Println(r.Body)
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		controller.PostTask(task)

	case "GET":
		tmp, err := template.ParseFiles("./pkg/views/form.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = tmp.Execute(w, nil)
	}

}
