package wallet

// SetDaemonRequest represents the request model for SetDaemon
type SetDaemonRequest struct {
	// The URL of the daemon to connect to.
	Address string `json:"address,omitempty"`
	// If false, some RPC wallet methods will be disabled.
	Trusted bool `json:"trusted,omitempty"`
	// Accepts: disabled, enabled, autodetect) Specifies whether the Daemon uses SSL encryption
	SSLSupport string `json:"ssl_support,omitempty"`
	// The file path location of the SSL key.
	SSLPrivateKeyPath string `json:"ssl_private_key_path,omitempty"`
	// The file path location of the SSL certificate.
	SSLCertificatePath string `json:"ssl_certificate_path,omitempty"`
	// The file path location of the certificate authority file.
	SSLCaFile string `json:"ssl_ca_file,omitempty"`
	// The SHA1 fingerprints accepted by the SSL certificate.
	SSLAllowedFingerprints string `json:"ssl_allowed_fingerprints,omitempty"`
	// If false, the certificate must be signed by a trusted certificate authority.
	SSLAllowAnyCert string `json:"ssl_allow_any_cert,omitempty"`
}

// GetBalanceRequest represents the request model for GetBalance
type GetBalanceRequest struct {
	// Return balance for this account.
	AccountIndex uint32 `json:"account_index"`
	// Return balance detail for those subaddresses.
	AddressIndices uint32 `json:"address_indices,omitempty"`
}

// PerSubaddress model
type PerSubaddress struct {
	AccountIndex uint32 `json:"account_index"`
	// Index of the subaddress in the account.
	AddressIndex uint32 `json:"address_index"`
	// Address at this index. Base58 representation of the public keys.
	Address string `json:"address"`
	// Balance for the subaddress (locked or unlocked).
	Balance uint64 `json:"balance"`
	// Unlocked balance for the subaddress.
	UnlockedBalance uint64 `json:"unlocked_balance"`
	// Label for the subaddress.
	Label string `json:"label"`
	// Number of unspent outputs available for the subaddress.
	NumUnspentOutputs uint64 `json:"num_unspent_outputs"`
	BlocksToUnlock    uint64 `json:"blocks_to_unlock"`
	TimeToUnlock      uint64 `json:"time_to_unlock"`
}

// GetBalanceResponse represents the response model for GetBalance
type GetBalanceResponse struct {
	// The total balance of the current monero-wallet-rpc in session.
	Balance uint64 `json:"balance"`
	// Unlocked funds are those funds that are sufficiently deep enough in the Monero blockchain to be considered safe to spend.
	UnlockedBalance uint64 `json:"unlocked_balance"`
	// True if importing multisig data is needed for returning a correct balance.
	MultisigImportNeeded bool `json:"multisig_import_needed"`
	// Array of subaddress information; Balance information for each subaddress in an account.
	PerSubaddress  []PerSubaddress `json:"per_subaddress"`
	BlocksToUnlock uint64          `json:"blocks_to_unlock"`
	TimeToUnlock   uint64          `json:"time_to_unlock"`
}

// GetAddressRequest represents the request model for GetAddress
type GetAddressRequest struct {
	// Return addresses for this account.
	AccountIndex uint32 `json:"account_index"`
	// (Optional, defaults to all) List of address indices to return from the
	// account. Index 0 is of account 0 is the primary address, all others
	// are subadddresses.
	AddressIndices []uint32 `json:"address_index,omitempty"`
}

// Address model
type Address struct {
	// The base58 (sub)address string.
	Address string `json:"address"`
	// Label of the (sub)address
	Label string `json:"label"`
	// index of the (sub)address
	AddressIndex uint32 `json:"address_index"`
	// states if the (sub)address has already received funds
	Used bool `json:"used"`
}

// GetAddressResponse represents the response model for GetAddress
type GetAddressResponse struct {
	// The primary address of the requested account index.
	Address string `json:"address"`
	// Array of address information entries
	Addresses []Address `json:"addresses"`
}

// GetAddressIndexRequest represents the request model for GetAddressIndex
type GetAddressIndexRequest struct {
	// (sub)address to look for.
	Address string `json:"address"`
}

// Index model
type Index struct {
	// Account index.
	Major uint32 `json:"major"`
	// Subaddress index.
	Minor uint32 `json:"minor"`
}

// GetAddressIndexResponse represents the response model for GetAddressIndex
type GetAddressIndexResponse struct {
	// subaddress information
	Index Index `json:"index"`
}

// CreateAddressRequest represents the request model for CreateAddress
type CreateAddressRequest struct {
	// Create the new subaddress(es) in this account.
	AccountIndex uint32 `json:"account_index"`
	// (Optional) Label for the new address(es).
	Label string `json:"label,omitempty"`
	// (Optional) Number of addresses to create (range: 1 to 64, defaults to 1).
	Count uint32 `json:"count,omitempty"`
}

// CreateAddressResponse represents the response model for CreateAddress
type CreateAddressResponse struct {
	// 1st newly created address. Base58 representation of the public keys.
	Address string `json:"address"`
	// Index of the first new address created in the requested account.
	AddressIndex uint32 `json:"address_index"`
	// List of all address indices created.
	AddressIndices []uint32 `json:"address_indices"`
	// List of all addresses created
	Addresses []string `json:"addresses"`
}

// LabelAddressRequest represents the request model for LabelAddress
type LabelAddressRequest struct {
	// subaddress index; JSON Object containing the major & minor address
	Index Index  `json:"index"`
	Label string `json:"label"`
}

// ValidateAddressRequest represents the request model for ValidateAddress
type ValidateAddressRequest struct {
	// The address to validate.
	Address string `json:"address"`
	// If true, consider addresses belonging to any of the three Monero networks (mainnet, stagenet, and testnet) valid.
	// Otherwise, only consider an address valid if it belongs to the network on which the rpc-wallet's current daemon is running (Defaults to false).
	AnyNetType bool `json:"any_net_type,omitempty"`
	// If true, consider OpenAlias-formatted addresses valid (Defaults to false).
	AllowOpenalias bool `json:"allow_openalias,omitempty"`
}

// ValidateAddressResponse represents the response model for ValidateAddress
type ValidateAddressResponse struct {
	// True if the input address is a valid Monero address.
	Valid bool `json:"valid"`
	// True if the given address is an integrated address.
	Integrated bool `json:"integrated"`
	// True if the given address is a subaddress
	Subaddress bool `json:"subaddress"`
	// Specifies which of the three Monero networks (mainnet, stagenet, and testnet) the address belongs to.
	Nettype string `json:"nettype"`
	// True if the address is OpenAlias-formatted.
	OpenaliasAddress string `json:"openalias_address"`
}

