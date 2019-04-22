package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/dlopes7/go-database-connector/connector"
	"github.com/op/go-logging"
	"os"
)

func main() {

	dbType := flag.String("type", "oracle", "Database type (oracle)")

	query := flag.String("query", "", "SELECT 1 FROM DUAL")

	dbHost := flag.String("hostname", "", "Hostname")
	dbPort := flag.Int("port", 1521, "Port")

	dbUser := flag.String("username", "", "Port")
	dbPassword := flag.String("password", "", "Port")

	oracleSID := flag.String("sid", "", "Oracle SID")
	databaseName := flag.String("database", "", "Database Name")

	connectionString := flag.String("connection-string", "", "Connection string, ex: oracle://user:password@localhost/ORCLPDB1")

	debug := flag.Bool("debug", false, "Enable debug logging")

	flag.Parse()

	var log = logging.MustGetLogger("go-database-connector")
	logging.SetLevel(logging.WARNING, "go-database-connector")
	var format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{level:.5s} ▶ %{longfunc}: %{color:reset} %{message}`,
	)

	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend1Formatter := logging.NewBackendFormatter(backend1, format)

	if *debug == true {

		logging.SetLevel(logging.DEBUG, "go-database-connector")
		logging.SetBackend(backend1Formatter)
	}

	response := new(connector.Response)
	if *connectionString == "" && *dbHost == "" {

		response.Error = true
		response.ErrorMessage = "Need a connectionString or a hostname"

	} else {

		if *dbType == "oracle" {
			log.Debug("Obtaining an Oracle database connection")

			if len(*oracleSID) == 0 {
				oracleSID = databaseName
			}

			oracleConnector := &connector.OracleConnector{
				Logger: *log,
			}
			db := oracleConnector.GetDB(dbHost, dbPort, dbUser, dbPassword, oracleSID, connectionString)

			response = oracleConnector.Query(*query, db)

		}

	}

	resJson, _ := json.Marshal(response)
	fmt.Println(string(resJson))

}
