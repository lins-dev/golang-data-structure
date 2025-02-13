package exercises

func ArrayAndSliceEx2() []string {
	bestMoviesDB := GetBestMovies();
	resultFromApi := []int{0,1,2,3,4,5,6,7,8,9}
	var movies []string

	for _, value := range resultFromApi {
		movie := bestMoviesDB[value]
		movies = append(movies, movie)
	}

	return movies
}