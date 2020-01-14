package main

import (
	"fmt"
	"github.com/LimeChain/SupplyChainPOCs"
	"github.com/LimeChain/SupplyChainPOCs/constants"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	err := shim.Start(new(poc01.POC1Chaincode))
	if err != nil {
		fmt.Printf(constants.ErrorStartChaincode, err)
	}
}
