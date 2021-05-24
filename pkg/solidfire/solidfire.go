package solidfire

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/apex/log"
	"github.com/spf13/viper"
)

type RPC string

const (
	RPCGetClusterCapacity      RPC = "GetClusterCapacity"
	RPCGetClusterFullThreshold RPC = "GetClusterFullThreshold"
	RPCGetClusterStats         RPC = "GetClusterStats"
	RPCListAllNodes            RPC = "ListAllNodes"
	RPCListClusterFaults       RPC = "ListClusterFaults"
	RPCListDrives              RPC = "ListDrives"
	RPCListISCSISessions       RPC = "ListISCSISessions"
	RPCListNodeStats           RPC = "ListNodeStats"
	RPCListVolumeQoSHistograms RPC = "ListVolumeQoSHistograms"
	RPCListVolumes             RPC = "ListVolumes"
	RPCListVolumeStats         RPC = "ListVolumeStats"
)

func NewSolidfireClient(logger *log.Entry) (*Client, error) {
	logger.Debug("initializing new solidfire client")

	insecure := viper.GetBool(InsecureSSL)
	if insecure {
		logger.Warn("TLS certificate verification is currently disabled - this is not recommended.")
	}
	rpcServer := viper.GetString(Endpoint)
	_, err := url.Parse(rpcServer)
	if err != nil {
		return nil, fmt.Errorf("error parsing RPC Server url: %s", err.Error())
	}
	logger.Debugf("RPC Server: %v", rpcServer)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
	}
	return &Client{
		HttpClient: &http.Client{
			Transport: tr,
			Timeout:   time.Duration(viper.GetInt64(HTTPClientTimeout)) * time.Second,
		},
		Username:    viper.GetString(Username),
		Password:    viper.GetString(Password),
		RPCEndpoint: rpcServer,
	}, nil
}

func doRpcCall(c *Client, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", c.RPCEndpoint, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("Error building RPC request to %v: %v", c.RPCEndpoint, err)
	}
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error making RPC call %v: %v", string(body), err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Received invalid status code from RPC call: %v", resp.StatusCode)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %v", err)
	}

	return body, nil
}

