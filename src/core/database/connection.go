package database

import (
	"database/sql"
	"github.com/parking_lot/src/constants"
	"log"
	_ "github.com/mattn/go-sqlite3/upgrade"
	_ "github.com/mattn/go-sqlite3"
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
	connector = newDbConnection(constants.DbDriverName, constants.DbSourcePath)
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