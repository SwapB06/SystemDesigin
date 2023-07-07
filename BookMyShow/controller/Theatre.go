package controller

type slcOfScreen []Screen
type slcOfShows []*Show

type Theatre struct {
	TheatreID int
	Address   string
	City      City
	Screen    slcOfScreen
	Shows     slcOfShows
}

func NewTheatre() Theatre {
	return Theatre{}
}
func (t Theatre) GetTheatreId() int {
	return t.TheatreID
}
func (t *Theatre) SetTheatreID(theatreID int) {
	t.TheatreID = theatreID
}
func (t Theatre) GetAddress() string {
	return t.Address
}
func (t *Theatre) SetAddress(address string) {
	t.Address = address
}
func (t Theatre) GetScreen() slcOfScreen {
	return t.Screen
}
func (t *Theatre) SetScreen(screen slcOfScreen) {
	t.Screen = screen
}
func (t Theatre) GetShows() slcOfShows {
	return t.Shows
}
func (t *Theatre) SetShows(shows slcOfShows) {
	t.Shows = shows
}
func (t Theatre) GetCity() City {
	return t.City
}
func (t *Theatre) SetCity(city City) {
	t.City = city
}
