package monerorpc

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/MarinX/monerorpc/daemon"
	"github.com/MarinX/monerorpc/wallet"
	"github.com/gabstv/httpdigest"
	"github.com/gorilla/rpc/v2/json2"
)

const (
	// TestnetURI local testnet monerod instance
	TestnetURI = "http://127.0.0.1:28081/json_rpc"
	// StagnetURI local stagnet monerod instance
	StagnetURI = "http://127.0.0.1:38080/json_rpc"
	// ProdnetURI local production monerod instance
	ProdnetURI = "http://127.0.0.1:18080/json_rpc"
)

// MoneroRPC holds json rpc http client for various monero calls
type MoneroRPC struct {
	client *http.Client
	uri    string
	Wallet wallet.Wallet
	Daemon daemon.Daemon
}

// New creates a new MoneroRPC client
func New(endpoint string, httpClient *http.Client) *MoneroRPC {
	cli := http.DefaultClient
	if httpClient != nil {
		cli = httpClient
	}
	client := &MoneroRPC{
		client: cli,
		uri:    endpoint,
	}
	client.Wallet = wallet.New(client)
	client.Daemon = daemon.New(client)
	return client
}

// SetAuth sets digest username and password to be used with client
func (c *MoneroRPC) SetAuth(username, password string) *MoneroRPC {
	c.client.Transport = httpdigest.New(username, password)
	return c
}

// Do calls monero json rpc server, usefull if you are calling undocumented API
func (c *MoneroRPC) Do(method string, req interface{}, res interface{}) error {
	buff, err := json2.EncodeClientRequest(method, req)
	if err != nil {
		return fmt.Errorf("error creating encoded request %v", err)
	}

	httpReq, err := http.NewRequest("POST", c.uri, bytes.NewReader(buff))
	if err != nil {
		return fmt.Errorf("error creating http request %v", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.client.Do(httpReq)
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
