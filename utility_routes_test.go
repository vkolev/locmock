package locmock

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"regexp"
	"runtime"
	"strings"
	"testing"
)

func TestPingRoute(t *testing.T) {
	t.Parallel()
	_, filename, _, _ := runtime.Caller(0)
	dataPath := strings.Replace(filename, "utility_routes_test.go", "data", 1)
	config := Config{
		DataPath: dataPath,
		Port:     ":8080",
	}
	router := setupRouter(config)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/l/ping", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestIp(t *testing.T) {
	t.Parallel()
	_, filename, _, _ := runtime.Caller(0)
	dataPath := strings.Replace(filename, "utility_routes_test.go", "data", 1)
	config := Config{
		DataPath: dataPath,
		Port:     ":8080",
	}
	router := setupRouter(config)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/l/ip", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestPersonProfile(t *testing.T) {
	t.Parallel()
	_, filename, _, _ := runtime.Caller(0)
	dataPath := strings.Replace(filename, "utility_routes_test.go", "data", 1)
	config := Config{
		DataPath: dataPath,
		Port:     ":8080",
	}
	router := setupRouter(config)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/l/person?gender=male", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.True(t, strings.Contains(w.Body.String(), "\"gender\": \"male\""))
	req, _ = http.NewRequest(http.MethodGet, "/l/person?gender=female", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.True(t, strings.Contains(w.Body.String(), "\"gender\": \"female\""))
}

func TestUserAgent(t *testing.T) {
	t.Parallel()
	_, filename, _, _ := runtime.Caller(0)
	dataPath := strings.Replace(filename, "utility_routes_test.go", "data", 1)
	config := Config{
		DataPath: dataPath,
		Port:     ":8080",
	}
	router := setupRouter(config)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/l/user-agent", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestUuid(t *testing.T) {
	t.Parallel()
	_, filename, _, _ := runtime.Caller(0)
	dataPath := strings.Replace(filename, "utility_routes_test.go", "data", 1)
	config := Config{
		DataPath: dataPath,
		Port:     ":8080",
	}
	router := setupRouter(config)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/l/uuid", nil)
	router.ServeHTTP(w, req)
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	assert.Equal(t, 200, w.Code)
	//fmt.Print(w.Body.String())
	assert.True(t, r.MatchString(w.Body.String()))
}
