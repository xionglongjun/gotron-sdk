package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	// "github.com/fbsobreira/gotron-sdk/pkg/keystore"
)

const (
	OWNER_ADDRESS = "TUDvBGVduVgAjKpApbsMstmDtJjT7HgkUY"
)

func TestGenerateAddress(t *testing.T) {
	address, err := wallet.GenerateAddress(context.Background(), &api.EmptyMessage{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(address)
}

func TestCreateAccount(t *testing.T) {
	address, err := wallet.GenerateAddress(context.Background(), &api.EmptyMessage{})
	if err != nil {
		t.Fatal(err)
	}
	transactionExtention, err := cli.CreateAccount(OWNER_ADDRESS, address.GetAddress())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(transactionExtention.String())
}

func TestGetAccount(t *testing.T) {
	account, err := cli.GetAccount("TUDvBGVduVgAjKpApbsMstmDtJjT7HgkUY")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(account.String())
	t.Log(account.GetType(), account.Balance)
}
