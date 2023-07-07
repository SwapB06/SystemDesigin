package controller

import "fmt"

type BookMyShow struct {
	MovieController
	TheatreController
}

func NewBookMyShow() *BookMyShow {
	return &BookMyShow{
		MovieController:   NewMovieController(),
		TheatreController: NewTheatreController(),
	}
}

func (bms BookMyShow) CreateBooking(userCity City, movieName string) {

	//1. search movie by my location
	movies := bms.MovieController.GetMovieByCity(userCity)

	//2. select the movie which you want to see. i want to see Baahubali
	var interestedMovie Movie
	for _, movie := range movies {

		if movie.GetMovieName() == movieName {
			interestedMovie = movie
		}
	}
	//3. get all show of this movie in Bangalore location
	//Map<Theatre, List<Show>>
	showsTheatreWise := bms.TheatreController.GetAllShow(interestedMovie, userCity)

	//4. select the particular show user is interested in
	//Map.Entry<Theatre,List<Show>>
	/*entry := showsTheatreWise.entrySet().iterator().next()
	  List<Show> runningShows = entry.getValue()
	  Show interestedShow = runningShows.get(0)*/

	var interestedShow *Show
	for _, runningshows := range showsTheatreWise {
		interestedShow = runningshows[0]
	}

	//5. select the seat
	seatNumber := 30
	//List<Integer>
	bookedSeats := interestedShow.GetBookedSeatIds()

	isSeatAvailable := true
	for _, val := range *bookedSeats {
		if seatNumber == val {
			isSeatAvailable = false
			break
		}
	}
	if !isSeatAvailable {
		//throw exception
		fmt.Println("seat already booked, try again")
		return
	} else {
		*bookedSeats = append(*bookedSeats, seatNumber)
		//startPayment
		booking := NewBooking()
		myBookedSeats := make(slcOfSeats, 0)
		for _, screenSeat := range interestedShow.GetScreen().GetSeats() {
			if screenSeat.GetSeatID() == seatNumber {
				myBookedSeats = append(myBookedSeats, screenSeat)
			}
		}
		booking.SetBookedSeats(myBookedSeats)
		booking.SetShow(interestedShow)

		interestedShow.SetBookedSeatIds([]int{seatNumber})
	}

	fmt.Println("BOOKING SUCCESSFUL")
}

func (bms *BookMyShow) Initialize() {
	//create movies
	bms.createMovies()

	//create theater with screens, seats and shows
	bms.createTheatre()
}

// creating 2 theatre
func (bms *BookMyShow) createTheatre() {

	avengerMovie := bms.MovieController.GetMovieByName("AVENGERS")
	baahubali := bms.MovieController.GetMovieByName("BAAHUBALI")

	inoxTheatre := NewTheatre()
	inoxTheatre.SetTheatreID(1)
	inoxTheatre.SetScreen(createScreen())
	inoxTheatre.SetCity(Bangalore)
	inoxShows := make(slcOfShows, 0)
	inoxMorningShow := createShows(1, inoxTheatre.GetScreen()[0], avengerMovie, 8)
	inoxEveningShow := createShows(2, inoxTheatre.GetScreen()[0], baahubali, 16)
	inoxShows = append(inoxShows, &inoxMorningShow)
	inoxShows = append(inoxShows, &inoxEveningShow)
	inoxTheatre.SetShows(inoxShows)

	pvrTheatre := NewTheatre()
	pvrTheatre.SetTheatreID(2)
	pvrTheatre.SetScreen(createScreen())
	pvrTheatre.SetCity(Delhi)
	pvrShows := make(slcOfShows, 0)
	pvrMorningShow := createShows(3, pvrTheatre.GetScreen()[0], avengerMovie, 13)
	pvrEveningShow := createShows(4, pvrTheatre.GetScreen()[0], baahubali, 20)
	pvrShows = append(pvrShows, &pvrMorningShow)
	pvrShows = append(pvrShows, &pvrEveningShow)
	pvrTheatre.SetShows(pvrShows)

	bms.TheatreController.AddTheatre(inoxTheatre, Bangalore)
	bms.TheatreController.AddTheatre(pvrTheatre, Delhi)

}

func createScreen() slcOfScreen {

	screens := make(slcOfScreen, 0)
	screen1 := NewScreen()
	screen1.SetScreenID(1)
	screen1.SetSeats(createSeats())
	screens = append(screens, screen1)

	return screens
}

func createShows(showId int, screen Screen, movie Movie, showStartTime int) Show {

	show := NewShow()
	show.SetShowID(showId)
	show.SetScreen(screen)
	show.SetMovie(movie)
	show.SetShowStartTime(showStartTime) //24 hrs time ex: 14 means 2pm and 8 means 8AM
	return show
}

// creating 100 seats
func createSeats() slcOfSeats {

	//creating 100 seats for testing purpose, this can be generalised
	seats := make(slcOfSeats, 0)

	//1 to 40 : SILVER
	for i := 0; i < 40; i++ {
		seat := NewSeat()
		seat.SetSeatID(i)
		seat.SetSeatCategory(SILVER)
		seats = append(seats, seat)
	}

	//41 to 70 : SILVER
	for i := 40; i < 70; i++ {
		seat := NewSeat()
		seat.SetSeatID(i)
		seat.SetSeatCategory(GOLD)
		seats = append(seats, seat)
	}

	//1 to 40 : SILVER
	for i := 70; i < 100; i++ {
		seat := NewSeat()
		seat.SetSeatID(i)
		seat.SetSeatCategory(PLATINUM)
		seats = append(seats, seat)
	}

	return seats
}

func (bms *BookMyShow) createMovies() {

	//create Movies1
	avengers := NewMovie()
	avengers.SetMovieID(1)
	avengers.SetMovieName("AVENGERS")
	avengers.SetMovieDuration(128)

	//create Movies2
	baahubali := NewMovie()
	baahubali.SetMovieID(2)
	baahubali.SetMovieName("BAAHUBALI")
	baahubali.SetMovieDuration(180)

	//add movies against the cities
	bms.MovieController.AddMovie(avengers, Bangalore)
	bms.MovieController.AddMovie(avengers, Delhi)
	bms.MovieController.AddMovie(baahubali, Bangalore)
	bms.MovieController.AddMovie(baahubali, Delhi)
}
