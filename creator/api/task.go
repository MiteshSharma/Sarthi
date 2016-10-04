package api

import (
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/MiteshSharma/Sarthi/dao"
	"github.com/MiteshSharma/Sarthi/utils"
	"github.com/MiteshSharma/Sarthi/creator/service"
)

type TaskResponse struct  {
	Task dao.Task
	Error string
}

func NewTaskResponse(task dao.Task, error string) TaskResponse  {
	taskResponse := &TaskResponse{}
	taskResponse.Task = task
	taskResponse.Error = error
	return *taskResponse
}

func InitTask(router *httprouter.Router) {
	router.GET("/tasks", getTask)
	router.POST("/tasks", createTasks)
}

func getTask(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryValues := r.URL.Query()
	taskId := queryValues.Get("taskId")
	if taskId == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("Incorrect task key received.")))
		return
	}

	var task dao.Task
	var err error
	if task,err = service.GetTask(taskId); err != nil {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(utils.ToJson(NewTaskResponse(task, "No task found for this key. Please verify this key."))))
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(utils.ToJson(NewTaskResponse(task, ""))))
}

func createTasks(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var tasks []dao.Task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tasks); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("Incorrect tasks received.")))
		return
	}

	taskResponses := make([]TaskResponse, len(tasks))

	for index:= 0; index < len(tasks); index++ {
		task := tasks[index]
		if err := task.IsValid(); err != nil {
			taskResponses[index] = NewTaskResponse(task, err.Error())
		} else if err := service.CreateTask(&task); err != nil {
			taskResponses[index] = NewTaskResponse(task, "Error happened during creating task object.")
		} else {
			taskResponses[index] = NewTaskResponse(task, "")
		}
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(utils.ToJson(taskResponses)))
}