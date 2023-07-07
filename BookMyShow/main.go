package main

import "bookmyshow/controller"

func main() {
	bookMyShow := controller.NewBookMyShow()

	bookMyShow.Initialize()

	//user1
	bookMyShow.CreateBooking(controller.Bangalore, "BAAHUBALI")
	//user2
	bookMyShow.CreateBooking(controller.Bangalore, "BAAHUBALI")
}
