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
	service, err := service.NewFromPath(servicePath)
	if err != nil {
		t.Errorf("Unable to load service: %v", err)
	}
	if service.Name != serviceName {
		t.Errorf("want: %v, got %v", serviceName, service.Name)
	}
	wantLen := 1
	gotLen := len(service.Actions)
	if wantLen != gotLen {
		t.Errorf("want len %d, got len %d", wantLen, gotLen)
	}
}
