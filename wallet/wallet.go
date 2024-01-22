package wallet

// Wallet interface is a list of the monero-wallet-rpc calls, their inputs and outputs, and examples of each.
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

// MoneroRPC interface for client
type MoneroRPC interface {
	Do(method string, req interface{}, res interface{}) error
}

type wallet struct {
	client MoneroRPC
}

// New creates a new wallet client
func New(client MoneroRPC) Wallet {
	return &wallet{
		client: client,
	}
}

// SetDaemon connects the RPC server to a Monero daemon.
func (w *wallet) SetDaemon(req *SetDaemonRequest) error {
	return w.client.Do("set_daemon", req, nil)
}

func (w *wallet) GetBalance(req *GetBalanceRequest) (*GetBalanceResponse, error) {
	res := new(GetBalanceResponse)
	err := w.client.Do("get_balance", req, res)
	return res, err
}

func (w *wallet) GetAddress(req *GetAddressRequest) (*GetAddressResponse, error) {
	res := new(GetAddressResponse)
	err := w.client.Do("get_address", req, res)
	return res, err
}

func (w *wallet) GetAddressIndex(req *GetAddressIndexRequest) (*GetAddressIndexResponse, error) {
	res := new(GetAddressIndexResponse)
	err := w.client.Do("get_address_index", req, res)
	return res, err
}

func (w *wallet) CreateAddress(req *CreateAddressRequest) (*CreateAddressResponse, error) {
	res := new(CreateAddressResponse)
	err := w.client.Do("create_address", req, res)
	return res, err
}

func (w *wallet) LabelAddress(req *LabelAddressRequest) error {
	return w.client.Do("label_address", req, nil)
}

func (w *wallet) ValidateAddress(req *ValidateAddressRequest) (*ValidateAddressResponse, error) {
	res := new(ValidateAddressResponse)
	err := w.client.Do("validate_address", req, res)
	return res, err
}

func (w *wallet) GetAccounts(req *GetAccountsRequest) (*GetAccountsResponse, error) {
	res := new(GetAccountsResponse)
	err := w.client.Do("get_accounts", req, res)
	return res, err
}

func (w *wallet) CreateAccount(req *CreateAccountRequest) (*CreateAccountResponse, error) {
	res := new(CreateAccountResponse)
	err := w.client.Do("create_account", req, res)
	return res, err
}

func (w *wallet) LabelAccount(req *LabelAccountRequest) error {
	return w.client.Do("label_account", req, nil)
}

func (w *wallet) GetAccountTags() (*GetAccountTagsResponse, error) {
	res := new(GetAccountTagsResponse)
	err := w.client.Do("get_account_tags", nil, res)
	return res, err
}

func (w *wallet) TagAccounts(req *TagAccountsRequest) error {
	return w.client.Do("tag_accounts", req, nil)
}

func (w *wallet) UntagAccounts(req *UntagAccountsRequest) error {
	return w.client.Do("untag_accounts", req, nil)
}

func (w *wallet) SetAccountTagDescription(req *SetAccountTagDescriptionRequest) error {
	return w.client.Do("set_account_tag_description", req, nil)
}

func (w *wallet) GetHeight() (*GetHeightResponse, error) {
	res := new(GetHeightResponse)
	err := w.client.Do("get_height", nil, res)
	return res, err
}

func (w *wallet) Transfer(req *TransferRequest) (*TransferResponse, error) {
	res := new(TransferResponse)
	err := w.client.Do("transfer", req, res)
	return res, err
}

func (w *wallet) TransferSplit(req *TransferSplitRequest) (*TransferSplitResponse, error) {
	res := new(TransferSplitResponse)
	err := w.client.Do("transfer_split", req, res)
	return res, err
}

func (w *wallet) SignTransfer(req *SignTransferRequest) (*SignTransferResponse, error) {
	res := new(SignTransferResponse)
	err := w.client.Do("sign_transfer", req, res)
	return res, err
}

func (w *wallet) SubmitTransfer(req *SubmitTransferRequest) (*SubmitTransferResponse, error) {
	res := new(SubmitTransferResponse)
	err := w.client.Do("submit_transfer", req, res)
	return res, err
}

