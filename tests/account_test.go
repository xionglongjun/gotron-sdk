package tests

import (
	"context"
	"testing"

	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	// "github.com/fbsobreira/gotron-sdk/pkg/keystore"
)

func testGenerateAddress(t *testing.T) *api.AddressPrKeyPairMessage {
	address, err := wallet.GenerateAddress(context.Background(), &api.EmptyMessage{})
	if err != nil {
		t.Fatal(err)
	}
	return address
}

func TestCreateAccount(t *testing.T) {
	address := testGenerateAddress(t)
	transactionExtention, err := cli.CreateAccount(OWNER_ADDRESS, address.GetAddress())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(transactionExtention.String())
}

func TestGetAccount(t *testing.T) {
	account, err := cli.GetAccount(TEST_ADDRESS)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(account.String())
	t.Log(account.GetType(), account.Balance)
}

func TestCreateTransaction(t *testing.T) {
	transaction, err := wallet.CreateTransaction(context.Background(), &core.TransferContract{
		OwnerAddress: []byte(OWNER_ADDRESS),
		ToAddress:    []byte(TEST_ADDRESS),
		Amount:       100,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(transaction)
}
