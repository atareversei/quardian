package capture

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type UDPHandler struct{}

func (h *UDPHandler) Handle(packet gopacket.Packet) {
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer == nil {
		return
	}

	ip, ok := ipLayer.(*layers.IPv4)
	if !ok {
		return
	}

	udpLayer := packet.Layer(layers.LayerTypeUDP)
	if udpLayer == nil {
		return
	}

	udp, ok := udpLayer.(*layers.UDP)
	if !ok {
		return
	}

	fmt.Printf("UDP Datagram: %s:%s -> %s:%s\n",
		ip.SrcIP, udp.SrcPort, ip.DstIP, udp.DstPort)
}
