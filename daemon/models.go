package daemon

// GetBlockCountResponse represents the response model for GetBlockCount
type GetBlockCountResponse struct {
	// Number of blocks in longest chain seen by the node.
	Count     uint64 `json:"count"`
	Untrusted bool   `json:"untrusted"`
}

// GetBlockTemplateRequest represents the request model for GetBlockTemplate
type GetBlockTemplateRequest struct {
	// Address of wallet to receive coinbase transactions if block is successfully mined.
	WalletAddress string `json:"wallet_address"`
	// Reserve size.
	ReserveSize uint64 `json:"reserve_size"`
}

// GetBlockTemplateResponse represents the response model for GetBlockTemplate
type GetBlockTemplateResponse struct {
	// Blob on which to try to mine a new block.
	BlocktemplateBlob string `json:"blocktemplate_blob"`
	// Blob on which to try to find a valid nonce.
	BlockhashingBlob string `json:"blockhashing_blob"`
	// Difficulty of next block.
	Difficulty uint64 `json:"difficulty"`
	// Coinbase reward expected to be received if block is successfully mined.
	ExpectedReward uint64 `json:"expected_reward"`
	// Height on which to mine.
	Height uint64 `json:"height"`
	// Hash of the most recent block on which to mine the next block.
	PrevHash string `json:"prev_hash"`
	// Reserved offset.
	ReservedOffset uint64 `json:"reserved_offset"`
	Untrusted      bool   `json:"untrusted"`
}

// SubmitBlockResponse represents the response model for SubmitBlock
type SubmitBlockResponse struct {
	// Block submit status.

}

// BlockHeader A structure containing block header information.
type BlockHeader struct {
	// The block size in bytes.
	BlockSize uint64 `json:"block_size"`
	// The number of blocks succeeding this block on the blockchain. A larger number means an older block.
	Depth uint64 `json:"depth"`
	// he strength of the Monero network based on mining power.
	Difficulty uint64 `json:"difficulty"`
	// The hash of this block.
	Hash string `json:"hash"`
	// The number of blocks preceding this block on the blockchain.
	Height uint64 `json:"height"`
	// The major version of the monero protocol at this block height.
	MajorVersion uint64 `json:"major_version"`
	// The minor version of the monero protocol at this block height.
	MinorVersion uint64 `json:"minor_version"`
	// a cryptographic random one-time number used in mining a Monero block.
	Nonce uint64 `json:"nonce"`
	// Number of transactions in the block, not counting the coinbase tx.
	NumTxes uint64 `json:"num_txes"`
	// Usually false. If true, this block is not part of the longest chain.
	OrphanStatus bool `json:"orphan_status"`
	// The hash of the block immediately preceding this block in the chain.
	PrevHash string `json:"prev_hash"`
	// The amount of new atomic units generated in this block and rewarded to the miner. Note: 1 XMR = 1e12 atomic units
	Reward uint64 `json:"reward"`
	// The unix time at which the block was recorded into the blockchain.
	Timestamp uint64 `json:"timestamp"`
}

// GetLastBlockHeaderResponse represents the response model for GetLastBlockHeader
type GetLastBlockHeaderResponse struct {
	// A structure containing block header information
	BlockHeader BlockHeader `json:"block_header"`
	Untrusted   bool        `json:"untrusted"`
}

// GetBlockHeaderByHashRequest represents the request model for GetBlockHeaderByHash
type GetBlockHeaderByHashRequest struct {
	// The block's sha256 hash.
	Hash string `json:"hash"`
}

// GetBlockHeaderByHashResponse represents the response model for GetBlockHeaderByHash
type GetBlockHeaderByHashResponse struct {
	// A structure containing block header information
	BlockHeader BlockHeader `json:"block_header"`
	Untrusted   bool        `json:"untrusted"`
}

// GetBlockHeaderByHeightRequest represents the request model for GetBlockHeaderByHeight
type GetBlockHeaderByHeightRequest struct {
	// The block's height.
	Height uint64 `json:"height"`
}

// GetBlockHeaderByHeightResponse represents the response model for GetBlockHeaderByHeight
type GetBlockHeaderByHeightResponse struct {
	// A structure containing block header information
	BlockHeader BlockHeader `json:"block_header"`
	Untrusted   bool        `json:"untrusted"`
}

