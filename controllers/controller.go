package controllers

import (
	"log"
	"net/http"

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
	details := domain.FlightDetail{}
	if err := c.ShouldBindJSON(&details); err != nil {
		log.Println("Invalid json body")
		return
	}
	domain.CreateFlightDetails(&details)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Created a new flight detail",
	})

}
