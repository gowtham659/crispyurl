package database

import (
	"database/sql"
	"errors"
	"log"

	"github.com/sijms/go-ora/v2"
)

//creates a connection and return a connection and an error , connection is returned if success else error if failed
func DbConnect()( *sql.DB, error){
	conString := go_ora.BuildUrl("localhost",1521, "orcl","C##GOWTHAM","1234",nil)

	dbConnect,er := sql.Open("oracle",conString)
	if er!=nil{
		log.Fatal("Connection Error: ",er)
	}

	if er := dbConnect.Ping(); er != nil{
		return nil, errors.New(er.Error())
	}
	return dbConnect, nil
}