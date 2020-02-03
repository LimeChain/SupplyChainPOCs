#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo "Input parameters must be 1 ..."
    exit 1
fi

PEER0_ORG3_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.example.com/peers/peer0.org3.example.com/tls/ca.crt
PEER0_ORG3_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.example.com/users/Admin@org3.example.com/msp
PEER0_ORG3_ADDRESS=peer0.org3.example.com:11051
PEER0_ORG3_LOCALMSPID=Org3MSP

CC_CHANNEL_NAME=mychannel

CA_ORDERER=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

CC_VERSION=v1.0
CC_NAME=pptsccc
CC_PATH=github.com/LimeChain/SupplyChainPOCs/examples/privacy-preserving-transparent-supply-chain/bin

ORDER_ID=$1

echo "Switching to Org3MSP ..."
export CORE_PEER_MSPCONFIGPATH=$PEER0_ORG3_MSPCONFIGPATH
export CORE_PEER_ADDRESS=$PEER0_ORG3_ADDRESS
export CORE_PEER_LOCALMSPID=$PEER0_ORG3_LOCALMSPID
export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG3_CA

export CHANNEL_NAME=$CC_CHANNEL_NAME
export ORDERER_CA=$CA_ORDERER

echo "Installing Chaincode on Org3MSP ..."
peer chaincode install -v $CC_VERSION -n $CC_NAME -p $CC_PATH

QUERY_PRIVATE_PARAMS=$(echo "{\"Collection\":\"Org1Org2PricePerUnit\", \"Key\": \"${ORDER_ID}\"}" | jq 'tojson')
QUERY_PRIVATE_ARGS="{\"Args\":[\"queryPrivate\", ${QUERY_PRIVATE_PARAMS}]}"

echo "Querying unauthorized private data ..."
peer chaincode query -n $CC_NAME -C $CC_CHANNEL_NAME -c "${QUERY_PRIVATE_ARGS}" --tls --cafile $ORDERER_CA
