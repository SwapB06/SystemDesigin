package controller

type Booking struct {
	Show      Show
	BookSeats slcOfSeats
	Payment   Payment
}

func NewBooking() *Booking {
	return &Booking{}
}

func (b Booking) GetShow() Show {
	return b.Show
}

func (b *Booking) SetShow(show *Show) {
	b.Show = *show
}

func (b Booking) GetBookedSeats() slcOfSeats {
	return b.BookSeats
}

func (b *Booking) SetBookedSeats(bookedSeats slcOfSeats) {
	b.BookSeats = bookedSeats
}

func (b Booking) GetPyment() Payment {
	return b.Payment
}

func (b *Booking) SetPayment(payment Payment) {
	b.Payment = payment
}
