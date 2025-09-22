package filters

import (
	"github.com/Julia-Marcal/packet-scope/internal/domain/model"
)

func FilterProtocolsType(protocols []model.AllowedProtocols) map[string]bool {
	filtered := make(map[string]bool)
	for _, p := range protocols {
		for proto := range p.Protocols {
			filtered[proto.String()] = true
		}
	}
	return filtered
}
