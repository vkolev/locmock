package locmock

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
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
	assert.True(t, r.MatchString(w.Body.String()))
}

func TestGetUserAgent(t *testing.T) {
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
	want := "Testing User Agent"
	req.Header.Set("User-Agent", want)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, want, w.Body.String())
}

func TestRedirectRequest(t *testing.T) {
	t.Parallel()
	_, filename, _, _ := runtime.Caller(0)
	dataPath := strings.Replace(filename, "utility_routes_test.go", "data", 1)
	config := Config{
		DataPath: dataPath,
		Port:     ":8080",
	}
	router := setupRouter(config)
	type testCase struct {
		arg1 string
		arg2 int
	}
	var testCases = []testCase{
		{"301", 301},
		{"302", 302},
		{"303", 303},
		{"304", 304},
		{"305", 305},
		{"307", 307},
	}
	for _, test := range testCases {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/l/redirect?status=%v", test.arg1), nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, test.arg2, w.Code)
	}
}

func TestRedirectRequestFailesForWrongCode(t *testing.T) {
	t.Parallel()
	_, filename, _, _ := runtime.Caller(0)
	dataPath := strings.Replace(filename, "utility_routes_test.go", "data", 1)
	config := Config{
		DataPath: dataPath,
		Port:     ":8080",
	}
	router := setupRouter(config)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/l/redirect?status=309", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "Status code 309 is not a redirect code", w.Body.String())
}

func TestGetHeaders(t *testing.T) {
	t.Parallel()
	_, filename, _, _ := runtime.Caller(0)
	dataPath := strings.Replace(filename, "utility_routes_test.go", "data", 1)
	config := Config{
		DataPath: dataPath,
		Port:     ":8080",
	}
	router := setupRouter(config)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/l/headers", nil)
	want := "testing 1 2 3"
	req.Header.Set("testing", want)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), want)
}

func TestFormRequest(t *testing.T) {
	t.Parallel()
	_, filename, _, _ := runtime.Caller(0)
	dataPath := strings.Replace(filename, "utility_routes_test.go", "data", 1)
	config := Config{
		DataPath: dataPath,
		Port:     ":8080",
	}
	router := setupRouter(config)
	w := httptest.NewRecorder()
	tc := map[string]string{
		"post":  http.MethodPost,
		"patch": http.MethodPatch,
		"put":   http.MethodPut,
	}

	for _, wantMethod := range tc {
		req, _ := http.NewRequest(wantMethod, "/l/form", nil)
		wantKey := "test_key"
		wantValue := "test_value"
		values := url.Values{}
		values.Set(wantKey, wantValue)
		req.Form = values

		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
		assert.Contains(t, w.Body.String(), wantKey)
		assert.Contains(t, w.Body.String(), wantValue)
		assert.Contains(t, w.Body.String(), wantMethod)
	}
}

func TestFormRequestWithInvalidMethod(t *testing.T) {
	t.Parallel()
	_, filename, _, _ := runtime.Caller(0)
	dataPath := strings.Replace(filename, "utility_routes_test.go", "data", 1)
	config := Config{
		DataPath: dataPath,
		Port:     ":8080",
	}
	router := setupRouter(config)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/l/form", nil)
	wantKey := "test_key"
	wantValue := "test_value"
	values := url.Values{}
	values.Set(wantKey, wantValue)
	req.Form = values

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestWriteStream(t *testing.T) {
	t.Parallel()
	_, filename, _, _ := runtime.Caller(0)
	dataPath := strings.Replace(filename, "utility_routes_test.go", "data", 1)
	config := Config{
		DataPath: dataPath,
		Port:     ":8080",
	}
	router := setupRouter(config)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/l/stream/2", nil)
	want := "{\"key_0\":\"value_0\"}\n{\"key_1\":\"value_1\"}\n"

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, want, w.Body.String())
}
