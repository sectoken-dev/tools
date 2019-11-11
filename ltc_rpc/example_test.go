package ltc_rpc

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	connCfg := &ConnConfig{
		Host:         "192.168.3.000:19332",
		User:         "ltc",
		Pass:         "ltc",
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := New(connCfg)
	if err != nil {

	}

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second)

	fmt.Println(client.GetBlockCount(ctx))

}
