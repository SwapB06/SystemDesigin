package controller

type Seat struct {
	SeatID       int
	Row          int
	SeatCategory SeatCategory
}

func NewSeat() Seat {
	return Seat{}
}
func (s Seat) GetSeatID() int {
	return s.SeatID
}
func (s *Seat) SetSeatID(seatID int) {
	s.SeatID = seatID
}
func (s Seat) GetRow() int {
	return s.Row
}
func (s *Seat) SetRow(row int) {
	s.Row = row
}
func (s Seat) GetSeatCategory() SeatCategory {
	return s.SeatCategory
}
func (s *Seat) SetSeatCategory(SeatCategory SeatCategory) {
	s.SeatCategory = SeatCategory
}
