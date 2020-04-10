package app

import (
	"github.com/vsivarajah/AirlineReservation/controllers"
)

func MapUrls() {
	router.GET("/test", controllers.GetFlights)
	router.POST("/create", controllers.CreateReservation)
	router.POST("/delete/:id", controllers.DeleteReservation)
	router.GET("/reservations", controllers.GetReservationDetails)
	router.POST("/payment/pay", controllers.CreatePayment)
	router.GET("/payments", controllers.GetPayment)
	router.PUT("/reservation/:id", controllers.UpdateReservation)
}
