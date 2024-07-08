package ms

type SyncLevelData struct {
	Modoverrides string `json:"modoverrides"`
	ClusterToken string `json:"cluster_token"`

	MasterPort uint   `json:"master_port"`
	MasterIp   string `json:"master_ip"`
}

var SYNC_LEVEL = "sync_level"
var SYNC_START = "sync_start"

var SERVER_SYNC_STATUS = "server_sync_status"
var CLIENT_STATUS = "client_status"
