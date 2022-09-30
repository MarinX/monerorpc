package daemon

import (
	"bytes"
	"fmt"
	"io/ioutil"
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
		buff, _ := ioutil.ReadAll(req.Body)
		t.Log(string(buff))
		rw.Write([]byte(output))
	}))
	return server
}

func TestDaemonGetBlockCount(t *testing.T) {
	output := `{  
		"id": "0",  
		"jsonrpc": "2.0",  
		"result": {  
		  "count": 993163,  
		  "status": "OK"  
		}  
	  }  `
	server := setupServer(t, "get_block_count", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	res, err := w.GetBlockCount()
	if err != nil {
		t.Error(err)
	}
	is.New(t).Equal(res, &GetBlockCountResponse{
		Count: 993163,
	})
}

func TestDaemonOnGetBlockHash(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": "e22cf75f39ae720e8b71b3d120a5ac03f0db50bba6379e2850975b4859190bc6"
	  }`
	server := setupServer(t, "on_get_block_hash", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	res, err := w.OnGetBlockHash([]uint64{1})
	if err != nil {
		t.Error(err)
	}
	is.New(t).Equal(res, "e22cf75f39ae720e8b71b3d120a5ac03f0db50bba6379e2850975b4859190bc6")
}

func TestDaemonGetBlockTemplate(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "blockhashing_blob": "070786a498d705f8dc58791266179087907a2ff4cd883615216749b97d2f12173171c725a6f84a00000000fc751ea4a94c2f840751eaa36138eee66dda15ef554e7d6594395827994e31da10",
		  "blocktemplate_blob": "070786a498d705f8dc58791266179087907a2ff4cd883615216749b97d2f12173171c725a6f84a0000000002aeab5f01fff2aa5f01e0a9d0f2f08a01028fdb3d5b5a2c363d36ea17a4add99a23a3ec7935b4c3e1e0364fcc4295c7a2ef5f01f912b15f5d17c1539d4722f79d8856d8654c5af87f54cfb3a4ff7f6b512b2a08023c000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000f1755090c809421d69873c161e7969b8bf33cee3b451dd4859bfc244a705f0b4900498f804b6023e13fa023a0fb759e8b7c9a39506a21442bc47077beeedc6b78d34c4ebdae91bd96097ccc9a882bc5056568b0d2f1f06559368fea4acba8e745444e883e53156d5083c1fd260edf05292934c8b40c098b81fe4e261720bdd272b209e317247a1d2c55dc4718891af0d16273c5a610f36f382a3bf50f54808aaa6a508e51d4601dd0d8fbf8b3b1685066ce121666a1409e8ac7a4d673c1cc36d10b825f764af647441f53230518e4d2efbcf8791c6060912c76e90db4982a66d51bbd96290bbb34db8080b216c2940cec407260bf5e2c3a5ee280835f15298f0801e9d98c4d414792282fbc2c28c3e20bc0fcb1829b5c3ad8f8d20847be8fdb2a949fd96f0205fbd6d271c880c5d8c83e9813606cd803a44d377fdeae45bfa67112132af601e9b3b0613ba7dff2ec3d4b935c447b47bfe39f7b950981b2f4c66c0d853e2218f1f69229a9b608c3d98be09b6d4d640a9f6ff0e920dbacf7e58b59554c0b398b1ae4b1d497104b4e4e745d850eed7eddb8aa93437427bf442ae5beb22cbf10a8fa738ea38cfa5d86dfd30675d4be11a38016e36936fd5601e52643e8b8bc433702ea7ae6149309c95b898cc854850e73fe0b95c5b8879b7325ecd4",
		  "difficulty": 61043624293,
		  "expected_reward": 4771949057248,
		  "height": 1561970,
		  "prev_hash": "f8dc58791266179087907a2ff4cd883615216749b97d2f12173171c725a6f84a",
		  "reserved_offset": 129,
		  "status": "OK",
		  "untrusted": false
		}
	  }`
	server := setupServer(t, "get_block_template", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	res, err := w.GetBlockTemplate(&GetBlockTemplateRequest{})
	if err != nil {
		t.Error(err)
	}
	is.New(t).Equal(res, &GetBlockTemplateResponse{
		BlockhashingBlob:  "070786a498d705f8dc58791266179087907a2ff4cd883615216749b97d2f12173171c725a6f84a00000000fc751ea4a94c2f840751eaa36138eee66dda15ef554e7d6594395827994e31da10",
		BlocktemplateBlob: "070786a498d705f8dc58791266179087907a2ff4cd883615216749b97d2f12173171c725a6f84a0000000002aeab5f01fff2aa5f01e0a9d0f2f08a01028fdb3d5b5a2c363d36ea17a4add99a23a3ec7935b4c3e1e0364fcc4295c7a2ef5f01f912b15f5d17c1539d4722f79d8856d8654c5af87f54cfb3a4ff7f6b512b2a08023c000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000f1755090c809421d69873c161e7969b8bf33cee3b451dd4859bfc244a705f0b4900498f804b6023e13fa023a0fb759e8b7c9a39506a21442bc47077beeedc6b78d34c4ebdae91bd96097ccc9a882bc5056568b0d2f1f06559368fea4acba8e745444e883e53156d5083c1fd260edf05292934c8b40c098b81fe4e261720bdd272b209e317247a1d2c55dc4718891af0d16273c5a610f36f382a3bf50f54808aaa6a508e51d4601dd0d8fbf8b3b1685066ce121666a1409e8ac7a4d673c1cc36d10b825f764af647441f53230518e4d2efbcf8791c6060912c76e90db4982a66d51bbd96290bbb34db8080b216c2940cec407260bf5e2c3a5ee280835f15298f0801e9d98c4d414792282fbc2c28c3e20bc0fcb1829b5c3ad8f8d20847be8fdb2a949fd96f0205fbd6d271c880c5d8c83e9813606cd803a44d377fdeae45bfa67112132af601e9b3b0613ba7dff2ec3d4b935c447b47bfe39f7b950981b2f4c66c0d853e2218f1f69229a9b608c3d98be09b6d4d640a9f6ff0e920dbacf7e58b59554c0b398b1ae4b1d497104b4e4e745d850eed7eddb8aa93437427bf442ae5beb22cbf10a8fa738ea38cfa5d86dfd30675d4be11a38016e36936fd5601e52643e8b8bc433702ea7ae6149309c95b898cc854850e73fe0b95c5b8879b7325ecd4",
		Difficulty:        61043624293,
		ExpectedReward:    4771949057248,
		Height:            1561970,
		PrevHash:          "f8dc58791266179087907a2ff4cd883615216749b97d2f12173171c725a6f84a",
		ReservedOffset:    129,
		Untrusted:         false,
	})
}

func TestDaemonSubmitBlock(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"error": {
		  "code": -7,
		  "message": "Block not accepted"
		}
	  }`
	server := setupServer(t, "submit_block", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	_, err := w.SubmitBlock([]string{"123"})
	if err == nil {
		t.Error("Submit block should fail")
		return
	}
	if err.Error() != "Block not accepted" {
		t.Error("unknown error, err", err, "expected", "Block not accepted")
	}
}

func TestDaemonGetLastBlockHeader(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "block_header": {
			"block_size": 62774,
			"depth": 0,
			"difficulty": 60097900840,
			"hash": "3a289b8fa88b10e2163826c230b45d79f2be37d14fa3153ee58ff8a427782d14",
			"height": 1562023,
			"major_version": 7,
			"minor_version": 7,
			"nonce": 3789681204,
			"num_txes": 5,
			"orphan_status": false,
			"prev_hash": "743e5d0a26849efe27b96086f2c4ecc39a0bc744bf21473dad6710221aff6ac3",
			"reward": 4724029079703,
			"timestamp": 1525029411
		  },
		  "status": "OK",
		  "untrusted": false
		}
	  }`
	server := setupServer(t, "get_last_block_header", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	res, err := w.GetLastBlockHeader()
	if err != nil {
		t.Error(err)
	}

	is.New(t).Equal(res, &GetLastBlockHeaderResponse{
		BlockHeader: BlockHeader{
			BlockSize:    62774,
			Depth:        0,
			Difficulty:   60097900840,
			Hash:         "3a289b8fa88b10e2163826c230b45d79f2be37d14fa3153ee58ff8a427782d14",
			Height:       1562023,
			MajorVersion: 7,
			MinorVersion: 7,
			Nonce:        3789681204,
			NumTxes:      5,
			OrphanStatus: false,
			PrevHash:     "743e5d0a26849efe27b96086f2c4ecc39a0bc744bf21473dad6710221aff6ac3",
			Reward:       4724029079703,
			Timestamp:    1525029411,
		},
		Untrusted: false,
	})
}

func TestDaemonGetBlockHeaderByHash(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "block_header": {
			"block_size": 210,
			"depth": 649717,
			"difficulty": 815625611,
			"hash": "e22cf75f39ae720e8b71b3d120a5ac03f0db50bba6379e2850975b4859190bc6",
			"height": 912345,
			"major_version": 1,
			"minor_version": 2,
			"nonce": 1646,
			"num_txes": 0,
			"orphan_status": false,
			"prev_hash": "b61c58b2e0be53fad5ef9d9731a55e8a81d972b8d90ed07c04fd37ca6403ff78",
			"reward": 7388968946286,
			"timestamp": 1452793716
		  },
		  "status": "OK",
		  "untrusted": false
		}
	  }`
	server := setupServer(t, "get_block_header_by_hash", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	res, err := w.GetBlockHeaderByHash(&GetBlockHeaderByHashRequest{})
	if err != nil {
		t.Error(err)
	}

	is.New(t).Equal(res, &GetBlockHeaderByHashResponse{
		BlockHeader: BlockHeader{
			BlockSize:    210,
			Depth:        649717,
			Difficulty:   815625611,
			Hash:         "e22cf75f39ae720e8b71b3d120a5ac03f0db50bba6379e2850975b4859190bc6",
			Height:       912345,
			MajorVersion: 1,
			MinorVersion: 2,
			Nonce:        1646,
			NumTxes:      0,
			OrphanStatus: false,
			PrevHash:     "b61c58b2e0be53fad5ef9d9731a55e8a81d972b8d90ed07c04fd37ca6403ff78",
			Reward:       7388968946286,
			Timestamp:    1452793716,
		},
		Untrusted: false,
	})
}

func TestDaemonGetBlockHeaderByHeight(t *testing.T) {
	output := `{
	"id": "0",
	"jsonrpc": "2.0",
	"result": {
	  "block_header": {
		"block_size": 210,
		"depth": 649721,
		"difficulty": 815625611,
		"hash": "e22cf75f39ae720e8b71b3d120a5ac03f0db50bba6379e2850975b4859190bc6",
		"height": 912345,
		"major_version": 1,
		"minor_version": 2,
		"nonce": 1646,
		"num_txes": 0,
		"orphan_status": false,
		"prev_hash": "b61c58b2e0be53fad5ef9d9731a55e8a81d972b8d90ed07c04fd37ca6403ff78",
		"reward": 7388968946286,
		"timestamp": 1452793716
	  },
	  "status": "OK",
	  "untrusted": false
	}
  }`
	server := setupServer(t, "get_block_header_by_height", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	res, err := w.GetBlockHeaderByHeight(&GetBlockHeaderByHeightRequest{})
	if err != nil {
		t.Error(err)
	}
	is.New(t).Equal(res, &GetBlockHeaderByHeightResponse{
		BlockHeader: BlockHeader{
			BlockSize:    210,
			Depth:        649721,
			Difficulty:   815625611,
			Hash:         "e22cf75f39ae720e8b71b3d120a5ac03f0db50bba6379e2850975b4859190bc6",
			Height:       912345,
			MajorVersion: 1,
			MinorVersion: 2,
			Nonce:        1646,
			NumTxes:      0,
			OrphanStatus: false,
			PrevHash:     "b61c58b2e0be53fad5ef9d9731a55e8a81d972b8d90ed07c04fd37ca6403ff78",
			Reward:       7388968946286,
			Timestamp:    1452793716,
		},
		Untrusted: false,
	})
}

func TestDaemonGetBlockHeadersRange(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "headers": [{
			"block_size": 301413,
			"depth": 16085,
			"difficulty": 134636057921,
			"hash": "86d1d20a40cefcf3dd410ff6967e0491613b77bf73ea8f1bf2e335cf9cf7d57a",
			"height": 1545999,
			"major_version": 6,
			"minor_version": 6,
			"nonce": 3246403956,
			"num_txes": 20,
			"orphan_status": false,
			"prev_hash": "0ef6e948f77b8f8806621003f5de24b1bcbea150bc0e376835aea099674a5db5",
			"reward": 5025593029981,
			"timestamp": 1523002893
		  },{
			"block_size": 13322,
			"depth": 16084,
			"difficulty": 134716086238,
			"hash": "b408bf4cfcd7de13e7e370c84b8314c85b24f0ba4093ca1d6eeb30b35e34e91a",
			"height": 1546000,
			"major_version": 7,
			"minor_version": 7,
			"nonce": 3737164176,
			"num_txes": 1,
			"orphan_status": false,
			"prev_hash": "86d1d20a40cefcf3dd410ff6967e0491613b77bf73ea8f1bf2e335cf9cf7d57a",
			"reward": 4851952181070,
			"timestamp": 1523002931
		  }],
		  "status": "OK",
		  "untrusted": false
		}
	  }`
	server := setupServer(t, "get_block_headers_range", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	res, err := w.GetBlockHeadersRange(&GetBlockHeadersRangeRequest{})
	if err != nil {
		t.Error(err)
	}
	is.New(t).Equal(res, &GetBlockHeadersRangeResponse{
		Headers: []BlockHeader{
			{
				BlockSize:    301413,
				Depth:        16085,
				Difficulty:   134636057921,
				Hash:         "86d1d20a40cefcf3dd410ff6967e0491613b77bf73ea8f1bf2e335cf9cf7d57a",
				Height:       1545999,
				MajorVersion: 6,
				MinorVersion: 6,
				Nonce:        3246403956,
				NumTxes:      20,
				OrphanStatus: false,
				PrevHash:     "0ef6e948f77b8f8806621003f5de24b1bcbea150bc0e376835aea099674a5db5",
				Reward:       5025593029981,
				Timestamp:    1523002893,
			},
			{
				BlockSize:    13322,
				Depth:        16084,
				Difficulty:   134716086238,
				Hash:         "b408bf4cfcd7de13e7e370c84b8314c85b24f0ba4093ca1d6eeb30b35e34e91a",
				Height:       1546000,
				MajorVersion: 7,
				MinorVersion: 7,
				Nonce:        3737164176,
				NumTxes:      1,
				OrphanStatus: false,
				PrevHash:     "86d1d20a40cefcf3dd410ff6967e0491613b77bf73ea8f1bf2e335cf9cf7d57a",
				Reward:       4851952181070,
				Timestamp:    1523002931,
			},
		},
		Untrusted: false,
	})
}

