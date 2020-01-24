package main

import (
	"fmt"
	"github.com/LimeChain/SupplyChainPOCs/constants"
	"github.com/LimeChain/SupplyChainPOCs/examples/transparent-supply-chain"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	err := shim.Start(new(transparent_supply_chain.TSCChaincode))
	if err != nil {
		fmt.Printf(constants.ErrorStartChaincode, err)
	}
}
