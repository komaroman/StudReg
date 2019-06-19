package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Student struct {
	ID            string `json:"id"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	MiddleName    string `json:"middleName"`
	PlaceOfBirth  string `json:"placeOfBirth"`
	DateOfBirth   string `json:"dateOfBirth"`
	PassportNum   string `json:"passportNum"`
	MaritalStatus string `json:"maritalStatus"`
	Gender        string `json:"gender"`
}

type StudentInputData struct {
	StudID            string `json:"studId"`
	StudFirstName     string `json:"studFirstName"`
	StudLastName      string `json:"studLastName"`
	StudMiddleName    string `json:"studMiddleName"`
	StudPlaceOfBirth  string `json:"studPlaceOfBirth"`
	StudDateOfBirth   string `json:"studDateOfBirth"`
	StudPassportNum   string `json:"studPassportNum"`
	StudMaritalStatus string `json:"studMaritalStatus"`
	StudGender        string `json:"studGender"`
}

func (student *Student) Store(stub shim.ChaincodeStubInterface) error {

	bytes, err := json.Marshal(student)
	if err != nil {
		return err
	}

	return stub.PutState(student.ID, bytes)
}

func CreateStudent(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) == 0 {
		return shim.Error("Not enough arguments")
	}

	var input StudentInputData
	err := json.Unmarshal([]byte(args[0]), &input)
	if err != nil {
		return shim.Error(err.Error())
	}

	newStudent := Student{
		ID:           input.StudID,
		FirstName:    input.StudFirstName,
		LastName:     input.StudLastName,
		MiddleName:   input.StudMiddleName,
		PlaceOfBirth: input.StudPlaceOfBirth,
		DateOfBirth:  input.StudDateOfBirth,
		PassportNum:  input.StudPassportNum,
		MaritalStatus: input.StudMaritalStatus,
		Gender:       input.StudGender,
	}

	err = newStudent.Store(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func QueryStudent(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) == 0 {
		return shim.Error("Not enough arguments")
	}

	id := args[0]

	bytes, err := stub.GetState(id)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(bytes)
}

func UpdateStudent(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) == 0 {
		return shim.Error("Not enough arguments")
	}

	var input StudentInputData
	err := json.Unmarshal([]byte(args[0]), &input)
	if err != nil {
		return shim.Error(err.Error())
	}

	bytes, err := stub.GetState(input.StudID)
	if err != nil {
		return shim.Error(err.Error())
	}

	var stud Student
	err = json.Unmarshal(bytes, &stud)
	if err != nil {
		return shim.Error(err.Error())
	}

	stud.FirstName = input.StudFirstName
	stud.LastName = input.StudLastName
	stud.MiddleName = input.StudMiddleName
	stud.PlaceOfBirth = input.StudPlaceOfBirth
	stud.DateOfBirth = input.StudDateOfBirth
	stud.PassportNum = input.StudPassportNum
	stud.MaritalStatus = input.StudMaritalStatus
	stud.Gender = input.StudGender

	err = stud.Store(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)

}