func TestDaemonGetBlock(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "blob": "0102f4bedfb405b61c58b2e0be53fad5ef9d9731a55e8a81d972b8d90ed07c04fd37ca6403ff786e0600000195d83701ffd9d73704ee84ddb42102378b043c1724c92c69d923d266fe86477d3a5ddd21145062e148c64c5767700880c0fc82aa020273733cbd6e6218bda671596462a4b062f95cfe5e1dbb5b990dacb30e827d02f280f092cbdd080247a5dab669770da69a860acde21616a119818e1a489bb3c4b1b6b3c50547bc0c80e08d84ddcb01021f7e4762b8b755e3e3c72b8610cc87b9bc25d1f0a87c0c816ebb952e4f8aff3d2b01fd0a778957f4f3103a838afda488c3cdadf2697b3d34ad71234282b2fad9100e02080000000bdfc2c16c00",
		  "block_header": {
			"block_size": 210,
			"depth": 649772,
			"difficulty": 815625611,
			"hash": "e22cf75f39ae720e8b71b3d120a5ac03f0db50bba6379e2850975b4859190bc6",
			"height": 912345,
			"major_version": 1,
			"minor_version": 2,
			"nonce": 1646,
			"num_txes": 0,
			"orphan_status": false,
			"prev_hash": "b61c58b2e0be53fad5ef9d9731a55e8a81d972b8d90ed07c04fd37ca6403ff78",
			"reward": 7388968946286,
			"timestamp": 1452793716
		  },
		  "json": "{\n  \"major_version\": 1, \n  \"minor_version\": 2, \n  \"timestamp\": 1452793716, \n  \"prev_id\": \"b61c58b2e0be53fad5ef9d9731a55e8a81d972b8d90ed07c04fd37ca6403ff78\", \n  \"nonce\": 1646, \n  \"miner_tx\": {\n    \"version\": 1, \n    \"unlock_time\": 912405, \n    \"vin\": [ {\n        \"gen\": {\n          \"height\": 912345\n        }\n      }\n    ], \n    \"vout\": [ {\n        \"amount\": 8968946286, \n        \"target\": {\n          \"key\": \"378b043c1724c92c69d923d266fe86477d3a5ddd21145062e148c64c57677008\"\n        }\n      }, {\n        \"amount\": 80000000000, \n        \"target\": {\n          \"key\": \"73733cbd6e6218bda671596462a4b062f95cfe5e1dbb5b990dacb30e827d02f2\"\n        }\n      }, {\n        \"amount\": 300000000000, \n        \"target\": {\n          \"key\": \"47a5dab669770da69a860acde21616a119818e1a489bb3c4b1b6b3c50547bc0c\"\n        }\n      }, {\n        \"amount\": 7000000000000, \n        \"target\": {\n          \"key\": \"1f7e4762b8b755e3e3c72b8610cc87b9bc25d1f0a87c0c816ebb952e4f8aff3d\"\n        }\n      }\n    ], \n    \"extra\": [ 1, 253, 10, 119, 137, 87, 244, 243, 16, 58, 131, 138, 253, 164, 136, 195, 205, 173, 242, 105, 123, 61, 52, 173, 113, 35, 66, 130, 178, 250, 217, 16, 14, 2, 8, 0, 0, 0, 11, 223, 194, 193, 108\n    ], \n    \"signatures\": [ ]\n  }, \n  \"tx_hashes\": [ ]\n}",
		  "miner_tx_hash": "c7da3965f25c19b8eb7dd8db48dcd4e7c885e2491db77e289f0609bf8e08ec30",
		  "status": "OK",
		  "untrusted": false
		}
	  }`
	server := setupServer(t, "get_block", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	res, err := w.GetBlock(&GetBlockRequest{})
	if err != nil {
		t.Error(err)
	}
	is := is.New(t)

	is.Equal(res.Blob, "0102f4bedfb405b61c58b2e0be53fad5ef9d9731a55e8a81d972b8d90ed07c04fd37ca6403ff786e0600000195d83701ffd9d73704ee84ddb42102378b043c1724c92c69d923d266fe86477d3a5ddd21145062e148c64c5767700880c0fc82aa020273733cbd6e6218bda671596462a4b062f95cfe5e1dbb5b990dacb30e827d02f280f092cbdd080247a5dab669770da69a860acde21616a119818e1a489bb3c4b1b6b3c50547bc0c80e08d84ddcb01021f7e4762b8b755e3e3c72b8610cc87b9bc25d1f0a87c0c816ebb952e4f8aff3d2b01fd0a778957f4f3103a838afda488c3cdadf2697b3d34ad71234282b2fad9100e02080000000bdfc2c16c00")
	is.Equal(res.BlockHeader.BlockSize, uint64(210))
	is.Equal(res.BlockHeader.Depth, uint64(649772))
	is.Equal(res.BlockHeader.Difficulty, uint64(815625611))
	is.Equal(res.BlockHeader.Hash, "e22cf75f39ae720e8b71b3d120a5ac03f0db50bba6379e2850975b4859190bc6")
	is.Equal(res.BlockHeader.Height, uint64(912345))
	is.Equal(res.BlockHeader.MajorVersion, uint64(1))
	is.Equal(res.BlockHeader.MinorVersion, uint64(2))
	is.Equal(res.BlockHeader.Nonce, uint64(1646))
	is.Equal(res.BlockHeader.NumTxes, uint64(0))
	is.Equal(res.BlockHeader.OrphanStatus, false)
	is.Equal(res.BlockHeader.PrevHash, "b61c58b2e0be53fad5ef9d9731a55e8a81d972b8d90ed07c04fd37ca6403ff78")
	is.Equal(res.BlockHeader.Reward, uint64(7388968946286))
	is.Equal(res.BlockHeader.Timestamp, uint64(1452793716))

}

func TestDaemonGetConnections(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "connections": [{
			"address": "173.90.69.136:62950",
			"avg_download": 0,
			"avg_upload": 2,
			"connection_id": "083c301a3030329a487adb12ad981d2c",
			"current_download": 0,
			"current_upload": 2,
			"height": 1562127,
			"host": "173.90.69.136",
			"incoming": true,
			"ip": "173.90.69.136",
			"live_time": 8,
			"local_ip": false,
			"localhost": false,
			"peer_id": "c959fbfbed9e44fb",
			"port": "62950",
			"recv_count": 259,
			"recv_idle_time": 8,
			"send_count": 24342,
			"send_idle_time": 8,
			"state": "state_normal",
			"support_flags": 0
		  }],
		  "status": "OK"
		}
	  }`
	server := setupServer(t, "get_connections", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	res, err := w.GetConnections()
	if err != nil {
		t.Error(err)
	}

	is.New(t).Equal(res, &GetConnectionsResponse{
		Connections: []Connection{
			{
				Address:         "173.90.69.136:62950",
				AvgDownload:     0,
				AvgUpload:       2,
				ConnectionID:    "083c301a3030329a487adb12ad981d2c",
				CurrentDownload: 0,
				CurrentUpload:   2,
				Height:          1562127,
				Host:            "173.90.69.136",
				Incoming:        true,
				IP:              "173.90.69.136",
				LiveTime:        8,
				LocalIP:         false,
				Localhost:       false,
				PeerID:          "c959fbfbed9e44fb",
				Port:            "62950",
				RecvCount:       259,
				RecvIdleTime:    8,
				SendCount:       24342,
				SendIdleTime:    8,
				State:           "state_normal",
				SupportFlags:    0,
			},
		},
	})
}

