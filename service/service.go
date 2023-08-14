package service

import (
	"github.com/vkolev/locmock/action"
	"os"
	"path/filepath"
)

type Service struct {
	Name    string
	Actions []action.Action
}

func NewFromPath(path string) (Service, error) {
	pathName := filepath.Base(path)
	service := Service{
		Name: pathName,
	}
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
			service.Actions = append(service.Actions, a)
		}
	}
	return service, nil
}
