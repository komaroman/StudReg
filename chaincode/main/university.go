package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type University struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type UniversityInputData struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (university *University) Store(stub shim.ChaincodeStubInterface) error {

	bytes, err := json.Marshal(university)
	if err != nil {
		return err
	}

	return stub.PutState(university.ID, bytes)
}
