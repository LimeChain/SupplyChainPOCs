package main

import (
	"fmt"
	"github.com/LimeChain/SupplyChainPOCs/constants"
	"github.com/LimeChain/SupplyChainPOCs/examples/poc01_1"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	err := shim.Start(new(poc01_1.POC1_1_Chaincode))
	if err != nil {
		fmt.Printf(constants.ErrorStartChaincode, err)
	}
}
