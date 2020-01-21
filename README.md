# SupplyKit

## Overview
SupplyKit is a library for supply chain based Hyperledger Fabric chaincode development. It provides
abstractions of components which developers can extend in order to make their development process easier.

## Architecture
We have provided abstract definitions of the most general entities of a supply chain. An **Asset** represents
the definition of an item, and **Record** is the actual item with its properties and quantity. An **Order** is
an entity, representing the transfer of records between different organizations.

![Example](examples/img/example.png)

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

## Examples

## [1. Transparent Supply Chain](examples/transparent-supply-chain)

### Goal
The goal of this example is to prove blockchain can be used for tracking of goods
without the need of external system integration or hidden private data.
It involves only chaincode invocation in order to govern digitalization of
assets, their records, orders, and their state transitions through the supply chain.

### Structure
   * [chaincode](examples/transparent-supply-chain/chaincode.go) - extends both [AssetBoundChaincode](cc/AssetBoundChaincode.go) and [ComposableChaincode](cc/ComposableChaincode.go),
   having the functionality to create assets, records, orders, and functionality to make complex queries and sell records
   * [chaincode_test](examples/transparent-supply-chain/chaincode_test.go) - tests for **chaincode**

## [2. Transparent Supply Chain 2](examples/transparent-supply-chain-2)
Each record has an additional list of quality certificates.

### Structure
   * [certifiedRecord](examples/transparent-supply-chain-2/certifiedRecord.go) - extends both [AssetBoundRecord](types/record/assetBoundRecord.go) and [ComposableRecord](types/record/composableRecord.go)
   * [chaincode](examples/transparent-supply-chain-2/chaincode.go) - extends both [AssetBoundChaincode](cc/AssetBoundChaincode.go) and [ComposableChaincode](cc/ComposableChaincode.go),
   having the functionality to create assets, records, orders, and functionality to make complex queries and sell records
   * [chaincode_test](examples/transparent-supply-chain-2/chaincode_test.go) - tests for **chaincode**
