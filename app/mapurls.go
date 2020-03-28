package app

import (
	"github.com/vsivarajah/AirlineReservation/controllers"
)

func MapUrls() {
	router.GET("/test", controllers.GetFlights)
	router.POST("/create", controllers.CreateReservation)
	router.GET("/reservations", controllers.GetReservationDetails)
}
