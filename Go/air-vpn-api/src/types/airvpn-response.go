package types

import (
	"encoding/json"
)

type Session struct {
	DeviceName         string      `json:"device_name"`
	DeviceDescription  string      `json:"device_description"`
	ExitIpv4           string      `json:"exit_ipv4"`
	ServerName         string      `json:"server_name"`
	ServerCountry      string      `json:"server_country"`
	BytesRead          json.Number `json:"bytes_read"`
	BytesWrite         json.Number `json:"bytes_write"`
	ConnectedSinceDate string      `json:"connected_since_date"`
	ConnectedSinceUnix int         `json:"connected_since_unix"`
}

type UserInfo struct {
	Sessions []Session `json:"sessions"`
}

type Status struct {
	Servers []ServerStatus `json:"servers"`
	Result  string         `json:"result"`
}

type ServerStatus struct {
	ServerName    string `json:"public_name"`
	BandwidthUsed int    `json:"bw"`
	BandwidthMax  int    `json:"bw_max"`
}
