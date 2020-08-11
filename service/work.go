package service

import "github.com/vaibhawj/golang-kata-1/v1/domain"

func Load() ([]domain.Author, []domain.Work) {
	authors := getAuthors()

	works := loadBooks()

	works = append(works, loadMagazines()...)

	return authors, works
}
