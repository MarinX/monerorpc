package daemon

// Daemon interface is a list of the monerod daemon RPC calls, their inputs and outputs, and examples of each.
// Many RPC calls use the daemon's JSON RPC interface while others use their own interfaces, as demonstrated below.
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
	SetBans(req *SetBansRequest) error
	// GetBans Get list of banned IPs.
	GetBans() (*GetBansResponse, error)
	// FlushTxpool Flush tx ids from transaction pool
	FlushTxpool(req *FlushTxpoolRequest) error
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
	RelayTx(req *RelayTxRequest) error
	// SyncInfo Get synchronisation informations
	SyncInfo() (*SyncInfoResponse, error)
	// GetTxpoolBacklog Get all transaction pool backlog
	GetTxpoolBacklog() (*GetTxpoolBacklogResponse, error)
	// GetOutputDistribution Alias: None.
	GetOutputDistribution(req *GetOutputDistributionRequest) (*GetOutputDistributionResponse, error)
}

// MoneroRPC interface for client
type MoneroRPC interface {
	Do(method string, req interface{}, res interface{}) error
}

type daemon struct {
	client MoneroRPC
}

// New creates a new daemon client
func New(client MoneroRPC) Daemon {
	return &daemon{
		client: client,
	}
}

func (d *daemon) GenerateBlocks(req *GenerateBlocksRequest) (*GenerateBlocksResponse, error) {
	res := new(GenerateBlocksResponse)
	err := d.client.Do("generateblocks", req, res)
	return res, err
}

func (d *daemon) GetBlockCount() (*GetBlockCountResponse, error) {
	res := new(GetBlockCountResponse)
	err := d.client.Do("get_block_count", nil, res)
	return res, err
}

func (d *daemon) OnGetBlockHash(req []uint64) (string, error) {
	var res string
	err := d.client.Do("on_get_block_hash", req, &res)
	return res, err
}

func (d *daemon) GetBlockTemplate(req *GetBlockTemplateRequest) (*GetBlockTemplateResponse, error) {
	res := new(GetBlockTemplateResponse)
	err := d.client.Do("get_block_template", req, res)
	return res, err
}

func (d *daemon) SubmitBlock(req []string) (*SubmitBlockResponse, error) {
	res := new(SubmitBlockResponse)
	err := d.client.Do("submit_block", &req, res)
	return res, err
}

func (d *daemon) GetLastBlockHeader() (*GetLastBlockHeaderResponse, error) {
	res := new(GetLastBlockHeaderResponse)
	err := d.client.Do("get_last_block_header", nil, res)
	return res, err
}

func (d *daemon) GetBlockHeaderByHash(req *GetBlockHeaderByHashRequest) (*GetBlockHeaderByHashResponse, error) {
	res := new(GetBlockHeaderByHashResponse)
	err := d.client.Do("get_block_header_by_hash", req, res)
	return res, err
}

func (d *daemon) GetBlockHeaderByHeight(req *GetBlockHeaderByHeightRequest) (*GetBlockHeaderByHeightResponse, error) {
	res := new(GetBlockHeaderByHeightResponse)
	err := d.client.Do("get_block_header_by_height", req, res)
	return res, err
}

func (d *daemon) GetBlockHeadersRange(req *GetBlockHeadersRangeRequest) (*GetBlockHeadersRangeResponse, error) {
	res := new(GetBlockHeadersRangeResponse)
	err := d.client.Do("get_block_headers_range", req, res)
	return res, err
}

func (d *daemon) GetBlock(req *GetBlockRequest) (*GetBlockResponse, error) {
	res := new(GetBlockResponse)
	err := d.client.Do("get_block", req, res)
	return res, err
}

func (d *daemon) GetConnections() (*GetConnectionsResponse, error) {
	res := new(GetConnectionsResponse)
	err := d.client.Do("get_connections", nil, res)
	return res, err
}

func (d *daemon) GetInfo() (*GetInfoResponse, error) {
	res := new(GetInfoResponse)
	err := d.client.Do("get_info", nil, res)
	return res, err
}

func (d *daemon) HardForkInfo() (*HardForkInfoResponse, error) {
	res := new(HardForkInfoResponse)
	err := d.client.Do("hard_fork_info", nil, res)
	return res, err
}

func (d *daemon) SetBans(req *SetBansRequest) error {
	err := d.client.Do("set_bans", req, nil)
	return err
}

func (d *daemon) GetBans() (*GetBansResponse, error) {
	res := new(GetBansResponse)
	err := d.client.Do("get_bans", nil, res)
	return res, err
}

func (d *daemon) FlushTxpool(req *FlushTxpoolRequest) error {
	err := d.client.Do("flush_txpool", req, nil)
	return err
}

func (d *daemon) GetOutputHistogram(req *GetOutputHistogramRequest) (*GetOutputHistogramResponse, error) {
	res := new(GetOutputHistogramResponse)
	err := d.client.Do("get_output_histogram", req, res)
	return res, err
}

func (d *daemon) GetVersion() (*GetVersionResponse, error) {
	res := new(GetVersionResponse)
	err := d.client.Do("get_version", nil, res)
	return res, err
}

func (d *daemon) GetCoinbaseTxSum(req *GetCoinbaseTxSumRequest) (*GetCoinbaseTxSumResponse, error) {
	res := new(GetCoinbaseTxSumResponse)
	err := d.client.Do("get_coinbase_tx_sum", req, res)
	return res, err
}

func (d *daemon) GetFeeEstimate(req *GetFeeEstimateRequest) (*GetFeeEstimateResponse, error) {
	res := new(GetFeeEstimateResponse)
	err := d.client.Do("get_fee_estimate", req, res)
	return res, err
}

func (d *daemon) GetAlternateChains() (*GetAlternateChainsResponse, error) {
	res := new(GetAlternateChainsResponse)
	err := d.client.Do("get_alternate_chains", nil, res)
	return res, err
}

func (d *daemon) RelayTx(req *RelayTxRequest) error {
	err := d.client.Do("relay_tx", req, nil)
	return err
}

func (d *daemon) SyncInfo() (*SyncInfoResponse, error) {
	res := new(SyncInfoResponse)
	err := d.client.Do("sync_info", nil, res)
	return res, err
}

func (d *daemon) GetTxpoolBacklog() (*GetTxpoolBacklogResponse, error) {
	res := new(GetTxpoolBacklogResponse)
	err := d.client.Do("get_txpool_backlog", nil, res)
	return res, err
}

func (d *daemon) GetOutputDistribution(req *GetOutputDistributionRequest) (*GetOutputDistributionResponse, error) {
	res := new(GetOutputDistributionResponse)
	err := d.client.Do("get_output_distribution", req, res)
	return res, err
}