// GetAccountsRequest represents the request model for GetAccounts
type GetAccountsRequest struct {
	// Tag for filtering accounts.
	Tag string `json:"tag,omitempty"`
}

// SubaddressAcount model
type SubaddressAcount struct {
	// Index of the account.
	AccountIndex uint32 `json:"account_index"`
	// Balance of the account (locked or unlocked).
	Balance uint64 `json:"balance"`
	// Base64 representation of the first subaddress in the account.
	BaseAddress string `json:"base_address"`
	// Label of the account.
	Label string `json:"label"`
	// Tag for filtering accounts.
	Tag string `json:"tag"`
	// Unlocked balance for the account.
	UnlockedBalance uint64 `json:"unlocked_balance"`
}

// GetAccountsResponse represents the response model for GetAccounts
type GetAccountsResponse struct {
	// array of subaddress account information
	SubaddressAccounts []SubaddressAcount `json:"subaddress_accounts"`
	// Total balance of the selected accounts (locked or unlocked).
	TotalBalance uint64 `json:"total_balance"`
	// Total unlocked balance of the selected accounts.
	TotalUnlockedBalance uint64 `json:"total_unlocked_balance"`
}

// CreateAccountRequest represents the request model for CreateAccount
type CreateAccountRequest struct {
	// Label for the account.
	Label string `json:"label,omitempty"`
}

// CreateAccountResponse represents the response model for CreateAccount
type CreateAccountResponse struct {
	// Index of the new account.
	AccountIndex uint32 `json:"account_index"`
	// Address for this account. Base58 representation of the public keys.
	Address string `json:"address"`
}

// LabelAccountRequest represents the request model for LabelAccount
type LabelAccountRequest struct {
	// Apply label to account at this index.
	AccountIndex uint32 `json:"account_index"`
	// Label for the account.
	Label string `json:"label"`
}

// AccountTag model
type AccountTag struct {
	// Filter tag.
	Tag string `json:"tag"`
	// Label for the tag.
	Label string `json:"label"`
	// List of tagged account indices.
	Accounts []int `json:"accounts"`
}

// GetAccountTagsResponse represents the response model for GetAccountTags
type GetAccountTagsResponse struct {
	// array of account tag information
	AccountTags []AccountTag `json:"account_tags"`
}

// TagAccountsRequest represents the request model for TagAccounts
type TagAccountsRequest struct {
	// Tag for the accounts.
	Tag string `json:"tag"`
	// Tag this list of accounts.
	Accounts []int `json:"accounts"`
}

// UntagAccountsRequest represents the request model for UntagAccounts
type UntagAccountsRequest struct {
	// Remove tag from this list of accounts.
	Accounts []int `json:"accounts"`
}

// SetAccountTagDescriptionRequest represents the request model for SetAccountTagDescription
type SetAccountTagDescriptionRequest struct {
	// Set a description for this tag.
	Tag string `json:"tag"`
	// Description for the tag.
	Description string `json:"description"`
}

// GetHeightResponse represents the response model for GetHeight
type GetHeightResponse struct {
	// The current monero-wallet-rpc's blockchain height.
	// If the wallet has been offline for a long time, it may need to catch up with the daemon.
	Height uint64 `json:"height"`
}

// Destination model
type Destination struct {
	// Amount to send to each destination, in atomic units.
	Amount uint64 `json:"amount"`
	// Destination public address.
	Address string `json:"address"`
}

// TransferRequest represents the request model for Transfer
type TransferRequest struct {
	// array of destinations to receive XMR
	Destinations []Destination `json:"destinations"`
	// Transfer from this account index. (Defaults to 0)
	AccountIndex uint32 `json:"account_index,omitempty"`
	// Transfer from this set of subaddresses. (Defaults to empty - all indices)
	SubaddrIndices []uint32 `json:"subaddr_indices,omitempty"`
	// Set a priority for the transaction. Accepted Values are: 0-3 for: default, unimportant, normal, elevated, priority.
	Priority uint64 `json:"priority,omitempty"`
	// Number of outputs from the blockchain to mix with (0 means no mixing).
	Mixin uint64 `json:"mixin"`
	// Number of outputs to mix in the transaction (this output + N decoys from the blockchain).
	RingSize uint64 `json:"ring_size"`
	// Number of blocks before the monero can be spent (0 to not add a lock).
	UnlockTime uint64 `json:"unlock_time"`
	// Return the transaction key after sending.
	GetTxKey bool `json:"get_tx_key,omitempty"`
	// If true, the newly created transaction will not be relayed to the monero network. (Defaults to false)
	DoNotRelay bool `json:"do_not_relay,omitempty"`
	// Return the transaction as hex string after sending (Defaults to false)
	GetTxHex bool `json:"get_tx_hex"`
	// Return the metadata needed to relay the transaction. (Defaults to false)
	GetTxMetadata bool `json:"get_tx_metadata"`
}

// TransferResponse represents the response model for Transfer
type TransferResponse struct {
	// Amount transferred for the transaction.
	Amount uint64 `json:"amount"`
	// Integer value of the fee charged for the txn.
	Fee uint64 `json:"fee"`
	// Set of multisig transactions in the process of being signed (empty for non-multisig).
	MultisigTxset string `json:"multisig_txset"`
	// Raw transaction represented as hex string, if get_tx_hex is true.
	TxBlob string `json:"tx_blob"`
	// String for the publically searchable transaction hash.
	TxHash string `json:"tx_hash"`
	// String for the transaction key if get_tx_key is true, otherwise, blank string.
	TxKey string `json:"tx_key"`
	// Set of transaction metadata needed to relay this transfer later, if get_tx_metadata is true.
	TxMetadata string `json:"tx_metadata"`
	// Set of unsigned tx for cold-signing purposes.
	UnsignedTxset string `json:"unsigned_txset"`
}

// TransferSplitRequest represents the request model for TransferSplit
type TransferSplitRequest struct {
	// array of destinations to receive XMR
	Destinations []Destination `json:"destinations"`
	// Transfer from this account index. (Defaults to 0)
	AccountIndex uint32 `json:"account_index,omitempty"`
	// Transfer from this set of subaddresses. (Defaults to empty - all indices)
	SubaddrIndices []uint32 `json:"subaddr_indices,omitempty"`
	// Number of outputs from the blockchain to mix with (0 means no mixing).
	Mixin uint64 `json:"mixin"`
	// Sets ringsize to n (mixin + 1).
	RingSize uint64 `json:"ring_size"`
	// Number of blocks before the monero can be spent (0 to not add a lock).
	UnlockTime uint64 `json:"unlock_time"`
	// Return the transaction keys after sending.
	GetTxKey bool `json:"get_tx_key,omitempty"`
	//  Set a priority for the transactions. Accepted Values are: 0-3 for: default, unimportant, normal, elevated, priority.
	Priority uint64 `json:"priority,omitempty"`
	// If true, the newly created transaction will not be relayed to the monero network. (Defaults to false)
	DoNotRelay bool `json:"do_not_relay,omitempty"`
	// Return the transactions as hex string after sending
	GetTxHex bool `json:"get_tx_hex"`
	// True to use the new transaction construction algorithm, defaults to false.
	NewAlgorithm bool `json:"new_algorithm"`
	// Return list of transaction metadata needed to relay the transfer later.
	GetTxMetadata bool `json:"get_tx_metadata"`
}

