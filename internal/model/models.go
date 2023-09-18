package model

import "time"

type Period struct {
	Semester  int       `json:"semester"`
	StartDate time.Time `json:"start_date"`
	YearStart int       `json:"year_start"`
	YearEnd   int       `json:"year_end"`
}

type Week struct {
	Week int `json:"week"`
}

type ErrorResponse struct {
	Detail string `json:"detail"`
}
