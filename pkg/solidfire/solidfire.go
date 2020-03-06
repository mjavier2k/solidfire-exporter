package solidfire

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	log "github.com/amoghe/distillog"
)

func NewSolidfireClient() *Client {
	log.Infof("initializing new solidfire client")

	insecure, err := strconv.ParseBool(os.Getenv("INSECURE_SKIP_VERIFY"))
	if err != nil {
		log.Warningln("Could not parse environment variable INSECURE_SKIP_VERIFY. Defaulting to INSECURE_SKIP_VERIFY=false")
		insecure = false
	}
	if insecure {
		log.Warningln("TLS certificate verification is currently disabled - This is not recommended.")
	}

	log.Infoln("RPC Server:", os.Getenv("SOLIDFIRE_RPC_ENDPOINT"))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
	}
	return &Client{
		HttpClient: &http.Client{
			Transport: tr,
			Timeout:   30 * time.Second,
		},
		Username:    os.Getenv("SOLIDFIRE_USER"),
		Password:    os.Getenv("SOLIDFIRE_PASS"),
		RPCEndpoint: os.Getenv("SOLIDFIRE_RPC_ENDPOINT"),
	}
}

func doRpcCall(c *Client, body []byte) ([]byte, error) {

	req, err := http.NewRequest("POST", c.RPCEndpoint, bytes.NewReader(body))
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error making RPC call: %v", err)
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
		Method: "ListVolumes",
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
		Method: "ListVolumeStats",
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
		Method: "GetClusterCapacity",
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

func (s *Client) ListClusterActiveFaults() (ListClusterFaultsResponse, error) {
	payload := &RPCBody{
		Method: "ListClusterActiveFaults",
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
		Method: "ListNodeStats",
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
		Method: "ListVolumeQoSHistograms",
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
		Method: "ListAllNodes",
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
		Method: "GetClusterStats",
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
