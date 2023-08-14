package action_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/vkolev/locmock/action"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestLoadAction(t *testing.T) {
	t.Parallel()
	_, filename, _, _ := runtime.Caller(0)
	dataPath := strings.Replace(filename, "action/action_test.go", "data", 1)
	serviceName := "test"
	servicePath := filepath.Join(dataPath, serviceName)
	want := action.Action{
		Name:   "hello",
		Method: "GET",
		Parameters: map[string]interface{}{
			"q": "test",
		},
		ParametersConfig: map[string]interface{}{
			"q": map[string]interface{}{
				"type":     "string",
				"required": false,
			},
		},
		Response:       "hello test",
		ResponseConfig: nil,
		ResponseType:   "plain/text",
		Paginate:       false,
		ResponseStatus: 200,
	}
	got, err := action.LoadAction(servicePath, "hello")
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if !cmp.Equal(got, want) {
		t.Errorf(cmp.Diff(want, got))
	}
}
