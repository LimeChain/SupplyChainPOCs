# Examples

## 1. Transparent Supply Chain

### Goals
The goal of this example is to prove blockchain can be used for tracking of goods
without the need of external system integration or hidden private data.
It involves only chaincode invocation in order to govern digitalization of
assets, their records, orders, and their state transitions through the supply chain.

### Structure
   * [chaincode](transparent-supply-chain/chaincode.go) - extends **AssetBoundChaincode** and **ComposableChaincode**,
   having the functionality to create assets, create records 
   * [chaincode_test](transparent-supply-chain/chaincode_test.go) - tests for chaincode

## 2. Transparent Supply Chain 2
Each record has an additional list of quality certificates.

### Structure
   * [chaincode](transparent-supply-chain-2/chaincode.go) - extends **AssetBoundChaincode** and **ComposableChaincode**,
   having the functionality to create assets, create records 
   * [chaincode_test](transparent-supply-chain-2/chaincode_test.go) - tests for chaincode
