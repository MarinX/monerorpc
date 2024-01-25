# Monero RPC Client in Go

[![Audit](https://github.com/MarinX/monerorpc/actions/workflows/audit.yml/badge.svg)](https://github.com/MarinX/monerorpc/actions/workflows/audit.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/MarinX/monerorpc)](https://goreportcard.com/report/github.com/MarinX/monerorpc)
[![GoDoc](https://godoc.org/github.com/MarinX/monerorpc?status.svg)](https://godoc.org/github.com/MarinX/monerorpc)
[![License MIT](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](LICENSE)

Full Monero RPC client(Wallet AND Daemon) written in go

## Version

Client was written per docs on getmonero.org.

## Preposition

Running monerod with RPC enabled

## Installation

```sh
go get github.com/MarinX/monerorpc
```

## Usage

```go
import "github.com/MarinX/monerorpc"
```

## Wallet example

```go
package main

import (
	"fmt"

	"github.com/MarinX/monerorpc"
)

func main() {
	// create a new client for Testnet
	client := monerorpc.New(monerorpc.TestnetURI, nil)

	// if your monerod is protected, set username/password
	client.SetAuth("username", "password")

	// call RPC endpoint
	ver, err := client.Wallet.GetVersion()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Wallet version %d\n", ver.Version)
}


```

## Wallet methods

```go
type Wallet interface {
	// SetDaemon connects the RPC server to a Monero daemon.
	SetDaemon(req *SetDaemonRequest) error
	// GetBalance Return the wallet's balance.
	GetBalance(req *GetBalanceRequest) (*GetBalanceResponse, error)
	// GetAddress Return the wallet's addresses for an account. Optionally filter for specific set of subaddresses.
	GetAddress(req *GetAddressRequest) (*GetAddressResponse, error)
	// GetAddressIndex Get account and address indexes from a specific (sub)address
	GetAddressIndex(req *GetAddressIndexRequest) (*GetAddressIndexResponse, error)
	// CreateAddress Create a new address for an account. Optionally, label the new address.
	CreateAddress(req *CreateAddressRequest) (*CreateAddressResponse, error)
	// LabelAddress Label an address.
	LabelAddress(req *LabelAddressRequest) error
	// ValidateAddress Analyzes a string to determine whether it is a valid monero wallet address and returns the result and the address specifications.
	ValidateAddress(req *ValidateAddressRequest) (*ValidateAddressResponse, error)
	// GetAccount Get all accounts for a wallet. Optionally filter accounts by tag.
	GetAccounts(req *GetAccountsRequest) (*GetAccountsResponse, error)
	// CreateAccount Create a new account with an optional label.
	CreateAccount(req *CreateAccountRequest) (*CreateAccountResponse, error)
	// LabelAccount Label an account.
	LabelAccount(req *LabelAccountRequest) error
	// GetAccountTags Get a list of user-defined account tags.
	GetAccountTags() (*GetAccountTagsResponse, error)
	// TagAccounts Apply a filtering tag to a list of accounts.
	TagAccounts(req *TagAccountsRequest) error
	// UntagAccount Remove filtering tag from a list of accounts.
	UntagAccounts(req *UntagAccountsRequest) error
	// SetAccountTagDescription Set description for an account tag.
	SetAccountTagDescription(req *SetAccountTagDescriptionRequest) error
	// GetHeight Returns the wallet's current block height.
	GetHeight() (*GetHeightResponse, error)
	// Transfer Send monero to a number of recipients.
	Transfer(req *TransferRequest) (*TransferResponse, error)
	// TransferSplit Same as transfer, but can split into more than one tx if necessary.
	TransferSplit(req *TransferSplitRequest) (*TransferSplitResponse, error)
	// SignTransfer Sign a transaction created on a read-only wallet (in cold-signing process)
	SignTransfer(req *SignTransferRequest) (*SignTransferResponse, error)
	// SubmitTransfer Submit a previously signed transaction on a read-only wallet (in cold-signing process)
	SubmitTransfer(req *SubmitTransferRequest) (*SubmitTransferResponse, error)
	// SweepDust Send all dust outputs back to the wallet's, to make them easier to spend (and mix).
	SweepDust(req *SweepDustRequest) (*SweepDustResponse, error)
	// SweepAll Send all unlocked balance to an address.
	SweepAll(req *SweepAllRequest) (*SweepAllResponse, error)
	// SweepSingle Send all of a specific unlocked output to an address.
	SweepSingle(req *SweepSingleRequest) (*SweepSingleResponse, error)
	// RelaxTx Relay a transaction previously created with "do_not_relay":true.
	RelayTx(req *RelayTxRequest) (*RelayTxResponse, error)
	// Store Save the wallet file.
	Store() error
	// GetPayments Get a list of incoming payments using a given payment id.
	GetPayments(req *GetPaymentsRequest) (*GetPaymentsResponse, error)
	// GetBulkPayments Get a list of incoming payments using a given payment id, or a list of payments ids, from a given height.
	// This method is the preferred method over get_payments because it has the same functionality but is more extendable.
	// Either is fine for looking up transactions by a single payment ID.
	GetBulkPayments(req *GetBulkPaymentsRequest) (*GetBulkPaymentsResponse, error)
	// IncomingTransfers Return a list of incoming transfers to the wallet.
	IncomingTransfers(req *IncomingTransfersRequest) (*IncomingTransfersResponse, error)
	// QueryKey Return the spend or view private key.
	QueryKey(req *QueryKeyRequest) (*QueryKeyResponse, error)
	// MakeIntegratedAddress Make an integrated address from the wallet address and a payment id.
	MakeIntegratedAddress(req *MakeIntegratedAddressRequest) (*MakeIntegratedAddressResponse, error)
	// SplitIntegratedAddress Retrieve the standard address and payment id corresponding to an integrated address.
	SplitIntegratedAddress(req *SplitIntegratedAddressRequest) (*SplitIntegratedAddressResponse, error)
	// StopWallet Stops the wallet, storing the current state.
	StopWallet() error
	// RescanBlockchain Rescan the blockchain from scratch, losing any information which can not be recovered from the blockchain itself.
	// This includes destination addresses, tx secret keys, tx notes, etc.
	RescanBlockchain() error
	// SetTxNotes Set arbitrary string notes for transactions.
	SetTxNotes(req *SetTxNotesRequest) error
	// GetTxNotes Get string notes for transactions.
	GetTxNotes(req *GetTxNotesRequest) (*GetTxNotesResponse, error)
	// SetAttribute Set arbitrary attribute.
	SetAttribute(req *SetAttributeRequest) error
	// GetAttribute Get attribute value by name.
	GetAttribute(req *GetAttributeRequest) (*GetAttributeResponse, error)
	// GetTxKey Get transaction secret key from transaction id.
	GetTxKey(req *GetTxKeyRequest) (*GetTxKeyResponse, error)
	// CheckTxKey Check a transaction in the blockchain with its secret key.
	CheckTxKey(req *CheckTxKeyRequest) (*CheckTxKeyResponse, error)
	// GetTxProof Get transaction signature to prove it.
	GetTxProof(req *GetTxProofRequest) (*GetTxProofResponse, error)
	// CheckTxProof Prove a transaction by checking its signature.
	CheckTxProof(req *CheckTxProofRequest) (*CheckTxProofResponse, error)
	// GetSpendProof Generate a signature to prove a spend. Unlike proving a transaction, it does not requires the destination public address.
	GetSpendProof(req *GetSpendProofRequest) (*GetSpendProofResponse, error)
	// CheckSpendProof Prove a spend using a signature. Unlike proving a transaction, it does not requires the destination public address.
	CheckSpendProof(req *CheckSpendProofRequest) (*CheckSpendProofResponse, error)
	// GetReserveProof Generate a signature to prove of an available amount in a wallet.
	GetReserveProof(req *GetReserveProofRequest) (*GetReserveProofResponse, error)
	// CheckReserveProof Proves a wallet has a disposable reserve using a signature.
	CheckReserveProof(req *CheckReserveProofRequest) (*CheckReserveProofResponse, error)
	// GetTransfers Returns a list of transfers.
	GetTransfers(req *GetTransfersRequest) (*GetTransfersResponse, error)
	// GetTransferByTxid Show information about a transfer to/from this address.
	GetTransferByTxid(req *GetTransferByTxidRequest) (*GetTransferByTxidResponse, error)
	// DescribeTransfer Returns details for each transaction in an unsigned or multisig transaction set.
	DescribeTransfer(req *DescribeTransferRequest) (*DescribeTransferResponse, error)
	// Sign a string.
	Sign(req *SignRequest) (*SignResponse, error)
	// Verify a signature on a string.
	Verify(req *VerifyRequest) (*VerifyResponse, error)
	// ExportOutputs Export all outputs in hex format.
	ExportOutputs(req *ExportOutputsRequest) (*ExportOutputsResponse, error)
	// ImportOutputs Import outputs in hex format.
	ImportOutputs(req *ImportOutputsRequest) (*ImportOutputsResponse, error)
	// ExportKeyImages Export a signed set of key images.
	ExportKeyImages(req *ExportKeyImagesRequest) (*ExportKeyImagesResponse, error)
	// ImportKeyImages Import signed key images list and verify their spent status.
	ImportKeyImages(req *ImportKeyImagesRequest) (*ImportKeyImagesResponse, error)
	// MakeURI Create a payment URI using the official URI spec.
	MakeURI(req *MakeURIRequest) (*MakeURIResponse, error)
	// ParseURI Parse a payment URI to get payment information.
	ParseURI(req *ParseURIRequest) (*ParseURIResponse, error)
	// GetAddressBook Retrieves entries from the address book.
	GetAddressBook(req *GetAddressBookRequest) (*GetAddressBookResponse, error)
	// AddAddressBook Add an entry to the address book.
	AddAddressBook(req *AddAddressBookRequest) (*AddAddressBookResponse, error)
	// EditAddressBook Edit an existing address book entry.
	EditAddressBook(req *EditAddressBookRequest) error
	// DeleteAddressBook Delete an entry from the address book
	DeleteAddressBook(req *DeleteAddressBookRequest) error
	// Refresh a wallet after openning.
	Refresh(req *RefreshRequest) (*RefreshResponse, error)
	// AutoRefresh Set whether and how often to automatically refresh the current wallet.
	AutoRefresh(req *AutoRefreshRequest) error
	// RescanSpent Rescan the blockchain for spent outputs.
	RescanSpent() error
	// StartMining Start mining in the Monero daemon.
	StartMining(req *StartMiningRequest) error
	// StopMining Stop mining in the Monero daemon.
	StopMining() error
	// GetLanguages Get a list of available languages for your wallet's seed.
	GetLanguages() (*GetLanguagesResponse, error)
	// CreateWallet Create a new wallet. You need to have set the argument "–wallet-dir" when launching monero-wallet-rpc to make this work.
	CreateWallet(req *CreateWalletRequest) error
	// GenerateFromKeys Restores a wallet from a given wallet address, view key, and optional spend key.
	GenerateFromKeys(req *GenerateFromKeysRequest) (*GenerateFromKeysResponse, error)
	// OpenWallet Open a wallet. You need to have set the argument "–wallet-dir" when launching monero-wallet-rpc to make this work.
	OpenWallet(req *OpenWalletRequest) error
	// RestoreDeterministicWallet Create and open a wallet on the RPC server from an existing mnemonic phrase and close the currently open wallet.
	RestoreDeterministicWallet(req *RestoreDeterministicWalletRequest) (*RestoreDeterministicWalletResponse, error)
	// CloseWallet Close the currently opened wallet, after trying to save it.
	CloseWallet() error
	// ChangeWalletPassword Change a wallet password.
	ChangeWalletPassword(req *ChangeWalletPasswordRequest) error
	// IsMultisig Check if a wallet is a multisig one.
	IsMultisig() (*IsMultisigResponse, error)
	// PrepareMultisig Prepare a wallet for multisig by generating a multisig string to share with peers.
	PrepareMultisig() (*PrepareMultisigResponse, error)
	// MakeMultisig Make a wallet multisig by importing peers multisig string.
	MakeMultisig(req *MakeMultisigRequest) (*MakeMultisigResponse, error)
	// ExportMultisigInfo Export multisig info for other participants.
	ExportMultisigInfo() (*ExportMultisigInfoResponse, error)
	// ImportMultisigInfo Import multisig info from other participants.
	ImportMultisigInfo(req *ImportMultisigInfoRequest) (*ImportMultisigInfoResponse, error)
	// FinalizeMultisig Turn this wallet into a multisig wallet, extra step for N-1/N wallets.
	FinalizeMultisig(req *FinalizeMultisigRequest) (*FinalizeMultisigResponse, error)
	// SignMultisig Sign a transaction in multisig.
	SignMultisig(req *SignMultisigRequest) (*SignMultisigResponse, error)
	// SubmitMultisig Submit a signed multisig transaction.
	SubmitMultisig(req *SubmitMultisigRequest) (*SubmitMultisigResponse, error)
	// GetVersion Get RPC version Major & Minor integer-format, where Major is the first 16 bits and Minor the last 16 bits.
	GetVersion() (*GetVersionResponse, error)
}
```

## Daemon example

```go
package main

import (
	"fmt"

	"github.com/MarinX/monerorpc"
)

func main() {
	// create a new client for Testnet
	client := monerorpc.New(monerorpc.TestnetURI, nil)

	// if your monerod is protected, set username/password
	client.SetAuth("username", "password")

	// call RPC endpoint
	block, err := client.Daemon.GetBlockCount()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Block count %d\n", block.Count)
}

```

## Daemon methods

```go
type Daemon interface {
	GenerateBlocks(req *GenerateBlocksRequest) (*GenerateBlocksResponse, error)
	// GetBlockCount Look up how many blocks are in the longest chain known to the node.
	GetBlockCount() (*GetBlockCountResponse, error)
	// OnGetBlockHash Look up a block's hash by its height.
	OnGetBlockHash(req []uint64) (string, error)
	// GetBlockTemplate Get a block template on which mining a new block.
	GetBlockTemplate(req *GetBlockTemplateRequest) (*GetBlockTemplateResponse, error)
	// SubmitBlock Submit a mined block to the network.
	SubmitBlock(req []string) (*SubmitBlockResponse, error)
	// GetLastBlockHeader Block header information for the most recent block is easily retrieved with this method. No inputs are needed.
	GetLastBlockHeader() (*GetLastBlockHeaderResponse, error)
	// GetBlockHeaderByHash Block header information can be retrieved using either a block's hash or height.
	// This method includes a block's hash as an input parameter to retrieve basic information about the block.
	GetBlockHeaderByHash(req *GetBlockHeaderByHashRequest) (*GetBlockHeaderByHashResponse, error)
	// GetBlockHeaderByHeight Similar to get_block_header_by_hash above.
	// This method includes a block's height as an input parameter to retrieve basic information about the block.
	GetBlockHeaderByHeight(req *GetBlockHeaderByHeightRequest) (*GetBlockHeaderByHeightResponse, error)
	// GetBlockHeadersRange Similar to get_block_header_by_height above, but for a range of blocks.
	// This method includes a starting block height and an ending block height as parameters to retrieve basic information about the range of blocks.
	GetBlockHeadersRange(req *GetBlockHeadersRangeRequest) (*GetBlockHeadersRangeResponse, error)
	// GetBlock Full block information can be retrieved by either block height or hash, like with the above block header calls.
	// For full block information, both lookups use the same method, but with different input parameters.
	GetBlock(req *GetBlockRequest) (*GetBlockResponse, error)
	// GetConnections Retrieve information about incoming and outgoing connections to your node.
	GetConnections() (*GetConnectionsResponse, error)
	// GetInfo Retrieve general information about the state of your node and the network.
	GetInfo() (*GetInfoResponse, error)
	// HardForkInfo Look up information regarding hard fork voting and readiness.
	HardForkInfo() (*HardForkInfoResponse, error)
	// SetBans Ban another node by IP.
	SetBans(req *SetBansRequest) (*SetBansResponse, error)
	// GetBans Get list of banned IPs.
	GetBans() (*GetBansResponse, error)
	// FlushTxpool Flush tx ids from transaction pool
	FlushTxpool(req *FlushTxpoolRequest) (*FlushTxpoolResponse, error)
	// GetOutputHistogram Get a histogram of output amounts. For all amounts (possibly filtered by parameters), gives the number of outputs on the chain for that amount.
	// RingCT outputs counts as 0 amount.
	GetOutputHistogram(req *GetOutputHistogramRequest) (*GetOutputHistogramResponse, error)
	// GetVersion Give the node current version
	GetVersion() (*GetVersionResponse, error)
	// GetCoinbaseTxSum Get the coinbase amount and the fees amount for n last blocks starting at particular height
	GetCoinbaseTxSum(req *GetCoinbaseTxSumRequest) (*GetCoinbaseTxSumResponse, error)
	// GetFeeEstimate Gives an estimation on fees per byte.
	GetFeeEstimate(req *GetFeeEstimateRequest) (*GetFeeEstimateResponse, error)
	// GetAlternateChains Display alternative chains seen by the node.
	GetAlternateChains() (*GetAlternateChainsResponse, error)
	// RelayTx Relay a list of transaction IDs.
	RelayTx(req *RelayTxRequest) (*RelayTxResponse, error)
	// SyncInfo Get synchronisation informations
	SyncInfo() (*SyncInfoResponse, error)
	// GetTxpoolBacklog Get all transaction pool backlog
	GetTxpoolBacklog() (*GetTxpoolBacklogResponse, error)
	// GetOutputDistribution Alias: None.
	GetOutputDistribution(req *GetOutputDistributionRequest) (*GetOutputDistributionResponse, error)
}
```

## Contributing

PR's are welcome. Please read [CONTRIBUTING.md](https://github.com/MarinX/monerorpc/blob/master/CONTRIBUTING.md) for more info

## FAQ

### There is a method missing, can I use this library?

Yes, client has a `Do` method which accepts any struct. Example:

```go
package main

import (
	"fmt"

	"github.com/MarinX/monerorpc"
)

func main() {
	// create a new client for Testnet
	client := monerorpc.New(monerorpc.TestnetURI, nil)

	// if your monerod is protected, set username/password
	client.SetAuth("username", "password")

	// define request / response model
	var req struct{}
	var res struct{}

	// call RPC endpoint that is not documented
	err := client.Do("rpc_method_name", &req, &req)
	if err != nil {
		fmt.Println("error from RPC endpoint", err)
	}

}

```

### I found a bug/issue

Please submit an issue on github or if you know how to fix it, PR's are welcome.

### Version of the docs has changed, what now?

I will try to update the client as per docs changes. You are free to create an issue to notify me. I dont monitor monero docs 24/7 :)
