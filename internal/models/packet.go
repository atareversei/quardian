package models

import (
	"net"
	"time"
)

type Packet struct {
	Timestamp time.Time
	Length    int

	Ethernet    EthernetInfo
	Network     NetworkInfo
	Transport   TransportInfo
	Application ApplicationData
}

type EthernetInfo struct {
	SrcMAC  net.HardwareAddr
	DstMAC  net.HardwareAddr
	EthType string
}

type NetworkInfo struct {
	SrcIP    net.IP
	DstIP    net.IP
	Protocol string
}

type TransportInfo struct {
	SrcPort  uint16
	DstPort  uint16
	Protocol string
}

type ApplicationData struct {
	DNS *DNSInfo
}

type DNSInfo struct {
	Query string
	Type  string
	ID    uint16
}