// TransferSplitResponse represents the response model for TransferSplit
type TransferSplitResponse struct {
	// The tx hashes of every transaction.
	TxHashList []string `json:"tx_hash_list"`
	// The transaction keys for every transaction.
	TxKeyList []string `json:"tx_key_list"`
	// The amount transferred for every transaction.
	AmountList []uint64 `json:"amount_list"`
	// The amount of fees paid for every transaction.
	FeeList []uint64 `json:"fee_list"`
	// The tx as hex string for every transaction.
	TxBlobList []string `json:"tx_blob_list"`
	// List of transaction metadata needed to relay the transactions later.
	TxMetadataList []string `json:"tx_metadata_list"`
	// The set of signing keys used in a multisig transaction (empty for non-multisig).
	MultisigTxset string `json:"multisig_txset"`
	// Set of unsigned tx for cold-signing purposes.
	UnsignedTxset string `json:"unsigned_txset"`
}

// SignTransferRequest represents the request model for SignTransfer
type SignTransferRequest struct {
	// Set of unsigned tx returned by "transfer" or "transfer_split" methods.
	UnsignedTxset string `json:"unsigned_txset"`
	// If true, return the raw transaction data. (Defaults to false)
	ExportRaw bool `json:"export_raw,omitempty"`
}

// SignTransferResponse represents the response model for SignTransfer
type SignTransferResponse struct {
	// Set of signed tx to be used for submitting transfer.
	SignedTxset string `json:"signed_txset"`
	// The tx hashes of every transaction.
	TxHashList []string `json:"tx_hash_list"`
	// The tx raw data of every transaction.
	TxRawList []string `json:"tx_raw_list"`
}

// SubmitTransferRequest represents the request model for SubmitTransfer
type SubmitTransferRequest struct {
	// Set of signed tx returned by "sign_transfer"
	TxDataHex string `json:"tx_data_hex"`
}

// SubmitTransferResponse represents the response model for SubmitTransfer
type SubmitTransferResponse struct {
	// The tx hashes of every transaction.
	TxHashList []string `json:"tx_hash_list"`
}

// SweepDustRequest represents the request model for SweepDust
type SweepDustRequest struct {
	// Return the transaction keys after sending.
	GetTxKeys bool `json:"get_tx_keys,omitempty"`
	// If true, the newly created transaction will not be relayed to the monero network. (Defaults to false)
	DoNotRelay bool `json:"do_not_relay,omitempty"`
	// Return the transactions as hex string after sending. (Defaults to false)
	GetTxHex bool `json:"get_tx_hex,omitempty"`
	// Return list of transaction metadata needed to relay the transfer later. (Defaults to false)
	GetTxMetadata bool `json:"get_tx_metadata,omitempty"`
}

// SweepDustResponse represents the response model for SweepDust
type SweepDustResponse struct {
	// The tx hashes of every transaction.
	TxHashList []string `json:"tx_hash_list"`
	// The transaction keys for every transaction.
	TxKeyList []string `json:"tx_key_list"`
	// The amount transferred for every transaction.
	AmountList []uint64 `json:"amount_list"`
	// The amount of fees paid for every transaction.
	FeeList []uint64 `json:"fee_list"`
	// The tx as hex string for every transaction.
	TxBlobList []string `json:"tx_blob_list"`
	// List of transaction metadata needed to relay the transactions later.
	TxMetadataList []string `json:"tx_metadata_list"`
	// The set of signing keys used in a multisig transaction (empty for non-multisig).
	MultisigTxset string `json:"multisig_txset"`
	// Set of unsigned tx for cold-signing purposes.
	UnsignedTxset string `json:"unsigned_txset"`
}

// SweepAllRequest represents the request model for SweepAll
type SweepAllRequest struct {
	// Destination public address.
	Address string `json:"address"`
	// Sweep transactions from this account.
	AccountIndex uint32 `json:"account_index"`
	// Sweep from this set of subaddresses in the account.
	SubaddrIndices []uint32 `json:"subaddr_indices,omitempty"`
	// Priority for sending the sweep transfer, partially determines fee.
	Priority uint64 `json:"priority,omitempty"`
	// Number of outputs from the blockchain to mix with (0 means no mixing).
	Mixin uint64 `json:"mixin"`
	// Sets ringsize to n (mixin + 1).
	RingSize uint64 `json:"ring_size"`
	// Number of blocks before the monero can be spent (0 to not add a lock).
	UnlockTime uint64 `json:"unlock_time"`
	// Return the transaction keys after sending.
	GetTxKeys bool `json:"get_tx_keys,omitempty"`
	// Include outputs below this amount.
	BelowAmount uint64 `json:"below_amount,omitempty"`
	// If true, do not relay this sweep transfer. (Defaults to false)
	DoNotRelay bool `json:"do_not_relay,omitempty"`
	// return the transactions as hex encoded string. (Defaults to false)
	GetTxHex bool `json:"get_tx_hex,omitempty"`
	// return the transaction metadata as a string. (Defaults to false)
	GetTxMetadata bool `json:"get_tx_metadata,omitempty"`
}

// SweepAllResponse represents the response model for SweepAll
type SweepAllResponse struct {
	TxHashList []string `json:"tx_hash_list"`
	// The transaction keys for every transaction.
	TxKeyList []string `json:"tx_key_list"`
	// The amount transferred for every transaction.
	AmountList []uint64 `json:"amount_list"`
	// The amount of fees paid for every transaction.
	FeeList []uint64 `json:"fee_list"`
	// The tx as hex string for every transaction.
	TxBlobList []string `json:"tx_blob_list"`
	// List of transaction metadata needed to relay the transactions later.
	TxMetadataList []string `json:"tx_metadata_list"`
	// The set of signing keys used in a multisig transaction (empty for non-multisig).
	MultisigTxset string `json:"multisig_txset"`
	// Set of unsigned tx for cold-signing purposes.
	UnsignedTxset string `json:"unsigned_txset"`
}

