package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type RequestStatus string

const (
	RequestStatusCreated RequestStatus = "CREATED"
	RequestStatusFailed  RequestStatus = "FAILED"
	RequestStatusSuccess RequestStatus = "SUCCESS"
)

type Examinations struct {
	ExamName string `json:"examName"`
	Mark     string `json:"mark"`
}

type EnrollmentRequest struct {
	ID             string         `json:"id"`
	UniversityID   string         `json:"universityId"`
	Status         RequestStatus  `json:"status"`
	CreatedAt      string         `json:"createdAt"`
	ModifiedAt     string         `json:"modifiedAt"`
	Specialization string         `json:"specialization"`
	Examinations   []Examinations `json:"examinations"`
}

type EnrollmentRequestInputData struct {
	ID             string         `json:"id"`
	UniversityID   string         `json:"universityId"`
	Specialization string         `json:"specialization"`
	Examinations   []Examinations `json:"examinations"`
}

func (enrollmentRequest *EnrollmentRequest) Store(stub shim.ChaincodeStubInterface) error {

	bytes, err := json.Marshal(enrollmentRequest)
	if err != nil {
		return err
	}

	return stub.PutState(enrollmentRequest.ID, bytes)
}
