package capture

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// Sniffer holds the configuration for a packet capture session.
type Sniffer struct {
	device      string
	snapshotLen int32
	promiscuous bool
	timeout     time.Duration
	handle      *pcap.Handle
	dispatcher  *Dispatcher
}

func New(device string, dispatcher *Dispatcher) *Sniffer {
	return &Sniffer{
		device:      device,
		snapshotLen: 65535,
		promiscuous: true,
		timeout:     1 * time.Second,
		dispatcher:  dispatcher,
	}
}

func (s *Sniffer) Start() {
	var err error
	s.handle, err = pcap.OpenLive(s.device, s.snapshotLen, s.promiscuous, s.timeout)
	if err != nil {
		log.Fatal(err)
	}

	defer s.handle.Close()

	packetSource := gopacket.NewPacketSource(s.handle, s.handle.LinkType())
	fmt.Printf("Starting packet capture on device %s...\n", s.device)

	for packet := range packetSource.Packets() {
		s.dispatcher.Dispatch(packet)
	}
}
