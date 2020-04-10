package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	payments "github.com/vsivarajah/AirlineReservation/domain/payments"
	"github.com/vsivarajah/AirlineReservation/domain/reservations"

	"github.com/vsivarajah/AirlineReservation/services"
)

func GetFlights(c *gin.Context) {
	flights := services.FlightService.GetFlights()
	fmt.Printf("%T\n", flights)
	c.JSON(200, gin.H{
		"message": flights,
	})
}

func GetReservationDetails(c *gin.Context) {
	reservation := services.ReservationService.GetReservationDetails()
	c.JSON(200, gin.H{
		"reservation": reservation,
	})
}

func CreateReservation(c *gin.Context) {

	log.Println("Creating new reservation")

	details := reservations.Reservation{}
	if err := c.ShouldBindJSON(&details); err != nil {
		log.Println("Invalid json body")
		return
	}

	log.WithFields(log.Fields{
		"json": details,
	}).Info("message")

	if services.FlightService.DoesFlightExist(details.FlightInfo.SourceAirport, details.FlightInfo.TargetAirport) {
		details.FlightInfo.FlightNumber, details.FlightInfo.OperatingAirlines = services.FlightService.AssignFlightNumber(details.FlightInfo.SourceAirport, details.FlightInfo.TargetAirport)
		details.IsValid = false

		var created bool = false
		created = services.ReservationService.CreateFlightDetails(&details)
		if created {
			c.JSON(http.StatusCreated, gin.H{
				"message": "Created a new flight detail",
			})
		} else {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Error, reservation id already exists",
			})
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "can not find a flight",
		})
	}

}

func CreatePayment(c *gin.Context) {
	payment := payments.Payment{}
	if err := c.ShouldBindJSON(&payment); err != nil {
		log.Println("Invalid json body")
	}
	services.PaymentService.CreatePayment(&payment)
	c.JSON(http.StatusCreated, gin.H{
		"message": "payment successful",
	})

}

func GetPayment(c *gin.Context) {
	payments := services.PaymentService.GetPayment()
	c.JSON(200, gin.H{
		"payment": payments,
	})
}

func UpdateReservation(c *gin.Context) {
	id := c.Param("id")
	reservationId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Can not convert to int from string")
	}
	reservation := reservations.Reservation{}
	paymentList := services.PaymentService.GetPayment()

	for _, payment := range paymentList {
		if reservationId == payment.PaymentID {
			reservation.IsValid = true
		}
	}
	err = services.ReservationService.UpdateReservation(reservationId, &reservation)
	if err != nil {
		log.Println("Can not update reservation")
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "reservation confirmed!",
	})
}

func DeleteReservation(c *gin.Context) {
	id := c.Param("id")
	reservationId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Can not convert to int from string")
	}

	log.Println(reservationId)

	var i int = 0
	i = services.ReservationService.DeleteReservation(reservationId)
	if i != -1 {
		log.Println("Reservation deleted: ", i)
		c.JSON(http.StatusOK, gin.H{
			"message": "Reservation deleted! ",
		})
	} else {
		log.Println("Error: ", err)
		c.JSON(http.StatusOK, gin.H{
			"message": "Reservation not found",
		})
	}
}
