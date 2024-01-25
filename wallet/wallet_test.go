package wallet

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/rpc/v2/json2"
	"github.com/matryer/is"
)

type MockMoneroRPC struct {
	client *http.Client
	uri    string
}

func (m *MockMoneroRPC) Do(method string, req interface{}, res interface{}) error {
	buff, err := json2.EncodeClientRequest(method, req)
	if err != nil {
		return fmt.Errorf("error creating encoded request %v", err)
	}

	httpReq, err := http.NewRequest("POST", m.uri, bytes.NewReader(buff))
	if err != nil {
		return fmt.Errorf("error creating http request %v", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := m.client.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode == http.StatusUnauthorized {
		return fmt.Errorf("unauthorized - invalid username or password")
	}

	if res != nil {
		err = json2.DecodeClientResponse(httpResp.Body, res)
		if err == json2.ErrNullResult {
			return nil
		}
	}
	return err
}

func getClient(uri string, client *http.Client) *MockMoneroRPC {
	return &MockMoneroRPC{
		uri:    uri,
		client: client,
	}
}

func setupServer(t *testing.T, method string, output string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		buff, _ := io.ReadAll(req.Body)
		t.Log(string(buff))
		rw.Write([]byte(output))
	}))
	return server
}

func TestWalletSetDaemon(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		}
	  }`
	server := setupServer(t, "set_daemon", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	err := w.SetDaemon(&SetDaemonRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletGetBalance(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
			"balance": 17981101058456048,
			"blocks_to_unlock": 59,
			"multisig_import_needed": false,
			"per_subaddress": [{
			"account_index": 0,
			"address": "4Ae8FJHMJsnfsghVcg1u3SWqcLyUKCzNVMq8JLZCzgpSKCu6X2mEWZJXAqQpyJsQ11KVXgJDJ24LaBWCjYK8jQXU4NUpiJn",
			"address_index": 0,
			"balance": 17981101058456048,
			"blocks_to_unlock": 59,
			"label": "Primary account",
			"num_unspent_outputs": 512,
			"time_to_unlock": 0,
			"unlocked_balance": 15909955156562018
			}],
			"time_to_unlock": 0,
			"unlocked_balance": 15909955156562018
		}
	  }`
	server := setupServer(t, "get_balance", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	res, err := w.GetBalance(&GetBalanceRequest{})
	if err != nil {
		t.Error(err)
	}

	is.New(t).Equal(res, &GetBalanceResponse{
		Balance:              17981101058456048,
		BlocksToUnlock:       59,
		MultisigImportNeeded: false,
		PerSubaddress: []PerSubaddress{
			{
				AccountIndex:      0,
				Address:           "4Ae8FJHMJsnfsghVcg1u3SWqcLyUKCzNVMq8JLZCzgpSKCu6X2mEWZJXAqQpyJsQ11KVXgJDJ24LaBWCjYK8jQXU4NUpiJn",
				AddressIndex:      0,
				Balance:           17981101058456048,
				BlocksToUnlock:    59,
				Label:             "Primary account",
				NumUnspentOutputs: 512,
				TimeToUnlock:      0,
				UnlockedBalance:   15909955156562018,
			},
		},
		TimeToUnlock:    0,
		UnlockedBalance: 15909955156562018,
	})
}