func (w *wallet) SweepDust(req *SweepDustRequest) (*SweepDustResponse, error) {
	res := new(SweepDustResponse)
	err := w.client.Do("sweep_dust", req, res)
	return res, err
}

func (w *wallet) SweepAll(req *SweepAllRequest) (*SweepAllResponse, error) {
	res := new(SweepAllResponse)
	err := w.client.Do("sweep_all", req, res)
	return res, err
}

func (w *wallet) SweepSingle(req *SweepSingleRequest) (*SweepSingleResponse, error) {
	res := new(SweepSingleResponse)
	err := w.client.Do("sweep_single", req, res)
	return res, err
}

func (w *wallet) RelayTx(req *RelayTxRequest) (*RelayTxResponse, error) {
	res := new(RelayTxResponse)
	err := w.client.Do("relay_tx", req, res)
	return res, err
}

func (w *wallet) Store() error {
	return w.client.Do("store", nil, nil)
}

func (w *wallet) GetPayments(req *GetPaymentsRequest) (*GetPaymentsResponse, error) {
	res := new(GetPaymentsResponse)
	err := w.client.Do("get_payments", req, res)
	return res, err
}

func (w *wallet) GetBulkPayments(req *GetBulkPaymentsRequest) (*GetBulkPaymentsResponse, error) {
	res := new(GetBulkPaymentsResponse)
	err := w.client.Do("get_bulk_payments", req, res)
	return res, err
}

func (w *wallet) IncomingTransfers(req *IncomingTransfersRequest) (*IncomingTransfersResponse, error) {
	res := new(IncomingTransfersResponse)
	err := w.client.Do("incoming_transfers", req, res)
	return res, err
}

func (w *wallet) QueryKey(req *QueryKeyRequest) (*QueryKeyResponse, error) {
	res := new(QueryKeyResponse)
	err := w.client.Do("query_key", req, res)
	return res, err
}

func (w *wallet) MakeIntegratedAddress(req *MakeIntegratedAddressRequest) (*MakeIntegratedAddressResponse, error) {
	res := new(MakeIntegratedAddressResponse)
	err := w.client.Do("make_integrated_address", req, res)
	return res, err
}

func (w *wallet) SplitIntegratedAddress(req *SplitIntegratedAddressRequest) (*SplitIntegratedAddressResponse, error) {
	res := new(SplitIntegratedAddressResponse)
	err := w.client.Do("split_integrated_address", req, res)
	return res, err
}

func (w *wallet) StopWallet() error {
	return w.client.Do("stop_wallet", nil, nil)
}

func (w *wallet) RescanBlockchain() error {
	return w.client.Do("rescan_blockchain", nil, nil)
}

func (w *wallet) SetTxNotes(req *SetTxNotesRequest) error {
	return w.client.Do("set_tx_notes", req, nil)
}

func (w *wallet) GetTxNotes(req *GetTxNotesRequest) (*GetTxNotesResponse, error) {
	res := new(GetTxNotesResponse)
	err := w.client.Do("get_tx_notes", req, res)
	return res, err
}

func (w *wallet) SetAttribute(req *SetAttributeRequest) error {
	return w.client.Do("set_attribute", req, nil)
}

func (w *wallet) GetAttribute(req *GetAttributeRequest) (*GetAttributeResponse, error) {
	res := new(GetAttributeResponse)
	err := w.client.Do("get_attribute", req, res)
	return res, err
}

func (w *wallet) GetTxKey(req *GetTxKeyRequest) (*GetTxKeyResponse, error) {
	res := new(GetTxKeyResponse)
	err := w.client.Do("get_tx_key", req, res)
	return res, err
}

func (w *wallet) CheckTxKey(req *CheckTxKeyRequest) (*CheckTxKeyResponse, error) {
	res := new(CheckTxKeyResponse)
	err := w.client.Do("check_tx_key", req, res)
	return res, err
}

func (w *wallet) GetTxProof(req *GetTxProofRequest) (*GetTxProofResponse, error) {
	res := new(GetTxProofResponse)
	err := w.client.Do("get_tx_proof", req, res)
	return res, err
}

