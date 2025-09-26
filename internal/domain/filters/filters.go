package filters

import (
	"github.com/Julia-Marcal/packet-scope/internal/domain/model"
)

func FilterProtocolsType(packetFilter model.PacketFilter) map[string]bool {
	filtered := make(map[string]bool)
	for _, protocol := range packetFilter.AllowedProtocols {
		switch protocol {
		case model.IPv4:
			filtered["IPv4"] = true
		case model.IPv6:
			filtered["IPv6"] = true
		case model.TCP:
			filtered["TCP"] = true
		case model.UDP:
			filtered["UDP"] = true
		}
	}
	return filtered
}
