# SupplyKit

## Overview
SupplyKit is a library for supply chain based Hyperledger Fabric chaincode development. It provides
abstractions of components which developers can extend in order to make their development process easier.

## Structure

* [chaincodes](cc)
    * [BaseSupplyChainChaincode](cc/BaseSupplyChainChaincode.go) - base chaincode, having the functionality to create records, create orders and fulfill them
    * [AssetBoundChaincode](cc/AssetBoundChaincode.go) - chaincode, having an additional functionality to create assets and extends **BaseSupplyChainChaincode**'s
    creation of records and orders
    * [ComposableChaincode](cc/ComposableChaincode.go) - chaincode, having the functionality to compose a record based on other records and extends **BaseSupplyChainChaincode**'s
    creation of records and orders
    * All **chaincodes** are written in a way that leaves developers decide how they want to store the data in the world state
* [constants](constants) - declarations of constants, used in chaincode development
* [examples](examples) - examples, based on [chaincodes](cc) & [types](types)
* [types](types)
    * [asset](types/asset) - package, storing all asset related structs
    * [dto](types/dto) - package, storing all DTOs
    * [order](types/order) - package, storing all order related structs
    * [record](types/record) - package, storing all record related structs 
* [utils](utils) - helper functions for querying and creating objects

## [Chaincode Examples](examples/README.md)