// SweepSingleRequest represents the request model for SweepSingle
type SweepSingleRequest struct {
	// Destination public address.
	Address string `json:"address"`
	// Sweep transactions from this account.
	AccountIndex uint32 `json:"account_index"`
	// Sweep from this set of subaddresses in the account.
	SubaddrIndices []uint32 `json:"subaddr_indices,omitempty"`
	// Priority for sending the sweep transfer, partially determines fee.
	Priority uint64 `json:"priority,omitempty"`
	// Number of outputs from the blockchain to mix with (0 means no mixing).
	Mixin uint64 `json:"mixin"`
	// Sets ringsize to n (mixin + 1).
	RingSize uint64 `json:"ring_size"`
	// Number of blocks before the monero can be spent (0 to not add a lock).
	UnlockTime uint64 `json:"unlock_time"`
	// Return the transaction keys after sending.
	GetTxKeys bool `json:"get_tx_keys,omitempty"`
	// Key image of specific output to sweep.
	KeyImage string `json:"key_image"`
	// Include outputs below this amount.
	BelowAmount uint64 `json:"below_amount,omitempty"`
	// If true, do not relay this sweep transfer. (Defaults to false)
	DoNotRelay bool `json:"do_not_relay,omitempty"`
	// return the transactions as hex encoded string. (Defaults to false)
	GetTxHex bool `json:"get_tx_hex,omitempty"`
	// return the transaction metadata as a string. (Defaults to false)
	GetTxMetadata bool `json:"get_tx_metadata,omitempty"`
}

// SweepSingleResponse represents the response model for SweepSingle
type SweepSingleResponse struct {
	TxHashList []string `json:"tx_hash_list"`
	// The transaction keys for every transaction.
	TxKeyList []string `json:"tx_key_list"`
	// The amount transferred for every transaction.
	AmountList []uint64 `json:"amount_list"`
	// The amount of fees paid for every transaction.
	FeeList []uint64 `json:"fee_list"`
	// The tx as hex string for every transaction.
	TxBlobList []string `json:"tx_blob_list"`
	// List of transaction metadata needed to relay the transactions later.
	TxMetadataList []string `json:"tx_metadata_list"`
	// The set of signing keys used in a multisig transaction (empty for non-multisig).
	MultisigTxset string `json:"multisig_txset"`
	// Set of unsigned tx for cold-signing purposes.
	UnsignedTxset string `json:"unsigned_txset"`
}

// RelayTxRequest represents the request model for RelayTx
type RelayTxRequest struct {
	// transaction metadata returned from a transfer method with get_tx_metadata set to true.
	Hex string `json:"hex"`
}

// RelayTxResponse represents the response model for RelayTx
type RelayTxResponse struct {
	// String for the publically searchable transaction hash.
	TxHash string `json:"tx_hash"`
}

// GetPaymentsRequest represents the request model for GetPayments
type GetPaymentsRequest struct {
	// Payment ID used to find the payments (16 characters hex).
	PaymentID string `json:"payment_id"`
}

// Payment model
type Payment struct {
	// Payment ID matching the input parameter.
	PaymentID string `json:"payment_id"`
	// ransaction hash used as the transaction ID.
	TxHash string `json:"tx_hash"`
	// Amount for this payment.
	Amount uint64 `json:"amount"`
	// Height of the block that first confirmed this payment.
	BlockHeight uint64 `json:"block_height"`
	// Time (in block height) until this payment is safe to spend.
	UnlockTime uint64 `json:"unlock_time"`
	// subaddress index
	SubaddrIndex Index `json:"subaddr_index"`
	// Address receiving the payment; Base58 representation of the public keys.
	Address string `json:"address"`
}

// GetPaymentsResponse represents the response model for GetPayments
type GetPaymentsResponse struct {
	// list of
	Payments []Payment `json:"payments"`
}

// GetBulkPaymentsRequest represents the request model for GetBulkPayments
type GetBulkPaymentsRequest struct {
	// Payment IDs used to find the payments (16 characters hex).
	PaymentID string `json:"payment_id"`
	// The block height at which to start looking for payments.
	MinBlockHeight uint64 `json:"min_block_height"`
}

// GetBulkPaymentsResponse represents the response model for GetBulkPayments
type GetBulkPaymentsResponse struct {
	// list of
	Payments []Payment `json:"payments"`
}

// IncomingTransfersRequest represents the request model for IncomingTransfers
type IncomingTransfersRequest struct {
	// "all": all the transfers, "available": only transfers which are not yet spent, OR "unavailable": only transfers which are already spent.
	TransferType string `json:"transfer_type"`
	// Return transfers for this account. (defaults to 0)
	AccountIndex uint32 `json:"account_index,omitempty"`
	// Return transfers sent to these subaddresses.
	SubaddrIndices *Index `json:"subaddr_indices,omitempty"`
}

// IncomingTransfer model
type IncomingTransfer struct {
	// Amount of this transfer.
	Amount uint64 `json:"amount"`
	// BlockHeight is the block height at which this transfer was confirmed.
	BlockHeight uint64 `json:"block_height"`
	// Frozen indicates whether the output key was frozen by freeze
	Frozen bool `json:"frozen"`
	// Mostly internal use, can be ignored by most users.
	GlobalIndex uint64 `json:"global_index"`
	// Key image for the incoming transfer's unspent output.
	KeyImage string `json:"key_image"`
	// Spent indicates if the output has been spent.
	Spent bool `json:"spent"`
	// Subaddress index for incoming transfer.
	SubaddrIndex Index `json:"subaddr_index"`
	// Several incoming transfers may share the same hash if they were in the same transaction.
	TxHash string `json:"tx_hash"`
	// Unlocked indicates if the output is spendable.
	Unlocked bool `json:"unlocked"`
}

// IncomingTransfersResponse represents the response model for IncomingTransfers
type IncomingTransfersResponse struct {
	// list of
	Transfers []IncomingTransfer `json:"transfers"`
}

// QueryKeyRequest represents the request model for QueryKey
type QueryKeyRequest struct {
	// Which key to retrieve: "mnemonic" - the mnemonic seed (older wallets do not have one) OR "view_key" - the view key
	KeyType string `json:"key_type"`
}

// QueryKeyResponse represents the response model for QueryKey
type QueryKeyResponse struct {
	// The view key will be hex encoded, while the mnemonic will be a string of words.
	Key string `json:"key"`
}

// MakeIntegratedAddressRequest represents the request model for MakeIntegratedAddress
type MakeIntegratedAddressRequest struct {
	// (Optional, defaults to primary address) Destination public address.
	StandardAddress string `json:"standard_address,omitempty"`
	// (Optional, defaults to a random ID) 16 characters hex encoded.
	PaymentID string `json:"payment_id"`
}

// MakeIntegratedAddressResponse represents the response model for MakeIntegratedAddress
type MakeIntegratedAddressResponse struct {
	IntegratedAddress string `json:"integrated_address"`
	// hex encoded
	PaymentID string `json:"payment_id"`
}

