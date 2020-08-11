package domain

import "fmt"

type Work struct {
	Title          string
	Isbn           string
	AuthorEmailIds []string
	OtherProps     map[string]string
}

func (work Work) String() string {
	return fmt.Sprintf("%v", work.Title)
}
