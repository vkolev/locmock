package locmock

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vkolev/locmock/action"
	"github.com/vkolev/locmock/service"
	"net/http"
	"os"
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

func createNewAction(c *gin.Context) {
	var newAction action.Action
	err := c.BindJSON(&newAction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	serviceName := c.Param("service")
	dataPath, _ := c.Get("dataPath")
	servicePath := filepath.Join(dataPath.(string), serviceName)

	_, err = action.CreateAction(servicePath, newAction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusCreated, newAction)
}

func updateAction(c *gin.Context) {
	var newAction action.Action
	err := c.BindJSON(&newAction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	serviceName := c.Param("service")
	dataPath, _ := c.Get("dataPath")
	servicePath := filepath.Join(dataPath.(string), serviceName)
	_, err = action.CreateAction(servicePath, newAction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, newAction)
}

func deleteAction(c *gin.Context) {
	serviceName := c.Param("service")
	actionName := c.Param("action")
	dataPath, _ := c.Get("dataPath")
	servicePath := filepath.Join(dataPath.(string), serviceName)
	actionPath := filepath.Join(servicePath, actionName)
	if _, err := os.Stat(actionPath); !os.IsExist(err) {
		// The requested action does not exist
		c.JSON(http.StatusNotFound,
			gin.H{"error": fmt.Sprintf("Cannot find action %q for service %q", actionName, serviceName)},
		)
		return
	}
	err := os.RemoveAll(actionPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("cannot delete action: %v", err)})
	}
	c.JSON(
		http.StatusOK,
		gin.H{"message": fmt.Sprintf("Action %q  for Service %q deleted", actionName, serviceName)},
	)
}
