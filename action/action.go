package action

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Action struct {
	Name             string                 `yaml:"name"`
	Method           string                 `yaml:"method"`
	Parameters       map[string]interface{} `yaml:"parameters"`
	ParametersConfig map[string]interface{} `yaml:"parameters_config"`
	Response         string                 `yaml:"response"`
	ResponseConfig   map[string]interface{} `yaml:"response_config"`
	ResponseType     string                 `yaml:"response_type"`
	Paginate         bool                   `yaml:"paginate"`
	ResponseStatus   int                    `yaml:"response_status"`
}

func LoadAction(servicePath string, name string) (Action, error) {
	yamlFile, err := os.Open(filepath.Join(servicePath, name, fmt.Sprintf("%v.yml", name)))
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

func (a *Action) RunAction() (int, any) {
	switch a.ResponseType {
	case "plain/text":
		return a.ResponseStatus, a.Response
	case "application/json":
		var jsonResponse map[string]interface{}
		json.Unmarshal([]byte(a.Response), &jsonResponse)
		return a.ResponseStatus, jsonResponse
	case "application/xml", "text/xml":
		return a.ResponseStatus, a.Response
	}
	return a.ResponseStatus, a.Response
}
