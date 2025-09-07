package capture

import (
	"github.com/atareversei/quardian/internal/models"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// TODO: add logs for easier debugs of system

type TCPHandler struct{}

func (h *TCPHandler) Handle(pkt *models.Packet, raw gopacket.Packet) {
	ipLayer := raw.Layer(layers.LayerTypeIPv4)
	if ipLayer == nil {
		return
	}

	_, ok := ipLayer.(*layers.IPv4)
	if !ok {
		return
	}

	tcpLayer := raw.Layer(layers.LayerTypeTCP)
	if tcpLayer == nil {
		return
	}

	tcp, ok := tcpLayer.(*layers.TCP)
	if !ok {
		return
	}

	pkt.Transport = models.TransportInfo{
		SrcPort:  uint16(tcp.SrcPort),
		DstPort:  uint16(tcp.DstPort),
		Protocol: "tcp",
	}
}