func (w *wallet) CheckTxProof(req *CheckTxProofRequest) (*CheckTxProofResponse, error) {
	res := new(CheckTxProofResponse)
	err := w.client.Do("check_tx_proof", req, res)
	return res, err
}

func (w *wallet) GetSpendProof(req *GetSpendProofRequest) (*GetSpendProofResponse, error) {
	res := new(GetSpendProofResponse)
	err := w.client.Do("get_spend_proof", req, res)
	return res, err
}

func (w *wallet) CheckSpendProof(req *CheckSpendProofRequest) (*CheckSpendProofResponse, error) {
	res := new(CheckSpendProofResponse)
	err := w.client.Do("check_spend_proof", req, res)
	return res, err
}

func (w *wallet) GetReserveProof(req *GetReserveProofRequest) (*GetReserveProofResponse, error) {
	res := new(GetReserveProofResponse)
	err := w.client.Do("get_reserve_proof", req, res)
	return res, err
}

func (w *wallet) CheckReserveProof(req *CheckReserveProofRequest) (*CheckReserveProofResponse, error) {
	res := new(CheckReserveProofResponse)
	err := w.client.Do("check_reserve_proof", req, res)
	return res, err
}

func (w *wallet) GetTransfers(req *GetTransfersRequest) (*GetTransfersResponse, error) {
	res := new(GetTransfersResponse)
	err := w.client.Do("get_transfers", req, res)
	return res, err
}

func (w *wallet) GetTransferByTxid(req *GetTransferByTxidRequest) (*GetTransferByTxidResponse, error) {
	res := new(GetTransferByTxidResponse)
	err := w.client.Do("get_transfer_by_txid", req, res)
	return res, err
}

func (w *wallet) DescribeTransfer(req *DescribeTransferRequest) (*DescribeTransferResponse, error) {
	res := new(DescribeTransferResponse)
	err := w.client.Do("describe_transfer", req, res)
	return res, err
}

func (w *wallet) Sign(req *SignRequest) (*SignResponse, error) {
	res := new(SignResponse)
	err := w.client.Do("sign", req, res)
	return res, err
}

func (w *wallet) Verify(req *VerifyRequest) (*VerifyResponse, error) {
	res := new(VerifyResponse)
	err := w.client.Do("verify", req, res)
	return res, err
}

func (w *wallet) ExportOutputs(req *ExportOutputsRequest) (*ExportOutputsResponse, error) {
	res := new(ExportOutputsResponse)
	err := w.client.Do("export_outputs", req, res)
	return res, err
}

func (w *wallet) ImportOutputs(req *ImportOutputsRequest) (*ImportOutputsResponse, error) {
	res := new(ImportOutputsResponse)
	err := w.client.Do("import_outputs", req, res)
	return res, err
}

func (w *wallet) ExportKeyImages(req *ExportKeyImagesRequest) (*ExportKeyImagesResponse, error) {
	res := new(ExportKeyImagesResponse)
	err := w.client.Do("export_key_images", req, res)
	return res, err
}

func (w *wallet) ImportKeyImages(req *ImportKeyImagesRequest) (*ImportKeyImagesResponse, error) {
	res := new(ImportKeyImagesResponse)
	err := w.client.Do("import_key_images", req, res)
	return res, err
}

func (w *wallet) MakeURI(req *MakeURIRequest) (*MakeURIResponse, error) {
	res := new(MakeURIResponse)
	err := w.client.Do("make_uri", req, res)
	return res, err
}

func (w *wallet) ParseURI(req *ParseURIRequest) (*ParseURIResponse, error) {
	res := new(ParseURIResponse)
	err := w.client.Do("parse_uri", req, res)
	return res, err
}

func (w *wallet) GetAddressBook(req *GetAddressBookRequest) (*GetAddressBookResponse, error) {
	res := new(GetAddressBookResponse)
	err := w.client.Do("get_address_book", req, res)
	return res, err
}

func (w *wallet) AddAddressBook(req *AddAddressBookRequest) (*AddAddressBookResponse, error) {
	res := new(AddAddressBookResponse)
	err := w.client.Do("add_address_book", req, res)
	return res, err
}

func (w *wallet) EditAddressBook(req *EditAddressBookRequest) error {
	return w.client.Do("edit_address_book", nil, nil)
}

