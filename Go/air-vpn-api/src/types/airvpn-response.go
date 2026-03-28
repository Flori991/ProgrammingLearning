package types

import (
	"encoding/json"
)

type Session struct {
	DeviceName         string      `json:"device_name"`
	DeviceDescription  string      `json:"device_description"`
	VpnIp              string      `json:"vpn_ip"`
	VpnIpv4            string      `json:"vpn_ipv4"`
	VpnIpv6            string      `json:"vpn_ipv6"`
	ExitIp             string      `json:"exit_ip"`
	ExitIpv4           string      `json:"exit_ipv4"`
	ExitIpv6           string      `json:"exit_ipv6"`
	EntryIp            string      `json:"entry_ip"`
	EntryIpv4          string      `json:"entry_ipv4"`
	EntryIpv6          string      `json:"entry_ipv6"`
	ServerName         string      `json:"server_name"`
	ServerCountry      string      `json:"server_country"`
	ServerCountryCode  string      `json:"server_country_code"`
	ServerContinent    string      `json:"server_continent"`
	ServerLocation     string      `json:"server_location"`
	ServerBw           int         `json:"server_bw"`
	BytesRead          json.Number `json:"bytes_read"`
	BytesWrite         json.Number `json:"bytes_write"`
	ConnectedSinceDate string      `json:"connected_since_date"`
	ConnectedSinceUnix int         `json:"connected_since_unix"`
	SpeedRead          int         `json:"speed_read"`
	SpeedWrite         int         `json:"speed_write"`
}

type UserInfo struct {
	User       json.RawMessage `json:"user"`
	Sessions   []Session       `json:"sessions"`
	Connection json.RawMessage `json:"connection"`
}
