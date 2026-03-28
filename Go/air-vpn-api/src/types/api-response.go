package types

import "encoding/json"

type SessionSummary struct {
	DeviceName         string      `json:"device_name"`
	DeviceDescription  string      `json:"device_description"`
	ExitIpv4           string      `json:"exit_ipv4"`
	ServerName         string      `json:"server_name"`
	ServerCountry      string      `json:"server_country"`
	BandwidthUsed      int         `json:"bandwidth_used"`
	BandwidthMax       int         `json:"bandwidth_max"`
	BytesRead          json.Number `json:"bytes_read"`
	BytesWrite         json.Number `json:"bytes_write"`
	ConnectedSinceDate string      `json:"connected_since_date"`
	ConnectedSinceUnix int         `json:"connected_since_unix"`
}

type SessionSummaries struct {
	Sessions []SessionSummary `json:"sessions"`
}
