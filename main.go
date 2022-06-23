package main

import (
	"strconv"

	"github.com/derickr/go-litra-driver"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Define API Endpoints
	router.GET("/on", apiOn)
	router.GET("/off", apiOff)
	router.GET("/brightness/:val", apiBrightness)
	router.GET("/temperature/:val", apiTemperature)
	// Run the Web Server
	router.Run(":8080")
}

func apiOn(c *gin.Context) {
	ld, err := litra.New()
	if err != nil {
		c.AbortWithStatus(500)
	} else {
		ld.TurnOn()
		ld.Close()
		c.Status(200)
	}
	return
}

func apiOff(c *gin.Context) {
	ld, err := litra.New()
	if err != nil {
		c.AbortWithStatus(500)
	} else {
		ld.TurnOff()
		ld.Close()
		c.Status(200)
	}
	return
}

func apiBrightness(c *gin.Context) {
	param, _ := strconv.Atoi(c.Param("val"))
	val := int(param)
	ld, err := litra.New()
	if err != nil {
		c.AbortWithStatus(500)
	} else {
		ld.SetBrightness(val)
		ld.Close()
		c.Status(200)
	}
}

func apiTemperature(c *gin.Context) {
	param, _ := strconv.Atoi(c.Param("val"))
	val := int16(param)
	ld, err := litra.New()
	if err != nil {
		c.AbortWithStatus(500)
	} else {
		ld.SetTemperature(val)
		ld.Close()
		c.Status(200)
	}
}
