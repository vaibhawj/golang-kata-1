package service

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/vaibhawj/golang-kata-1/v1/domain"
)

func getAuthors() []domain.Author {
	csvfile, err := os.Open("resources/authors.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)
	r.Comma = ';'
	authors := make([]domain.Author, 0)
	var count int = 0
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		count++
		if count == 1 {
			continue
		}
		author := domain.Author{
			Email:     record[0],
			FirstName: record[1],
			LastName:  record[2],
		}
		authors = append(authors, author)
	}
	return authors
}
