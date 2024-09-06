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
}

func (p *PaymentContract) SetPaymentLimit(ctx contractapi.TransactionContextInterface, payerID string, limit float64) error {
	payer := Payer{
		PayerID: payerID,
		Limit:   limit,
	}

	payerAsBytes, err := json.Marshal(payer)
	if err != nil {
		return fmt.Errorf("failed to marshal payer: %s", err.Error())
	}

	return ctx.GetStub().PutState(payerID, payerAsBytes)
}

func (p *PaymentContract) AuthorizePayment(ctx contractapi.TransactionContextInterface, payerID string, paymentAmount float64) (bool, error) {
	payerAsBytes, err := ctx.GetStub().GetState(payerID)
	if err != nil {
		return false, fmt.Errorf("failed to get payer: %s", err.Error())
	}

	if payerAsBytes == nil {
		return false, fmt.Errorf("payer %s does not exist", payerID)
	}

	payer := new(Payer)
	err = json.Unmarshal(payerAsBytes, payer)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal payer: %s", err.Error())
	}

	if paymentAmount > payer.Limit {
		return false, fmt.Errorf("payment amount exceeds limit for payer %s", payerID)
	}

	return true, nil
}

func (p *PaymentContract) QueryPaymentLimit(ctx contractapi.TransactionContextInterface, payerID string) (float64, error) {
	payerAsBytes, err := ctx.GetStub().GetState(payerID)
	if err != nil {
		return 0, fmt.Errorf("failed to get payer: %s", err.Error())
	}

	if payerAsBytes == nil {
		return 0, fmt.Errorf("payer %s does not exist", payerID)
	}

	payer := new(Payer)
	err = json.Unmarshal(payerAsBytes, payer)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal payer: %s", err.Error())
	}

	return payer.Limit, nil
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
