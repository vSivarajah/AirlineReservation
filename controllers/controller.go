package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vsivarajah/AirlineReservation/domain"
)

func GetFlights(c *gin.Context) {
	flights := domain.GetFlights()
	fmt.Printf("%T\n", flights)
	c.JSON(200, gin.H{
		"message": flights,
	})
}

func GetReservationDetails(c *gin.Context) {
	reservation := domain.GetReservationDetails()
	c.JSON(200, gin.H{
		"reservation": reservation,
	})
}

func CreateReservation(c *gin.Context) {

	//if domain.DoesFlightExist(sourceairport, targetairport) {
	details := domain.Reservation{}
	if err := c.ShouldBindJSON(&details); err != nil {
		log.Println("Invalid json body")
		return
	}
	if domain.DoesFlightExist(details.FlightInfo.SourceAirport, details.FlightInfo.TargetAirport) {
		details.FlightInfo.FlightNumber, details.FlightInfo.OperatingAirlines = domain.AssignFlightNumber(details.FlightInfo.SourceAirport, details.FlightInfo.TargetAirport)
		details.IsValid = false
		domain.CreateFlightDetails(&details)
		c.JSON(http.StatusCreated, gin.H{
			"message": "Created a new flight detail",
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "can not find a flight",
		})
	}

}

func CreatePayment(c *gin.Context) {
	payment := domain.Payment{}
	if err := c.ShouldBindJSON(&payment); err != nil {
		log.Println("Invalid json body")
	}
	domain.CreatePayment(&payment)
	c.JSON(http.StatusCreated, gin.H{
		"message": "payment successful",
	})

}

func GetPayment(c *gin.Context) {
	payments := domain.GetPayment()
	c.JSON(200, gin.H{
		"payment": payments,
	})
}
