package locmock

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vkolev/locmock/service"
	"net/http"
	"path/filepath"
)

func listServices(c *gin.Context) {
	dataPath, ok := c.Get("dataPath")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Config path not supplied in context"})
	}
	services, err := service.GetServices(dataPath.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, services)
}

func listServiceActions(c *gin.Context) {
	dataPath, ok := c.Get("dataPath")
	serviceName := c.Param("service")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Config path not supplied in context"})
	}
	rService, err := service.NewFromPath(filepath.Join(dataPath.(string), serviceName))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Service %q, was not found", serviceName)})
	}
	serviceActions := rService.GetActions()

	c.JSON(http.StatusOK, serviceActions)
}
