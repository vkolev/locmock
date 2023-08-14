package service

import (
	"github.com/vkolev/locmock/action"
)

type Service struct {
	Name    string
	Actions []action.Action
}

func (s Service) NewFromPath(path string) (Service, error) {
	return Service{}, nil
}
