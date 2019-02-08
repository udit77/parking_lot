package database

import (
	"database/sql"
	"fmt"
)

func dropTable(db *sql.DB) error{
	statement, err := db.Prepare("DROP TABLE IF EXISTS PARKING")
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil{
		return err
	}
	return nil
}

func createTable(db *sql.DB)error{
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS PARKING (id INTEGER PRIMARY KEY AUTOINCREMENT, registration_number VARCHAR(64), colour VARCHAR(64), status VARCHAR(64))")
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		return err
	}
	return nil
}

func checkCarStatus(db *sql.DB, number string) (bool, error){
	var count int
	statement, err :=  db.Prepare("SELECT COUNT(*) AS COUNT FROM PARKING WHERE registration_number = ? AND STATUS = ?")
	if err != nil {
		return false, err
	}
	rows , err := statement.Query(number,"OCCUPIED")
	if err != nil{
		return false, err
	}
	for rows.Next(){
		rows.Scan(&count)
	}
	if count == 0{
		return false, nil
	}
	return true, nil
}


func getOccupancy(db *sql.DB)(int, error){
	var count int
	statement, err := db.Prepare("SELECT COUNT(*) AS COUNT FROM PARKING WHERE STATUS = ?")
	if err != nil {
		return count, err
	}
	rows , err := statement.Query("OCCUPIED")
	if err != nil{
		return count, err
	}
	for rows.Next(){
		rows.Scan(&count)
	}
	return count, nil
}


func getVacantSlot(db *sql.DB)(int,error){
	var id int
	statement, err := db.Prepare("SELECT id FROM PARKING WHERE STATUS = ? ORDER BY id ASC LIMIT 1")
	if err != nil {
		return id, err
	}
	rows , err := statement.Query("VACANT")
	if err != nil{
		return id, err
	}
	for rows.Next(){
		rows.Scan(&id)
	}
	return id,nil
}


func park(db *sql.DB, number string, color string)error{
	statement, err := db.Prepare("INSERT INTO PARKING (registration_number, colour, status) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_ , err = statement.Exec(number,color,"OCCUPIED")
	if err != nil{
		return err
	}
	return nil
}


func parkAtVacant(db *sql.DB, number string, color string, id int)error{
	statement, err := db.Prepare("UPDATE PARKING SET registration_number = ?, colour = ?, status = ? where id = ?")
	if err != nil {
		return err
	}
	_ , err = statement.Exec(number, color, "OCCUPIED", id)
	if err != nil{
		return err
	}
	return nil
}


func vacant(db *sql.DB, id string)error{
	statement, err := db.Prepare("UPDATE PARKING set status=? where id=?")
	if err != nil {
		return err
	}
	_, err = statement.Exec("VACANT", id)
	if err != nil {
		return err
	}
	return nil
}


func getParkingStatus(db *sql.DB)([]ParkingDbRowData, error){
	var (
		id int
		number string
		colour string
		status string
		result []ParkingDbRowData
	)
	statement, err := db.Prepare("SELECT * FROM PARKING WHERE STATUS = ?")
	if err != nil {
		return result, err
	}
	rows , err := statement.Query("OCCUPIED")
	if err != nil {
		return result, err
	}
	for rows.Next(){
		_ = rows.Scan(&id, &number, &colour, &status)
		result = append(result, ParkingDbRowData{id,number,colour})
	}
	return result,nil
}


func getRegistrationForColour(db *sql.DB, colour string)(string,error){
	var (
		result string
		number string
	)
	statement, err := db.Prepare("SELECT registration_number FROM PARKING WHERE COLOUR = ? AND STATUS = ?")
	if err != nil {
		return result, err
	}
	rows , err := statement.Query(colour,"OCCUPIED")
	if err != nil {
		return result, err
	}
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


func getSlotForNumber(db *sql.DB, registration string)(string, error){
	var slot int
	statement, err := db.Prepare("SELECT id FROM PARKING WHERE registration_number = ? AND STATUS = ?")
	if err != nil {
		return "",err
	}
	rows , err := statement.Query(registration,"OCCUPIED")
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


func getSlotsForColor(db *sql.DB, colour string)(string,error){
	var (
		slot int
		result string
	)
	statement, err := db.Prepare("SELECT id FROM PARKING WHERE colour = ? AND STATUS = ?")
	if err != nil {
		return "",err
	}
	rows , err := statement.Query(colour,"OCCUPIED")
	if err != nil {
		return "",err
	}
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