package service_test

import (
	"github.com/vkolev/locmock/service"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestService_NewFromPath(t *testing.T) {
	t.Parallel()
	_, filename, _, _ := runtime.Caller(0)
	dataPath := strings.Replace(filename, "service/service_test.go", "data", 1)
	serviceName := "test"
	servicePath := filepath.Join(dataPath, serviceName)
	tService, err := service.NewFromPath(servicePath)
	if err != nil {
		t.Errorf("Unable to load service: %v", err)
	}
	if tService.Name != serviceName {
		t.Errorf("want: %v, got %v", serviceName, tService.Name)
	}
	wantLen := 2
	gotLen := len(tService.Actions)
	if wantLen != gotLen {
		t.Errorf("want len %d, got len %d", wantLen, gotLen)
	}
}

func TestGetServices(t *testing.T) {
	t.Parallel()
	_, filename, _, _ := runtime.Caller(0)
	dataPath := strings.Replace(filename, "service/service_test.go", "data", 1)
	services, err := service.GetServices(dataPath)
	if err != nil {
		t.Fatal(err)
	}
	wantLen := 1
	gotLen := len(services)
	if wantLen != gotLen {
		t.Errorf("want len %d, got len %d", wantLen, gotLen)
	}

}

func TestService_GetActions(t *testing.T) {
	t.Parallel()
	_, filename, _, _ := runtime.Caller(0)
	dataPath := strings.Replace(filename, "service/service_test.go", "data", 1)
	tService, err := service.NewFromPath(filepath.Join(dataPath, "test"))
	if err != nil {
		t.Fatal(err)
	}
	wantLen := 2
	actionsResponse := tService.GetActions()
	gotLen := len(actionsResponse.Actions)
	if wantLen != gotLen {
		t.Errorf("want len %d, got len %d", wantLen, gotLen)
	}

}
