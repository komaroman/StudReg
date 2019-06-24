package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type Achievements struct {
	ID                 string `json:"id"`
	Olympiads          string `json:"olympiads"`
	Publications       string `json:"publications"`
	PublicPerformances string `json:"publicPerformances"`
	GTO                string `json:"gto"`
	Sports             string `json:"sports"`
	HonorDiplomas      string `json:"honorDiplomas"`
	AverageDiplomaRank string `json:"averageDiplomaRank"`
	EnglishCert        string `json:"englishCert"`
}

type AchievementsInputData struct {
	ID                 string `json:"achievmentid"`
	Olympiads          string `json:"achievmentolympiads"`
	Publications       string `json:"achievmentpublications"`
	PublicPerformances string `json:"achievmentpublicPerformances"`
	GTO                string `json:"achievmentgto"`
	Sports             string `json:"achievmentsports"`
	HonorDiplomas      string `json:"achievmenthonorDiplomas"`
	AverageDiplomaRank string `json:"achievmentaverageDiplomaRank"`
	EnglishCert        string `json:"achievmentenglishCert"`
}

func (achievements *Achievements) Store(stub shim.ChaincodeStubInterface) error {

	bytes, err := json.Marshal(achievements)
	if err != nil {
		return err
	}

	return stub.PutState(achievements.ID, bytes)
}
