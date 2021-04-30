package tests

import (
	"context"
	"os"
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
	t.Log(string(account.GetAddress()), account)
	t.Log(account.GetType(), account.Balance)
}

func testCreateTransaction(t *testing.T) *core.Transaction {
	transaction, err := wallet.CreateTransaction(context.Background(), &core.TransferContract{
		OwnerAddress: []byte(OWNER_ADDRESS),
		ToAddress:    []byte(TEST_ADDRESS),
		Amount:       100,
	})
	if err != nil {
		t.Fatal(err)
		os.Exit(0)
	}
	return transaction
}

func TestBroadcastTransaction(t *testing.T) {
	transaction, err := wallet.BroadcastTransaction(context.Background(), testCreateTransaction(t))
	if err != nil {
		t.Fatal(err)
		os.Exit(0)
	}
	t.Log(transaction, "转账返回")
}

func TestEasyTransferByPrivate(t *testing.T) {
	transaction, err := wallet.EasyTransferByPrivate(context.Background(), &api.EasyTransferByPrivateMessage{
		PrivateKey: []byte("41928C9AF0651632157EF27A2CF17CA72C575A4D21"),
		ToAddress:  []byte(TEST_ADDRESS),
	})
	if err != nil {
		t.Error(err)
		t.Fatal(err)
	}
	t.Log(transaction, "转账返回")
}
