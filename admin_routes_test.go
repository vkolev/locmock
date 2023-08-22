package locmock

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"runtime"
	"strings"
	"testing"
)

func TestListServicesFailsWithWrongConfig(t *testing.T) {
	t.Parallel()
	_, filename, _, _ := runtime.Caller(0)
	dataPath := strings.Replace(filename, "the-most-non-existing-path-ever-123654", "data", 1)
	config := Config{
		DataPath: dataPath,
		Port:     ":8080",
	}
	router := setupRouter(config)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/admin/services", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Code)
}

func TestListServices(t *testing.T) {
	t.Parallel()
	_, filename, _, _ := runtime.Caller(0)
	dataPath := strings.Replace(filename, "admin_routes_test.go", "data", 1)
	config := Config{
		DataPath: dataPath,
		Port:     ":8080",
	}
	router := setupRouter(config)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/admin/services", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}
