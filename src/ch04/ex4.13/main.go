package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

const APIURL =  "http://omdbapi.com/?apikey=da71b552&t=%s"

type Movie struct {
	Title  string
	Year   string
	Poster string
}


func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: ex4.13 movie_title")
		os.Exit(1)
	}

	title := os.Args[1]

	movie, err := getMovie(title)
	if err != nil || movie == nil || len(movie.Poster) == 0 {
		log.Fatalf("Failed to find movie %s", title)
		os.Exit(1)
	}

	err = movie.writePoster()
	if err != nil {
		log.Fatalf("Failed to write poster %s", movie.Poster)
		os.Exit(1)
	}

}

func getMovie(title string) (*Movie , error) {

	var movie Movie

	url_ := fmt.Sprintf(APIURL, title)
	resp, err := http.Get(url_)
	if err != nil {
		return &movie, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = fmt.Errorf("%d response from %s", resp.StatusCode, url_)
		return &movie, err
	}
	err = json.NewDecoder(resp.Body).Decode(&movie)
	return &movie, err

}

func (m Movie) writePoster() error {
	url_ := m.Poster

	resp, err := http.Get(url_)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = fmt.Errorf("%d response from %s", resp.StatusCode, url_)
		return err
	}

	pwd, _ := os.Getwd()

	file, err := os.Create(path.Join(pwd, "sample.jpg"))
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.ReadFrom(resp.Body)

	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil

}
