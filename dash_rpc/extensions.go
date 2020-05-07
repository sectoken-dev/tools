// Copyright (c) 2014-2017 The ltcsuite developers
// Copyright (c) 2015-2017 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package dash_rpc

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sectoken-dev/godash/btcjson"
	"github.com/sectoken-dev/godash/wire"
)

// // FutureDebugLevelResult is a future promise to deliver the result of a
// // DebugLevelAsync RPC invocation (or an applicable error).
// type FutureDebugLevelResult chan *response
// 
// // Receive waits for the response promised by the future and returns the result
// // of setting the debug logging level to the passed level specification or the
// // list of of the available subsystems for the special keyword 'show'.
// func (r FutureDebugLevelResult) Receive() (string, error) {
// 	res, err := receiveFuture(r)
// 	if err != nil {
// 		return "", err
// 	}
// 
// 	// Unmashal the result as a string.
// 	var result string
// 	err = json.Unmarshal(res, &result)
// 	if err != nil {
// 		return "", err
// 	}
// 	return result, nil
// }
// 
// // DebugLevelAsync returns an instance of a type that can be used to get the
// // result of the RPC at some future time by invoking the Receive function on
// // the returned instance.
// //
// // See DebugLevel for the blocking version and more details.
// //
// // NOTE: This is a ltcd extension.
// func (c *Client) DebugLevelAsync(ctx context.Context, levelSpec string) FutureDebugLevelResult {
// 	cmd := btcjson.NewDebugLevelCmd(levelSpec)
// 	return c.sendCmd(ctx, cmd)
// }

// DebugLevel dynamically sets the debug logging level to the passed level
// specification.
//
// The levelspec can be either a debug level or of the form:
// 	<subsystem>=<level>,<subsystem2>=<level2>,...
//
// Additionally, the special keyword 'show' can be used to get a list of the
// available subsystems.
//
// // NOTE: This is a ltcd extension.
// func (c *Client) DebugLevel(ctx context.Context, levelSpec string) (string, error) {
// 	return c.DebugLevelAsync(ctx, levelSpec).Receive()
// }

// FutureCreateEncryptedWalletResult is a future promise to deliver the error
// result of a CreateEncryptedWalletAsync RPC invocation.
type FutureCreateEncryptedWalletResult chan *response

// Receive waits for and returns the error response promised by the future.
func (r FutureCreateEncryptedWalletResult) Receive() error {
	_, err := receiveFuture(r)
	return err
}

// CreateEncryptedWalletAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// See CreateEncryptedWallet for the blocking version and more details.
//
// NOTE: This is a ltcwallet extension.
func (c *Client) CreateEncryptedWalletAsync(ctx context.Context, passphrase string) FutureCreateEncryptedWalletResult {
	cmd := btcjson.NewCreateEncryptedWalletCmd(passphrase)
	return c.sendCmd(ctx, cmd)
}

// CreateEncryptedWallet requests the creation of an encrypted wallet.  Wallets
// managed by ltcwallet are only written to disk with encrypted private keys,
// and generating wallets on the fly is impossible as it requires user input for
// the encryption passphrase.  This RPC specifies the passphrase and instructs
// the wallet creation.  This may error if a wallet is already opened, or the
// new wallet cannot be written to disk.
//
// NOTE: This is a ltcwallet extension.
func (c *Client) CreateEncryptedWallet(ctx context.Context, passphrase string) error {
	return c.CreateEncryptedWalletAsync(ctx, passphrase).Receive()
}

// FutureListAddressTransactionsResult is a future promise to deliver the result
// of a ListAddressTransactionsAsync RPC invocation (or an applicable error).
type FutureListAddressTransactionsResult chan *response

// Receive waits for the response promised by the future and returns information
// about all transactions associated with the provided addresses.
func (r FutureListAddressTransactionsResult) Receive() ([]btcjson.ListTransactionsResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result as an array of listtransactions objects.
	var transactions []btcjson.ListTransactionsResult
	err = json.Unmarshal(res, &transactions)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// ListAddressTransactionsAsync returns an instance of a type that can be used
// to get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// See ListAddressTransactions for the blocking version and more details.
//
// // NOTE: This is a ltcd extension.
// func (c *Client) ListAddressTransactionsAsync(ctx context.Context, addresses []ltcutil.Address, account string) FutureListAddressTransactionsResult {
// 	// Convert addresses to strings.
// 	addrs := make([]string, 0, len(addresses))
// 	for _, addr := range addresses {
// 		addrs = append(addrs, addr.EncodeAddress())
// 	}
// 	cmd := btcjson.NewListAddressTransactionsCmd(addrs, &account)
// 	return c.sendCmd(ctx, cmd)
// }

// ListAddressTransactions returns information about all transactions associated
// with the provided addresses.
//
// // NOTE: This is a ltcwallet extension.
// func (c *Client) ListAddressTransactions(ctx context.Context, addresses []ltcutil.Address, account string) ([]btcjson.ListTransactionsResult, error) {
// 	return c.ListAddressTransactionsAsync(ctx, addresses, account).Receive()
// }

// FutureGetBestBlockResult is a future promise to deliver the result of a
// GetBestBlockAsync RPC invocation (or an applicable error).
type FutureGetBestBlockResult chan *response

// Receive waits for the response promised by the future and returns the hash
// and height of the block in the longest (best) chain.
func (r FutureGetBestBlockResult) Receive() (*wire.ShaHash, int32, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, 0, err
	}

	// Unmarshal result as a getbestblock result object.
	var bestBlock btcjson.GetBestBlockResult
	err = json.Unmarshal(res, &bestBlock)
	if err != nil {
		return nil, 0, err
	}

	// Convert to hash from string.
	hash, err := wire.NewShaHashFromStr(bestBlock.Hash)
	if err != nil {
		return nil, 0, err
	}

	return hash, bestBlock.Height, nil
}

// GetBestBlockAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See GetBestBlock for the blocking version and more details.
//
// NOTE: This is a ltcd extension.
func (c *Client) GetBestBlockAsync(ctx context.Context) FutureGetBestBlockResult {
	cmd := btcjson.NewGetBestBlockCmd()
	return c.sendCmd(ctx, cmd)
}

// GetBestBlock returns the hash and height of the block in the longest (best)
// chain.
//
// NOTE: This is a ltcd extension.
func (c *Client) GetBestBlock(ctx context.Context) (*wire.ShaHash, int32, error) {
	return c.GetBestBlockAsync(ctx).Receive()
}

// FutureGetCurrentNetResult is a future promise to deliver the result of a
// GetCurrentNetAsync RPC invocation (or an applicable error).
type FutureGetCurrentNetResult chan *response

// Receive waits for the response promised by the future and returns the network
// the server is running on.
func (r FutureGetCurrentNetResult) Receive() (wire.BitcoinNet, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return 0, err
	}

	// Unmarshal result as an int64.
	var net int64
	err = json.Unmarshal(res, &net)
	if err != nil {
		return 0, err
	}

	return wire.BitcoinNet(net), nil
}

// GetCurrentNetAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See GetCurrentNet for the blocking version and more details.
//
// NOTE: This is a ltcd extension.
func (c *Client) GetCurrentNetAsync(ctx context.Context) FutureGetCurrentNetResult {
	cmd := btcjson.NewGetCurrentNetCmd()
	return c.sendCmd(ctx, cmd)
}

// GetCurrentNet returns the network the server is running on.
//
// NOTE: This is a ltcd extension.
func (c *Client) GetCurrentNet(ctx context.Context) (wire.BitcoinNet, error) {
	return c.GetCurrentNetAsync(ctx).Receive()
}

//
// // FutureGetHeadersResult is a future promise to deliver the result of a
// // getheaders RPC invocation (or an applicable error).
// //
// // NOTE: This is a ltcsuite extension ported from
// // github.com/decred/dcrrpcclient.
// type FutureGetHeadersResult chan *response
//
// // Receive waits for the response promised by the future and returns the
// // getheaders result.
// //
// // NOTE: This is a ltcsuite extension ported from
// // github.com/decred/dcrrpcclient.
// func (r FutureGetHeadersResult) Receive() ([]wire.BlockHeader, error) {
// 	res, err := receiveFuture(r)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// Unmarshal result as a slice of strings.
// 	var result []string
// 	err = json.Unmarshal(res, &result)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// Deserialize the []string into []wire.BlockHeader.
// 	headers := make([]wire.BlockHeader, len(result))
// 	for i, headerHex := range result {
// 		serialized, err := hex.DecodeString(headerHex)
// 		if err != nil {
// 			return nil, err
// 		}
// 		err = headers[i].Deserialize(bytes.NewReader(serialized))
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	return headers, nil
// }
//
// // GetHeadersAsync returns an instance of a type that can be used to get the result
// // of the RPC at some future time by invoking the Receive function on the returned instance.
// //
// // See GetHeaders for the blocking version and more details.
// //
// // NOTE: This is a ltcsuite extension ported from
// // github.com/decred/dcrrpcclient.
// func (c *Client) GetHeadersAsync(ctx context.Context, blockLocators []wire.ShaHash, hashStop *wire.ShaHash) FutureGetHeadersResult {
// 	locators := make([]string, len(blockLocators))
// 	for i := range blockLocators {
// 		locators[i] = blockLocators[i].String()
// 	}
// 	hash := ""
// 	if hashStop != nil {
// 		hash = hashStop.String()
// 	}
// 	cmd := btcjson.NewGetHeadersCmd(locators, hash)
// 	return c.sendCmd(ctx, cmd)
// }
//
// // GetHeaders mimics the wire protocol getheaders and headers messages by
// // returning all headers on the main chain after the first known block in the
// // locators, up until a block hash matches hashStop.
// //
// // NOTE: This is a ltcsuite extension ported from
// // github.com/decred/dcrrpcclient.
// func (c *Client) GetHeaders(ctx context.Context, blockLocators []wire.ShaHash, hashStop *wire.ShaHash) ([]wire.BlockHeader, error) {
// 	return c.GetHeadersAsync(ctx, blockLocators, hashStop).Receive()
// }

