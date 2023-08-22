package locmock

import (
	"fmt"
	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
)

func getIp(c *gin.Context) {
	ip := net.ParseIP(c.ClientIP())
	if ip.IsLoopback() {
		ip = net.ParseIP("127.0.0.1")
	}
	if c.Query("ipv6") != "" {
		c.JSON(http.StatusOK, ip.To16().String())
		return
	}
	c.String(http.StatusOK, ip.To4().String())
}

func getPing(c *gin.Context) {
	// returns the string "pong" on GET request
	c.String(http.StatusOK, "pong")
}

func personProfile(c *gin.Context) {
	// Return random Person Profile in JSON format
	gender := c.Query("gender")
	var genderType int
	switch gender {
	case "male":
		genderType = randomdata.Male
	case "female":
		genderType = randomdata.Female
	default:
		genderType = randomdata.RandomGender
	}
	c.IndentedJSON(http.StatusOK, randomdata.GenerateProfile(genderType))
}

func userAgent(c *gin.Context) {
	// Return the requester User-Agent
	userAgent := c.Request.Header.Get("User-Agent")
	c.String(http.StatusOK, userAgent)
}

func uuidResponse(c *gin.Context) {
	// Return a random UUIDv4 string
	c.String(http.StatusOK, uuid.NewV4().String())
}

func formRequest(c *gin.Context) {
	_ = c.Request.ParseMultipartForm(2048)
	requestedForm := c.Request.Form
	var file map[string][]*multipart.FileHeader
	if c.Request.MultipartForm != nil {
		file = c.Request.MultipartForm.File
	}
	requestHeaders := c.Request.Header
	c.IndentedJSON(http.StatusOK, gin.H{
		"form":    requestedForm,
		"headers": requestHeaders,
		"method":  c.Request.Method,
		"file":    file,
	})
}

func redirectRequest(c *gin.Context) {
	redirectCodes := map[string]int{
		"301": 301,
		"302": 302,
		"303": 303,
		"304": 304,
		"305": 305,
		"307": 307,
	}
	statusCode, ok := redirectCodes[c.Query("status")]
	if !ok {
		c.String(http.StatusBadRequest, fmt.Sprintf("Status code %v is not a redirect code", c.Query("status")))
		return
	}
	redirectUrl := c.Query("url")

	_, err := url.Parse(redirectUrl)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("URL %q is not valid", redirectUrl))
		return
	}

	c.Redirect(statusCode, redirectUrl)
}

func gzipRequest(c *gin.Context) {
	headers := c.Request.Header
	requestBody := c.Request.Body
	query := c.Request.URL.Query()
	c.IndentedJSON(http.StatusOK, gin.H{
		"headers":          headers,
		"request_body":     requestBody,
		"query_parameters": query,
	})
}

func genericRouteResponse(c *gin.Context) {
	allowedMethod := map[string]string{
		"post":    "POST",
		"get":     "GET",
		"put":     "PUT",
		"patch":   "PATCH",
		"head":    "HEAD",
		"delete":  "DELETE",
		"options": "OPTIONS",
	}
	method := c.Param("method")
	if m, ok := allowedMethod[method]; !ok || m != c.Request.Method {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Request is not accepted request method: %v real request method: %v", m, c.Request.Method),
		})
		return
	}

	requestHeaders := c.Request.Header
	requestBody := c.Request.Body
	requestQuery := c.Request.URL.Query()
	c.IndentedJSON(http.StatusOK, gin.H{
		"headers": requestHeaders,
		"body":    requestBody,
		"qurey":   requestQuery,
	})
}

func getHeaders(c *gin.Context) {
	requestHeaders := c.Request.Header

	c.IndentedJSON(http.StatusOK, gin.H{
		"headers": requestHeaders,
	})
}
