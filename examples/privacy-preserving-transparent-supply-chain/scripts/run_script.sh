#!/bin/bash

PEER0_ORG1_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
PEER0_ORG1_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
PEER0_ORG1_ADDRESS=peer0.org1.example.com:7051
PEER0_ORG1_LOCALMSPID=Org1MSP

PEER0_ORG2_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
PEER0_ORG2_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
PEER0_ORG2_ADDRESS=peer0.org2.example.com:9051
PEER0_ORG2_LOCALMSPID=Org2MSP

CC_CHANNEL_NAME=mychannel

CA_ORDERER=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

CC_VERSION=v1.0
CC_NAME=pptsccc
CC_PATH=github.com/LimeChain/SupplyChainPOCs/examples/privacy-preserving-transparent-supply-chain/bin

PDC_CONFIG=/opt/gopath/src/github.com/LimeChain/SupplyChainPOCs/examples/privacy-preserving-transparent-supply-chain/collections_config.json

echo "Switching to Org1MSP ..."
export CORE_PEER_MSPCONFIGPATH=$PEER0_ORG1_MSPCONFIGPATH
export CORE_PEER_ADDRESS=$PEER0_ORG1_ADDRESS
export CORE_PEER_LOCALMSPID=$PEER0_ORG1_LOCALMSPID
export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA

export CHANNEL_NAME=$CC_CHANNEL_NAME
export ORDERER_CA=$CA_ORDERER

INIT_PARAMS='{"Args": ["init", "Org1MSP", "Org2MSP", "Org3MSP"]}'

echo "Installing chaincode on Org1MSP ..."
peer chaincode install -n $CC_NAME -v $CC_VERSION -p $CC_PATH
sleep 10

INIT_ARGS="{\"Args\":[\"init\", \"Org1MSP\", \"Org2MSP\", \"Org3MSP\"]}"
echo "Instantiating chaincode on Org1MSP ..."
peer chaincode instantiate -n $CC_NAME -v $CC_VERSION -c "${INIT_ARGS}" --collections-config /opt/gopath/src/github.com/LimeChain/SupplyChainPOCs/examples/privacy-preserving-transparent-supply-chain/collections_config.json --tls --cafile $ORDERER_CA -C mychannel

echo "Waiting for instantiation request to be committed ..."
sleep 10

ASSET_PARAMS=$(echo "{\"description\": \"saddle\", \"isActive\": true}" | jq 'tojson')
ADD_ASSET_PARAMS="{\"Args\":[\"addAsset\", ${ASSET_PARAMS}]}"

peer chaincode invoke -n $CC_NAME -C $CC_CHANNEL_NAME -c "${ADD_ASSET_PARAMS}" --tls --cafile $ORDERER_CA >& log.txt
ASSET_ID=$(cat log.txt | awk -F "payload:" '{print $2}' | jq 'fromjson | .id')
rm log.txt

RECORD_PARAMS=$(echo "{\"assetId\": ${ASSET_ID}, \"batchId\": \"123\", \"quantity\": 2}" | jq 'tojson')
MANUFACTURE_PARAMS="{\"Args\":[\"manufacture\", ${RECORD_PARAMS}]}"

sleep 10
peer chaincode invoke -n $CC_NAME -C $CC_CHANNEL_NAME -c "${MANUFACTURE_PARAMS}" --tls --cafile $ORDERER_CA >& log.txt

sleep 10

RECORD_ID=$(cat log.txt | awk -F "payload:" '{print $2}' | jq 'fromjson | .id')
rm log.txt

# Switch to Org2
echo "Switching to Org2MSP ..."

export CORE_PEER_MSPCONFIGPATH=$PEER0_ORG2_MSPCONFIGPATH
export CORE_PEER_ADDRESS=$PEER0_ORG2_ADDRESS
export CORE_PEER_LOCALMSPID=$PEER0_ORG2_LOCALMSPID
export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG2_CA

echo "Installing Chaincode on Org2MSP ..."
peer chaincode install -n $CC_NAME -v $CC_VERSION -p $CC_PATH
sleep 10

ORDER_PARAMS=$(echo "{\"assetId\": ${ASSET_ID}, \"sellerId\": \"Org1MSP\", \"quantity\": 2, \"pricePerUnit\": 50.0}" | jq 'tojson')
PLACE_ORDER_PARAMS="{\"Args\":[\"placeOrder\", ${ORDER_PARAMS}]}"

echo "Invoking \"Place Order\" ..."
peer chaincode invoke -n $CC_NAME -C $CC_CHANNEL_NAME -c "${PLACE_ORDER_PARAMS}" --tls --cafile $ORDERER_CA >& log.txt
sleep 10

ORDER_ID=$(cat log.txt | awk -F "payload": '{print $2}' | jq 'fromjson | .id')
rm log.txt
echo $ORDER_ID

QUERY_PRIVATE_PARAMS=$(echo "{\"Collection\": \"Org1Org2PricePerUnit\", \"Key\": ${ORDER_ID}}" | jq 'tojson')
QUERY_PRIVATE_ARGS="{\"Args\":[\"queryPrivate\", ${QUERY_PRIVATE_PARAMS}]}"

sleep 10
echo "Querying private data ..."
peer chaincode query -n $CC_NAME -C $CC_CHANNEL_NAME -c "${QUERY_PRIVATE_ARGS}" --tls --cafile $ORDERER_CA
sleep 10

# Switch to Org1
echo "Switching to Org1MSP ..."
export CORE_PEER_MSPCONFIGPATH=$PEER0_ORG1_MSPCONFIGPATH
export CORE_PEER_ADDRESS=$PEER0_ORG1_ADDRESS
export CORE_PEER_LOCALMSPID=$PEER0_ORG1_LOCALMSPID
export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA

echo "Querying private data ..."
peer chaincode query -n $CC_NAME -C $CC_CHANNEL_NAME -c "${QUERY_PRIVATE_ARGS}" --tls --cafile $ORDERER_CA
sleep 10

FULFILL_ORDER_PARAMS=$(echo "{\"id\": ${ORDER_ID}, \"status\": true, \"records\":[{\"id\": ${RECORD_ID}, \"quantity\": 2}]}" | jq 'tojson')
FULFILL_ORDER_ARGS="{\"Args\": [\"fulfillOrder\", ${FULFILL_ORDER_PARAMS}]}"

echo "Fulfilling Order ..."
peer chaincode invoke -n $CC_NAME -C $CC_CHANNEL_NAME -c "${FULFILL_ORDER_ARGS}" --tls --cafile $ORDERER_CA
sleep 10

# Query Records (owner is Org2MSP)
echo "Switching back to Org2MSP ..."
export CORE_PEER_MSPCONFIGPATH=$PEER0_ORG1_MSPCONFIGPATH
export CORE_PEER_ADDRESS=$PEER0_ORG1_ADDRESS
export CORE_PEER_LOCALMSPID=$PEER0_ORG1_LOCALMSPID
export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA

QUERY_ORG2_RECORDS_PARAMS=$(echo "{\"selector\": {\"owner\": \"Org2MSP\"}}" | jq 'tojson')
QUERY_ORG2_RECORDS_ARGS="{\"Args\": [\"query\", ${QUERY_ORG2_RECORDS_PARAMS}]}"

echo "Querying Org2MSP records ..."
peer chaincode query -n $CC_NAME -C $CC_CHANNEL_NAME -c "${QUERY_ORG2_RECORDS_ARGS}" --tls --cafile $ORDERER_CA
