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

func bulkCreateActions(c *gin.Context) {
	var newActions []action.Action
	err := c.BindJSON(&newActions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	serviceName := c.Param("service")
	dataPath, _ := c.Get("dataPath")
	servicePath := filepath.Join(dataPath.(string), serviceName)

	var errors []error

	for index, act := range newActions {
		_, err = action.CreateAction(servicePath, act)
		if err != nil {
			// If we get an error creating an Action we remove it from the slice
			// this way we don't return errored actions
			newActions = append(newActions[:index], newActions[index+1:]...)
			errors = append(errors, err)
		}
	}
	c.JSON(http.StatusCreated, gin.H{"actions": newActions, "errors": errors})
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

func bulkDeleteActions(c *gin.Context) {
	var actionNames map[string][]string
	err := c.BindJSON(&actionNames)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	serviceName := c.Param("service")
	dataPath, _ := c.Get("dataPath")
	servicePath := filepath.Join(dataPath.(string), serviceName)
	var errors []error
	for index, actionName := range actionNames["actionNames"] {
		actionPath := filepath.Join(servicePath, actionName)
		if _, err := os.Stat(actionPath); !os.IsExist(err) {
			errors = append(errors, err)
			actionNames["actionNames"] = append(actionNames["actionName"][:index], actionNames["actionNames"][:index+1]...)
			continue
		}
		err := os.RemoveAll(actionPath)
		if err != nil {
			errors = append(errors, err)
			actionNames["actionNames"] = append(actionNames["actionName"][:index], actionNames["actionNames"][:index+1]...)
		}
	}
	c.JSON(http.StatusOK, gin.H{"removed": actionNames["actionNames"], "errors": errors})
}
