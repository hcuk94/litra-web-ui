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
	ld, _ := litra.New()
	ld.TurnOn()
	ld.Close()
}

func apiOff(c *gin.Context) {
	ld, _ := litra.New()
	ld.TurnOff()
	ld.Close()
}

func apiBrightness(c *gin.Context) {
	param, _ := strconv.Atoi(c.Param("val"))
	val := int(param)
	ld, err := litra.New()
	if err == nil {
		ld.SetBrightness(val)
		ld.Close()
	}
}

func apiTemperature(c *gin.Context) {
	param, _ := strconv.Atoi(c.Param("val"))
	val := int16(param)
	ld, err := litra.New()
	if err == nil {
		ld.SetTemperature(val)
		ld.Close()
	}
}
