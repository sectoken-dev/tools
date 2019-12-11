package bch_rpc

import (
	"context"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	connCfg := &ConnConfig{
		Host:         "192.168.3.196:18332",
		User:         "bchuer",
		Pass:         "bchpass",
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := New(connCfg)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	fmt.Println(client.GetBlockCount(ctx))

}
