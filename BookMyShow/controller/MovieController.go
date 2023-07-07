package controller

type slcOfMovies []Movie
type CityVsMovies map[City][]Movie

type MovieController struct {
	CityVsMovies
	AllMovies slcOfMovies
}

func NewMovieController() MovieController {
	return MovieController{
		CityVsMovies: make(CityVsMovies, 0),
		AllMovies:    make(slcOfMovies, 0),
	}
}

// ADD movie to a particular city, make use of cityVsMovies map
func (m *MovieController) AddMovie(movie Movie, city City) {
	m.AllMovies = append(m.AllMovies, movie)

	movies := make([]Movie, 0)
	if _, ok := m.CityVsMovies[city]; ok {
		movies = m.CityVsMovies[city]
	} else {
		m.CityVsMovies[city] = make([]Movie, 0)
	}
	movies = append(movies, movie)
	m.CityVsMovies[city] = movies
}

func (m MovieController) GetMovieByName(movieName string) Movie {
	for _, movie := range m.AllMovies {
		if movie.GetMovieName() == movieName {
			return movie
		}
	}
	return Movie{}
}

func (m MovieController) GetMovieByCity(city City) slcOfMovies {
	return m.CityVsMovies[city]
}

//REMOVE movie from a particular city, make use of cityVsMovies map

//UPDATE movie of a particular city, make use of cityVsMovies map

//CRUD operation based on Movie ID, make use of allMovies list
