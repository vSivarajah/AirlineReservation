package app

import (
	"github.com/vsivarajah/AirlineReservation/controllers"
)

func MapUrls() {
	router.GET("/test", controllers.GetReservation)
	router.POST("/create", controllers.Create)
}
