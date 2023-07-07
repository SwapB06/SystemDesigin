package controller

type slcOfSeats []Seat
type Screen struct {
	ScreenID int
	Seats    slcOfSeats
}

func NewScreen() Screen {
	return Screen{}
}
func (s Screen) GetScreenID() int {
	return s.ScreenID
}
func (s *Screen) SetScreenID(screenID int) {
	s.ScreenID = screenID
}
func (s Screen) GetSeats() slcOfSeats {
	return s.Seats
}
func (s *Screen) SetSeats(seats slcOfSeats) {
	s.Seats = seats
}
