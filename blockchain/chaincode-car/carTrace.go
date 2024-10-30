package main

import (
	"carlife-chaincode-go/chaincode"
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	carChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating car chaincode: %v", err)
	}
	if err := carChaincode.Start(); err != nil {
		log.Panicf("Error starting car chaincode: %v", err)
	}
}
