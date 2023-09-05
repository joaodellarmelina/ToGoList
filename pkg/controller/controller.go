package controller

import (
	"github.com/RamsesMartins/ToGoList/pkg/database"
	"github.com/RamsesMartins/ToGoList/pkg/model"
)

var db = database.GetDB()

func GetTask( task *model.Task, id int) model.Task{
	db.First(task, id)

	return *task
}

func GetAllTasks(tasks *[]model.Task) []model.Task{
	db.Find(tasks,"status = ?", 0)
	return *tasks
}

func PostTask(task model.Task){
	db.Create(task)
}

func DoneTask(task *model.Task, id int){
	db.Model(task).Where("id = ?", id).Update("status",1)
}