// GetBlockHeadersRangeRequest represents the request model for GetBlockHeadersRange
type GetBlockHeadersRangeRequest struct {
	// The starting block's height.
	StartHeight uint64 `json:"start_height"`
	// The ending block's height.
	EndHeight uint64 `json:"end_height"`
}

// GetBlockHeadersRangeResponse represents the response model for GetBlockHeadersRange
type GetBlockHeadersRangeResponse struct {
	// (a structure containing block header information. See get_last_block_header).
	Headers   []BlockHeader `json:"headers"`
	Untrusted bool          `json:"untrusted"`
}

// GetBlockRequest represents the request model for GetBlock
type GetBlockRequest struct {
	// The block's height.
	Height uint64 `json:"height"`
	// The block's hash.
	Hash string `json:"string"`
}

// GetBlockResponse represents the response model for GetBlock
type GetBlockResponse struct {
	// Hexadecimal blob of block information.
	Blob string `json:"blob"`
	// A structure containing block header information
	BlockHeader BlockHeader `json:"block_header"`
	// SON formatted block details
	JSON      string `json:"json"`
	Untrusted bool   `json:"untrusted"`
}

// Connection model
type Connection struct {
	// The peer's address, actually IPv4 & port
	Address string `json:"address"`
	// Average bytes of data downloaded by node.
	AvgDownload uint64 `json:"avg_download"`
	// Average bytes of data uploaded by node.
	AvgUpload uint64 `json:"avg_upload"`
	// The connection ID
	ConnectionID string `json:"connection_id"`
	// Current bytes downloaded by node.
	CurrentDownload uint64 `json:"current_download"`
	// Current bytes uploaded by node.
	CurrentUpload uint64 `json:"current_upload"`
	// The peer height
	Height uint64 `json:"height"`
	// The peer host
	Host string `json:"host"`
	// Is the node getting information from your node?
	Incoming bool `json:"incoming"`
	// The node's IP address.
	IP string `json:"ip"`
	// LiveTime
	LiveTime uint64 `json:"live_time"`
	// LocalIP
	LocalIP bool `json:"local_ip"`
	// Localhost
	Localhost bool `json:"localhost"`
	// The node's ID on the network.
	PeerID string `json:"peer_id"`
	// The port that the node is using to connect to the network.
	Port string `json:"port"`
	// RecvCount
	RecvCount uint64 `json:"recv_count"`
	// RecvIdleTime
	RecvIdleTime uint64 `json:"recv_idle_time"`
	// SendCount
	SendCount uint64 `json:"send_count"`
	// SendIdleTime
	SendIdleTime uint64 `json:"send_idle_time"`
	// State
	State string `json:"state"`
	// SupportFlags
	SupportFlags uint64 `json:"support_flags"`
}

// GetConnectionsResponse represents the response model for GetConnections
type GetConnectionsResponse struct {
	Connections []Connection `json:"connections"`
}

