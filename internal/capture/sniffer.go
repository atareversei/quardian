package capture

import (
	"fmt"
	"log"
	"time"

	"github.com/atareversei/ids/pkg/cli"
	"github.com/google/gopacket"
	_ "github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

// Sniffer holds the configuration for a packet capture session.
type Sniffer struct {
	device      string
	snapshotLen int32
	promiscuous bool
	timeout     time.Duration
	handle      *pcap.Handle
}

func New(device string) *Sniffer {
	return &Sniffer{
		device:      device,
		snapshotLen: 1024,
		promiscuous: false,
		timeout:     30 * time.Second,
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
	cli.Info(fmt.Sprintf("Starting packet capture on device %s...\n", s.device))

	for packet := range packetSource.Packets() {
		// TODO: Print TCP packet info.
		fmt.Println(packet)
	}
}
