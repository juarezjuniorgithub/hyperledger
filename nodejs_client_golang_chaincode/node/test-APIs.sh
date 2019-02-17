#!/bin/bash
#
# Get end2end network config from user

#default config
CHANNEL_NAME="default"
CHAINCODE_NAME="golang_cc"
CHAINCDDE_VER="v0"
CHAINCODE_PATH="github.com"
ORG_NAME="default"
#example peers:"peer1.org1.com,peer2.org1.com"
TARGET_PEER="default"

# echo "We need the following details to initialize the ledger config."
# echo  "Enter name of channel [ENTER]: "
# read -e CHANNEL_NAME
# echo -n "Enter name of chaincode [ENTER]: "
# read -e CHAINCODE_NAME

starttime=$(date +%s)

echo "POST Install chaincode on peer $TARGET_PEER"
echo
node end2end.js "install" '{"channelName":"'"$CHANNEL_NAME"'","chaincodeId":"'"$CHAINCODE_NAME"'","chaincodePath":"'"$CHAINCODE_PATH"'","chaincodeVersion":"'"$CHAINCDDE_VER"'","peers":"'"$TARGET_PEER"'","org":"'"$ORG_NAME"'"}'
echo
echo

echo "POST instantiate chaincode on channel $CHANNEL_NAME"
echo
node end2end.js "instantiate" '{"channelName":"'"$CHANNEL_NAME"'","chaincodeId":"'"$CHAINCODE_NAME"'","chaincodePath":"'"$CHAINCODE_PATH"'","chaincodeVersion":"'"$CHAINCDDE_VER"'","peers":"'"$TARGET_PEER"'","org":"'"$ORG_NAME"'","cfcn":"init","cargs":"a,1000,b,1000"}'
echo
echo

echo "POST invoke chaincode on peers of $ORG_NAME"
echo
node end2end.js "invoke" '{"channelName":"'"$CHANNEL_NAME"'","chaincodeId":"'"$CHAINCODE_NAME"'","peers":"'"$TARGET_PEER"'","org":"'"$ORG_NAME"'","cfcn":"invoke","cargs":"move,b,a,1"}'
echo
echo


echo "GET query chaincode on peer1 of $ORG_NAME"
echo
node end2end.js "invoke" '{"channelName":"'"$CHANNEL_NAME"'","chaincodeId":"'"$CHAINCODE_NAME"'","peers":"'"$TARGET_PEER"'","org":"'"$ORG_NAME"'","cfcn":"invoke","cargs":"query,a"}'
echo
echo

echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
