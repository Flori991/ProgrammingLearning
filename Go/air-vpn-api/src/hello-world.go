package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const API_URL = "https://airvpn.org/api/userinfo/?key="
const API_KEY = "bf8ab1d1386c4854d0f580fede6f36c448d6e9e808cf6aabc9eb2b48166bdf6f"

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

type Response struct {
	User       json.RawMessage `json:"user"`
	Sessions   []Session       `json:"sessions"`
	Connection json.RawMessage `json:"connection"`
}

func main() {
	body := sendGetRequest()
	response := parseResponse(body)

	if len(response.Sessions) > 0 {
		log.Printf("%+v\n", response.Sessions[0])
	}
}

func sendGetRequest() []byte {
	log.Println("Starting API call...")
	resp, err := http.Get(API_URL + API_KEY)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Println("Response status:", resp.Status)
	body, err := io.ReadAll(resp.Body)
	return body
}

func parseResponse(body []byte) Response {
	var resp Response
	if err := json.Unmarshal(body, &resp); err != nil {
		log.Fatal(err)
	}
	return resp
}