func (s *Client) ListVolumes() (ListVolumesResponse, error) {
	payload := &RPCBody{
		Method: RPCListVolumes,
		Params: ListVolumesRPCParams{
			IncludeVirtualVolumes: true,
		},
		ID: 1,
	}
	payloadBytes, err := json.Marshal(&payload)
	r := ListVolumesResponse{}
	bodyBytes, err := doRpcCall(s, payloadBytes)
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(bodyBytes, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}

func (s *Client) ListVolumeStats() (ListVolumeStatsResponse, error) {
	payload := &RPCBody{
		Method: RPCListVolumeStats,
		Params: ListVolumeStatsRPCParams{
			VolumeIDs:             []int{}, // blank gives us all of them
			IncludeVirtualVolumes: true,
		},
		ID: 1,
	}
	payloadBytes, err := json.Marshal(&payload)
	r := ListVolumeStatsResponse{}
	bodyBytes, err := doRpcCall(s, payloadBytes)
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(bodyBytes, &r)

	if err != nil {
		return r, err
	}
	return r, nil
}

func (s *Client) GetClusterCapacity() (GetClusterCapacityResponse, error) {
	payload := &RPCBody{
		Method: RPCGetClusterCapacity,
		Params: GetClusterCapacityRPCParams{},
		ID:     1,
	}
	payloadBytes, err := json.Marshal(&payload)
	r := GetClusterCapacityResponse{}
	bodyBytes, err := doRpcCall(s, payloadBytes)
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(bodyBytes, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}

func (s *Client) ListClusterFaults() (ListClusterFaultsResponse, error) {
	payload := &RPCBody{
		Method: RPCListClusterFaults,
		Params: ListClusterFaultsRPCParams{
			FaultTypes:    "current",
			BestPractices: true,
		},
		ID: 1,
	}

	payloadBytes, err := json.Marshal(&payload)
	r := ListClusterFaultsResponse{}
	bodyBytes, err := doRpcCall(s, payloadBytes)
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(bodyBytes, &r)

	if err != nil {
		return r, err
	}
	return r, nil
}

func (s *Client) ListNodeStats() (ListNodeStatsResponse, error) {
	payload := &RPCBody{
		Method: RPCListNodeStats,
		Params: ListNodeStatsRPCParams{},
		ID:     1,
	}

	payloadBytes, err := json.Marshal(&payload)
	r := ListNodeStatsResponse{}
	bodyBytes, err := doRpcCall(s, payloadBytes)
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(bodyBytes, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}

func (s *Client) ListVolumeQoSHistograms() (ListVolumeQoSHistogramsResponse, error) {
	payload := &RPCBody{
		Method: RPCListVolumeQoSHistograms,
		Params: ListVolumeQoSHistogramsRPCParams{
			VolumeIDs: []int{}, // blank gives us all of them
		},
		ID: 1,
	}

	payloadBytes, err := json.Marshal(&payload)
	r := ListVolumeQoSHistogramsResponse{}
	bodyBytes, err := doRpcCall(s, payloadBytes)

	if err != nil {
		return r, err
	}
	err = json.Unmarshal(bodyBytes, &r)

	if err != nil {
		return r, err
	}
	return r, nil
}

func (s *Client) ListAllNodes() (ListAllNodesResponse, error) {
	payload := &RPCBody{
		Method: RPCListAllNodes,
		Params: ListAllNodesRPCParams{},
		ID:     1,
	}

	payloadBytes, err := json.Marshal(&payload)
	r := ListAllNodesResponse{}
	bodyBytes, err := doRpcCall(s, payloadBytes)

	if err != nil {
		return r, err
	}
	err = json.Unmarshal(bodyBytes, &r)

	if err != nil {
		return r, err
	}
	return r, nil
}

func (s *Client) GetClusterStats() (GetClusterStatsResponse, error) {
	payload := &RPCBody{
		Method: RPCGetClusterStats,
		Params: GetClusterStatsRPCParams{},
		ID:     1,
	}

	payloadBytes, err := json.Marshal(&payload)
	r := GetClusterStatsResponse{}
	bodyBytes, err := doRpcCall(s, payloadBytes)

	if err != nil {
		return r, err
	}
	err = json.Unmarshal(bodyBytes, &r)

	if err != nil {
		return r, err
	}
	return r, nil
}

func (s *Client) GetClusterFullThreshold() (GetClusterFullThresholdResponse, error) {
	payload := &RPCBody{
		Method: RPCGetClusterFullThreshold,
		Params: GetClusterFullThresholdParams{},
		ID:     1,
	}

	payloadBytes, err := json.Marshal(&payload)
	r := GetClusterFullThresholdResponse{}
	bodyBytes, err := doRpcCall(s, payloadBytes)

	if err != nil {
		return r, err
	}
	err = json.Unmarshal(bodyBytes, &r)

	if err != nil {
		return r, err
	}
	return r, nil
}

func (s *Client) ListDrives() (ListDrivesResponse, error) {
	payload := &RPCBody{
		Method: RPCListDrives,
		Params: ListDrivesParams{},
		ID:     1,
	}

	payloadBytes, err := json.Marshal(&payload)
	r := ListDrivesResponse{}
	bodyBytes, err := doRpcCall(s, payloadBytes)

	if err != nil {
		return r, err
	}
	err = json.Unmarshal(bodyBytes, &r)

	if err != nil {
		return r, err
	}
	return r, nil
}

func (s *Client) ListISCSISessions() (ListISCSISessionsResponse, error) {
	payload := &RPCBody{
		Method: RPCListISCSISessions,
		Params: ListISCSISessionsParams{},
		ID:     1,
	}

	payloadBytes, err := json.Marshal(&payload)
	r := ListISCSISessionsResponse{}
	bodyBytes, err := doRpcCall(s, payloadBytes)

	if err != nil {
		return r, err
	}
	err = json.Unmarshal(bodyBytes, &r)

	if err != nil {
		return r, err
	}
	return r, nil
}
