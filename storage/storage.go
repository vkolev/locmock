package storage

import (
	"github.com/spf13/afero"
	"io/fs"
	"path/filepath"
)

type Storage struct {
	Fs afero.Afero
}

func NewOsStorage(fsType afero.Afero) Storage {
	return Storage{Fs: fsType}
}

func (s *Storage) GetServiceNames(baseDir string) []string {
	elements, err := s.Fs.ReadDir(baseDir)
	if err != nil {
		panic(err)
	}
	var response []string
	for _, element := range elements {
		if element.IsDir() {
			response = append(response, element.Name())
		}
	}
	return response
}

func hasActionFile(path string, s Storage) (bool, error) {
	return s.Fs.Exists(filepath.Join(path, "action.yml"))
}

func (s *Storage) GetActionsForService(baseDir, serviceName string) []string {
	var response []string
	_ = s.Fs.Walk(filepath.Join(baseDir, serviceName), func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		if info.IsDir() {
			if ok, _ := hasActionFile(path, *s); ok == true {
				response = append(response, path[len(serviceName)+2:])
			}
		}
		return nil
	})
	return response
}

func (s *Storage) CreateActionFile(actionFilePath string, data []byte) error {
	file, err := s.Fs.Create(actionFilePath)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetActionData(actionFilePath string) ([]byte, error) {
	return s.Fs.ReadFile(actionFilePath)
}
