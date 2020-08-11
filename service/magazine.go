package service

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"

	"github.com/vaibhawj/golang-kata-1/v1/domain"
)

func loadMagazines() []domain.Work {
	csvfile, err := os.Open("resources/magazines.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)
	r.Comma = ';'
	works := make([]domain.Work, 0)
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

		work := domain.Work{
			Title:          record[0],
			Isbn:           record[1],
			AuthorEmailIds: strings.Split(record[2], ","),
			OtherProps:     map[string]string{"publishedAt": record[3]},
		}
		works = append(works, work)
	}
	return works
}
