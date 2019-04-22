package connector

import (
	"database/sql"
	"fmt"
	"github.com/op/go-logging"
	"gopkg.in/goracle.v2"
	"log"
)

type OracleConnector struct {
	Logger logging.Logger
}

func (o *OracleConnector) GetDB(hostname *string, port *int, username *string, password *string, sid *string, connectionString *string) *sql.DB {

	var db *sql.DB
	var err error

	var cs string

	if connectionString != nil && len(*connectionString) > 0 {
		o.Logger.Debug("Detected connectionString, using it")
		cs = *connectionString

	} else {

		connParams := goracle.ConnectionParams{
			Password:    *password,
			Username:    *username,
			SID:         fmt.Sprintf("%s/%s", *hostname, *sid),
			MaxSessions: 5,
		}
		o.Logger.Debugf("Built connectionString \"%s\"\n", connParams.String())
		cs = connParams.StringWithPassword()

	}

	o.Logger.Debug("Attempting to connect to database")
	db, err = sql.Open("goracle", cs)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err.Error())
	}
	return db

}

func (o *OracleConnector) Query(query string, db *sql.DB) *Response {

	response := new(Response)

	if db != nil {

		o.Logger.Debugf("Attempting to execute query \"%s\"\n", query)
		rows, err := db.Query(query)
		if err != nil {
			response.Error = true
			response.ErrorMessage = err.Error()
			return response
		}
		defer rows.Close()

		columnNames, err := rows.Columns()
		if err != nil {
			response.Error = true
			response.ErrorMessage = err.Error()
			return response
		}

		rc := NewMapStringScan(columnNames)
		for rows.Next() {
			err := rc.Update(rows)
			if err != nil {
				response.Error = true
				response.ErrorMessage = err.Error()
				return response
			}
			cv := rc.Get()
			row := new(Row)

			for k, v := range cv {
				o.Logger.Debugf("Creating column %s: %s", k, v)
				col := new(Column)
				col.Index = k

				for key, value := range v {
					col.Name = key
					col.Value = value
				}

				row.Columns = append(row.Columns, *col)
			}
			response.Rows = append(response.Rows, *row)

		}

	} else {
		response.Error = true
		response.ErrorMessage = "Received a nil object as a DB connection "
		return response
	}
	return response
}

type mapStringScan struct {
	cp []interface{}

	row      map[int]map[string]string
	colCount int
	colNames []string
}

func NewMapStringScan(columnNames []string) *mapStringScan {
	lenCN := len(columnNames)
	s := &mapStringScan{
		cp:       make([]interface{}, lenCN),
		row:      make(map[int]map[string]string, lenCN),
		colCount: lenCN,
		colNames: columnNames,
	}
	for i := 0; i < lenCN; i++ {
		s.cp[i] = new(sql.RawBytes)
	}
	return s
}

func (s *mapStringScan) Update(rows *sql.Rows) error {
	if err := rows.Scan(s.cp...); err != nil {
		return err
	}

	for i := 0; i < s.colCount; i++ {
		if rb, ok := s.cp[i].(*sql.RawBytes); ok {

			s.row[i] = map[string]string{s.colNames[i]: string(*rb)}
			*rb = nil // reset pointer to discard current value to avoid a bug
		} else {
			return fmt.Errorf("Cannot convert index %d column %s to type *sql.RawBytes", i, s.colNames[i])
		}
	}
	return nil
}

func (s *mapStringScan) Get() map[int]map[string]string {
	return s.row
}