func TestDaemonGetInfo(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "adjusted_time": 1612090533,
		  "alt_blocks_count": 2,
		  "block_size_limit": 600000,
		  "block_size_median": 300000,
		  "block_weight_limit": 600000,
		  "block_weight_median": 300000,
		  "bootstrap_daemon_address": "",
		  "busy_syncing": false,
		  "credits": 0,
		  "cumulative_difficulty": 86168732847545368,
		  "cumulative_difficulty_top64": 0,
		  "database_size": 34329849856,
		  "difficulty": 225889137349,
		  "difficulty_top64": 0,
		  "free_space": 10795802624,
		  "grey_peerlist_size": 4999,
		  "height": 2286472,
		  "height_without_bootstrap": 2286472,
		  "incoming_connections_count": 85,
		  "mainnet": true,
		  "nettype": "mainnet",
		  "offline": false,
		  "outgoing_connections_count": 16,
		  "rpc_connections_count": 1,
		  "stagenet": false,
		  "start_time": 1611915662,
		  "status": "OK",
		  "synchronized": true,
		  "target": 120,
		  "target_height": 2286464,
		  "testnet": false,
		  "top_block_hash": "b92720d8315b96e32020d04e14a0c54cc13e057d4a5beb4501be490d306fdd8f",
		  "top_hash": "",
		  "tx_count": 11239803,
		  "tx_pool_size": 21,
		  "untrusted": false,
		  "update_available": false,
		  "version": "0.17.1.9-release",
		  "was_bootstrap_ever_used": false,
		  "white_peerlist_size": 1000,
		  "wide_cumulative_difficulty": "0x1322201881f9c18",
		  "wide_difficulty": "0x34980ab2c5"
		}
	  }`
	server := setupServer(t, "get_info", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	res, err := w.GetInfo()
	if err != nil {
		t.Error(err)
	}
	is.New(t).Equal(res, &GetInfoResponse{
		AdjustedTime:              1612090533,
		AltBlocksCount:            2,
		BlockSizeLimit:            600000,
		BlockSizeMedian:           300000,
		BlockWeightLimit:          600000,
		BlockWeightMedian:         300000,
		BootstrapDaemonAddress:    "",
		BusySyncing:               false,
		Credits:                   0,
		CumulativeDifficulty:      86168732847545368,
		CumulativeDifficultyTop64: 0,
		DatabaseSize:              34329849856,
		Difficulty:                225889137349,
		DifficultyTop64:           0,
		FreeSpace:                 10795802624,
		GreyPeerlistSize:          4999,
		Height:                    2286472,
		HeightWithoutBootstrap:    2286472,
		IncomingConnectionsCount:  85,
		Mainnet:                   true,
		NetType:                   "mainnet",
		Offline:                   false,
		OutgoingConnectionsCount:  16,
		RPCConnectionsCount:       1,
		Stagenet:                  false,
		StartTime:                 1611915662,
		Synchronized:              true,
		Target:                    120,
		TargetHeight:              2286464,
		Testnet:                   false,
		TopBlockHash:              "b92720d8315b96e32020d04e14a0c54cc13e057d4a5beb4501be490d306fdd8f",
		TopHash:                   "",
		TxCount:                   11239803,
		TxPoolSize:                21,
		Untrusted:                 false,
		UpdateAvailable:           false,
		Version:                   "0.17.1.9-release",
		WasBootstrapEverUsed:      false,
		WhitePeerlistSize:         1000,
		WideCumulativeDifficulty:  "0x1322201881f9c18",
		WideDifficulty:            "0x34980ab2c5",
	})
}

func TestDaemonHardForkInfo(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "earliest_height": 1009827,
		  "enabled": false,
		  "state": 2,
		  "status": "OK",
		  "threshold": 0,
		  "version": 1,
		  "votes": 7277,
		  "voting": 2,
		  "window": 10080
		}
	  }`
	server := setupServer(t, "hard_fork_info", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	res, err := w.HardForkInfo()
	if err != nil {
		t.Error(err)
	}

	is.New(t).Equal(res, &HardForkInfoResponse{
		EarliestHeight: 1009827,
		Enabled:        false,
		State:          2,
		Threshold:      0,
		Version:        1,
		Votes:          7277,
		Voting:         2,
		Window:         10080,
	})
}

