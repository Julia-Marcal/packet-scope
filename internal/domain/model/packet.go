package model

import (
	"time"

	"github.com/google/gopacket"
)

type PacketInfo struct {
	DeviceName    string
	Timestamp     time.Time
	Length        int
	SourceIP      string
	DestinationIP string
	Protocol      string
	SourcePort    uint16
	DestPort      uint16
	RawPacket     gopacket.Packet
}

type NetworkInterface struct {
	Name        string
	Description string
	IsActive    bool
}

type CaptureConfig struct {
	Interfaces    []NetworkInterface
	LocalIP       string
	SnapshotLen   int32
	Promiscuous   bool
	Timeout       time.Duration
	FilterPackets bool
}