// GetInfoResponse represents the response model for GetInfo
type GetInfoResponse struct {
	// Current time approximated from chain data, as Unix time.
	AdjustedTime uint64 `json:"adjusted_time"`
	// Number of alternative blocks to main chain.
	AltBlocksCount uint64 `json:"alt_blocks_count"`
	// Maximum allowed block size
	BlockSizeLimit uint64 `json:"block_size_limit"`
	// Median block size of latest 100 blocks
	BlockSizeMedian uint64 `json:"block_size_median"`
	// Maximum allowed adjusted block size based on latest 100000 blocks
	BlockWeightLimit uint64 `json:"block_weight_limit"`
	// Median adjusted block size of latest 100000 blocks
	BlockWeightMedian uint64 `json:"block_weight_median"`
	// bootstrap node to give immediate usability to wallets while syncing by proxying RPC to it. (Note: the replies may be untrustworthy).
	BootstrapDaemonAddress string `json:"bootstrap_daemon_address"`
	// States if new blocks are being added (true) or not (false).
	BusySyncing bool `json:"busy_syncing"`
	// If payment for RPC is enabled, the number of credits available to the requesting client. Otherwise, 0.
	Credits uint64 `json:"credits"`
	// Cumulative difficulty of all blocks in the blockchain.
	CumulativeDifficulty uint64 `json:"cumulative_difficulty"`
	// Most-significant 64 bits of the 128-bit cumulative difficulty.
	CumulativeDifficultyTop64 uint64 `json:"cumulative_difficulty_top64"`
	// The size of the blockchain database, in bytes.
	DatabaseSize uint64 `json:"database_size"`
	// Network difficulty (analogous to the strength of the network)
	Difficulty uint64 `json:"difficulty"`
	// Most-significant 64 bits of the 128-bit network difficulty.
	DifficultyTop64 uint64 `json:"difficulty_top64"`
	// Available disk space on the node.
	FreeSpace uint64 `json:"free_space"`
	// Grey Peerlist Size
	GreyPeerlistSize uint64 `json:"grey_peerlist_size"`
	// Current length of longest chain known to daemon.
	Height uint64 `json:"height"`
	// Current length of the local chain of the daemon.
	HeightWithoutBootstrap uint64 `json:"height_without_bootstrap"`
	// Number of peers connected to and pulling from your node.
	IncomingConnectionsCount uint64 `json:"incoming_connections_count"`
	// States if the node is on the mainnet (true) or not (false).
	Mainnet bool `json:"mainnet"`
	// Network type (one of mainnet, stagenet or testnet).
	NetType string `json:"nettype"`
	// States if the node is offline (true) or online (false).
	Offline bool `json:"offline"`
	// Number of peers that you are connected to and getting information from.
	OutgoingConnectionsCount uint64 `json:"outgoing_connections_count"`
	// Number of RPC client connected to the daemon (Including this RPC request).
	RPCConnectionsCount uint64 `json:"rpc_connections_count"`
	// States if the node is on the stagenet (true) or not (false).
	Stagenet bool `json:"stagenet"`
	// Start time of the daemon, as UNIX time.
	StartTime uint64 `json:"start_time"`
	// States if the node is synchronized (true) or not (false).
	Synchronized bool `json:"synchronized"`
	// Current target for next proof of work.
	Target uint64 `json:"target"`
	// The height of the next block in the chain.
	TargetHeight uint64 `json:"target_height"`
	// States if the node is on the testnet (true) or not (false).
	Testnet bool `json:"testnet"`
	// Hash of the highest block in the chain.
	TopBlockHash string `json:"top_block_hash"`
	// If payment for RPC is enabled, the hash of the highest block in the chain. Otherwise, empty.
	TopHash string `json:"top_hash"`
	// Total number of non-coinbase transaction in the chain.
	TxCount uint64 `json:"tx_count"`
	// Number of transactions that have been broadcast but not included in a block.
	TxPoolSize uint64 `json:"tx_pool_size"`
	// States if the result is obtained using the bootstrap mode, and is therefore not trusted (true), or when the daemon is fully synced (false).
	Untrusted bool `json:"untrusted"`
	//  States if a newer Monero software version is available.
	UpdateAvailable bool `json:"update_available"`
	// The version of the Monero software the node is running.
	Version string `json:"version"`
	// States if a bootstrap node has ever been used since the daemon started.
	WasBootstrapEverUsed bool `json:"was_bootstrap_ever_used"`
	// White Peerlist Size
	WhitePeerlistSize uint64 `json:"white_peerlist_size"`
	// Cumulative difficulty of all blocks in the blockchain as a hexadecimal string representing a 128-bit number.
	WideCumulativeDifficulty string `json:"wide_cumulative_difficulty"`
	// Network difficulty (analogous to the strength of the network) as a hexadecimal string representing a 128-bit number.
	WideDifficulty string `json:"wide_difficulty"`
}

