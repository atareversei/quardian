package capture

import (
	"github.com/atareversei/quardian/internal/models"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type IPHandler struct{}

func (h *IPHandler) Handle(pkt *models.Packet, raw gopacket.Packet) {
	ipLayer := raw.Layer(layers.LayerTypeIPv4)
	if ipLayer == nil {
		return
	}

	ip, ok := ipLayer.(*layers.IPv4)
	if !ok {
		return
	}

	pkt.Network = models.NetworkInfo{
		SrcIP:    ip.SrcIP,
		DstIP:    ip.DstIP,
		Protocol: "ip",
	}
}
