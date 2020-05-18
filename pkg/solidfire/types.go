package solidfire

import (
	"net/http"
	"time"
)

var (
	ListenPortFlag     = "listenPort"
	UsernameFlag       = "username"
	PasswordFlag       = "password"
	EndpointFlag       = "endpoint"
	InsecureSSLFlag    = "insecure"
	ListenPortFlagEnv  = "SOLIDFIRE_PORT"
	UsernameFlagEnv    = "SOLIDFIRE_USER"
	PasswordFlagEnv    = "SOLIDFIRE_PASS"
	EndpointFlagEnv    = "SOLIDFIRE_RPC_ENDPOINT"
	InsecureSSLFlagEnv = "INSECURE_SKIP_VERIFY"
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

type ListVolumeQoSHistogramsRPCParams struct {
	VolumeIDs []int `json:"volumeIDs"`
}

type ListAllNodesRPCParams struct {
	// No params needed
}
type GetClusterStatsRPCParams struct {
	// No params needed
}

type GetClusterFullThresholdParams struct {
	// No params needed
}

type ListDrivesParams struct {
	// No params needed
}

type ListISCSISessionsParams struct {
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
			NodeID              int           `json:"nodeID"`
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

type ListVolumeQoSHistogramsResponse struct {
	ID     int `json:"id"`
	Result struct {
		QosHistograms []struct {
			Histograms struct {
				BelowMinIopsPercentages struct {
					Bucket1To19   uint64 `json:"Bucket1To19"`
					Bucket20To39  uint64 `json:"Bucket20To39"`
					Bucket40To59  uint64 `json:"Bucket40To59"`
					Bucket60To79  uint64 `json:"Bucket60To79"`
					Bucket80To100 uint64 `json:"Bucket80To100"`
				} `json:"belowMinIopsPercentages"`
				MinToMaxIopsPercentages struct {
					Bucket101Plus uint64 `json:"Bucket101Plus"`
					Bucket1To19   uint64 `json:"Bucket1To19"`
					Bucket20To39  uint64 `json:"Bucket20To39"`
					Bucket40To59  uint64 `json:"Bucket40To59"`
					Bucket60To79  uint64 `json:"Bucket60To79"`
					Bucket80To100 uint64 `json:"Bucket80To100"`
				} `json:"minToMaxIopsPercentages"`
				ReadBlockSizes struct {
					Bucket131072Plus    uint64 `json:"Bucket131072Plus"`
					Bucket16384To32767  uint64 `json:"Bucket16384To32767"`
					Bucket32768To65535  uint64 `json:"Bucket32768To65535"`
					Bucket4096To8191    uint64 `json:"Bucket4096To8191"`
					Bucket65536To131071 uint64 `json:"Bucket65536To131071"`
					Bucket8192To16383   uint64 `json:"Bucket8192To16383"`
				} `json:"readBlockSizes"`
				TargetUtilizationPercentages struct {
					Bucket0       uint64 `json:"Bucket0"`
					Bucket101Plus uint64 `json:"Bucket101Plus"`
					Bucket1To19   uint64 `json:"Bucket1To19"`
					Bucket20To39  uint64 `json:"Bucket20To39"`
					Bucket40To59  uint64 `json:"Bucket40To59"`
					Bucket60To79  uint64 `json:"Bucket60To79"`
					Bucket80To100 uint64 `json:"Bucket80To100"`
				} `json:"targetUtilizationPercentages"`
				ThrottlePercentages struct {
					Bucket0       uint64 `json:"Bucket0"`
					Bucket1To19   uint64 `json:"Bucket1To19"`
					Bucket20To39  uint64 `json:"Bucket20To39"`
					Bucket40To59  uint64 `json:"Bucket40To59"`
					Bucket60To79  uint64 `json:"Bucket60To79"`
					Bucket80To100 uint64 `json:"Bucket80To100"`
				} `json:"throttlePercentages"`
				WriteBlockSizes struct {
					Bucket131072Plus    uint64 `json:"Bucket131072Plus"`
					Bucket16384To32767  uint64 `json:"Bucket16384To32767"`
					Bucket32768To65535  uint64 `json:"Bucket32768To65535"`
					Bucket4096To8191    uint64 `json:"Bucket4096To8191"`
					Bucket65536To131071 uint64 `json:"Bucket65536To131071"`
					Bucket8192To16383   uint64 `json:"Bucket8192To16383"`
				} `json:"writeBlockSizes"`
			} `json:"histograms"`
			Timestamp time.Time `json:"timestamp"`
			VolumeID  int       `json:"volumeID"`
		} `json:"qosHistograms"`
	} `json:"result"`
}
type ListAllNodesResponse struct {
	ID     int `json:"id"`
	Result struct {
		Nodes []struct {
			AssociatedFServiceID      int `json:"associatedFServiceID"`
			AssociatedMasterServiceID int `json:"associatedMasterServiceID"`
			Attributes                struct {
			} `json:"attributes"`
			ChassisName                 string      `json:"chassisName"`
			Cip                         string      `json:"cip"`
			Cipi                        string      `json:"cipi"`
			FibreChannelTargetPortGroup interface{} `json:"fibreChannelTargetPortGroup"`
			Mip                         string      `json:"mip"`
			Mipi                        string      `json:"mipi"`
			Name                        string      `json:"name"`
			NodeID                      int         `json:"nodeID"`
			NodeSlot                    string      `json:"nodeSlot"`
			PlatformInfo                struct {
				ChassisType           string  `json:"chassisType"`
				CPUModel              string  `json:"cpuModel"`
				NodeMemoryGB          float64 `json:"nodeMemoryGB"`
				NodeType              string  `json:"nodeType"`
				PlatformConfigVersion string  `json:"platformConfigVersion"`
			} `json:"platformInfo"`
			Sip             string        `json:"sip"`
			Sipi            string        `json:"sipi"`
			SoftwareVersion string        `json:"softwareVersion"`
			UUID            string        `json:"uuid"`
			VirtualNetworks []interface{} `json:"virtualNetworks"`
		} `json:"nodes"`
		PendingActiveNodes []interface{} `json:"pendingActiveNodes"`
		PendingNodes       []interface{} `json:"pendingNodes"`
	} `json:"result"`
}
type GetClusterStatsResponse struct {
	ID     int `json:"id"`
	Result struct {
		ClusterStats struct {
			ActualIOPS           float64 `json:"actualIOPS"`
			AverageIOPSize       float64 `json:"averageIOPSize"`
			ClientQueueDepth     float64 `json:"clientQueueDepth"`
			ClusterUtilization   float64 `json:"clusterUtilization"`
			LatencyUSec          float64 `json:"latencyUSec"`
			NormalizedIOPS       float64 `json:"normalizedIOPS"`
			ReadBytes            float64 `json:"readBytes"`
			ReadBytesLastSample  float64 `json:"readBytesLastSample"`
			ReadLatencyUSec      float64 `json:"readLatencyUSec"`
			ReadLatencyUSecTotal float64 `json:"readLatencyUSecTotal"`
			ReadOps              float64 `json:"readOps"`
			ReadOpsLastSample    float64 `json:"readOpsLastSample"`
			SamplePeriodMsec     float64 `json:"samplePeriodMsec"`
			ServicesCount        float64 `json:"servicesCount"`
			ServicesTotal        float64 `json:"servicesTotal"`
			// Timestamp             time.Time `json:"timestamp"`
			UnalignedReads        float64 `json:"unalignedReads"`
			UnalignedWrites       float64 `json:"unalignedWrites"`
			WriteBytes            float64 `json:"writeBytes"`
			WriteBytesLastSample  float64 `json:"writeBytesLastSample"`
			WriteLatencyUSec      float64 `json:"writeLatencyUSec"`
			WriteLatencyUSecTotal float64 `json:"writeLatencyUSecTotal"`
			WriteOps              float64 `json:"writeOps"`
			WriteOpsLastSample    float64 `json:"writeOpsLastSample"`
		} `json:"clusterStats"`
	} `json:"result"`
}

type GetClusterFullThresholdResponse struct {
	ID     int `json:"id"`
	Result struct {
		BlockFullness                  string  `json:"blockFullness"`
		Fullness                       string  `json:"fullness"`
		MaxMetadataOverProvisionFactor float64 `json:"maxMetadataOverProvisionFactor"`
		MetadataFullness               string  `json:"metadataFullness"`
		SliceReserveUsedThresholdPct   float64 `json:"sliceReserveUsedThresholdPct"`
		Stage2AwareThreshold           float64 `json:"stage2AwareThreshold"`
		Stage2BlockThresholdBytes      float64 `json:"stage2BlockThresholdBytes"`
		Stage3BlockThresholdBytes      float64 `json:"stage3BlockThresholdBytes"`
		Stage3BlockThresholdPercent    float64 `json:"stage3BlockThresholdPercent"`
		Stage3LowThreshold             float64 `json:"stage3LowThreshold"`
		Stage4BlockThresholdBytes      float64 `json:"stage4BlockThresholdBytes"`
		Stage4CriticalThreshold        float64 `json:"stage4CriticalThreshold"`
		Stage5BlockThresholdBytes      float64 `json:"stage5BlockThresholdBytes"`
		SumTotalClusterBytes           float64 `json:"sumTotalClusterBytes"`
		SumTotalMetadataClusterBytes   float64 `json:"sumTotalMetadataClusterBytes"`
		SumUsedClusterBytes            float64 `json:"sumUsedClusterBytes"`
		SumUsedMetadataClusterBytes    float64 `json:"sumUsedMetadataClusterBytes"`
	} `json:"result"`
}

type ListDrivesResponse struct {
	ID     int `json:"id"`
	Result struct {
		Drives []struct {
			Attributes struct {
			} `json:"attributes"`
			Capacity float64 `json:"capacity"`
			DriveID  int     `json:"driveID"`
			NodeID   int     `json:"nodeID"`
			Serial   string  `json:"serial"`
			Slot     int     `json:"slot"`
			Status   string  `json:"status"`
			Type     string  `json:"type"`
		} `json:"drives"`
	} `json:"result"`
}

type ListISCSISessionsResponse struct {
	ID     int `json:"id"`
	Result struct {
		Sessions []struct {
			AccountID              int         `json:"accountID"`
			AccountName            string      `json:"accountName"`
			CreateTime             time.Time   `json:"createTime"`
			DriveID                int         `json:"driveID"`
			DriveIDs               []int       `json:"driveIDs"`
			Initiator              interface{} `json:"initiator"`
			InitiatorIP            string      `json:"initiatorIP"`
			InitiatorName          string      `json:"initiatorName"`
			InitiatorPortName      string      `json:"initiatorPortName"`
			InitiatorSessionID     float64     `json:"initiatorSessionID"`
			MsSinceLastIscsiPDU    int         `json:"msSinceLastIscsiPDU"`
			MsSinceLastScsiCommand int         `json:"msSinceLastScsiCommand"`
			NodeID                 int         `json:"nodeID"`
			ServiceID              int         `json:"serviceID"`
			SessionID              int64       `json:"sessionID"`
			TargetIP               string      `json:"targetIP"`
			TargetName             string      `json:"targetName"`
			TargetPortName         string      `json:"targetPortName"`
			VirtualNetworkID       int         `json:"virtualNetworkID"`
			VolumeID               int         `json:"volumeID"`
			VolumeInstance         int64       `json:"volumeInstance"`
		} `json:"sessions"`
	} `json:"result"`
}
