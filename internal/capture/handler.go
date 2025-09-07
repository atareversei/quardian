package capture

import (
	"github.com/atareversei/quardian/internal/models"
	"github.com/google/gopacket"
)

type PacketHandler interface {
	Handle(pkt *models.Packet, raw gopacket.Packet)
}