func TestDaemonSetBans(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "status": "OK"
		}
	  }`
	server := setupServer(t, "set_bans", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	err := w.SetBans(&SetBansRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestDaemonGetBans(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "bans": [{
			"host": "102.168.1.51",
			"ip": 855746662,
			"seconds": 22
		  },{
			"host": "192.168.1.50",
			"ip": 838969536,
			"seconds": 28
		  }],
		  "status": "OK"
		}
	  }`
	server := setupServer(t, "get_bans", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	res, err := w.GetBans()
	if err != nil {
		t.Error(err)
	}
	is.New(t).Equal(res, &GetBansResponse{
		Bans: []Ban{
			{
				Host:    "102.168.1.51",
				IP:      855746662,
				Seconds: 22,
			},
			{
				Host:    "192.168.1.50",
				IP:      838969536,
				Seconds: 28,
			},
		},
	})
}

func TestDaemonFlushTxpool(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "status": "OK"
		}
	  }`
	server := setupServer(t, "flush_txpool", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	err := w.FlushTxpool(&FlushTxpoolRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestDaemonGetOutputHistogram(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "histogram": [{
			"amount": 20000000000,
			"recent_instances": 0,
			"total_instances": 381458,
			"unlocked_instances": 0
		  }],
		  "status": "OK",
		  "untrusted": false
		}
	  }`
	server := setupServer(t, "get_output_histogram", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	res, err := w.GetOutputHistogram(&GetOutputHistogramRequest{})
	if err != nil {
		t.Error(err)
	}
	is.New(t).Equal(res, &GetOutputHistogramResponse{
		Histogram: []Histogram{
			{
				Amount:            20000000000,
				RecentInstances:   0,
				TotalInstances:    381458,
				UnlockedInstances: 0,
			},
		},
		Untrusted: false,
	})
}

