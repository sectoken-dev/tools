package dash_rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/sectoken-dev/godash/wire"
)

func TestNew(t *testing.T) {
	connCfg := &ConnConfig{
		Host:         "192.168.3.138:19997",
		User:         "dash",
		Pass:         "pass",
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
	// _, _ = client.GetBlockCount(ctx)
	bh, _ := wire.NewShaHashFromStr("0000029dcece0728de9f0362e57f0ebda98e59612e86c72c26ffff0af6d6ddd0")
	ret, err  := client.GetBlockVerbose(ctx, bh)
	fmt.Println(ret, err)

	j, _ := json.Marshal(ret)
	fmt.Println(string(j))

	h, _ := wire.NewShaHashFromStr("1b1b0839e9b0f31ff7ab203ad6514977d42c1f122fdfa767a96d44ab2677ac5f")
	tx, _ := client.GetRawTransaction(ctx, h)
	fmt.Println(len(tx.MsgTx().TxIn))

}
