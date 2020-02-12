package solidfire

import (
	"net/http"
	"time"
)

type Client struct {
	Username    string
	Password    string
	RPCEndpoint string
	HttpClient  *http.Client
}
type RPCBody struct {
	Method string    `json:"method"`
	Params RPCParams `json:"params"`
	ID     int       `json:"id"`
}
type RPCParams interface{}
type ListVolumesRPCParams struct {
	IncludeVirtualVolumes bool `json:"includeVirtualVolumes"`
}
type ListVolumeStatsRPCParams struct {
	VolumeIDs             []int `json:"volumeIDs"`
	IncludeVirtualVolumes bool  `json:"includeVirtualVolumes"`
}
type GetClusterCapacityRPCParams struct {
	// No params needed
}

type ListClusterFaultsRPCParams struct {
	FaultTypes    string `json:"faultTypes"`
	BestPractices bool   `json:"bestPractices"`
}

type ListNodeStatsRPCParams struct {
	// No params needed
}

type ListVolumesResponse struct {
	ID     int `json:"id"`
	Result struct {
		Volumes []struct {
			Access     string `json:"access"`
			AccountID  int    `json:"accountID"`
			Attributes struct {
			} `json:"attributes"`
			BlockSize                   int         `json:"blockSize"`
			CreateTime                  time.Time   `json:"createTime"`
			CurrentProtectionScheme     string      `json:"currentProtectionScheme"`
			DeleteTime                  string      `json:"deleteTime"`
			Enable512E                  bool        `json:"enable512e"`
			EnableSnapMirrorReplication bool        `json:"enableSnapMirrorReplication"`
			Iqn                         string      `json:"iqn"`
			LastAccessTime              time.Time   `json:"lastAccessTime"`
			LastAccessTimeIO            time.Time   `json:"lastAccessTimeIO"`
			Name                        string      `json:"name"`
			PreviousProtectionScheme    interface{} `json:"previousProtectionScheme"`
			PurgeTime                   string      `json:"purgeTime"`
			Qos                         struct {
				BurstIOPS int `json:"burstIOPS"`
				BurstTime int `json:"burstTime"`
				Curve     struct {
					Num4096    int `json:"4096"`
					Num8192    int `json:"8192"`
					Num16384   int `json:"16384"`
					Num32768   int `json:"32768"`
					Num65536   int `json:"65536"`
					Num131072  int `json:"131072"`
					Num262144  int `json:"262144"`
					Num524288  int `json:"524288"`
					Num1048576 int `json:"1048576"`
				} `json:"curve"`
				MaxIOPS int `json:"maxIOPS"`
				MinIOPS int `json:"minIOPS"`
			} `json:"qos"`
			QosPolicyID                interface{}   `json:"qosPolicyID"`
			ScsiEUIDeviceID            string        `json:"scsiEUIDeviceID"`
			ScsiNAADeviceID            string        `json:"scsiNAADeviceID"`
			SliceCount                 int           `json:"sliceCount"`
			Status                     string        `json:"status"`
			TotalSize                  int64         `json:"totalSize"`
			VirtualVolumeID            interface{}   `json:"virtualVolumeID"`
			VolumeAccessGroups         []interface{} `json:"volumeAccessGroups"`
			VolumeConsistencyGroupUUID string        `json:"volumeConsistencyGroupUUID"`
			VolumeID                   int           `json:"volumeID"`
			VolumePairs                []interface{} `json:"volumePairs"`
			VolumeUUID                 string        `json:"volumeUUID"`
		} `json:"volumes"`
	} `json:"result"`
}

