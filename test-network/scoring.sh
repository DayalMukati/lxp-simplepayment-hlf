#!/bin/bash

# Initialize score
score=0
source ./scripts/setOrgPeerContext.sh 1
export PATH=${PWD}/../bin:${PWD}:$PATH
export FABRIC_CFG_PATH=${PWD}/configtx

# Query the payment limit for payer1
echo "Querying the payment limit for payer1..."
QUERY_LIMIT_OUTPUT=$(peer chaincode query -C mychannel -n simplepayment -c '{"Args":["QueryPaymentLimit","payer1"]}' 2>&1)

# Check if the payment limit was set correctly
if [[ $QUERY_LIMIT_OUTPUT == *"5000"* ]]; then
    echo "Payment limit for payer1 is set correctly at 5000."
    score=$((score + 20))
else
    echo "Payment limit not set correctly for payer1."
    echo "Final Score: $score/50"
    exit 0
fi

# Authorize a payment of 3000 for payer1 (this should succeed)
echo "Authorizing a payment of 3000 for payer1..."
AUTH_SUCCESS_OUTPUT=$(peer chaincode query -C mychannel -n simplepayment -c '{"Args":["QueryPaymentStatus","payer1","3000"]}' 2>&1)

if [[ $AUTH_SUCCESS_OUTPUT == *"authorized"* ]]; then
    echo "Payment authorization for 3000 successful."
    score=$((score + 10))
else
    echo "Payment authorization for 3000 failed."
fi

# Attempt to authorize a payment of 6000 for payer1 (this should fail)
echo "Attempting to authorize payment of 6000 (should fail)..."
AUTH_FAIL_OUTPUT=$(peer chaincode invoke -C mychannel -n simplepayment -c '{"Args":["AuthorizePayment","payer1","6000"]}' 2>&1)

if [[ $AUTH_FAIL_OUTPUT == *"status:500"* && $AUTH_FAIL_OUTPUT == *"payment amount exceeds limit"* ]]; then
    echo "Payment authorization for 6000 correctly failed."
    score=$((score + 30))
else
    echo "Payment authorization for 6000 did not fail as expected."
fi

# Query the payment status to verify payment was authorized
echo "Querying payment status..."
QUERY_PAYMENT_OUTPUT=$(peer chaincode query -C mychannel -n simplepayment -c '{"Args":["QueryPaymentStatus","payer1"]}' 2>&1)

if [[ $QUERY_PAYMENT_OUTPUT == *"authorized"* ]]; then
    echo "Payment query successful, payment status found."
    score=$((score + 50))
else
    echo "Payment query failed."
fi

# Final score output
echo "Final Score: $score/50"

# Exit with success
exit 0
