package model

type NetworkProtocol int

const (
	IPv4 NetworkProtocol = iota
	IPv6
	TCP
	UDP
)

type PacketFilter struct {
	AllowedProtocols []NetworkProtocol
}
