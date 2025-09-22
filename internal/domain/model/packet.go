package model

import (
	"github.com/google/gopacket"
)

type AllowedProtocols struct {
	Protocols map[gopacket.LayerType]bool
}
