package connector

import "database/sql"

type Connector interface {
	Query(db *sql.DB) (*Response, error)
	GetDB(*string, *int, *string, *string, *string, *string) *sql.DB
}

type Response struct {
	Error        bool   `json:"error"`
	ErrorMessage string `json:"errorMessage"`
	Rows         []Row  `json:"rows"`
}

type Row struct {
	Columns []Column `json:"columns"`
}

type Column struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
