package database

import (
	"github.com/parking_lot/src/commands"
	"strconv"
		"errors"
	"fmt"
)

type QueryBuilder interface {
	CreateLot()
	Park()
	Leave()
	GetStatus()
	GetNumbersWithColor()
	GetSlotWithNumber()
	GetSlotsWithColor()
}

type ParkingDbRowData struct{
	Id int
	Number string
	Color string
}

func (connector *DBConnector) CreateLot(command *commands.CommandBuilder) error{
	statement, err := connector.db.Prepare("DROP TABLE IF EXISTS PARKING")
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil{
		return err
	}
	statement, err = connector.db.Prepare("CREATE TABLE IF NOT EXISTS PARKING (id INTEGER PRIMARY KEY AUTOINCREMENT, registration_number VARCHAR(64), colour VARCHAR(64), status VARCHAR(64))")
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		return err
	}
	connector.size,err  = strconv.Atoi(command.Parameters[0])
	if err != nil{
		return err
	}
	return nil
}

func (connector *DBConnector) Park(command *commands.CommandBuilder) (string, error){
	if connector.size == 0 {
		return "parking lot is not initialized", errors.New("parking lot is not initialized")
	}
	var count int
	statement, err := connector.db.Prepare("SELECT COUNT(*) AS COUNT FROM PARKING WHERE STATUS = ?")
	if err != nil {
		return "", err
	}
	rows , err := statement.Query("OCCUPIED")
	if err != nil{
		return "", err
	}
	for rows.Next(){
		rows.Scan(&count)
	}
	if count == connector.size{ //Fully occupied
		return "Sorry, parking lot is full", nil
	}else{
		var id int
		statement, err := connector.db.Prepare("SELECT id FROM PARKING WHERE STATUS = ? ORDER BY id ASC LIMIT 1")
		if err != nil {
			return "", err
		}
		rows , err := statement.Query("VACANT")
		if err != nil{
			return "", err
		}
		for rows.Next(){
			rows.Scan(&id)
		}
		if id == 0 {
			statement, err = connector.db.Prepare("INSERT INTO PARKING (registration_number, colour, status) VALUES (?, ?, ?)")
			if err != nil {
				return "", err
			}
			_ , err := statement.Exec(command.Parameters[0], command.Parameters[1], "OCCUPIED")
			if err != nil{
				return "", err
			}
			return fmt.Sprintf("%v%v","Allocated slot number: ",count+1), nil
		}else{
			statement, err = connector.db.Prepare("UPDATE PARKING SET registration_number = ?, colour = ?, status = ? where id = ?")
			if err != nil {
				return "", err
			}
			_ , err := statement.Exec(command.Parameters[0], command.Parameters[1], "OCCUPIED", id)
			if err != nil{
				return "", err
			}
			return fmt.Sprintf("%v%v","Allocated slot number: ",id), nil
		}
	}
}

func (connector *DBConnector) Leave(command *commands.CommandBuilder) (string,error){
	if connector.size == 0 {
		return "parking lot is not initialized", errors.New("parking lot is not initialized")
	}
	statement, err := connector.db.Prepare("UPDATE PARKING set status=? where id=?")
	if err != nil {
		return "", err
	}
	_, err = statement.Exec("VACANT", command.Parameters[0])
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v%v%v","Slot number ", command.Parameters[0], " is free"), nil
}

func (connector *DBConnector) GetStatus(command *commands.CommandBuilder)([]ParkingDbRowData, string, error){
	var (
		id int
		number string
		colour string
		status string
	    result []ParkingDbRowData
	)
	if connector.size == 0 {
		return []ParkingDbRowData{}, "parking lot is not initialized", errors.New("parking lot is not initialized")
	}
	statement, err := connector.db.Prepare("SELECT * FROM PARKING WHERE STATUS = ?")
	if err != nil {
		return []ParkingDbRowData{}, "", err
	}
	rows , err := statement.Query("OCCUPIED")
	if err != nil {
		return []ParkingDbRowData{}, "", err
	}
	for rows.Next(){
		_ = rows.Scan(&id, &number, &colour, &status)
		result = append(result, ParkingDbRowData{id,number,colour})
	}
	return result, "", nil
}

func (connector *DBConnector) GetNumbersWithColor(command *commands.CommandBuilder) (string, error){
	if connector.size == 0 {
		return "parking lot is not initialized", errors.New("parking lot is not initialized")
	}
	var (
		number string
	)
	statement, err := connector.db.Prepare("SELECT registration_number FROM PARKING WHERE COLOUR = ? AND STATUS = ?")
	if err != nil {
		return "", err
	}
	rows , err := statement.Query(command.Parameters[0],"OCCUPIED")
	if err != nil {
		return "", err
	}
	result := ""
	for rows.Next(){
		_ = rows.Scan(&number)
		if result == ""{
			result = fmt.Sprintf("%v",number)
		}else {
			result = fmt.Sprintf("%v , %v", result, number)
		}
	}
	if result == ""{
		return "Not found",  nil
	}
	return result, nil
}

func (connector *DBConnector) GetSlotWithNumber(command *commands.CommandBuilder)(string,error){
	if connector.size == 0 {
		return "parking lot is not initialized", errors.New("parking lot is not initialized")
	}
	var (
		slot int
	)
	statement, err := connector.db.Prepare("SELECT id FROM PARKING WHERE registration_number = ? AND STATUS = ?")
	if err != nil {
		return "",err
	}
	rows , err := statement.Query(command.Parameters[0],"OCCUPIED")
	if err != nil {
		return "",err
	}
	for rows.Next(){
		_ = rows.Scan(&slot)
	}
	if slot == 0{
		return "Not found" , nil
	}
	return fmt.Sprintf("%v",slot), nil
}

func (connector *DBConnector) GetSlotsWithColor(command *commands.CommandBuilder)(string, error){
	if connector.size == 0 {
		return "parking lot is not initialized", errors.New("parking lot is not initialized")
	}
	var (
		slot int
	)
	statement, err := connector.db.Prepare("SELECT id FROM PARKING WHERE colour = ? AND STATUS = ?")
	if err != nil {
		return "",err
	}
	rows , err := statement.Query(command.Parameters[0],"OCCUPIED")
	if err != nil {
		return "",err
	}
	result := ""
	for rows.Next(){
		_ = rows.Scan(&slot)
		if result == ""{
			result = fmt.Sprintf("%v",slot)
		}else {
			result = fmt.Sprintf("%v , %v", result, slot)
		}
	}
	if slot == 0{
		return "Not found" , nil
	}
	return fmt.Sprintf("%v",slot), nil
}