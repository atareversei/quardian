package capture

import "github.com/google/gopacket"

type PacketHandler interface {
	Handle(packet gopacket.Packet)
}