// SplitIntegratedAddressRequest represents the request model for SplitIntegratedAddress
type SplitIntegratedAddressRequest struct {
	IntegratedAddress string `json:"integrated_address"`
}

// SplitIntegratedAddressResponse represents the response model for SplitIntegratedAddress
type SplitIntegratedAddressResponse struct {
	// States if the address is a subaddress
	IsSubaddress bool `json:"is_subaddress"`
	// hex encoded
	Payment         string `json:"payment"`
	StandardAddress string `json:"standard_address"`
}

// SetTxNotesRequest represents the request model for SetTxNotes
type SetTxNotesRequest struct {
	// transaction ids
	Txids []string `json:"txids"`
	// notes for the transactions
	Notes []string `json:"notes"`
}

// GetTxNotesRequest represents the request model for GetTxNotes
type GetTxNotesRequest struct {
	// transaction ids
	Txids []string `json:"txids"`
}

// GetTxNotesResponse represents the response model for GetTxNotes
type GetTxNotesResponse struct {
	// notes for the transactions
	Notes []string `json:"notes"`
}

// SetAttributeRequest represents the request model for SetAttribute
type SetAttributeRequest struct {
	// attribute name
	Key string `json:"key"`
	// attribute value
	Value string `json:"value"`
}

// GetAttributeRequest represents the request model for GetAttribute
type GetAttributeRequest struct {
	// attribute name
	Key string `json:"key"`
}

// GetAttributeResponse represents the response model for GetAttribute
type GetAttributeResponse struct {
	// attribute value
	Value string `json:"value"`
}

// GetTxKeyRequest represents the request model for GetTxKey
type GetTxKeyRequest struct {
	// transaction id
	Txid string `json:"txid"`
}

// GetTxKeyResponse represents the response model for GetTxKey
type GetTxKeyResponse struct {
	// transaction secret key.
	TxKey string `json:"tx_key"`
}

// CheckTxKeyRequest represents the request model for CheckTxKey
type CheckTxKeyRequest struct {
	// transaction id
	Txid string `json:"txid"`
	// transaction secret key.
	TxKey string `json:"tx_key"`
	// destination public address of the transaction.
	Address string `json:"address"`
}

// CheckTxKeyResponse represents the response model for CheckTxKey
type CheckTxKeyResponse struct {
	// Number of block mined after the one with the transaction.
	Confirmations uint64 `json:"confirmations"`
	// States if the transaction is still in pool or has been added to a block.
	InPool bool `json:"in_pool"`
	// Amount of the transaction.
	Received uint64 `json:"received"`
}

// GetTxProofRequest represents the request model for GetTxProof
type GetTxProofRequest struct {
	// transaction id.
	Txid string `json:"txid"`
	// destination public address of the transaction.
	Address string `json:"address"`
	// add a message to the signature to further authenticate the prooving process.
	Message string `json:"messag,omitempty"`
}

// GetTxProofResponse represents the response model for GetTxProof
type GetTxProofResponse struct {
	// transaction signature.
	Signature string `json:"signature"`
}

// CheckTxProofRequest represents the request model for CheckTxProof
type CheckTxProofRequest struct {
	// Transaction id.
	TxID string `json:"txid"`
	// Destination public address of the transaction.
	Address string `json:"address"`
	// (Optional) Should be the same message used in get_tx_proof.
	Message string `json:"message,omitempty"`
	// Transaction signature to confirm.
	Signature string `json:"signature"`
}

// CheckTxProofResponse represents the response model for CheckTxProof
type CheckTxProofResponse struct {
	Confirmations uint64 `json:"confirmations"`
	// States if the inputs proves the transaction.
	Good bool `json:"good"`
	// States if the transaction is still in pool or has been added to a block.
	InPool bool `json:"in_pool"`
	// Amount of the transaction.
	Received uint64 `json:"received"`
}

// GetSpendProofRequest represents the request model for GetSpendProof
type GetSpendProofRequest struct {
	// Transaction id.
	TxID string `json:"txid"`
	// (Optional) add a message to the signature to further authenticate the prooving process.
	Message string `json:"message,omitempty"`
}

// GetSpendProofResponse represents the response model for GetSpendProof
type GetSpendProofResponse struct {
	// Spend signature.
	Signature string `json:"signature"`
}

// CheckSpendProofRequest represents the request model for CheckSpendProof
type CheckSpendProofRequest struct {
	// Transaction id.
	TxID string `json:"txid"`
	// (Optional) Should be the same message used in get_spend_proof.
	Message string `json:"message,omitempty"`
	// Spend signature to confirm.
	Signature string `json:"signature"`
}

// CheckSpendProofResponse represents the response model for CheckSpendProof
type CheckSpendProofResponse struct {
	// States if the inputs proves the spend.
	Good bool `json:"good"`
}

// GetReserveProofRequest represents the request model for GetReserveProof
type GetReserveProofRequest struct {
	// Proves all wallet balance to be disposable.
	All bool `json:"all"`
	// Specify the account from witch to prove reserve. (ignored if all is set to true)
	AccountIndex uint32 `json:"account_index"`
	// Amount (in atomic units) to prove the account has for reserve. (ignored if all is set to true)
	Amount uint64 `json:"amount"`
	// (Optional) add a message to the signature to further authenticate the prooving process.
	Message string `json:"message,omitempty"`
}

// GetReserveProofResponse represents the response model for GetReserveProof
type GetReserveProofResponse struct {
	// Reserve signature.
	Signature string `json:"signature"`
}

// CheckReserveProofRequest represents the request model for CheckReserveProof
type CheckReserveProofRequest struct {
	// Public address of the wallet.
	Address string `json:"address"`
	// (Optional) Should be the same message used in get_reserve_proof.
	Message string `json:"message,omitempty"`
	// Reserve signature to confirm.
	Signature string `json:"signature"`
}

// CheckReserveProofResponse represents the response model for CheckReserveProof
type CheckReserveProofResponse struct {
	// States if the inputs proves the reserve.
	Good bool `json:"good"`
}

// GetTransfersRequest represents the request model for GetTransfers
type GetTransfersRequest struct {
	// (Optional) Include incoming transfers.
	In bool `json:"in,omitempty"`
	// (Optional) Include outgoing transfers.
	Out bool `json:"out,omitempty"`
	// (Optional) Include pending transfers.
	Pending bool `json:"pending,omitempty"`
	// (Optional) Include failed transfers.
	Failed bool `json:"failed,omitempty"`
	// (Optional) Include transfers from the daemon's transaction pool.
	Pool bool `json:"pool,omitempty"`
	// (Optional) Filter transfers by block height.
	FilterByHeight bool `json:"filter_by_height,omitempty"`
	// (Optional) Minimum block height to scan for transfers, if filtering by height is enabled.
	MinHeight uint64 `json:"min_height,omitempty"`
	// (Opional) Maximum block height to scan for transfers, if filtering by height is enabled (defaults to max block height).
	MaxHeight uint64 `json:"max_height,omitempty"`
	// (Optional) Index of the account to query for transfers. (defaults to 0)
	AccountIndex uint32 `json:"account_index,omitempty"`
	// (Optional) List of subaddress indices to query for transfers. (Defaults to empty - all indices)
	SubaddrIndices []uint32 `json:"subaddr_indices,omitempty"`
}

