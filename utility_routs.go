package locmock

import (
	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
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

func uuidResponse(c *gin.Context) {
	uuidVersion := c.Query("type")
	switch uuidVersion {
	case "v1":
		c.String(http.StatusOK, uuid.NewV1().String())
	case "v4":
		c.String(http.StatusOK, uuid.NewV4().String())
	case "v3":
		c.String(http.StatusOK, uuid.NewV3(uuid.NameSpaceURL).String())
	case "v5":
		// Add namspace and arguments from Query Parameters
		c.String(http.StatusOK, uuid.NewV5(uuid.NameSpaceURL).String())
	}
}
