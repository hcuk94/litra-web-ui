package main

import (
	"strconv"

	"github.com/derickr/go-litra-driver"
	"github.com/gin-gonic/gin"
)

var ld *litra.LitraDevice

func main() {
	// Initialise Litra Device
	var ld_err error
	ld, ld_err = litra.New()
	if ld_err == nil {
		// Define Gin Router
		router := gin.Default()
		// Define API Endpoints
		router.GET("/api/on", apiOn)
		router.GET("/api/off", apiOff)
		router.GET("/api/brightness/:val", apiBrightness)
		router.GET("/api/temperature/:val", apiTemperature)
		// Define Static Web Endpoint
		router.Static("/", "/html")
		// Run the Web Server
		router.Run(":8080")
	}
	// Close Litra Device
	ld.Close()
}

func apiOn(c *gin.Context) {
	ld.TurnOn()
	c.Status(200)
	return
}

func apiOff(c *gin.Context) {
	ld.TurnOff()
	c.Status(200)
	return
}

func apiBrightness(c *gin.Context) {
	param, _ := strconv.Atoi(c.Param("val"))
	val := int(param)
	ld.SetBrightness(val)
	c.Status(200)
	return
}

func apiTemperature(c *gin.Context) {
	param, _ := strconv.Atoi(c.Param("val"))
	val := int16(param)
	ld.SetTemperature(val)
	c.Status(200)
	return
}
