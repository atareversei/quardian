package flow

import "time"

type Flow struct {
	srcIP       string
	srcPort     uint16
	dstIP       string
	dstPort     uint16
	protocol    string
	packetCount int
	byteCount   int
	startTime   time.Time
	endTime     time.Time
}
