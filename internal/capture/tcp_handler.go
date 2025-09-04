package capture

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// TODO: add logs for easier debugs of system

type TCPHandler struct{}

func (h *TCPHandler) Handle(packet gopacket.Packet) {
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer == nil {
		return
	}

	ip, ok := ipLayer.(*layers.IPv4)
	if !ok {
		return
	}

	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer == nil {
		return
	}

	tcp, ok := tcpLayer.(*layers.TCP)
	if !ok {
		return
	}

	fmt.Printf("TCP Packet: %s:%s -> %s:%s\n",
		ip.SrcIP, tcp.SrcPort, ip.DstIP, tcp.DstPort)
}
