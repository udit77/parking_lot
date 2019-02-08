package database

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3/upgrade"
	_ "github.com/mattn/go-sqlite3"
	"github.com/parking_lot/src/config"
)

type DBConnector struct{
	db *sql.DB
	size int
}
var connector *DBConnector

func init(){
	initConnection()
}

func initConnection(){
	connector = newDbConnection(config.Get().DataBase.DriverName, config.Get().DataBase.SourcePath)
}

func GetConnector() *DBConnector{
	if connector.db == nil{
		initConnection()
	}
	return connector
}

func newDbConnection(dbDriver string, dbSource string) *DBConnector{
	database, err := sql.Open(dbDriver, dbSource)
	if err != nil{
		log.Fatalln("fatal [DataBase][Init] db connection initialization failed ", err)
	}
	connector := &DBConnector{
		db:database,
	}
	return connector
}