package locmock

import (
	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
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
	c.String(http.StatusOK, "pong")
}

func personProfile(c *gin.Context) {
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
	c.String(http.StatusOK, randomdata.UserAgentString())
}