// Transfer model
type Transfer struct {
	// Public address of the transfer.
	Address string `json:"address"`
	// Amount transferred.
	Amount uint64 `json:"amount"`
	// Individual amounts if multiple where received
	Amounts []uint64 `json:"amounts"`
	// Number of block mined since the block containing this transaction (or block height at which the transaction should be added to a block if not yet confirmed).
	Confirmations uint64 `json:"confirmations"`
	// JSON objects containing transfer destinations:
	Destinations []Destination `json:"destinations"`
	// True if the key image(s) for the transfer have been seen before.
	DoubleSpendSeen bool `json:"double_spend_seen"`
	// Transaction fee for this transfer.
	Fee uint64 `json:"fee"`
	// Height of the first block that confirmed this transfer (0 if not mined yet).
	Height uint64 `json:"height"`
	Locked bool   `json:"locked"`
	// Note about this transfer.
	Note string `json:"note"`
	// Payment ID for this transfer.
	PaymentID string `json:"payment_id"`
	// JSON object containing the major & minor subaddress index:
	SubaddrIndex Index `json:"subaddr_index"`
	// Estimation of the confirmations needed for the transaction to be included in a block.
	SuggestedConfirmationsThreshold uint64 `json:"suggested_confirmations_threshold"`
	// POSIX timestamp for when this transfer was first confirmed in a block (or timestamp submission if not mined yet).
	Timestamp uint64 `json:"timestamp"`
	// Transaction ID for this transfer.
	TxID string `json:"txid"`
	// Transfer type: "in/out/pending/failed/pool"
	Type string `json:"type"`
	// Number of blocks until transfer is safely spendable.
	UnlockTime uint64 `json:"unlock_time"`
}

// GetTransfersResponse represents the response model for GetTransfers
type GetTransfersResponse struct {
	// Array of transfers:
	In      []*Transfer `json:"in"`
	Out     []*Transfer `json:"out"`
	Pending []*Transfer `json:"pending"`
	Failed  []*Transfer `json:"failed"`
	Pool    []*Transfer `json:"pool"`
}

// GetTransferByTxidRequest represents the request model for GetTransferByTxid
type GetTransferByTxidRequest struct {
	// Transaction ID used to find the transfer.
	TxID string `json:"txid"`
	// (Optional) Index of the account to query for the transfer.
	AccountIndex uint32 `json:"account_index,omitempty"`
}

// GetTransferByTxidResponse represents the response model for GetTransferByTxid
type GetTransferByTxidResponse struct {
	// JSON object containing payment information:
	Transfer Transfer `json:"transfer"`
	// If the list length is > 1 then multiple outputs where received in this transaction, each of which has its own transfer JSON object
	Transfers []Transfer `json:"transfers"`
}

// DescribeTransferRequest represents the request model for DescribeTransfer
type DescribeTransferRequest struct {
	// A hexadecimal string representing a set of unsigned transactions (empty for multisig transactions; non-multisig signed transactions are not supported).
	UnsignedTxset string `json:"unsigned_txset,omitempty"`
	// A hexadecimal string representing the set of signing keys used in a multisig transaction (empty for unsigned transactions; non-multisig signed transactions are not supported).
	MultisigTxset string `json:"multisig_txset"`
}

// Desc model
type Desc struct {
	// The sum of the inputs spent by the transaction in atomic units.
	AmountIn uint64 `json:"amount_in"`
	// The sum of the outputs created by the transaction in atomic units.
	AmountOut uint64 `json:"amount_out"`
	// list of:
	Recipients []Recipient `json:"recipients"`

	// The address of the change recipient.
	ChangeAddress string `json:"change_address"`
	// The amount sent to the change address in atomic units.
	ChangeAmount uint64 `json:"change_amount"`
	// The fee charged for the transaction in atomic units.
	Fee uint64 `json:"fee"`
	// payment ID for this transfer (empty if not provided.)
	PaymentID string `json:"payment_id"`
	// The number of inputs in the ring (1 real output + the number of decoys from the blockchain).
	RingSize uint64 `json:"ring_size"`
	// The number of blocks before the monero can be spent (0 for no lock).
	UnlockTime uint64 `json:"unlock_time"`
	// The number of fake outputs added to single-output transactions. Fake outputs have 0 amount and are sent to a random address.
	DummyOutputs uint64 `json:"dummy_outputs"`
	// Arbitrary transaction data in hexadecimal format.
	Extra string `json:"extra"`
}

// Recipient model
type Recipient struct {
	// The public address of the recipient.
	Address string `json:"address"`
	// The amount sent to the recipient in atomic units.
	Amount uint64 `json:"amount"`
}

// DescribeTransferResponse represents the response model for DescribeTransfer
type DescribeTransferResponse struct {
	Desc []Desc `json:"desc"`
}

// SignRequest represents the request model for Sign
type SignRequest struct {
	// Anything you need to sign.
	Data string `json:"data"`
}

// SignResponse represents the response model for Sign
type SignResponse struct {
	// Signature generated against the "data" and the account public address.
	Signature string `json:"signature"`
}

// VerifyRequest represents the request model for Verify
type VerifyRequest struct {
	// What should have been signed.
	Data string `json:"data"`
	// Public address of the wallet used to sign the data.
	Address string `json:"address"`
	// Signature generated by sign method.
	Signature string `json:"signature"`
}

// VerifyResponse represents the response model for Verify
type VerifyResponse struct {
	// True if signature is valid.
	Good bool `json:"good"`
}

// ExportOutputsRequest represents the request model for ExportOutputs
type ExportOutputsRequest struct {
	// If true, export all outputs. Otherwise, export outputs since the last export. (default = false)
	All bool `json:"all"`
}

// ExportOutputsResponse represents the response model for ExportOutputs
type ExportOutputsResponse struct {
	// Wallet outputs in hex format.
	OutputsDataHex string `json:"outputs_data_hex"`
}

// ImportOutputsRequest represents the request model for ImportOutputs
type ImportOutputsRequest struct {
	// Wallet outputs in hex format.
	OutputsDataHex string `json:"outputs_data_hex"`
}

