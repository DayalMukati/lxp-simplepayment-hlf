package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type PaymentContract struct {
	contractapi.Contract
}

type Payer struct {
	PayerID string  `json:"payerID"`
	Limit   float64 `json:"limit"`
	Status  string  `json:"status"` // Only keeps track of the last status (PaymentAuthorized, PaymentFailed)
}

// SetPaymentLimit sets the payment limit for a payer
func (p *PaymentContract) SetPaymentLimit(ctx contractapi.TransactionContextInterface, payerID string, limit float64) error {
	// Write the logic to set the payment limit for a payer
}

// AuthorizePayment authorizes a payment for a payer if it doesn't exceed the limit
func (p *PaymentContract) AuthorizePayment(ctx contractapi.TransactionContextInterface, payerID string, paymentAmount float64) error {
	// write the logic to authorize a payment for a payer if it doesn't exceed the limit

	// Check if payment exceeds the limit, if so, set status to PaymentFailed
	

	// If payment is within the limit, set status to authorized
	
}

// QueryPaymentStatus returns the last status of the payer
func (p *PaymentContract) QueryPaymentStatus(ctx contractapi.TransactionContextInterface, payerID string) (string, error) {
	// Get the status from the world state
	// Return the current status of the payer
	return payer.Status, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(PaymentContract))
	if err != nil {
		fmt.Printf("Error creating payment contract chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting payment contract chaincode: %s", err.Error())
	}
}
