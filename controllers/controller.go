package controllers

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vsivarajah/AirlineReservation/domain"
)

func GET(c *gin.Context) {
	message := domain.GetFlightDetails()
	c.JSON(200, gin.H{
		"message": message,
	})
}

func Create(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatal("cannot read body")
	}
	fmt.Println(string(value))
}
