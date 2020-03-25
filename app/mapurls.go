package app

import (
	"github.com/vsivarajah/AirlineReservation/controllers"
)

func MapUrls() {
	router.GET("/test", controllers.GET)
	router.POST("/create", controllers.Create)
}
