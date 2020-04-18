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
		"flights": flights,
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

		if details.FlightInfo.FlightNumber == "full" {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Flight is full",
			})
			return
		}

		created := services.ReservationService.CreateReservation(&details)
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
	err, created := services.PaymentService.CreatePayment(&payment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not process the payment",
		})
	}
	_, i, _ := services.ReservationService.FindReservationById(payment.PaymentID)
	if i == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "reservation does not exist",
		})
	} else if created {
		c.JSON(http.StatusCreated, gin.H{
			"message": "payment successful",
		})
	} else {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Payment exists",
		})
	}
}

func GetPayment(c *gin.Context) {
	payments, err := services.PaymentService.GetPayment()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
	}
	c.JSON(http.StatusOK, gin.H{
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
	paymentList, err := services.PaymentService.GetPayment()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not retrieve payments, bad request",
		})
	}

	for _, payment := range paymentList {
		if reservationId == payment.PaymentID {
			reservation.IsValid = true
		}
	}
	err = services.ReservationService.UpdateReservation(reservationId, &reservation)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not update reservation",
		})
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
