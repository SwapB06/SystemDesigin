package controller

type slcOfTheatre []Theatre

type CityVsTheatre map[City]slcOfTheatre

type TheatreVsShows map[*Theatre]slcOfShows

type TheatreController struct {
	CityVsTheatre
	AllTheatre slcOfTheatre
}

func NewTheatreController() TheatreController {
	//theatres := make([]Theatre, 0)
	return TheatreController{
		CityVsTheatre: make(CityVsTheatre, 0),
		AllTheatre:    make([]Theatre, 0),
	}
}
func (t *TheatreController) AddTheatre(theatre Theatre, city City) {
	t.AllTheatre = append(t.AllTheatre, theatre)
	theatres := make(slcOfTheatre, 0)
	if ts, ok := t.CityVsTheatre[city]; ok {
		theatres = ts
	}
	theatres = append(theatres, theatre)
	t.CityVsTheatre[city] = theatres
}

func (t TheatreController) GetAllShow(movie Movie, city City) TheatreVsShows {
	theatrevsshows := make(TheatreVsShows, 0)
	theatres := t.CityVsTheatre[city]

	for _, t := range theatres {
		givenMoviesShows := make([]*Show, 0)
		shows := t.GetShows()

		for _, show := range shows {
			if show.Movie.GetMovieID() == movie.GetMovieID() {
				givenMoviesShows = append(givenMoviesShows, show)
			}
		}
		if len(givenMoviesShows) != 0 {
			theatrevsshows[&t] = givenMoviesShows
		}
	}
	return theatrevsshows
}
