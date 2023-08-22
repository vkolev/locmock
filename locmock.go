package locmock

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/vkolev/locmock/service"
	"gopkg.in/yaml.v3"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	DataPath string `yaml:"data_path"`
	Port     string `yaml:"port"`
}

func getDefaultDataPath() string {
	path, _ := os.Getwd()
	return filepath.Join(path, "data")
}

func IsValidConfigExtension(ext string) bool {
	switch ext {
	case ".yml", ".yaml":
		return true
	}
	return false
}

func LoadConfig(path string) (Config, error) {
	// Loads the configuration from locmock.yml file
	// If this file doesn't exist default one is created
	if !IsValidConfigExtension(filepath.Ext(path)) {
		return Config{}, fmt.Errorf("Invalid configuration extension. Should be '.yml' or '.yaml'")
	}
	file, err := os.ReadFile(path)
	if err != nil {
		config := Config{
			DataPath: getDefaultDataPath(),
			Port:     ":8080",
		}
		var data []byte
		data, _ = yaml.Marshal(config)
		err = os.WriteFile(path, data, 0666)
		if err != nil {
			panic(err)
		}
		return config, nil
	}
	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func configMiddleware(config *Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dataPath", config.DataPath)
		c.Next()
	}
}

func setupRouter(config Config) *gin.Engine {
	router := gin.New()

	// Add middlewares
	router.Use(cors.Default())
	router.Use(configMiddleware(&config))

	// Admin routes to create/delete/update services and actions
	router.GET("/admin/services", listServices)

	router.GET("/admin/service/:service/actions", listServiceActions)

	router.POST("/admin/service/:service/action", createNewAction)

	router.POST("/admin/service/:service/actions", bulkCreateActions)

	router.PUT("admin/service/:service/action/:action", updateAction)

	router.DELETE("admin/service/:service/action/:action", deleteAction)

	router.POST("admin/service/:service/actions/delete", bulkDeleteActions)

	// Add the utility routes to the router
	addUtilityRoads(&router)

	// The magic catcher method - catch all requests and call appropriate service/action and return response
	router.Any("/:service/*action", func(c *gin.Context) {
		serviceName := c.Param("service")
		actionName := c.Param("action")
		values := c.Request.URL.Query()
		method := c.Request.Method

		service, err := service.NewFromPath(filepath.Join(config.DataPath, serviceName))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"method":  method,
				"service": serviceName,
				"action":  actionName,
				"query":   values,
			})
			return
		}
		action, ok := service.Actions[strings.TrimLeft(actionName, "/")]
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"method":  method,
				"service": serviceName,
				"action":  actionName,
				"query":   values,
			})
			return
		}
		if action.Method != c.Request.Method {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported request method"})
		}
		c.JSON(action.RunAction())
	})
	return router
}

func Run(config Config) {
	// Create default router
	router := setupRouter(config)
	_ = router.Run(config.Port)
}

func addUtilityRoads(i **gin.Engine) {
	router := (*i).Group("/l")
	router.GET("/ping", getPing)
	router.GET("/ip", getIp)
	router.GET("/person", personProfile)
	router.GET("/user-agent", userAgent)
	router.GET("/uuid", uuidResponse)
	router.POST("/form", formRequest)
	router.PUT("/form", formRequest)
	router.PATCH("/form", formRequest)
	router.Any("/headers", getHeaders)
	router.GET("/redirect", redirectRequest)
	router.Any("/gzip", gzip.Gzip(gzip.DefaultCompression), gzipRequest)
	router.Any("/:method", genericRouteResponse)
}