// HardForkInfoResponse represents the response model for HardForkInfo
type HardForkInfoResponse struct {
	// Block height at which hard fork would be enabled if voted in.
	EarliestHeight uint64 `json:"earliest_height"`
	// Tells if hard fork is enforced.
	Enabled bool `json:"enabled"`
	// Current hard fork state: 0 (There is likely a hard fork), 1 (An update is needed to fork properly), or 2 (Everything looks good).
	State uint64 `json:"state"`
	// Minimum percent of votes to trigger hard fork. Default is 80.
	Threshold uint64 `json:"threshold"`
	// The major block version for the fork.
	Version uint64 `json:"version"`
	// Number of votes towards hard fork.
	Votes uint64 `json:"votes"`
	// Hard fork voting status.
	Voting uint64 `json:"voting"`
	// Number of blocks over which current votes are cast. Default is 10080 blocks.
	Window uint64 `json:"window"`
}

// Ban model
type Ban struct {
	// Host to ban (IP in A.B.C.D form - will support I2P address in the future).
	Host string `json:"host"`
	// IP address to ban, in Int format.
	IP uint64 `json:"ip"`
	// Set true to ban.
	Ban bool `json:"ban"`
	// Number of seconds to ban node.
	Seconds uint64 `json:"seconds"`
}

// SetBansRequest represents the request model for SetBans
type SetBansRequest struct {
	// A list of nodes to ban:
	Bans []Ban `json:"bans"`
}

// GetBansResponse represents the response model for GetBans
type GetBansResponse struct {
	// List of banned nodes
	Bans []Ban `json:"bans"`
}

// FlushTxpoolRequest represents the request model for FlushTxpool
type FlushTxpoolRequest struct {
	// Optional, list of transactions IDs to flush from pool (all tx ids flushed if empty).
	TxIDs []string `json:"txids,omitempty"`
}

// GetOutputHistogramRequest represents the request model for GetOutputHistogram
type GetOutputHistogramRequest struct {
	// list of unsigned int
	Amounts      []uint64 `json:"amounts"`
	MinCount     uint64   `json:"min_count"`
	MaxCount     uint64   `json:"max_count"`
	Unlocked     bool     `json:"unlocked"`
	RecentCutoff uint64   `json:"recent_cutoff"`
}

// Histogram model
type Histogram struct {
	// Output amount in atomic units
	Amount            uint64 `json:"amount"`
	TotalInstances    uint64 `json:"total_instances"`
	UnlockedInstances uint64 `json:"unlocked_instances"`
	RecentInstances   uint64 `json:"recent_instances"`
}

// GetOutputHistogramResponse represents the response model for GetOutputHistogram
type GetOutputHistogramResponse struct {
	// list of histogram entries
	Histogram []Histogram `json:"histogram"`
	// States if the result is obtained using the bootstrap mode, and is therefore not trusted (true), or when the daemon is fully synced (false).
	Untrusted bool `json:"untrusted"`
}

// GetVersionResponse represents the response model for GetVersion
type GetVersionResponse struct {
	Untrusted bool `json:"untrusted"`
	// Version
	Version uint64 `json:"version"`
}

// GetCoinbaseTxSumRequest represents the request model for GetCoinbaseTxSum
type GetCoinbaseTxSumRequest struct {
	// Block height from which getting the amounts
	Height uint64 `json:"height"`
	// number of blocks to include in the sum
	Count uint64 `json:"count"`
}

// GetCoinbaseTxSumResponse represents the response model for GetCoinbaseTxSum
type GetCoinbaseTxSumResponse struct {
	// amount of coinbase reward in atomic units
	EmissionAmount uint64 `json:"emission_amount"`
	// amount of fees in atomic units
	FeeAmount uint64 `json:"fee_amount"`
}

// GetFeeEstimateRequest represents the request model for GetFeeEstimate
type GetFeeEstimateRequest struct {
	// Optional
	GraceBlocks uint64 `json:"grace_blocks,omitempty"`
}

// GetFeeEstimateResponse represents the response model for GetFeeEstimate
type GetFeeEstimateResponse struct {
	// Amount of fees estimated per byte in atomic units
	Fee uint64 `json:"fee"`
	// Final fee should be rounded up to an even multiple of this value
	QuantizationMask uint64 `json:"quantization_mask"`
	// States if the result is obtained using the bootstrap mode, and is therefore not trusted (true), or when the daemon is fully synced (false).
	Untrusted bool `json:"untrusted"`
}

