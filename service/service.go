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

type ServiceResponse struct {
	Name         string
	ActionsCount int
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

func filterDirectories(files []os.DirEntry) []os.DirEntry {
	filtered := make([]os.DirEntry, 0)
	for _, file := range files {
		if file.IsDir() {
			filtered = append(filtered, file)
		}
	}
	return filtered
}

func GetServices(rootPath string) ([]ServiceResponse, error) {
	files, err := os.ReadDir(rootPath)
	if err != nil {
		return make([]ServiceResponse, 0), err
	}

	var response []ServiceResponse

	for _, file := range files {
		if file.IsDir() {
			serviceFiles, _ := os.ReadDir(filepath.Join(rootPath, file.Name()))
			serviceFiles = filterDirectories(serviceFiles)
			response = append(response, ServiceResponse{
				Name:         filepath.Base(file.Name()),
				ActionsCount: len(serviceFiles),
			})
		}
	}

	return response, nil
}
