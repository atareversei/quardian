package capture

import (
	"github.com/atareversei/quardian/internal/models"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type EthernetHandler struct{}

func (h *EthernetHandler) Handle(pkt *models.Packet, raw gopacket.Packet) {
	ethLayer := raw.Layer(layers.LayerTypeEthernet)
	if ethLayer == nil {
		return
	}

	eth, ok := ethLayer.(*layers.Ethernet)
	if !ok {
		return
	}

	pkt.Ethernet = models.EthernetInfo{
		SrcMAC:  eth.SrcMAC,
		DstMAC:  eth.DstMAC,
		EthType: eth.EthernetType.String(),
	}

	md := raw.Metadata()

	pkt.Timestamp = md.Timestamp
	pkt.Length = md.Length
}
