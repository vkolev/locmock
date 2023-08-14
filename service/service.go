package service

import (
	"github.com/vkolev/locmock/action"
	"os"
	"path/filepath"
)

type Service struct {
	Name    string
	Actions map[string]action.Action
}

func NewFromPath(path string) (Service, error) {
	pathName := filepath.Base(path)
	service := Service{
		Name: pathName,
	}
	serviceActions := make(map[string]action.Action)
	files, err := os.ReadDir(path)
	if err != nil {
		return Service{}, err
	}
	for _, file := range files {
		if file.IsDir() {
			a, err := action.LoadAction(path, file.Name())
			if err != nil {
				return Service{}, nil
				break
			}
			serviceActions[a.Name] = a
		}
	}
	service.Actions = serviceActions
	return service, nil
}

func GetServices(rootPath string) ([]Service, error) {
	return make([]Service, 0), nil
}