func TestDaemonGetVersion(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "status": "OK",
		  "untrusted": false,
		  "version": 65555
		}
	  }`
	server := setupServer(t, "get_version", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	res, err := w.GetVersion()
	if err != nil {
		t.Error(err)
	}
	is.New(t).Equal(res, &GetVersionResponse{
		Version:   65555,
		Untrusted: false,
	})
}

func TestDaemonGetCoinbaseTxSum(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "emission_amount": 9387854817320,
		  "fee_amount": 83981380000,
		  "status": "OK"
		}
	  }`
	server := setupServer(t, "get_coinbase_tx_sum", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	res, err := w.GetCoinbaseTxSum(&GetCoinbaseTxSumRequest{})
	if err != nil {
		t.Error(err)
	}
	is.New(t).Equal(res, &GetCoinbaseTxSumResponse{
		EmissionAmount: 9387854817320,
		FeeAmount:      83981380000,
	})
}

func TestDaemonGetFeeEstimate(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "fee": 187610000,
		  "status": "OK",
		  "untrusted": false
		}
	  }`
	server := setupServer(t, "get_fee_estimate", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	res, err := w.GetFeeEstimate(&GetFeeEstimateRequest{})
	if err != nil {
		t.Error(err)
	}
	is.New(t).Equal(res, &GetFeeEstimateResponse{
		Fee:       187610000,
		Untrusted: false,
	})
}

func TestDaemonGetAlternateChains(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "chains": [{
			"block_hash": "697cf03c89a9b118f7bdf11b1b3a6a028d7b3617d2d0ed91322c5709acf75625",
			"difficulty": 14114729638300280,
			"height": 1562062,
			"length": 2
		  }],
		  "status": "OK"
		}
	  }`
	server := setupServer(t, "get_alternate_chains", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	res, err := w.GetAlternateChains()
	if err != nil {
		t.Error(err)
	}
	is.New(t).Equal(res, &GetAlternateChainsResponse{
		Chains: []Chain{
			{
				BlockHash:  "697cf03c89a9b118f7bdf11b1b3a6a028d7b3617d2d0ed91322c5709acf75625",
				Difficulty: 14114729638300280,
				Height:     1562062,
				Length:     2,
			},
		},
	})
}

