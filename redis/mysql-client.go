package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var MoviesName = []string{"Avengers", "The Eternals", "Overwatch", "Elden ring", "Dark souls",
	"The Last of Us", "Honkai impact", "FlyMe2theMoon", "Stardew Valley", "Doom"}
var MoviesGenre = []string{"Action", "Comedy", "Romance", "Horror", "Harem",
	"Isekai", "Sci-fi", "Gender bender", "Slice of life", "Fantasy"}
var MoviesDirector = []string{"Kevin Kaslana", "Raiden Mei", "Bronya Zaychik", "Seele", "Otto Apocalypse",
	"Murata Himeko", "Rita Rossweisse", "Hidetaka Miyazaki", "George Martin", "Michael Bay"}
var MoviesRating = []float64{9.1, 9.5, 9.3, 9.2, 8.7, 7.2, 5.4, 3.1, 9.9, 8.6}

type Movies struct {
	MovieName  string  `db:"MovieName"`
	MovieGenre string  `db:"MovieGenre"`
	Director   string  `db:"Director"`
	Rating     float32 `db:"Rating"`
}

func GetAllMovies(db *sqlx.DB) {
	var movies []Movies
	rows, err := db.Queryx("SELECT * FROM Movies ORDER BY Rating DESC")
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var movie Movies
		err = rows.StructScan(&movie)
		if err != nil {
			panic(err.Error())
		}
		movies = append(movies, movie)
	}
	for _, movie := range movies {
		fmt.Println("Movie name: ", movie.MovieName, "| Movie genre: ", movie.MovieGenre, "| Movie Director: ", movie.Director, "| Movie Rating: ", movie.Rating)
	}
}

func main() {
	db, err := sqlx.Open("mysql", "root:1@tcp(127.0.0.1:3306)/Movies")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// Query the database
	GetAllMovies(db)
}
