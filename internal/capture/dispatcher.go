package capture

import (
	"fmt"
	"github.com/atareversei/quardian/internal/models"
	"github.com/google/gopacket"
)

type Dispatcher struct {
	handlers map[gopacket.LayerType]PacketHandler
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		handlers: make(map[gopacket.LayerType]PacketHandler),
	}
}

func (d *Dispatcher) Register(layer gopacket.LayerType, handler PacketHandler) {
	d.handlers[layer] = handler
}

func (d *Dispatcher) Dispatch(raw gopacket.Packet) {
	pkt := new(models.Packet)
	for _, layer := range raw.Layers() {
		h, ok := d.handlers[layer.LayerType()]
		if ok {
			h.Handle(pkt, raw)
		}
	}
	fmt.Println(pkt)
}
