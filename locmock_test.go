package locmock_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/vkolev/locmock"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestIsValidConfigExtension(t *testing.T) {
	t.Parallel()
	want := true
	got := locmock.IsValidConfigExtension(".yml")
	if want != got {
		t.Errorf("want %t, got %t for .yml", want, got)
	}
	got = locmock.IsValidConfigExtension(".yaml")
	if want != got {
		t.Errorf("want %t, got %t for .yaml", want, got)
	}
}

func TestIsValidConfigExtensionReturnsFalseForConf(t *testing.T) {
	t.Parallel()
	want := false
	got := locmock.IsValidConfigExtension(".conf")
	if want != got {
		t.Errorf("Got true for invalid extension .conf")
	}
}

func TestLoadConfig(t *testing.T) {
	t.Parallel()
	_, filename, _, _ := runtime.Caller(0)
	dataPath := strings.Replace(filename, "locmock_test.go", "data", 1)
	configFilePath := filepath.Join(dataPath, "locmock.yml")
	want := locmock.Config{
		DataPath: "/Users/vladi/GolandProjects/locmock/data/",
		Port:     ":8080",
	}
	got, err := locmock.LoadConfig(configFilePath)
	if err != nil {
		t.Errorf("Want configuration, got error %v", err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestLoadConfigInvalidPathCreatesNewFile(t *testing.T) {
	t.Parallel()
	_, filename, _, _ := runtime.Caller(0)
	dataPath := strings.Replace(filename, "locmock_test.go", "data", 1)
	configFilePath := filepath.Join(dataPath, "locmock_test.yml")
	want := locmock.Config{
		DataPath: "/Users/vladi/GolandProjects/locmock/data",
		Port:     ":8080",
	}
	got, err := locmock.LoadConfig(configFilePath)
	if err != nil {
		t.Errorf("Want configuration, got error %v", err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
	_ = os.Remove(configFilePath)
}
