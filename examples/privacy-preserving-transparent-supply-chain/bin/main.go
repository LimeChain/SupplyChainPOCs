package main

import (
	"fmt"
	"github.com/LimeChain/SupplyChainPOCs/constants"
	"github.com/LimeChain/SupplyChainPOCs/examples/privacy-preserving-transparent-supply-chain"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	err := shim.Start(new(privacy_preserving_transparent_supply_chain.PPTSCChaincode))
	if err != nil {
		fmt.Printf(constants.ErrorStartChaincode, err)
	}
}
