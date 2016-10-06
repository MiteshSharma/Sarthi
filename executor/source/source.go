package source

import (
	"github.com/MiteshSharma/Sarthi/executor/work"
	"github.com/MiteshSharma/Sarthi/utils"
	"github.com/MiteshSharma/Sarthi/executor/logs"
)

type Source interface  {
	Get() work.Work
}

var source Source

func GetSource() Source {
	if source != nil {
		return source
	} else {
		config := utils.ConfigParam.ExecutorConfig
		sourceType, ok := config["mq"].(string)
		if !ok {
			logs.Logger.Error("No source type found to read data for execution.")
		}
		switch sourceType {
		case "mq":
			source = NewMqSource()
		default:
			// Default is Mq
			source = NewMqSource()
		}

		return source
	}
}