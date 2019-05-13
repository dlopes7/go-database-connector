package connector

import (
	"database/sql"
	"fmt"
	_ "github.com/ibmdb/go_ibm_db"
	"github.com/op/go-logging"
	"log"
)

type DB2Connector struct {
	Logger logging.Logger
}

func (o *DB2Connector) GetDB(hostname *string, port *int, username *string, password *string, database *string, connectionString *string) *sql.DB {

	var db *sql.DB
	var err error

	var cs string

	if connectionString != nil && len(*connectionString) > 0 {
		o.Logger.Debug("Detected connectionString, using it")
		cs = *connectionString

	} else {

		cs = fmt.Sprintf("HOSTNAME=%s;DATABASE=%s;PORT=%d;UID=%s;PWD=%s", *hostname, *database, *port, *username, *password)

	}

	o.Logger.Debugf("Attempting to connect to database %s\n", cs)
	db, err = sql.Open("go_ibm_db", cs)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err.Error())
	}
	return db

}
