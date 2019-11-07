package ethclient

import (
	"context"
	"fmt"
	"testing"
)

func TestDial(t *testing.T) {
	client, err := Dial("http://eth.sectoken.io", "", "")

	if err != nil {
		panic(err)
	}
	var res string
	err = client.CallContext(context.TODO(), &res, "eth_blockNumber")
	fmt.Println(err)

	fmt.Println(res)
}
