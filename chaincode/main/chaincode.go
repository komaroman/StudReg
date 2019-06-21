package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type MainChaincode struct {
}

func (cc *MainChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	return shim.Success(nil)
}

func (cc *MainChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	function, args := stub.GetFunctionAndParameters()

	switch function {
	case "createStudent":
		return CreateStudent(stub, args)
	case "queryStudent":
		return QueryStudent(stub, args)
	case "updateStudent":
		return UpdateStudent(stub, args)
	default:
		return shim.Error("Unknown function")
	}
}
