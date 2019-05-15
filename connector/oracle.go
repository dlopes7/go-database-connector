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