// Chain model
type Chain struct {
	// the block hash of the first diverging block of this alternative chain.
	BlockHash string `json:"block_hash"`
	// the cumulative difficulty of all blocks in the alternative chain.
	Difficulty uint64 `json:"difficulty"`
	// the block height of the first diverging block of this alternative chain.
	Height uint64 `json:"height"`
	// the length in blocks of this alternative chain, after divergence.
	Length uint64 `json:"length"`
}

// GetAlternateChainsResponse represents the response model for GetAlternateChains
type GetAlternateChainsResponse struct {
	// array of chains
	Chains []Chain `json:"chains"`
}

// RelayTxRequest represents the request model for RelayTx
type RelayTxRequest struct {
	// list of transaction IDs to relay
	TxIDs []string `json:"txids"`
}

// Peer model
type Peer struct {
	Info Connection `json:"info"`
}

// Span model
type Span struct {
	// Id of connection
	ConnectionID string `json:"connection_id"`
	// number of blocks in that span
	Nblocks uint64 `json:"nblocks"`
	// connection rate
	Rate uint64 `json:"rate"`
	// peer address the node is downloading (or has downloaded) than span from
	RemoteAddress string `json:"remote_address"`
	// total number of bytes in that span's blocks (including txes)
	Size uint64 `json:"size"`
	// connection speed
	Speed uint64 `json:"speed"`
	// block height of the first block in that span
	StartBlockHeight uint64 `json:"start_block_height"`
}

// SyncInfoResponse represents the response model for SyncInfo
type SyncInfoResponse struct {
	Height uint64 `json:"height"`
	// array of peer structure
	Peers []Peer `json:"peers"`
	// array of span structure, defined as follows (optional, absent if node is fully synced)
	Spans []Span `json:"spans"`
	// target height the node is syncing from (will be undefined if node is fully synced)
	TargetHeight uint64 `json:"target_height"`
}

// TXBacklogEntry model
type TXBacklogEntry struct {
	// (in binary form)
	BlobSize uint64 `json:"blob_size"`
	// (in binary form)
	Fee uint64 `json:"fee"`
	// (in binary form)
	TimeInPool uint64 `json:"time_in_pool"`
}

// GetTxpoolBacklogResponse represents the response model for GetTxpoolBacklog
type GetTxpoolBacklogResponse struct {
	// array of structures tx_backlog_entry (in binary form)
	Backlog []TXBacklogEntry `json:"backlog"`
	// States if the result is obtained using the bootstrap mode, and is therefore not trusted (true), or when the daemon is fully synced (false).
	Untrusted bool `json:"untrusted"`
}

// GetOutputDistributionRequest represents the request model for GetOutputDistribution
type GetOutputDistributionRequest struct {
	// amounts to look for
	Amounts []uint64 `json:"amounts"`
	// (optional, default is false) States if the result should be cumulative (true) or not (false)
	Cumulative bool `json:"cumulative,omitempty"`
	// (optional, default is 0) starting height to check from
	FromHeight uint64 `json:"from_height,omitempty"`
	// (optional, default is 0) ending height to check up to
	ToHeight uint64 `json:"to_height,omitempty"`
}

// Distribution model
type Distribution struct {
	Amount       uint64   `json:"amount"`
	Base         uint64   `json:"base"`
	Distribution []uint64 `json:"distribution"`
	StartHeight  uint64   `json:"start_height"`
}

// GetOutputDistributionResponse represents the response model for GetOutputDistribution
type GetOutputDistributionResponse struct {
	// array of structure distribution
	Distributions []Distribution `json:"distributions"`
}

type GenerateBlocksRequest struct {
	// AmountOfBlocks is the number of blocks to be generated.
	AmountOfBlocks uint64 `json:"amount_of_blocks,omitempty"`
	// WalletAddress is the address of the wallet that will get the rewards
	// of the coinbase transaction for such the blocks generates.
	WalletAddress string `json:"wallet_address,omitempty"`
	PreviousBlock string `json:"prev_block,omitempty"`
	StartingNonce uint64 `json:"starting_nonce,omitempty"`
}
type GenerateBlocksResponse struct {
	Blocks    []string `json:"blocks"`
	Height    uint64   `json:"height"`
	Untrusted bool     `json:"untrusted"`
}