// FutureExportWatchingWalletResult is a future promise to deliver the result of
// an ExportWatchingWalletAsync RPC invocation (or an applicable error).
type FutureExportWatchingWalletResult chan *response

// Receive waits for the response promised by the future and returns the
// exported wallet.
func (r FutureExportWatchingWalletResult) Receive() ([]byte, []byte, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, nil, err
	}

	// Unmarshal result as a JSON object.
	var obj map[string]interface{}
	err = json.Unmarshal(res, &obj)
	if err != nil {
		return nil, nil, err
	}

	// Check for the wallet and tx string fields in the object.
	base64Wallet, ok := obj["wallet"].(string)
	if !ok {
		return nil, nil, fmt.Errorf("unexpected response type for "+
			"exportwatchingwallet 'wallet' field: %T\n",
			obj["wallet"])
	}
	base64TxStore, ok := obj["tx"].(string)
	if !ok {
		return nil, nil, fmt.Errorf("unexpected response type for "+
			"exportwatchingwallet 'tx' field: %T\n",
			obj["tx"])
	}

	walletBytes, err := base64.StdEncoding.DecodeString(base64Wallet)
	if err != nil {
		return nil, nil, err
	}

	txStoreBytes, err := base64.StdEncoding.DecodeString(base64TxStore)
	if err != nil {
		return nil, nil, err
	}

	return walletBytes, txStoreBytes, nil

}

// ExportWatchingWalletAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// See ExportWatchingWallet for the blocking version and more details.
//
// NOTE: This is a ltcwallet extension.
func (c *Client) ExportWatchingWalletAsync(ctx context.Context, account string) FutureExportWatchingWalletResult {
	cmd := btcjson.NewExportWatchingWalletCmd(&account, btcjson.Bool(true))
	return c.sendCmd(ctx, cmd)
}

// ExportWatchingWallet returns the raw bytes for a watching-only version of
// wallet.bin and tx.bin, respectively, for the specified account that can be
// used by ltcwallet to enable a wallet which does not have the private keys
// necessary to spend funds.
//
// NOTE: This is a ltcwallet extension.
func (c *Client) ExportWatchingWallet(ctx context.Context, account string) ([]byte, []byte, error) {
	return c.ExportWatchingWalletAsync(ctx, account).Receive()
}

// FutureSessionResult is a future promise to deliver the result of a
// SessionAsync RPC invocation (or an applicable error).
type FutureSessionResult chan *response

// Receive waits for the response promised by the future and returns the
// session result.
func (r FutureSessionResult) Receive() (*btcjson.SessionResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a session result object.
	var session btcjson.SessionResult
	err = json.Unmarshal(res, &session)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

// SessionAsync returns an instance of a type that can be used to get the result
// of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See Session for the blocking version and more details.
//
// NOTE: This is a ltcsuite extension.
func (c *Client) SessionAsync(ctx context.Context) FutureSessionResult {
	// Not supported in HTTP POST mode.
	if c.config.HTTPPostMode {
		return newFutureError(errors.New("ErrWebsocketsRequired"))
	}

	cmd := btcjson.NewSessionCmd()
	return c.sendCmd(ctx, cmd)
}

//
// // FutureVersionResult is a future promise to deliver the result of a version
// // RPC invocation (or an applicable error).
// //
// // NOTE: This is a ltcsuite extension ported from
// // github.com/decred/dcrrpcclient.
// type FutureVersionResult chan *response
//
// // Receive waits for the response promised by the future and returns the version
// // result.
// //
// // NOTE: This is a ltcsuite extension ported from
// // github.com/decred/dcrrpcclient.
// func (r FutureVersionResult) Receive() (map[string]btcjson.VersionResult,
// 	error) {
// 	res, err := receiveFuture(r)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// Unmarshal result as a version result object.
// 	var vr map[string]btcjson.VersionResult
// 	err = json.Unmarshal(res, &vr)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return vr, nil
// }
//
// // VersionAsync returns an instance of a type that can be used to get the result
// // of the RPC at some future time by invoking the Receive function on the
// // returned instance.
// //
// // See Version for the blocking version and more details.
// //
// // NOTE: This is a ltcsuite extension ported from
// // github.com/decred/dcrrpcclient.
// func (c *Client) VersionAsync(ctx context.Context) FutureVersionResult {
// 	cmd := btcjson.NewVersionCmd()
// 	return c.sendCmd(ctx, cmd)
// }
//
// // Version returns information about the server's JSON-RPC API versions.
// //
// // NOTE: This is a ltcsuite extension ported from
// // github.com/decred/dcrrpcclient.
// func (c *Client) Version(ctx context.Context) (map[string]btcjson.VersionResult, error) {
// 	return c.VersionAsync(ctx).Receive()
// }
