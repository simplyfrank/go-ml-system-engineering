package csvformat

import (
	"encoding/csv"
	"io"
	"strconv"
)

// Define type to holld movie data
type Movie struct {
	Title string
	Director string
	Year int
}

// ReadMovieDataFromCSV reads Movie entries from a csv file
// using a Reader converting marshaling to movie struct and
// returning a list of movies
func ReadMovieDataFromCSV(b io.Reader) ([]Movie, error) {
	// Get a csv Reader value
	r := csv.NewReader(b)

	// Set optional configuration options
	// which can be propagated to function API
	r.Comma = ';'
	r.Comment = '-'

	var movies []Movie

	// ignoring the header
	_, err := r.Read()
	if err != nil && err != io.EOF {
		return nil, err
	}

	// loop till end of Reader
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		year, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			return nil, err
		}

		// Create movie value from parsed line
		m := Movies{record[0], record[1], int(year)}
		movies = append(movies, m)

	}

	return movies, nil
}