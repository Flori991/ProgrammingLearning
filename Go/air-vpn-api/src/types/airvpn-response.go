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

type Status struct {
	DeprecatedWarning string          `json:"deprecated_warning"`
	Servers           []ServerStatus  `json:"servers"`
	Routing           json.RawMessage `json:"routing"`
	Countries         json.RawMessage `json:"countries"`
	Continents        json.RawMessage `json:"continents"`
	Planets           json.RawMessage `json:"planets"`
	Result            string          `json:"result"`
}

type ServerStatus struct {
	ServerName    string `json:"public_name"`
	CountryName   string `json:"country_name"`
	CountryCode   string `json:"country_code"`
	Location      string `json:"location"`
	Continent     string `json:"continent"`
	BandwidthUsed int    `json:"bw"`
	BandwidthMax  int    `json:"bw_max"`
	Users         int    `json:"users"`
	CurrentLoad   int    `json:"load"`
	IPV4In1       string `json:"ip_v4_in1"`
	IPV4In2       string `json:"ip_v4_in2"`
	IPV4In3       string `json:"ip_v4_in3"`
	IPV4In4       string `json:"ip_v4_in4"`
	IPV6In1       string `json:"ip_v6_in1"`
	IPV6In2       string `json:"ip_v6_in2"`
	IPV6In3       string `json:"ip_v6_in3"`
	IPV6In4       string `json:"ip_v6_in4"`
	Health        string `json:"health"`
}
