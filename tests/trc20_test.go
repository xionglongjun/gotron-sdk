package tests

import (
	"testing"
)

const (
	// 合约地址
	USDT_CONTRACT = "TXYZopYRdj2D9XRtbG411XZZ3kM5VkAeBf"
)

func TestTRC20GetName(t *testing.T) {
	contract, err := cli.TRC20GetName(USDT_CONTRACT)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(contract)
}

func TestTRC20ContractBalance(t *testing.T) {
	balance, err := cli.TRC20ContractBalance("TUDvBGVduVgAjKpApbsMstmDtJjT7HgkUY", USDT_CONTRACT)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(balance)
}