func TestDaemonRelayTx(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "status": "OK"
		}
	  }`
	server := setupServer(t, "relay_tx", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	err := w.RelayTx(&RelayTxRequest{})
	if err != nil {
		t.Error(err)
	}
}

func TestDaemonSyncInfo(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "height": 1563543,
		  "peers": [{
			"info": {
			  "address": "70.109.53.128:60064",
			  "avg_download": 0,
			  "avg_upload": 5,
			  "connection_id": "204067223b9b3415c265dd25ad29ee48",
			  "current_download": 0,
			  "current_upload": 1,
			  "height": 1559975,
			  "host": "70.109.53.128",
			  "incoming": true,
			  "ip": "70.109.53.128",
			  "live_time": 38,
			  "local_ip": false,
			  "localhost": false,
			  "peer_id": "96b8545dbc7a8866",
			  "port": "60064",
			  "recv_count": 1580,
			  "recv_idle_time": 28,
			  "send_count": 203603,
			  "send_idle_time": 8,
			  "state": "state_normal",
			  "support_flags": 1
			}
		  }],
		  "status": "OK",
		  "target_height": 1564067
		}
	  }`
	server := setupServer(t, "sync_info", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	res, err := w.SyncInfo()
	if err != nil {
		t.Error(err)
	}

	is.New(t).Equal(res, &SyncInfoResponse{
		Height: 1563543,
		Peers: []Peer{
			{
				Info: Connection{
					Address:         "70.109.53.128:60064",
					AvgDownload:     0,
					AvgUpload:       5,
					ConnectionID:    "204067223b9b3415c265dd25ad29ee48",
					CurrentDownload: 0,
					CurrentUpload:   1,
					Height:          1559975,
					Host:            "70.109.53.128",
					Incoming:        true,
					IP:              "70.109.53.128",
					LiveTime:        38,
					LocalIP:         false,
					Localhost:       false,
					PeerID:          "96b8545dbc7a8866",
					Port:            "60064",
					RecvCount:       1580,
					RecvIdleTime:    28,
					SendCount:       203603,
					SendIdleTime:    8,
					State:           "state_normal",
					SupportFlags:    1,
				},
			},
		},
		TargetHeight: 1564067,
	})
}

