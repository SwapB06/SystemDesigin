package controller

type slcOfInts []int

type Show struct {
	ShowID        int
	Movie         Movie
	Screen        Screen
	ShowStartTime int
	BookedSeatIDs slcOfInts
}

func NewShow() Show {
	return Show{}
}
func (s Show) GetShowID() int {
	return s.ShowID
}

func (s *Show) SetShowID(showID int) {
	s.ShowID = showID
}

func (s Show) getMovie() Movie {
	return s.Movie
}

func (s *Show) SetMovie(movie Movie) {
	s.Movie = movie
}

func (s Show) GetScreen() Screen {
	return s.Screen
}
func (s *Show) SetScreen(screen Screen) {
	s.Screen = screen
}
func (s Show) GetShowStartTime() int {
	return s.ShowStartTime
}
func (s *Show) SetShowStartTime(showStartTime int) {
	s.ShowStartTime = showStartTime
}
func (s *Show) GetBookedSeatIds() *slcOfInts {
	return &s.BookedSeatIDs
}
func (s *Show) SetBookedSeatIds(bookedSeatIds slcOfInts) {
	s.BookedSeatIDs = bookedSeatIds
}
