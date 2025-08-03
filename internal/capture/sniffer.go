package capture

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
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
	fmt.Printf("Starting packet capture on device %s...\n", s.device)

	for packet := range packetSource.Packets() {
		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		if ipLayer == nil {
			continue
		}
		ip, _ := ipLayer.(*layers.IPv4)

		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		if tcpLayer == nil {
			continue
		}
		tcp, _ := tcpLayer.(*layers.TCP)

		fmt.Printf("TCP Packet: %s:%s -> %s:%s\n",
			ip.SrcIP, tcp.SrcPort, ip.DstIP, tcp.DstPort)
	}
}
