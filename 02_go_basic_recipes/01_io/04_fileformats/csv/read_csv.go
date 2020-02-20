/// DESCRIPTION:

// In order to learn how to read a CSV format, we first represent our data as a structure. In Go, it's very useful to
//format data as a structure, as it makes things such as marshaling and encoding relatively simple. Our read example
//uses movies as our data type. The function takes an io.Reader interface that holds our CSV data as an input.
//This could be a file or a buffer. We then use that data to create and populate a Movie structure, including converting
//the year into an integer. We also add options to the CSV parser to use ; (semi-colon) as the separator and - (hyphen)
//as a comment line.
//
//Next, we explore the same idea, but in reverse. Novels are represented with a title and an author. We initialize an
//array of novels and then write specific novels in the CSV format to an io.Writer interface. Once again, this can be a
//file, stdout, or a buffer.
//
//The CSV package is an excellent example of why you'd want to think of data flows in Go as implementing common interfaces. It's easy to change the source and destination of our data with small one-line tweaks, and we can easily manipulate CSV data without using an excessive amount of memory or time. For example, it would be possible to read from a stream of data one record at a time and write to a separate stream in a modified format one record at a time. Doing this would not incur significant memory or processor usage.
//Later, when we explore data pipelines and worker pools, you'll see how these ideas can be combined and how to handle these streams in parallel.

package csv

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
// using a Reader creating to movie values and
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



