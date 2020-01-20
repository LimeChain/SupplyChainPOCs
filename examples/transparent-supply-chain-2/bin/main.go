package main

import (
	"fmt"
	"github.com/LimeChain/SupplyChainPOCs/constants"
	"github.com/LimeChain/SupplyChainPOCs/examples/transparent-supply-chain-2"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	err := shim.Start(new(transparent_supply_chain_2.TSCChaincode_2))
	if err != nil {
		fmt.Printf(constants.ErrorStartChaincode, err)
	}
}