func TestWalletGetAddress(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "address": "55LTR8KniP4LQGJSPtbYDacR7dz8RBFnsfAKMaMuwUNYX6aQbBcovzDPyrQF9KXF9tVU6Xk3K8no1BywnJX6GvZX8yJsXvt",
		  "addresses": [{
			"address": "55LTR8KniP4LQGJSPtbYDacR7dz8RBFnsfAKMaMuwUNYX6aQbBcovzDPyrQF9KXF9tVU6Xk3K8no1BywnJX6GvZX8yJsXvt",
			"address_index": 0,
			"label": "Primary account",
			"used": true
		  },{
			"address": "7BnERTpvL5MbCLtj5n9No7J5oE5hHiB3tVCK5cjSvCsYWD2WRJLFuWeKTLiXo5QJqt2ZwUaLy2Vh1Ad51K7FNgqcHgjW85o",
			"address_index": 1,
			"label": "",
			"used": true
		  },{
			"address": "77xa6Dha7kzCQuvmd8iB5VYoMkdenwCNRU9khGhExXQ8KLL3z1N1ZATBD1sFPenyHWT9cm4fVFnCAUApY53peuoZFtwZiw5",
			"address_index": 4,
			"label": "test2",
			"used": true
		  }]
		}
	  }`
	server := setupServer(t, "get_address", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.GetAddress(&GetAddressRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletGetAddressIndex(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "index": {
			"major": 0,
			"minor": 1
		  }
		}
	  }`
	server := setupServer(t, "get_address_index", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.GetAddressIndex(&GetAddressIndexRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletCreateAddress(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "address": "7BG5jr9QS5sGMdpbBrZEwVLZjSKJGJBsXdZLt8wiXyhhLjy7x2LZxsrAnHTgD8oG46ZtLjUGic2pWc96GFkGNPQQDA3Dt7Q",
		  "address_index": 5
		}
	  }`
	server := setupServer(t, "create_address", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.CreateAddress(&CreateAddressRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletLabelAddress(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		}
	  }`
	server := setupServer(t, "label_address", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	err := w.LabelAddress(&LabelAddressRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletValidateAddress(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "valid": true,
		  "integrated": false,
		  "subaddress": false,
		  "nettype": "mainnet",
		  "openalias_address": ""
		}
	  }`
	server := setupServer(t, "validate_address", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.ValidateAddress(&ValidateAddressRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletGetAccounts(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "subaddress_accounts": [{
			"account_index": 0,
			"balance": 157663195572433688,
			"base_address": "55LTR8KniP4LQGJSPtbYDacR7dz8RBFnsfAKMaMuwUNYX6aQbBcovzDPyrQF9KXF9tVU6Xk3K8no1BywnJX6GvZX8yJsXvt",
			"label": "Primary account",
			"tag": "myTag",
			"unlocked_balance": 157443303037455077
		  },{
			"account_index": 1,
			"balance": 0,
			"base_address": "77Vx9cs1VPicFndSVgYUvTdLCJEZw9h81hXLMYsjBCXSJfUehLa9TDW3Ffh45SQa7xb6dUs18mpNxfUhQGqfwXPSMrvKhVp",
			"label": "Secondary account",
			"tag": "myTag",
			"unlocked_balance": 0
		  }],
		  "total_balance": 157663195572433688,
		  "total_unlocked_balance": 157443303037455077
		}
	  }`
	server := setupServer(t, "get_accounts", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.GetAccounts(&GetAccountsRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletCreateAccount(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "account_index": 1,
		  "address": "77Vx9cs1VPicFndSVgYUvTdLCJEZw9h81hXLMYsjBCXSJfUehLa9TDW3Ffh45SQa7xb6dUs18mpNxfUhQGqfwXPSMrvKhVp"
		}
	  }`
	server := setupServer(t, "create_account", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.CreateAccount(&CreateAccountRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletLabelAccount(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "account_tags": [{
			"accounts": [0,1],
			"label": "",
			"tag": "myTag"
		  }]
		}
	  }`
	server := setupServer(t, "label_account", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	err := w.LabelAccount(&LabelAccountRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletGetAccountTags(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "account_tags": [{
			"accounts": [0],
			"label": "Test tag",
			"tag": "myTag"
		  }]
		}
	  }`
	server := setupServer(t, "get_account_tags", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.GetAccountTags()
	if err != nil {
		t.Error(err)
	}
}

func TestWalletTagAccounts(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		}
	  }`
	server := setupServer(t, "tag_accounts", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	err := w.TagAccounts(&TagAccountsRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletUntagAccounts(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		}
	  }`
	server := setupServer(t, "untag_accounts", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	err := w.UntagAccounts(&UntagAccountsRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletSetAccountTagDescription(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		}
	  }`
	server := setupServer(t, "set_account_tag_description", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	err := w.SetAccountTagDescription(&SetAccountTagDescriptionRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletGetHeight(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "height": 145545
		}
	  }`
	server := setupServer(t, "get_height", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.GetHeight()
	if err != nil {
		t.Error(err)
	}
}

func TestWalletTransfer(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "amount": 300000000000,
		  "fee": 86897600000,
		  "multisig_txset": "",
		  "tx_blob": "",
		  "tx_hash": "7663438de4f72b25a0e395b770ea9ecf7108cd2f0c4b75be0b14a103d3362be9",
		  "tx_key": "25c9d8ec20045c80c93d665c9d3684aab7335f8b2cd02e1ba2638485afd1c70e236c4bdd7a2f1cb511dbf466f13421bdf8df988b7b969c448ca6239d7251490e4bf1bbf9f6ffacffdcdc93b9d1648ec499eada4d6b4e02ce92d4a1c0452e5d009fbbbf15b549df8856205a4c7bda6338d82c823f911acd00cb75850b198c5803",
		  "tx_metadata": "",
		  "unsigned_txset": ""
		}
	  }`
	server := setupServer(t, "transfer", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.Transfer(&TransferRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletTransferSplit(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "amount_list": [3000000000000],
		  "fee_list": [473710000],
		  "multisig_txset": "",
		  "tx_hash_list": ["4adcdc1af3f665770cdf8fb7a380887cd07ac53c2b771bd18df5ca375d5e7540"],
		  "tx_key_list": ["5b455c0f97168be652a2c03c5c68a064bb84cdae4ddef01b5c48d73a0bbb27075fb714f2ca19ea6c8ff592417e606addea6deb1d6530e2969f75681ffcbfc4075677b94a8c9197963ae38fa6f543ee68f0a4c4bbda4c453f39538f00b28e980ea08509730b51c004960101ba2f3adbc34cbbdff0d5af9dba061b523090debd06"],
		  "unsigned_txset": ""
		}
	  }`
	server := setupServer(t, "transfer_split", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.TransferSplit(&TransferSplitRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletSignTransfer(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "amount": 1000000000000,
		  "fee": 15202740000,
		  "multisig_txset": "",
		  "tx_blob": "...long_hex...",
		  "tx_hash": "c648ba0a049e5ce4ec21361dbf6e4b21eac0f828eea9090215de86c76b31d0a4",
		  "tx_key": "",
		  "tx_metadata": "",
		  "unsigned_txset": "...long_hex..."
		}
	  }`
	server := setupServer(t, "sign_transfer", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.SignTransfer(&SignTransferRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletSubmitTransfer(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "tx_hash_list": ["40fad7c828bb383ac02648732f7afce9adc520ba5629e1f5d9c03f584ac53d74"]
		}
	  }`
	server := setupServer(t, "submit_transfer", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.SubmitTransfer(&SubmitTransferRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletSweepDust(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "multisig_txset": "",
		  "unsigned_txset": ""
		}
	  }`
	server := setupServer(t, "sweep_dust", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.SweepDust(&SweepDustRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletSweepAll(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "amount_list": [9985885770000],
		  "fee_list": [14114230000],
		  "multisig_txset": "",
		  "tx_hash_list": ["ab4b6b65cc8cd8c9dd317d0b90d97582d68d0aa1637b0065b05b61f9a66ea5c5"],
		  "tx_key_list": ["b9b4b39d3bb3062ddb85ec0266d4df39058f4c86077d99309f218ce4d76af607"],
		  "unsigned_txset": ""
		}
	  }`
	server := setupServer(t, "sweep_all", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.SweepAll(&SweepAllRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletSweepSingle(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "amount": 27126892247503,
		  "fee": 14111630000,
		  "multisig_txset": "",
		  "tx_blob": "",
		  "tx_hash": "106d4391a031e5b735ded555862fec63233e34e5fa4fc7edcfdbe461c275ae5b",
		  "tx_key": "",
		  "tx_metadata": "",
		  "unsigned_txset": ""
		}
	  }`
	server := setupServer(t, "sweep_single", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.SweepSingle(&SweepSingleRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletRelayTx(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "tx_hash": "1c42dcc5672bb09bccf33fb1e9ab4a498af59a6dbd33b3d0cfb289b9e0e25fa5"
		}
	  }`
	server := setupServer(t, "relay_tx", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.RelayTx(&RelayTxRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletStore(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		}
	  }`
	server := setupServer(t, "store", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	err := w.Store()
	if err != nil {
		t.Error(err)
	}
}

func TestWalletGetPayments(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "payments": [{
			"address": "55LTR8KniP4LQGJSPtbYDacR7dz8RBFnsfAKMaMuwUNYX6aQbBcovzDPyrQF9KXF9tVU6Xk3K8no1BywnJX6GvZX8yJsXvt",
			"amount": 1000000000000,
			"block_height": 127606,
			"payment_id": "60900e5603bf96e3",
			"subaddr_index": {
			  "major": 0,
			  "minor": 0
			},
			"tx_hash": "3292e83ad28fc1cc7bc26dbd38862308f4588680fbf93eae3e803cddd1bd614f",
			"unlock_time": 0
		  }]
		}
	  }`
	server := setupServer(t, "get_payments", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.GetPayments(&GetPaymentsRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletGetBulkPayments(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "payments": [{
			"address": "55LTR8KniP4LQGJSPtbYDacR7dz8RBFnsfAKMaMuwUNYX6aQbBcovzDPyrQF9KXF9tVU6Xk3K8no1BywnJX6GvZX8yJsXvt",
			"amount": 1000000000000,
			"block_height": 127606,
			"payment_id": "60900e5603bf96e3",
			"subaddr_index": {
			  "major": 0,
			  "minor": 0
			},
			"tx_hash": "3292e83ad28fc1cc7bc26dbd38862308f4588680fbf93eae3e803cddd1bd614f",
			"unlock_time": 0
		  }]
		}
	  }`
	server := setupServer(t, "get_bulk_payments", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.GetBulkPayments(&GetBulkPaymentsRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletIncomingTransfers(t *testing.T) {
	output := `{
	"id": "0",
	"jsonrpc": "2.0",
	"result": {
		"transfers": [
			{
				"amount": 65523760000,
				"block_height": 1555542,
				"frozen": false,
				"global_index": 5352374,
				"key_image": "727bfeb8c28dab760c9c5f646072bd55477557b0c4a7bfd0bdadd751f8120d96",
				"pubkey": "b3a500acfa163763ccf6d2c10262294f8c287df3e5cac6309408adf94dc50335",
				"spent": true,
				"subaddr_index": {
					"major": 0,
					"minor": 1
				},
				"tx_hash": "4b540773ddf9e819f0df47708f3d3c9f7f62933150b90edc89103d36d42ca4b7",
				"unlocked": true
			},
			{
				"amount": 7989160000,
				"block_height": 1764640,
				"frozen": false,
				"global_index": 8765446,
				"key_image": "fd70f985df78f13c8b77836cea91322f20c6f4ea80cff8d3274826008e74cde0",
				"pubkey": "8e19b5426dbfe6501f607d68fb1ca1a717f87c950af538cdff96f7e97746822b",
				"spent": true,
				"subaddr_index": {
					"major": 0,
					"minor": 0
				},
				"tx_hash": "2aa6843cb5de53f2260bcd222cbf9b90c724b7250d6c0ef039ed1d5ad43fa829",
				"unlocked": true
			},
			{
				"amount": 100000,
				"block_height": 2363583,
				"frozen": false,
				"global_index": 32533306,
				"key_image": "0d99656358a8499a78e9baff5e3a5ca5aa19ebf8a9fd21e6edd5d54b3d45c667",
				"pubkey": "2feb0deabe3b6cf36b0a5d7addf843c9d585e0f95702e7462b0fac36a2ef7df4",
				"spent": false,
				"subaddr_index": {
					"major": 0,
					"minor": 1
				},
				"tx_hash": "e53e9cfb1fe15495e0d6edc94eceb7b9e9471a4eb3a31ddd3e020be12fad5ea1",
				"unlocked": true
			},
			{
				"amount": 10000000000,
				"block_height": 2663277,
				"frozen": false,
				"global_index": 56874042,
				"key_image": "4638a83f2d792204575930919ee36f6c4bd039e10bcabaa0fc74dd24e4aab7c9",
				"pubkey": "b328b6d97bc6de9ef17b7dfe16147c5a1e395d012f238cee7395c91666ee1822",
				"spent": true,
				"subaddr_index": {
					"major": 0,
					"minor": 3
				},
				"tx_hash": "a5fdd2da7f23574a620b0bcddd921b806a397186e61fbe736d5ef7522411aa5f",
				"unlocked": true
			}
		]
	}
}`
	server := setupServer(t, "incoming_transfers", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.IncomingTransfers(&IncomingTransfersRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletQueryKey(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "key": "vocal either anvil films dolphin zeal bacon cuisine quote syndrome rejoices envy okay pancakes tulips lair greater petals organs enmity dedicated oust thwart tomorrow tomorrow"
		}
	  }`
	server := setupServer(t, "query_key", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.QueryKey(&QueryKeyRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletMakeIntegratedAddress(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "integrated_address": "5F38Rw9HKeaLQGJSPtbYDacR7dz8RBFnsfAKMaMuwUNYX6aQbBcovzDPyrQF9KXF9tVU6Xk3K8no1BywnJX6GvZXCkbHUXdPHyiUeRyokn",
		  "payment_id": "420fa29b2d9a49f5"
		}
	  }`
	server := setupServer(t, "make_integrated_address", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.MakeIntegratedAddress(&MakeIntegratedAddressRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletSplitIntegratedAddress(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "is_subaddress": false,
		  "payment_id": "420fa29b2d9a49f5",
		  "standard_address": "55LTR8KniP4LQGJSPtbYDacR7dz8RBFnsfAKMaMuwUNYX6aQbBcovzDPyrQF9KXF9tVU6Xk3K8no1BywnJX6GvZX8yJsXvt"
		}
	  }`
	server := setupServer(t, "split_integrated_address", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.SplitIntegratedAddress(&SplitIntegratedAddressRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletStopWallet(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		}
	  }`
	server := setupServer(t, "stop_wallet", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	err := w.StopWallet()
	if err != nil {
		t.Error(err)
	}
}

func TestWalletRescanBlockchain(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		}
	  }`
	server := setupServer(t, "rescan_blockchain", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	err := w.RescanBlockchain()
	if err != nil {
		t.Error(err)
	}
}

func TestWalletSetTxNotes(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		}
	  }`
	server := setupServer(t, "set_tx_notes", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	err := w.SetTxNotes(&SetTxNotesRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletGetTxNotes(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "notes": ["This is an example"]
		}
	  }`
	server := setupServer(t, "get_tx_notes", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.GetTxNotes(&GetTxNotesRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletSetAttribute(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		}
	  }`
	server := setupServer(t, "set_attribute", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	err := w.SetAttribute(&SetAttributeRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletGetAttribute(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "value": "my_value"
		}
	  }`
	server := setupServer(t, "get_attribute", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.GetAttribute(&GetAttributeRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletGetTxKey(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "tx_key": "feba662cf8fb6d0d0da18fc9b70ab28e01cc76311278fdd7fe7ab16360762b06"
		}
	  }`
	server := setupServer(t, "get_tx_key", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.GetTxKey(&GetTxKeyRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletCheckTxKey(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "confirmations": 0,
		  "in_pool": false,
		  "received": 1000000000000
		}
	  }`
	server := setupServer(t, "check_tx_key", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.CheckTxKey(&CheckTxKeyRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletGetTxProof(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "signature": "InProofV13vqBCT6dpSAXkypZmSEMPGVnNRFDX2vscUYeVS4WnSVnV5BwLs31T9q6Etfj9Wts6tAxSAS4gkMeSYzzLS7Gt4vvCSQRh9niGJMUDJsB5hTzb2XJiCkUzWkkcjLFBBRVD5QZ"
		}
	  }`
	server := setupServer(t, "get_tx_proof", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.GetTxProof(&GetTxProofRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletCheckTxProof(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "confirmations": 482,
		  "good": true,
		  "in_pool": false,
		  "received": 1000000000000
		}
	  }`
	server := setupServer(t, "check_tx_proof", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.CheckTxProof(&CheckTxProofRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletGetSpendProof(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "signature": "SpendProofV1aSh8Todhk54736iXgV6vJAFP7egxByuMWZeyNDaN2JY737S95X5zz5mNMQSuCNSLjjhi5HJCsndpNWSNVsuThxwv285qy1KkUrLFRkxMSCjfL6bbycYN33ScZ5UB4Fzseceo1ndpL393T1q638VmcU3a56dhNHF1RPZFiGPS61FA78nXFSqE9uoKCCoHkEz83M1dQVhxZV5CEPF2P6VioGTKgprLCH9vvj9k1ivd4SX19L2VSMc3zD1u3mkR24ioETvxBoLeBSpxMoikyZ6inhuPm8yYo9YWyFtQK4XYfAV9mJ9knz5fUPXR8vvh7KJCAg4dqeJXTVb4mbMzYtsSZXHd6ouWoyCd6qMALdW8pKhgMCHcVYMWp9X9WHZuCo9rsRjRpg15sJUw7oJg1JoGiVgj8P4JeGDjnZHnmLVa5bpJhVCbMhyM7JLXNQJzFWTGC27TQBbthxCfQaKdusYnvZnKPDJWSeceYEFzepUnsWhQtyhbb73FzqgWC4eKEFKAZJqT2LuuSoxmihJ9acnFK7Ze23KTVYgDyMKY61VXADxmSrBvwUtxCaW4nQtnbMxiPMNnDMzeixqsFMBtN72j5UqhiLRY99k6SE7Qf5f29haNSBNSXCFFHChPKNTwJrehkofBdKUhh2VGPqZDNoefWUwfudeu83t85bmjv8Q3LrQSkFgFjRT5tLo8TMawNXoZCrQpyZrEvnodMDDUUNf3NL7rxyv3gM1KrTWjYaWXFU2RAsFee2Q2MTwUW7hR25cJvSFuB1BX2bfkoCbiMk923tHZGU2g7rSKF1GDDkXAc1EvFFD4iGbh1Q5t6hPRhBV8PEncdcCWGq5uAL5D4Bjr6VXG8uNeCy5oYWNgbZ5JRSfm7QEhPv8Fy9AKMgmCxDGMF9dVEaU6tw2BAnJavQdfrxChbDBeQXzCbCfep6oei6n2LZdE5Q84wp7eoQFE5Cwuo23tHkbJCaw2njFi3WGBbA7uGZaGHJPyB2rofTWBiSUXZnP2hiE9bjJghAcDm1M4LVLfWvhZmFEnyeru3VWMETnetz1BYLUC5MJGFXuhnHwWh7F6r74FDyhdswYop4eWPbyrXMXmUQEccTGd2NaT8g2VHADZ76gMC6BjWESvcnz2D4n8XwdmM7ZQ1jFwhuXrBfrb1dwRasyXxxHMGAC2onatNiExyeQ9G1W5LwqNLAh9hvcaNTGaYKYXoceVzLkgm6e5WMkLsCwuZXvB"
		}
	  }`
	server := setupServer(t, "get_spend_proof", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.GetSpendProof(&GetSpendProofRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletCheckSpendProof(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "good": true
		}
	  }`
	server := setupServer(t, "check_spend_proof", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.CheckSpendProof(&CheckSpendProofRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletGetReserveProof(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "signature": "ReserveProofV11BZ23sBt9sZJeGccf84mzyAmNCP3KzYbE1111112VKmH111118NfCYJQjZ6c46gT2kXgcHCaSSZeL8sRdzqjqx7i1e7FQfQGu2o113UYFVdwzHQi3iENDPa76Kn1BvywbKz3bMkXdZkBEEhBSF4kjjGaiMJ1ucKb6wvMVC4A8sA4nZEdL2Mk3wBucJCYTZwKqA8i1M113kqakDkG25FrjiDqdQTCYz2wDBmfKxF3eQiV5FWzZ6HmAyxnqTWUiMWukP9A3Edy3ZXqjP1b23dhz7Mbj39bBxe3ZeDNu9HnTSqYvHNRyqCkeUMJpHyQweqjGUJ1DSfFYr33J1E7MkhMnEi1o7trqWjVix32XLetYfePG73yvHbS24837L7Q64i5n1LSpd9yMiQZ3Dyaysi5y6jPx7TpAvnSqBFtuCciKoNzaXoA3dqt9cuVFZTXzdXKqdt3cXcVJMNxY8RvKPVQHhUur94Lpo1nSpxf7BN5a5rHrbZFqoZszsZmiWikYPkLX72XUdw6NWjLrTBxSy7KuPYH86c6udPEXLo2xgN6XHMBMBJzt8FqqK7EcpNUBkuHm2AtpGkf9CABY3oSjDQoRF5n4vNLd3qUaxNsG4XJ12L9gJ7GrK273BxkfEA8fDdxPrb1gpespbgEnCTuZHqj1A"
		}
	  }`
	server := setupServer(t, "get_reserve_proof", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.GetReserveProof(&GetReserveProofRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletCheckReserveProof(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "good": true,
		  "spent": 0,
		  "total": 100000000000
		}
	  }`
	server := setupServer(t, "check_reserve_proof", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.CheckReserveProof(&CheckReserveProofRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletGetTransfers(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "in": [{
			"address": "77Vx9cs1VPicFndSVgYUvTdLCJEZw9h81hXLMYsjBCXSJfUehLa9TDW3Ffh45SQa7xb6dUs18mpNxfUhQGqfwXPSMrvKhVp",
			"amount": 200000000000,
			"confirmations": 1,
			"double_spend_seen": false,
			"fee": 21650200000,
			"height": 153624,
			"note": "",
			"payment_id": "0000000000000000",
			"subaddr_index": {
			  "major": 1,
			  "minor": 0
			},
			"suggested_confirmations_threshold": 1,
			"timestamp": 1535918400,
			"txid": "c36258a276018c3a4bc1f195a7fb530f50cd63a4fa765fb7c6f7f49fc051762a",
			"type": "in",
			"unlock_time": 0
		  }]
		}
	  }`
	server := setupServer(t, "get_transfers", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.GetTransfers(&GetTransfersRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletGetTransferByTxid(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "transfer": {
			"address": "53zii2WaqQwZU4oUsCUcrHgaSv2CrUGCSFJLdQnkLPyH7ZLPYHjtoHhi14dqjF6jywNRknYLwbate2eGv8TuZcS7GuR7wMY",
			"amount": 100000000000,
			"amounts": [100000000000],
			"confirmations": 19,
			"double_spend_seen": false,
			"fee": 53840000,
			"height": 1140109,
			"locked": false,
			"note": "",
			"payment_id": "0000000000000000",
			"subaddr_index": {
			  "major": 0,
			  "minor": 0
			},
			"subaddr_indices": [{
			  "major": 0,
			  "minor": 0
			}],
			"suggested_confirmations_threshold": 1,
			"timestamp": 1658360753,
			"txid": "765f7124d01bd2eb2d4e7e59aa44a28c24339a41e4009f463955b087017b0ca3",
			"type": "in",
			"unlock_time": 0
		  },
		  "transfers": [{
			"address": "53zii2WaqQwZU4oUsCUcrHgaSv2CrUGCSFJLdQnkLPyH7ZLPYHjtoHhi14dqjF6jywNRknYLwbate2eGv8TuZcS7GuR7wMY",
			"amount": 100000000000,
			"amounts": [100000000000],
			"confirmations": 19,
			"double_spend_seen": false,
			"fee": 53840000,
			"height": 1140109,
			"locked": false,
			"note": "",
			"payment_id": "0000000000000000",
			"subaddr_index": {
			  "major": 0,
			  "minor": 0
			},
			"subaddr_indices": [{
			  "major": 0,
			  "minor": 0
			}],
			"suggested_confirmations_threshold": 1,
			"timestamp": 1658360753,
			"txid": "765f7124d01bd2eb2d4e7e59aa44a28c24339a41e4009f463955b087017b0ca3",
			"type": "in",
			"unlock_time": 0
		  }]
		}
	  }`
	server := setupServer(t, "get_transfer_by_txid", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	res, err := w.GetTransferByTxid(&GetTransferByTxidRequest{})
	if err != nil {
		t.Error(err)
	}
	is := is.New(t)
	is.Equal(res.Transfer.Address, "53zii2WaqQwZU4oUsCUcrHgaSv2CrUGCSFJLdQnkLPyH7ZLPYHjtoHhi14dqjF6jywNRknYLwbate2eGv8TuZcS7GuR7wMY")
	is.Equal(res.Transfer.Amount, uint64(100000000000))
	is.Equal(res.Transfer.Amounts, []uint64{100000000000})
	is.Equal(res.Transfer.Confirmations, uint64(19))
	is.Equal(res.Transfer.DoubleSpendSeen, false)
	is.Equal(res.Transfer.Fee, uint64(53840000))
	is.Equal(res.Transfer.Height, uint64(1140109))
	is.Equal(res.Transfer.Locked, false)
	is.Equal(res.Transfer.Note, "")
	is.Equal(res.Transfer.PaymentID, "0000000000000000")
	is.Equal(res.Transfer.SubaddrIndex.Major, uint32(0))
	is.Equal(res.Transfer.SubaddrIndex.Minor, uint32(0))
	is.Equal(res.Transfer.SuggestedConfirmationsThreshold, uint64(1))
	is.Equal(res.Transfer.Timestamp, uint64(1658360753))
	is.Equal(res.Transfer.TxID, "765f7124d01bd2eb2d4e7e59aa44a28c24339a41e4009f463955b087017b0ca3")
	is.Equal(res.Transfer.Type, "in")
	is.Equal(res.Transfer.UnlockTime, uint64(0))
	is.Equal(res.Transfers[0].Address, "53zii2WaqQwZU4oUsCUcrHgaSv2CrUGCSFJLdQnkLPyH7ZLPYHjtoHhi14dqjF6jywNRknYLwbate2eGv8TuZcS7GuR7wMY")
	is.Equal(res.Transfers[0].Amount, uint64(100000000000))
	is.Equal(res.Transfers[0].Amounts, []uint64{100000000000})
	is.Equal(res.Transfers[0].Confirmations, uint64(19))
	is.Equal(res.Transfers[0].DoubleSpendSeen, false)
	is.Equal(res.Transfers[0].Fee, uint64(53840000))
	is.Equal(res.Transfers[0].Height, uint64(1140109))
	is.Equal(res.Transfers[0].Locked, false)
	is.Equal(res.Transfers[0].Note, "")
	is.Equal(res.Transfers[0].PaymentID, "0000000000000000")
	is.Equal(res.Transfers[0].SubaddrIndex.Major, uint32(0))
	is.Equal(res.Transfers[0].SubaddrIndex.Minor, uint32(0))
	is.Equal(res.Transfers[0].SuggestedConfirmationsThreshold, uint64(1))
	is.Equal(res.Transfers[0].Timestamp, uint64(1658360753))
	is.Equal(res.Transfers[0].TxID, "765f7124d01bd2eb2d4e7e59aa44a28c24339a41e4009f463955b087017b0ca3")
	is.Equal(res.Transfers[0].Type, "in")
	is.Equal(res.Transfers[0].UnlockTime, uint64(0))

}

func TestWalletDescribeTransfer(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "desc": [{
			"amount_in": 886489038634812,
			"amount_out": 886455352051344,
			"change_address": "5BqWeZrK944YesCy5VdmBneWeaSZutEijFVAKjpVHeVd4unsCSM55CjgViQsK9WFNHK1eZgcCuZ3fRqYpzKDokqSUmQfJzvswQs13AAidJ",
			"change_amount": 4976287087263,
			"dummy_outputs": 0,
			"extra": "01b998f16459bcbac9c92074d3128d10724f10b74f5a7b1ec8e5a1e7f1150544020209010000000000000000",
			"fee": 33686583468, 
			"payment_id": "0000000000000000000000000000000000000000000000000000000000000000",
			"recipients": [{
			  "address": "0b057f69acc1552014cb157138e5c4dd495347d333f68ff0af70494b979aed10",
			  "amount": 881479064964081
			}],       
			"ring_size": 11,
			"unlock_time": 0
			}]
		}
	  }`
	server := setupServer(t, "describe_transfer", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.DescribeTransfer(&DescribeTransferRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletSign(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "signature": "SigV14K6G151gycjiGxjQ74tKX6A2LwwghvuHjcDeuRFQio5LS6Gb27BNxjYQY1dPuUvXkEbGQUkiHSVLPj4nJAHRrrw3"
		}
	  }`
	server := setupServer(t, "sign", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.Sign(&SignRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletVerify(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "good": true
		}
	  }`
	server := setupServer(t, "verify", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.Verify(&VerifyRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletExportOutputs(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "outputs_data_hex": "...outputs..."
		}
	  }`
	server := setupServer(t, "export_outputs", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.ExportOutputs(&ExportOutputsRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletImportOutputs(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "num_imported": 6400
		}
	  }`
	server := setupServer(t, "import_outputs", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.ImportOutputs(&ImportOutputsRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletExportKeyImages(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "signed_key_images": [{
			"key_image": "cd35239b72a35e26a57ed17400c0b66944a55de9d5bda0f21190fed17f8ea876",
			"signature": "c9d736869355da2538ab4af188279f84138c958edbae3c5caf388a63cd8e780b8c5a1aed850bd79657df659422c463608ea4e0c730ba9b662c906ae933816d00"
		  },{
			"key_image": "65158a8ee5a3b32009b85a307d85b375175870e560e08de313531c7dbbe6fc19",
			"signature": "c96e40d09dfc45cfc5ed0b76bfd7ca793469588bb0cf2b4d7b45ef23d40fd4036057b397828062e31700dc0c2da364f50cd142295a8405b9fe97418b4b745d0c"
		  }]
		}
	  }`
	server := setupServer(t, "export_key_images", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.ExportKeyImages(&ExportKeyImagesRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletImportKeyImages(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "height": 76428,
		  "spent": 62708953408711,
		  "unspent": 0
		}
	  }`
	server := setupServer(t, "import_key_images", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.ImportKeyImages(&ImportKeyImagesRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletMakeUri(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "uri": "monero:55LTR8KniP4LQGJSPtbYDacR7dz8RBFnsfAKMaMuwUNYX6aQbBcovzDPyrQF9KXF9tVU6Xk3K8no1BywnJX6GvZX8yJsXvt?tx_payment_id=420fa29b2d9a49f5&tx_amount=0.000000000010&recipient_name=el00ruobuob%20Stagenet%20wallet&tx_description=Testing%20out%20the%20make_uri%20function."
		}
	  }`
	server := setupServer(t, "make_uri", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.MakeURI(&MakeURIRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletParseUri(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "uri": {
			"address": "55LTR8KniP4LQGJSPtbYDacR7dz8RBFnsfAKMaMuwUNYX6aQbBcovzDPyrQF9KXF9tVU6Xk3K8no1BywnJX6GvZX8yJsXvt",
			"amount": 10,
			"payment_id": "420fa29b2d9a49f5",
			"recipient_name": "el00ruobuob Stagenet wallet",
			"tx_description": "Testing out the make_uri function."
		  }
		}
	  }`
	server := setupServer(t, "parse_uri", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.ParseURI(&ParseURIRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletGetAddressBook(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "entries": [{
			"address": "77Vx9cs1VPicFndSVgYUvTdLCJEZw9h81hXLMYsjBCXSJfUehLa9TDW3Ffh45SQa7xb6dUs18mpNxfUhQGqfwXPSMrvKhVp",
			"description": "Second account",
			"index": 0,
			"payment_id": "0000000000000000000000000000000000000000000000000000000000000000"
		  },{
			"address": "78P16M3XmFRGcWFCcsgt1WcTntA1jzcq31seQX1Eg92j8VQ99NPivmdKam4J5CKNAD7KuNWcq5xUPgoWczChzdba5WLwQ4j",
			"description": "Third account",
			"index": 1,
			"payment_id": "0000000000000000000000000000000000000000000000000000000000000000"
		  }]
		}
	  }`
	server := setupServer(t, "get_address_book", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.GetAddressBook(&GetAddressBookRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletAddAddressBook(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "index": 1
		}
	  }`
	server := setupServer(t, "add_address_book", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.AddAddressBook(&AddAddressBookRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletEditAddressBook(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		}
	  }`
	server := setupServer(t, "edit_address_book", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	err := w.EditAddressBook(&EditAddressBookRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletDeleteAddressBook(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		}
	  }`
	server := setupServer(t, "delete_address_book", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	err := w.DeleteAddressBook(&DeleteAddressBookRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletRefresh(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "blocks_fetched": 24,
		  "received_money": true
		}
	  }`
	server := setupServer(t, "refresh", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.Refresh(&RefreshRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletAutoRefresh(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		}
	  }`
	server := setupServer(t, "auto_refresh", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	err := w.AutoRefresh(&AutoRefreshRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletRescanSpent(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		}
	  }`
	server := setupServer(t, "rescan_spent", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	err := w.RescanSpent()
	if err != nil {
		t.Error(err)
	}
}

func TestWalletStartMining(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		}
	  }`
	server := setupServer(t, "start_mining", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	err := w.StartMining(&StartMiningRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletStopMining(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		}
	  }`
	server := setupServer(t, "stop_mining", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	err := w.StopMining()
	if err != nil {
		t.Error(err)
	}
}

func TestWalletGetLanguages(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "languages": ["Deutsch","English","Español","Français","Italiano","Nederlands","Português","русский язык","日本語","简体中文 (中国)","Esperanto","Lojban"]
		}
	  }`
	server := setupServer(t, "get_languages", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.GetLanguages()
	if err != nil {
		t.Error(err)
	}
}

func TestWalletCreateWallet(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		}
	  }`
	server := setupServer(t, "create_wallet", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	err := w.CreateWallet(&CreateWalletRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletGenerateFromKeys(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
			"address":"42gt8cXJSHAL4up8XoZh7fikVuswDU7itAoaCjSQyo6fFoeTQpAcAwrQ1cs8KvFynLFSBdabhmk7HEe3HS7UsAz4LYnVPYM",
		  "info":"Wallet has been generated successfully."   
		}
	  }`
	server := setupServer(t, "generate_from_keys", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.GenerateFromKeys(&GenerateFromKeysRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletOpenWallet(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		}
	  }`
	server := setupServer(t, "open_wallet", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	err := w.OpenWallet(&OpenWalletRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletRestoreDeterministicWallet(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "address": "9wB1Jc5fy5hjTkFBnv4UNY3WfhUswhx8M7uWjZrwRBzH2uatJcn8AqiKEHWuSNrnapApCzzTxP4iSiV3y3pqYcRbDHNboJK",
		  "info": "Wallet has been restored successfully.",
		  "seed": "awkward vogue odometer amply bagpipe kisses poker aspire slug eluded hydrogen selfish later toolbox enigma wolf tweezers eluded gnome soprano ladder broken jukebox lordship aspire",
		  "was_deprecated": false
		}
	  }`
	server := setupServer(t, "restore_deterministic_wallet", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.RestoreDeterministicWallet(&RestoreDeterministicWalletRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletCloseWallet(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		}
	  }`
	server := setupServer(t, "close_wallet", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	err := w.CloseWallet()
	if err != nil {
		t.Error(err)
	}
}

func TestWalletChangeWalletPassword(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		}
	  }`
	server := setupServer(t, "change_wallet_password", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	err := w.ChangeWalletPassword(&ChangeWalletPasswordRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletIsMultisig(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "multisig": false,
		  "ready": false,
		  "threshold": 0,
		  "total": 0
		}
	  }`
	server := setupServer(t, "is_multisig", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.IsMultisig()
	if err != nil {
		t.Error(err)
	}
}

func TestWalletPrepareMultisig(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "multisig_info": "MultisigV1BFdxQ653cQHB8wsj9WJQd2VdnjxK89g5M94dKPBNw22reJnyJYKrz6rJeXdjFwJ3Mz6n4qNQLd6eqUZKLiNzJFi3UPNVcTjtkG2aeSys9sYkvYYKMZ7chCxvoEXVgm74KKUcUu4V8xveCBFadFuZs8shnxBWHbcwFr5AziLr2mE7KHJT"
		}
	  }`
	server := setupServer(t, "prepare_multisig", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.PrepareMultisig()
	if err != nil {
		t.Error(err)
	}
}

func TestWalletMakeMultisig(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "address": "55SoZTKH7D39drxfgT62k8T4adVFjmDLUXnbzEKYf1MoYwnmTNKKaqGfxm4sqeKCHXQ5up7PVxrkoeRzXu83d8xYURouMod",
		  "multisig_info": ""
		}
	  }`
	server := setupServer(t, "make_multisig", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.MakeMultisig(&MakeMultisigRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletExportMultisigInfo(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "info": "4d6f6e65726f206d756c7469736967206578706f72740105cf6442b09b75f5eca9d846771fe1a879c9a97ab0553ffbcec64b1148eb7832b51e7898d7944c41cee000415c5a98f4f80dc0efdae379a98805bb6eacae743446f6f421cd03e129eb5b27d6e3b73eb6929201507c1ae706c1a9ecd26ac8601932415b0b6f49cbbfd712e47d01262c59980a8f9a8be776f2bf585f1477a6df63d6364614d941ecfdcb6e958a390eb9aa7c87f056673d73bc7c5f0ab1f74a682e902e48a3322c0413bb7f6fd67404f13fb8e313f70a0ce568c853206751a334ef490068d3c8ca0e"
		}
	  }`
	server := setupServer(t, "export_multisig_info", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.ExportMultisigInfo()
	if err != nil {
		t.Error(err)
	}
}

func TestWalletImportMultisigInfo(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "n_outputs": 35
		}
	}`
	server := setupServer(t, "import_multisig_info", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.ImportMultisigInfo(&ImportMultisigInfoRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletFinalizeMultisig(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "address": "5B9gZUTDuHTcGGuY3nL3t8K2tDnEHeRVHSBQgLZUTQxtFYVLnho5JJjWJyFp5YZgZRQ44RiviJi1sPHgLVMbckRsDqDx1gV"
		}
	  }`
	server := setupServer(t, "finalize_multisig", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.FinalizeMultisig(&FinalizeMultisigRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletSignMultisig(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "tx_data_hex": "...multisig_txset...",
		  "tx_hash_list": ["4996091b61c1be112c1097fd5e97d8ff8b28f0e5e62e1137a8c831bacf034f2d"]
		}
	  }`
	server := setupServer(t, "sign_multisig", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.SignMultisig(&SignMultisigRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletSubmitMultisig(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "tx_hash_list": ["4996091b61c1be112c1097fd5e97d8ff8b28f0e5e62e1137a8c831bacf034f2d"]
		}
	  }`
	server := setupServer(t, "submit_multisig", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.SubmitMultisig(&SubmitMultisigRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestWalletGetVersion(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "version": 65539
		}
	  }`
	server := setupServer(t, "get_version", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.GetVersion()
	if err != nil {
		t.Error(err)
	}
}
