package api

import (
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/MiteshSharma/Sarthi/dao"
)

func InitTask(router *httprouter.Router)  {
	router.GET("/tasks", getTask)
	router.POST("/tasks", createTask)
}

func getTask(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tasks := ServerObj.Repository.Task().Get()
	js, err := json.Marshal(tasks)
	if err != nil {
		rw.Write([]byte("{}"))
	} else {
		rw.Write([]byte(string(js[:])))
	}
	rw.Header().Set("Content-Type", "application/json")
}

func createTask(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var task dao.Task
	if err := json.Unmarshal([]byte(r.FormValue("task")), &task); err != nil {
		rw.Write([]byte("{}"))
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write([]byte("{}"))
}