func TestDaemonGetTxpoolBacklog(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "backlog": [],
		  "status": "OK",
		  "untrusted": false
		}
	  }`
	server := setupServer(t, "get_txpool_backlog", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	res, err := w.GetTxpoolBacklog()
	if err != nil {
		t.Error(err)
	}
	is.New(t).Equal(res, &GetTxpoolBacklogResponse{
		Backlog:   []TXBacklogEntry{},
		Untrusted: false,
	})
}

func TestDaemonGetOutputDistribution(t *testing.T) {
	output := `{
		"id": "0",
		"jsonrpc": "2.0",
		"result": {
		  "distributions": [{
			"amount": 2628780000,
			"base": 0,
			"distribution": [],
			"start_height": 1462078
		  }],
		  "status": "OK"
		}
	  }`
	server := setupServer(t, "get_output_distribution", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	res, err := w.GetOutputDistribution(&GetOutputDistributionRequest{})
	if err != nil {
		t.Error(err)
	}
	is.New(t).Equal(res, &GetOutputDistributionResponse{
		Distributions: []Distribution{
			{
				Amount:       2628780000,
				Base:         0,
				Distribution: []uint64{},
				StartHeight:  1462078,
			},
		},
	})
}

func TestDaemonGenerateBlocks(t *testing.T) {
	output := `{
	"id": "0",
	"jsonrpc": "2.0",
	"result": {
			"blocks": ["49b712db7760e3728586f8434ee8bc8d7b3d410dac6bb6e98bf5845c83b917e4"],
			"height": 9783,
			"status": "OK",
			"untrusted": false
		}
	}
	`
	server := setupServer(t, "generateblocks", output)
	defer server.Close()

	w := New(getClient(server.URL, server.Client()))

	res, err := w.GenerateBlocks(&GenerateBlocksRequest{})
	if err != nil {
		t.Error(err)
	}

	is.New(t).Equal(res, &GenerateBlocksResponse{
		Blocks:    []string{"49b712db7760e3728586f8434ee8bc8d7b3d410dac6bb6e98bf5845c83b917e4"},
		Height:    9783,
		Untrusted: false,
	})
}
