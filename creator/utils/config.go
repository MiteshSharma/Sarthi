package utils

import (
	"os"
	"encoding/json"
	"path/filepath"
	"github.com/MiteshSharma/Sarthi/creator/model"
)

var Config *model.Config = &model.Config{}

func findConfigFile(fileName string) string {
	if _, error:= os.Stat("./"+fileName); error == nil {
		fileName,_ = filepath.Abs("./" + fileName)
	} else if _, error:= os.Stat("./"+fileName); error == nil {
		fileName,_ = filepath.Abs("./config/" + fileName)
	} else if _, error:= os.Stat("./creator/config/"+fileName); error == nil {
		fileName,_ = filepath.Abs("./creator/config/" + fileName)
	}
	return fileName;
}

func LoadConfig(fileName string)  {
	filePath := findConfigFile(fileName)

	file, error := os.Open(filePath)

	if error != nil {
		panic("Error occured during config file reading "+error.Error())
	}

	jsonParser := json.NewDecoder(file)

	config := model.Config{};

	if jsonErr := jsonParser.Decode(&config); jsonErr != nil {
		panic("Json parsing error"+ jsonErr.Error())
	}

	config.SaveDefaultConfigParams()

	Config = &config;
}