// ImportOutputsResponse represents the response model for ImportOutputs
type ImportOutputsResponse struct {
	// Number of outputs imported.
	NumImported uint64 `json:"num_imported"`
}

// SignedImage model
type SignedImage struct {
	KeyImage  string `json:"key_image"`
	Signature string `json:"signature"`
}

// ExportKeyImagesRequest represent the request model for ExportKeyImages
type ExportKeyImagesRequest struct {
	// If true, export all key images. Otherwise, export key images since the last export. (default = false)
	All bool `json:"all"`
}

// ExportKeyImagesResponse represent the response model for ExportKeyImages
type ExportKeyImagesResponse struct {
	// Array of signed key images:
	SignedKeyImages []SignedImage `json:"signed_key_images"`
}

// ImportKeyImagesRequest represents the request model for ImportKeyImages
type ImportKeyImagesRequest struct {
	// Array of signed key images:
	SignedKeyImages []SignedImage `json:"signed_key_images"`
}

// ImportKeyImagesResponse represents the response model for ImportKeyImages
type ImportKeyImagesResponse struct {
	Height uint64 `json:"height"`
	// Amount (in atomic units) spent from those key images.
	Spent uint64 `json:"spent"`
	// Amount (in atomic units) still available from those key images.
	Unspent uint64 `json:"unspent"`
}

// MakeURIRequest represents the request model for MakeUri
type MakeURIRequest struct {
	// Wallet address
	Address string `json:"address"`
	// (Optional) the integer amount to receive, in atomic units
	Amount uint64 `json:"amount,omitempty"`
	// (Optional) 16 or 64 character hexadecimal payment id
	PaymentID string `json:"payment_id,omitempty"`
	// (Optional) name of the payment recipient
	RecipientName string `json:"recipient_name,omitempty"`
	// (Optional) Description of the reason for the tx
	TxDescription string `json:"tx_description,omitempty"`
}

// MakeURIResponse represents the response model for MakeUri
type MakeURIResponse struct {
	// This contains all the payment input information as a properly formatted payment URI
	URI string `json:"uri"`
}

// ParseURIRequest represents the request model for ParseUri
type ParseURIRequest struct {
	// This contains all the payment input information as a properly formatted payment URI
	URI string `json:"uri"`
}

// URI model
type URI struct {
	// Wallet address
	Address string `json:"address"`
	// Integer amount to receive, in atomic units (0 if not provided)
	Amount uint64 `json:"amount"`
	// 16 or 64 character hexadecimal payment id (empty if not provided)
	PaymentID string `json:"payment_id"`
	// Name of the payment recipient (empty if not provided)
	RecipientName string `json:"recipient_name"`
	// Description of the reason for the tx (empty if not provided)
	TxDescription string `json:"tx_description"`
}

// ParseURIResponse represents the response model for ParseUri
type ParseURIResponse struct {
	// JSON object containing payment information:
	URI URI `json:"uri"`
}

// GetAddressBookRequest represents the request model for GetAddressBook
type GetAddressBookRequest struct {
	// Indices of the requested address book entries
	Entries []uint64 `json:"entries"`
}

// Entry model
type Entry struct {
	// Public address of the entry
	Address string `json:"address"`
	// Description of this address entry
	Description string `json:"description"`
	Index       uint64 `json:"index"`
	PaymentID   string `json:"payment_id"`
}

// GetAddressBookResponse represents the response model for GetAddressBook
type GetAddressBookResponse struct {
	// Array of entries:
	Entries []Entry `json:"entries"`
}

// AddAddressBookRequest represents the request model for AddAddressBook
type AddAddressBookRequest struct {
	Address string `json:"address"`
	// (Optional) string, defaults to "0000000000000000000000000000000000000000000000000000000000000000";
	PaymentID string `json:"payment_id,omitempty"`
	// (Optional) string, defaults to "";
	Description string `json:"description,omitempty"`
}

// AddAddressBookResponse represents the response model for AddAddressBook
type AddAddressBookResponse struct {
	// The index of the address book entry.
	Index uint64 `json:"index"`
}

// EditAddressBookRequest represents the request model for EditAddressBook
type EditAddressBookRequest struct {
	// Index of the address book entry to edit.
	Index uint64 `json:"index"`
	// If true, set the address for this entry to the value of "address".
	SetAddress bool `json:"set_address"`
	// (Optional) The 95-character public address to set.
	Address string `json:"address,omitempty"`
	// If true, set the description for this entry to the value of "description".
	SetDescription bool `json:"set_description"`
	// (Optional) Human-readable description for this entry.
	Description string `json:"description,omitempty"`
	// If true, set the payment ID for this entry to the value of "payment_id".
	SetPaymentID bool `json:"set_payment_id"`
	// (Optional) Payment ID for this address.
	PaymentID string `json:"payment_id,omitempty"`
}

// DeleteAddressBookRequest represents the request model for DeleteAddressBook
type DeleteAddressBookRequest struct {
	// The index of the address book entry.
	Index uint64 `json:"index"`
}

// RefreshRequest represents the request model for Refresh
type RefreshRequest struct {
	// (Optional) The block height from which to start refreshing.
	StartHeight *uint64 `json:"start_height,omitempty"`
}

// RefreshResponse represents the response model for Refresh
type RefreshResponse struct {
	// Number of new blocks scanned.
	BlocksFetched uint64 `json:"blocks_fetched"`
	// States if transactions to the wallet have been found in the blocks.
	ReceivedMoney bool `json:"received_money"`
}

// AutoRefreshRequest represents the request model for AutoRefresh
type AutoRefreshRequest struct {
	// Enable or disable automatic refreshing (default = true).
	Enable *bool `json:"enable,omitempty"`
	// The period of the wallet refresh cycle (i.e. time between refreshes) in seconds.
	Period uint64 `json:"period,omitempty"`
}

// StartMiningRequest represents the response model for StartMining
type StartMiningRequest struct {
	// Number of threads created for mining.
	ThreadsCount uint64 `json:"threads_count"`
	// Allow to start the miner in smart mining mode.
	DoBackgroundMining bool `json:"do_background_mining"`
	// Ignore battery status (for smart mining only)
	IgnoreBattery bool `json:"ignore_battery"`
}

// GetLanguagesResponse represents the response model for GetLanguages
type GetLanguagesResponse struct {
	// List of available languages
	Languages []string `json:"languages"`
}

// CreateWalletRequest represents the request model for CreateWallet
type CreateWalletRequest struct {
	// Wallet file name.
	Filename string `json:"filename"`
	// (Optional) password to protect the wallet.
	Password string `json:"password,omitempty"`
	// Language for your wallets' seed.
	Language string `json:"language"`
}

