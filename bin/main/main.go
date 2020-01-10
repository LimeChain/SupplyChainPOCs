package main

import (
	"fmt"
	"github.com/LimeChain/SupplyChainPOCs"
	"github.com/LimeChain/SupplyChainPOCs/constants"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	err := shim.Start(new(chaincode.SupplyChainChaincode))
	if err != nil {
		fmt.Printf(constants.ErrorStartChaincode, err)
	}
}