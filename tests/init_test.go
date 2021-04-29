package tests

import (
	"log"
	"testing"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"google.golang.org/grpc"
)

var (
	// addr = "47.252.3.238:50051"
	addr   = "127.0.0.1:50051"
	cli    *client.GrpcClient
	wallet api.WalletClient

	// 主账户
	OWNER_ADDRESS = "TPL66VK2gCXNCD7EJg9pgJRfqcRazjhUZY"
	// 测试账户 TronLink
	TEST_ADDRESS = "TUDvBGVduVgAjKpApbsMstmDtJjT7HgkUY"
)

func TestMain(t *testing.M) {
	cli = client.NewGrpcClient(addr)
	if err := cli.Start(grpc.WithInsecure()); err != nil {
		log.Fatal(err)
	}
	wallet = cli.Client
	t.Run()
}
