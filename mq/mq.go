package mq

import (
	"fmt"

	"github.com/MiteshSharma/Sarthi/errors"
	"github.com/MiteshSharma/Sarthi/utils"
)

type MqAgent interface {
	Read() (*MqObject, error)
	Write([]byte) error
	Delete(string) error
}

type MqObject struct {
	Id      string
	Message []byte
}

var agent MqAgent

func GetAgent() MqAgent {
	if agent != nil {
		return agent
	} else {
		config := utils.ConfigParam.ExecutorConfig
		mqType, ok := config["mq"].(string)
		if !ok {
			panic(&errors.MqNotConfiguredError{})
		}

		var err error
		switch mqType {
		case "sqs":
			agent, err = NewSqsAgent()
		//case "kafka":
		// TODO
		default:
			panic(&errors.MqTypeNotSupportedError{
				Type: mqType,
			})
		}

		if err != nil {
			fmt.Println("Error initializing message queue agent.")
			panic(err)
		}

		return agent
	}
}
