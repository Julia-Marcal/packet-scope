package application

import (
	"fmt"
	"log"
	"net"

	"github.com/Julia-Marcal/packet-scope/internal/domain/filters"
	"github.com/Julia-Marcal/packet-scope/internal/domain/model"
	"github.com/Julia-Marcal/packet-scope/internal/infrastructure/capture"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func StartAnalysis() error {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		return fmt.Errorf("failed to find network interfaces: %w", err)
	}

	if len(devices) == 0 {
		return fmt.Errorf("no network interfaces found - try running with elevated privileges")
	}

	localIP, err := getLocalIP()
	if err != nil {
		return fmt.Errorf("failed to get local IP: %w", err)
	}

	log.Printf("Starting packet capture on %d interfaces with local IP: %s", len(devices), localIP)

	allowed := []model.AllowedProtocols{
		{Protocols: map[gopacket.LayerType]bool{layers.LayerTypeIPv4: true}},
		{Protocols: map[gopacket.LayerType]bool{layers.LayerTypeIPv6: true}},
		{Protocols: map[gopacket.LayerType]bool{layers.LayerTypeTCP: true}},
		{Protocols: map[gopacket.LayerType]bool{layers.LayerTypeUDP: true}},
	}
	capture.StartCapture(devices, localIP, filters.FilterProtocolsType(allowed))
	return nil
}

func getLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}