// GenerateFromKeysRequest represents the request model for GenerateFromKeys
type GenerateFromKeysRequest struct {
	// (Optional; defaults to 0) The block height to restore the wallet from.
	RestoreHeight uint64 `json:"restore_height,omitempty"`
	// The wallet's file name on the RPC server.
	Filename string `json:"filename"`
	// The wallet's primary address.
	Address string `json:"address"`
	// (Optional; omit to create a view-only wallet) The wallet's private spend key.
	Spendkey string `json:"spendkey,omitempty"`
	// The wallet's private view key.
	Viewkey string `json:"viewkey"`
	// The wallet's password.
	Password string `json:"password"`
	//(Defaults to true) If true, save the current wallet before generating the new wallet.
	AutosaveCurrent *bool `json:"autosave_current,omitempty"`
}

// GenerateFromKeysResponse represents the response model for GenerateFromKeys
type GenerateFromKeysResponse struct {
	// The wallet's address.
	Address string `json:"address"`
	// Verification message indicating that the wallet was generated successfully and whether or not it is a view-only wallet.
	Info string `json:"info"`
}

// OpenWalletRequest represents the request model for OpenWallet
type OpenWalletRequest struct {
	// Wallet name stored in –wallet-dir.
	Filename string `json:"filename"`
	// (Optional) only needed if the wallet has a password defined.
	Password string `json:"password,omitempty"`
}

// RestoreDeterministicWalletRequest represents the request model for RestoreDeterministicWallet
type RestoreDeterministicWalletRequest struct {
	// Name of the wallet.
	Name string `json:"filename"`
	// Password of the wallet.
	Password string `json:"password"`
	// Mnemonic phrase of the wallet to restore.
	Seed string `json:"seed"`
	// Block height to restore the wallet from (default = 0).
	RestoreHeight uint64 `json:"restore_height,omitempty"`
	// Language of the mnemonic phrase in case the old language is invalid.
	Language string `json:"language,omitempty"`
	// Offset used to derive a new seed from the given mnemonic to recover a secret wallet from the mnemonic phrase.
	SeedOffset string `json:"seed_offset,omitempty"`
	//Whether to save the currently open RPC wallet before closing it (Defaults to true).
	AutosaveCurrent *bool `json:"autosave_current,omitempty"`
}

// RestoreDeterministicWalletResponse represents the response model for RestoreDeterministicWallet
type RestoreDeterministicWalletResponse struct {
	// 95-character hexadecimal address of the restored wallet as a string.
	Address string `json:"address"`
	// Message describing the success or failure of the attempt to restore the wallet.
	Info string `json:"info"`
	// Mnemonic phrase of the restored wallet, which is updated if the wallet was restored from a deprecated-style mnemonic phrase.
	Seed string `json:"seed"`
	// Indicates if the restored wallet was created from a deprecated-style mnemonic phrase.
	WasDeprecated bool `json:"was_deprecated"`
}

// ChangeWalletPasswordRequest represents the request model for ChangeWalletPassword
type ChangeWalletPasswordRequest struct {
	// (Optional) Current wallet password, if defined.
	OldPassword string `json:"old_password,omitempty"`
	// (Optional) New wallet password, if not blank.
	NewPassword string `json:"new_password,omitempty"`
}

// IsMultisigResponse represents the response model for IsMultisig
type IsMultisigResponse struct {
	// States if the wallet is multisig
	Multisig bool `json:"multisig"`
	Ready    bool `json:"ready"`
	// Amount of signature needed to sign a transfer.
	Threshold uint64 `json:"threshold"`
	// Total amount of signature in the multisig wallet.
	Total uint64 `json:"total"`
}

// PrepareMultisigResponse represents the response model for PrepareMultisig
type PrepareMultisigResponse struct {
	// Multisig string to share with peers to create the multisig wallet.
	MultisigInfo string `json:"multisig_info"`
}

// MakeMultisigRequest represents the request model for MakeMultisig
type MakeMultisigRequest struct {
	// List of multisig string from peers.
	MultisigInfo []string `json:"multisig_info"`
	// Amount of signatures needed to sign a transfer. Must be less or equal than the amount of signature in multisig_info.
	Threshold uint64 `json:"threshold"`
	// Wallet password
	Password string `json:"password"`
}

// MakeMultisigResponse represents the response model for MakeMultisig
type MakeMultisigResponse struct {
	// Multisig wallet address.
	Address string `json:"address"`
	// Multisig string to share with peers to create the multisig wallet (extra step for N-1/N wallets).
	MultisigInfo string `json:"multisig_info"`
}

// ExportMultisigInfoResponse represents the response model for ExportMultisigInfo
type ExportMultisigInfoResponse struct {
	// Multisig info in hex format for other participants.
	Info string `json:"info"`
}

// ImportMultisigInfoRequest represents the request model for ImportMultisigInfo
type ImportMultisigInfoRequest struct {
	// List of multisig info in hex format from other participants.
	Info []string `json:"info"`
}

// ImportMultisigInfoResponse represents the response model for ImportMultisigInfo
type ImportMultisigInfoResponse struct {
	// Number of outputs signed with those multisig info.
	NOutputs uint64 `json:"n_outputs"`
}

// FinalizeMultisigRequest represents the request model for FinalizeMultisig
type FinalizeMultisigRequest struct {
	// List of multisig string from peers.
	MultisigInfo []string `json:"multisig_info"`
	// Wallet password
	Password string `json:"password"`
}

// FinalizeMultisigResponse represents the response model for FinalizeMultisig
type FinalizeMultisigResponse struct {
	// Multisig wallet address.
	Address string `json:"address"`
}

// SignMultisigRequest represents the request model for SignMultisig
type SignMultisigRequest struct {
	// Multisig transaction in hex format, as returned by transfer under multisig_txset.
	TxDataHex string `json:"tx_data_hex"`
}

// SignMultisigResponse represents the response model for SignMultisig
type SignMultisigResponse struct {
	// Multisig transaction in hex format.
	TxDataHex string `json:"tx_data_hex"`
	// List of transaction Hash.
	TxHashList []string `json:"tx_hash_list"`
}

// SubmitMultisigRequest represents the request model for SubmitMultisig
type SubmitMultisigRequest struct {
	// Multisig transaction in hex format, as returned by sign_multisig under tx_data_hex.
	TxDataHex string `json:"tx_data_hex"`
}

// SubmitMultisigResponse represents the response model for SubmitMultisig
type SubmitMultisigResponse struct {
	// List of transaction Hash.
	TxHashList []string `json:"tx_hash_list"`
}

// GetVersionResponse represents the response model for GetVersion
type GetVersionResponse struct {
	// RPC version, formatted with Major * 2^16 + Minor (Major encoded over the first 16 bits, and Minor over the last 16 bits).
	Version uint64 `json:"version"`
}
