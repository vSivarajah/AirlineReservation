package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vsivarajah/AirlineReservation/domain"
)

func GetReservation(c *gin.Context) {
	message := domain.GetReservation()
	fmt.Println(message)
	c.JSON(200, gin.H{
		"message": message,
	})
}

func Create(c *gin.Context) {
	details := domain.Reservation{}
	if err := c.ShouldBindJSON(&details); err != nil {
		log.Println("Invalid json body")
		return
	}
	domain.CreateFlightDetails(&details)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Created a new flight detail",
	})

}
