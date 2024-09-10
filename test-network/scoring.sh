#!/bin/bash

# Initialize score
score=0
source ./scripts/setOrgPeerContext.sh 1
export PATH=${PWD}/../bin:${PWD}:$PATH
export FABRIC_CFG_PATH=${PWD}/configtx

# Step 1: Query the payment status for payer1
echo "Querying the payment status for payer1..."
QUERY_STATUS_OUTPUT=$(peer chaincode query -C mychannel -n simplepayment -c '{"Args":["QueryPaymentStatus","payer1"]}' 2>&1)

# Check if the payment limit was set ("LimitSet")
if [[ $QUERY_STATUS_OUTPUT == *"LimitSet"* ]]; then
    echo "Payment limit for payer1 was set successfully."
    score=$((score + 15))
else
    echo "Payment limit not set correctly for payer1."
fi

# Step 2: Check if a payment of 3000 was authorized ("PaymentAuthorized")
echo "Checking if payment of 3000 was authorized..."
if [[ $QUERY_STATUS_OUTPUT == *"PaymentAuthorized"* ]]; then
    echo "Payment of 3000 was successfully authorized."
    score=$((score + 30))
else
    echo "Payment of 3000 was not authorized."
fi

# Step 3: Check if a payment failed due to exceeding the limit ("PaymentFailed")
echo "Checking if any payment failed due to exceeding the limit..."
if [[ $QUERY_STATUS_OUTPUT == *"PaymentFailed"* ]]; then
    echo "A payment attempt failed due to exceeding the limit."
    score=$((score + 50))
else
    echo "No payment failures due to exceeding the limit."
fi

# Final score output
echo "Final Score: $score/50"

# Exit with success
exit 0
