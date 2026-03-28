package types

type SessionSummary struct {
	DeviceName         string `json:"device_name"`
	DeviceDescription  string `json:"device_description"`
	ConnectedSinceDate string `json:"connected_since_date"`
	ExitIp             string `json:"exit_ip"`
}
