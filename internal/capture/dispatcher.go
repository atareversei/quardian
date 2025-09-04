package capture

import "github.com/google/gopacket"

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

func (d *Dispatcher) Dispatch(packet gopacket.Packet) {
	for _, layer := range packet.Layers() {
		h, ok := d.handlers[layer.LayerType()]
		if ok {
			h.Handle(packet)
		}
	}
}
