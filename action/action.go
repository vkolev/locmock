package action

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Action struct {
	Name             string                 `yaml:"name" json:"name"`
	Description      string                 `yaml:"description" json:"description"`
	Method           string                 `yaml:"method" json:"method"`
	Parameters       map[string]interface{} `yaml:"parameters" json:"parameters"`
	ParametersConfig map[string]interface{} `yaml:"parameters_config" json:"parameters_config"`
	Response         string                 `yaml:"response" json:"response"`
	ResponseConfig   map[string]interface{} `yaml:"response_config" json:"response_config"`
	ResponseType     string                 `yaml:"response_type" json:"response_type"`
	Paginate         bool                   `yaml:"paginate" json:"paginate"`
	ResponseStatus   int                    `yaml:"response_status" json:"response_status"`
}

func LoadAction(servicePath string, name string) (Action, error) {
	yamlFile, err := os.Open(filepath.Join(servicePath, name, fmt.Sprintf("action.yml")))
	if err != nil {
		return Action{}, err
	}
	defer yamlFile.Close()

	var data []byte
	data, err = os.ReadFile(yamlFile.Name())

	if err != nil {
		return Action{}, err
	}

	config := make(map[string]Action)
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Action{}, err
	}
	return config["action"], nil
}

func createActionFile(path string, newAction Action) string {
	dirPath := filepath.Join(path, newAction.Name)
	_ = os.MkdirAll(dirPath, 0750)
	filePath := filepath.Join(dirPath, "action.yml")
	_, _ = os.Create(filePath)
	return filePath
}

func CreateAction(servicePath string, newAction Action) (Action, error) {
	actionFilePath := createActionFile(servicePath, newAction)
	yamlFile, err := os.OpenFile(actionFilePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return Action{}, err
	}
	defer yamlFile.Close()
	config := make(map[string]Action)
	config["action"] = newAction
	yamlData, err := yaml.Marshal(&config)
	if err != nil {
		return Action{}, err
	}
	_, err = yamlFile.Write(yamlData)
	if err != nil {
		return Action{}, err
	}
	return newAction, nil
}

func (a *Action) RunAction() (int, any) {
	statusCode := a.ResponseStatus
	var response any
	switch a.ResponseType {
	case "plain/text":
		response = a.Response
	case "application/json":
		var jsonResponse map[string]interface{}
		_ = json.Unmarshal([]byte(a.Response), &jsonResponse)
		response = jsonResponse
	case "application/xml", "text/xml":
		response = a.Response
	}
	return statusCode, response
}