type ListVolumeStatsResponse struct {
	ID     int `json:"id"`
	Result struct {
		VolumeStats []struct {
			AccountID            int64       `json:"accountID"`
			ActualIOPS           float64     `json:"actualIOPS"`
			AsyncDelay           interface{} `json:"asyncDelay"`
			AverageIOPSize       float64     `json:"averageIOPSize"`
			BurstIOPSCredit      float64     `json:"burstIOPSCredit"`
			ClientQueueDepth     float64     `json:"clientQueueDepth"`
			DesiredMetadataHosts interface{} `json:"desiredMetadataHosts"`
			LatencyUSec          float64     `json:"latencyUSec"`
			// MetadataHosts        struct {
			// 	DeadSecondaries []interface{} `json:"deadSecondaries"`
			// 	LiveSecondaries []int         `json:"liveSecondaries"`
			// 	Primary         int           `json:"primary"`
			// } `json:"metadataHosts"`
			NonZeroBlocks         float64       `json:"nonZeroBlocks"`
			NormalizedIOPS        float64       `json:"normalizedIOPS"`
			ReadBytes             float64       `json:"readBytes"`
			ReadBytesLastSample   float64       `json:"readBytesLastSample"`
			ReadLatencyUSec       float64       `json:"readLatencyUSec"`
			ReadLatencyUSecTotal  float64       `json:"readLatencyUSecTotal"`
			ReadOps               float64       `json:"readOps"`
			ReadOpsLastSample     float64       `json:"readOpsLastSample"`
			SamplePeriodMSec      float64       `json:"samplePeriodMSec"`
			Throttle              float64       `json:"throttle"`
			Timestamp             time.Time     `json:"timestamp"`
			UnalignedReads        float64       `json:"unalignedReads"`
			UnalignedWrites       float64       `json:"unalignedWrites"`
			VolumeAccessGroups    []interface{} `json:"volumeAccessGroups"`
			VolumeID              int           `json:"volumeID"`
			VolumeSize            float64       `json:"volumeSize"`
			VolumeUtilization     float64       `json:"volumeUtilization"`
			WriteBytes            float64       `json:"writeBytes"`
			WriteBytesLastSample  float64       `json:"writeBytesLastSample"`
			WriteLatencyUSec      float64       `json:"writeLatencyUSec"`
			WriteLatencyUSecTotal float64       `json:"writeLatencyUSecTotal"`
			WriteOps              float64       `json:"writeOps"`
			WriteOpsLastSample    float64       `json:"writeOpsLastSample"`
			ZeroBlocks            float64       `json:"zeroBlocks"`
		} `json:"volumeStats"`
	} `json:"result"`
}
type GetClusterCapacityResponse struct {
	ID     int `json:"id"`
	Result struct {
		ClusterCapacity struct {
			ActiveBlockSpace             float64   `json:"activeBlockSpace"`
			ActiveSessions               float64   `json:"activeSessions"`
			AverageIOPS                  float64   `json:"averageIOPS"`
			ClusterRecentIOSize          float64   `json:"clusterRecentIOSize"`
			CurrentIOPS                  float64   `json:"currentIOPS"`
			MaxIOPS                      float64   `json:"maxIOPS"`
			MaxOverProvisionableSpace    float64   `json:"maxOverProvisionableSpace"`
			MaxProvisionedSpace          float64   `json:"maxProvisionedSpace"`
			MaxUsedMetadataSpace         float64   `json:"maxUsedMetadataSpace"`
			MaxUsedSpace                 float64   `json:"maxUsedSpace"`
			NonZeroBlocks                float64   `json:"nonZeroBlocks"`
			PeakActiveSessions           float64   `json:"peakActiveSessions"`
			PeakIOPS                     float64   `json:"peakIOPS"`
			ProvisionedSpace             float64   `json:"provisionedSpace"`
			SnapshotNonZeroBlocks        float64   `json:"snapshotNonZeroBlocks"`
			Timestamp                    time.Time `json:"timestamp"`
			TotalOps                     float64   `json:"totalOps"`
			UniqueBlocks                 float64   `json:"uniqueBlocks"`
			UniqueBlocksUsedSpace        float64   `json:"uniqueBlocksUsedSpace"`
			UsedMetadataSpace            float64   `json:"usedMetadataSpace"`
			UsedMetadataSpaceInSnapshots float64   `json:"usedMetadataSpaceInSnapshots"`
			UsedSpace                    float64   `json:"usedSpace"`
			ZeroBlocks                   float64   `json:"zeroBlocks"`
		} `json:"clusterCapacity"`
	} `json:"result"`
}
type ListClusterFaultsResponse struct {
	ID     int `json:"id"`
	Result struct {
		Faults []struct {
			BlocksUpgrade       bool          `json:"blocksUpgrade"`
			ClusterFaultID      float64       `json:"clusterFaultID"`
			Code                string        `json:"code"`
			Data                interface{}   `json:"data"`
			Date                interface{}   `json:"date"`
			Details             string        `json:"details"`
			DriveID             float64       `json:"driveID"`
			DriveIDs            []interface{} `json:"driveIDs"`
			ExternalSource      string        `json:"externalSource"`
			NetworkInterface    string        `json:"networkInterface"`
			NodeHardwareFaultID float64       `json:"nodeHardwareFaultID"`
			NodeID              float64       `json:"nodeID"`
			Resolved            bool          `json:"resolved"`
			ResolvedDate        interface{}   `json:"resolvedDate"`
			ServiceID           float64       `json:"serviceID"`
			Severity            string        `json:"severity"`
			Type                string        `json:"type"`
		} `json:"faults"`
	} `json:"result"`
}

type ListNodeStatsResponse struct {
	ID     int `json:"id"`
	Result struct {
		NodeStats struct {
			Nodes []struct {
				CBytesIn                  float64 `json:"cBytesIn"`
				CBytesOut                 float64 `json:"cBytesOut"`
				Count                     uint64  `json:"count"`
				CPU                       float64 `json:"cpu"`
				CPUTotal                  float64 `json:"cpuTotal"`
				MBytesIn                  float64 `json:"mBytesIn"`
				MBytesOut                 float64 `json:"mBytesOut"`
				NetworkUtilizationCluster float64 `json:"networkUtilizationCluster"`
				NetworkUtilizationStorage float64 `json:"networkUtilizationStorage"`
				NodeID                    int     `json:"nodeID"`
				ReadLatencyUSecTotal      float64 `json:"readLatencyUSecTotal"`
				ReadOps                   float64 `json:"readOps"`
				SBytesIn                  float64 `json:"sBytesIn"`
				SBytesOut                 float64 `json:"sBytesOut"`
				SsLoadHistogram           struct {
					Bucket0       uint64 `json:"Bucket0"`
					Bucket1To19   uint64 `json:"Bucket1To19"`
					Bucket20To39  uint64 `json:"Bucket20To39"`
					Bucket40To59  uint64 `json:"Bucket40To59"`
					Bucket60To79  uint64 `json:"Bucket60To79"`
					Bucket80To100 uint64 `json:"Bucket80To100"`
				} `json:"ssLoadHistogram"`
				Timestamp             time.Time `json:"timestamp"`
				UsedMemory            float64   `json:"usedMemory"`
				WriteLatencyUSecTotal float64   `json:"writeLatencyUSecTotal"`
				WriteOps              float64   `json:"writeOps"`
			} `json:"nodes"`
		} `json:"nodeStats"`
	} `json:"result"`
}
