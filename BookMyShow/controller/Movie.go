package controller

type Movie struct {
	MovieID                int
	MovieName              string
	MovieDurationInMinutes int
	//other details like genre, language, etc.
}

func NewMovie() Movie {
	return Movie{}
}
func (m Movie) GetMovieID() int {
	return m.MovieID
}
func (m *Movie) SetMovieID(movieID int) {
	m.MovieID = movieID
}
func (m Movie) GetMovieName() string {
	return m.MovieName
}
func (m *Movie) SetMovieName(moviewName string) {
	m.MovieName = moviewName
}
func (m Movie) GetMovieDuration() int {
	return m.MovieDurationInMinutes
}
func (m *Movie) SetMovieDuration(movieDuration int) {
	m.MovieDurationInMinutes = movieDuration
}