func (w *wallet) DeleteAddressBook(req *DeleteAddressBookRequest) error {
	return w.client.Do("delete_address_book", req, nil)
}

func (w *wallet) Refresh(req *RefreshRequest) (*RefreshResponse, error) {
	res := new(RefreshResponse)
	err := w.client.Do("refresh", req, res)
	return res, err
}

func (w *wallet) AutoRefresh(req *AutoRefreshRequest) error {
	return w.client.Do("auto_refresh", req, nil)
}

func (w *wallet) RescanSpent() error {
	return w.client.Do("rescan_spent", nil, nil)
}

func (w *wallet) StartMining(req *StartMiningRequest) error {
	return w.client.Do("start_mining", req, nil)
}

func (w *wallet) StopMining() error {
	return w.client.Do("stop_mining", nil, nil)
}

func (w *wallet) GetLanguages() (*GetLanguagesResponse, error) {
	res := new(GetLanguagesResponse)
	err := w.client.Do("get_languages", nil, res)
	return res, err
}

func (w *wallet) CreateWallet(req *CreateWalletRequest) error {
	return w.client.Do("create_wallet", req, nil)
}

func (w *wallet) GenerateFromKeys(req *GenerateFromKeysRequest) (*GenerateFromKeysResponse, error) {
	res := new(GenerateFromKeysResponse)
	err := w.client.Do("generate_from_keys", req, res)
	return res, err
}

func (w *wallet) OpenWallet(req *OpenWalletRequest) error {
	return w.client.Do("open_wallet", req, nil)
}

func (w *wallet) RestoreDeterministicWallet(req *RestoreDeterministicWalletRequest) (*RestoreDeterministicWalletResponse, error) {
	res := new(RestoreDeterministicWalletResponse)
	err := w.client.Do("restore_deterministic_wallet", req, res)
	return res, err
}

func (w *wallet) CloseWallet() error {
	return w.client.Do("close_wallet", nil, nil)
}

func (w *wallet) ChangeWalletPassword(req *ChangeWalletPasswordRequest) error {
	return w.client.Do("change_wallet_password", req, nil)
}

func (w *wallet) IsMultisig() (*IsMultisigResponse, error) {
	res := new(IsMultisigResponse)
	err := w.client.Do("is_multisig", nil, res)
	return res, err
}

func (w *wallet) PrepareMultisig() (*PrepareMultisigResponse, error) {
	res := new(PrepareMultisigResponse)
	err := w.client.Do("prepare_multisig", nil, res)
	return res, err
}

func (w *wallet) MakeMultisig(req *MakeMultisigRequest) (*MakeMultisigResponse, error) {
	res := new(MakeMultisigResponse)
	err := w.client.Do("make_multisig", req, res)
	return res, err
}

func (w *wallet) ExportMultisigInfo() (*ExportMultisigInfoResponse, error) {
	res := new(ExportMultisigInfoResponse)
	err := w.client.Do("export_multisig_info", nil, res)
	return res, err
}

func (w *wallet) ImportMultisigInfo(req *ImportMultisigInfoRequest) (*ImportMultisigInfoResponse, error) {
	res := new(ImportMultisigInfoResponse)
	err := w.client.Do("import_multisig_info", req, res)
	return res, err
}

func (w *wallet) FinalizeMultisig(req *FinalizeMultisigRequest) (*FinalizeMultisigResponse, error) {
	res := new(FinalizeMultisigResponse)
	err := w.client.Do("finalize_multisig", req, res)
	return res, err
}

func (w *wallet) SignMultisig(req *SignMultisigRequest) (*SignMultisigResponse, error) {
	res := new(SignMultisigResponse)
	err := w.client.Do("sign_multisig", req, res)
	return res, err
}

func (w *wallet) SubmitMultisig(req *SubmitMultisigRequest) (*SubmitMultisigResponse, error) {
	res := new(SubmitMultisigResponse)
	err := w.client.Do("submit_multisig", req, res)
	return res, err
}

func (w *wallet) GetVersion() (*GetVersionResponse, error) {
	res := new(GetVersionResponse)
	err := w.client.Do("get_version", nil, res)
	return res, err
}
