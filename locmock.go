package locmock

import (
	"fmt"
	"github.com/gin-contrib/cors"
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
		data, err := yaml.Marshal(config)
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

func Run(config Config) {
	// Create default router
	router := gin.New()

	// Add middlewares
	router.Use(cors.Default())

	// Admin routes to create/delete/update services and actions
	router.GET("/admin/services", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"admin": "get all services"})
	})

	router.GET("/admin/service/:service/actions", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"admin": fmt.Sprintf("get %v service actions", c.Param("service"))})
	})

	router.POST("/admin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"admin": "create service"})
	})

	router.DELETE("/admin/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"admin": fmt.Sprintf("delete %v", c.Param("id"))})
	})

	// The megic catcher method - catch all requests and call appropriate service/action and return response
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
		c.JSON(action.RunAction())
	})

	router.Run(config.Port)
}
