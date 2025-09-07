package capture

import (
	"github.com/atareversei/quardian/internal/models"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type UDPHandler struct{}

func (h *UDPHandler) Handle(pkt *models.Packet, raw gopacket.Packet) {
	ipLayer := raw.Layer(layers.LayerTypeIPv4)
	if ipLayer == nil {
		return
	}

	_, ok := ipLayer.(*layers.IPv4)
	if !ok {
		return
	}

	udpLayer := raw.Layer(layers.LayerTypeUDP)
	if udpLayer == nil {
		return
	}

	udp, ok := udpLayer.(*layers.UDP)
	if !ok {
		return
	}

	pkt.Transport = models.TransportInfo{
		SrcPort:  uint16(udp.SrcPort),
		DstPort:  uint16(udp.DstPort),
		Protocol: "udp",
	}
}
