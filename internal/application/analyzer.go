package application

import (
	"fmt"
	"log"

	"github.com/Julia-Marcal/packet-scope/internal/domain/filters"
	"github.com/Julia-Marcal/packet-scope/internal/domain/model"
	"github.com/Julia-Marcal/packet-scope/internal/infrastructure/capture"
	"github.com/Julia-Marcal/packet-scope/pkg/utils"
)

func StartAnalysis() error {
	devices, err := capture.FindDevices()
	if err != nil {
		return fmt.Errorf("failed to find network interfaces: %w", err)
	}

	if len(devices) == 0 {
		return fmt.Errorf("no network interfaces found - try running with elevated privileges")
	}

	localIP, err := utils.GetLocalIP()
	if err != nil {
		return fmt.Errorf("failed to get local IP: %w", err)
	}

	log.Printf("Starting packet capture on %d interfaces with local IP: %s", len(devices), localIP)

	packetFilter := model.PacketFilter{
		AllowedProtocols: []model.NetworkProtocol{
			model.IPv4,
			model.IPv6,
			model.TCP,
			model.UDP,
		},
	}
	capture.StartCapture(devices, localIP, filters.FilterProtocolsType(packetFilter))
	return nil
}
