package source

import (
	"bytes"
	"encoding/json"
	"github.com/MiteshSharma/Sarthi/mq"
	"github.com/MiteshSharma/Sarthi/dao"
	"github.com/MiteshSharma/Sarthi/executor/work"
)

type MqSource struct {
}

func NewMqSource() *MqSource  {
	return &MqSource{}
}

func (mqs *MqSource) Get() work.Work  {
	mq_agent := mq.GetAgent()
	mqObject, err := mq_agent.Read()
	if err != nil {
		return &dao.Task{}
	}
	var task dao.Task
	decoder := json.NewDecoder(bytes.NewReader(mqObject.Message))
	if err := decoder.Decode(&task); err != nil {
		return &dao.Task{}
	}
	return &task
}
