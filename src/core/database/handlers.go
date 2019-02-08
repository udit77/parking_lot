package database

import (
	"github.com/parking_lot/src/commands"
	"strconv"
	"errors"
	"fmt"
)

type ParkingDbRowData struct{
	Id int
	Number string
	Color string
}


func (connector *DBConnector) CreateLot(command *commands.CommandBuilder) error{
	err := dropTable(connector.db)
	if err != nil {
		return err
	}
	err = createTable(connector.db)
	if err != nil{
		return err
	}
	connector.size, err = strconv.Atoi(command.Parameters[0])
	if err != nil{
		return err
	}
	return nil
}


func (connector *DBConnector) Park(command *commands.CommandBuilder) (string, error){
	if connector.size == 0 {
		return "parking lot is not initialized", errors.New("parking lot is not initialized")
	}
	carExists , err := checkCarStatus(connector.db, command.Parameters[0])
	if err != nil {
		return "", err
	}
	if carExists {
		return "car with a same number already exists in parking", nil
	}
	occupancy, err := getOccupancy(connector.db)
	if err != nil {
		return "", err
	}
	if occupancy == connector.size{ //Fully occupied
		return "Sorry, parking lot is full", nil
	}else{
		id, err := getVacantSlot(connector.db)
		if err != nil {
			return "", err
		}
		if id == 0 {
			err := park(connector.db,command.Parameters[0], command.Parameters[1])
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("%v%v","Allocated slot number: ",occupancy+1), nil
		}else{
			err = parkAtVacant(connector.db,command.Parameters[0], command.Parameters[1], id)
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
	err := vacant(connector.db,command.Parameters[0])
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v%v%v","Slot number ", command.Parameters[0], " is free"), nil
}


func (connector *DBConnector) GetStatus(command *commands.CommandBuilder)([]ParkingDbRowData, string, error){
	if connector.size == 0 {
		return []ParkingDbRowData{}, "parking lot is not initialized", errors.New("parking lot is not initialized")
	}
	result, err := getParkingStatus(connector.db)
	return result, "", err
}


func (connector *DBConnector) GetNumbersWithColor(command *commands.CommandBuilder) (string, error){
	if connector.size == 0 {
		return "parking lot is not initialized", errors.New("parking lot is not initialized")
	}
	result, err := getRegistrationForColour(connector.db,command.Parameters[0])
	return result,err
}


func (connector *DBConnector) GetSlotWithNumber(command *commands.CommandBuilder)(string,error){
	if connector.size == 0 {
		return "parking lot is not initialized", errors.New("parking lot is not initialized")
	}
	slot, err := getSlotForNumber(connector.db,command.Parameters[0])
	return slot,err
}


func (connector *DBConnector) GetSlotsWithColor(command *commands.CommandBuilder)(string, error){
	if connector.size == 0 {
		return "parking lot is not initialized", errors.New("parking lot is not initialized")
	}
	slots, err := getSlotsForColor(connector.db,command.Parameters[0])
	return slots,err
}
