package capture

import (
	"time"

	"github.com/Julia-Marcal/packet-scope/pkg/logger"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func StartCapture(devices []pcap.Interface, localIP string, allowedProtocols map[string]bool) {
	for _, device := range devices {
		go func(deviceName string) {
			handle, err := pcap.OpenLive(deviceName, 65536, true, pcap.BlockForever)

			if err != nil {
				logger.Error("Error opening device %s: %v", deviceName, err)
				return
			}

			defer handle.Close()

			packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

			for packet := range packetSource.Packets() {
				networkLayer := packet.NetworkLayer()

				if networkLayer == nil {
					continue
				}

				src, _ := networkLayer.NetworkFlow().Endpoints()

				if ignorePacket(src.String(), localIP, networkLayer) {
					continue
				}

				logPacketInfo(deviceName, packet, allowedProtocols)
			}
		}(device.Name)
	}

	select {}
}

func ignorePacket(src string, localIP string, networkLayer gopacket.NetworkLayer) bool {
	if networkLayer == nil {
		return true
	}

	if src != localIP {
		return true
	}

	return false
}

func logPacketInfo(deviceName string, packet gopacket.Packet, allowedProtocols map[string]bool) {
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	udpLayer := packet.Layer(layers.LayerTypeUDP)

	msg := "[%s] %s | Len=%d | "
	args := []interface{}{deviceName, time.Now().Format("15:04:05"), len(packet.Data())}

	if len(packet.Data()) > 1024 {
		logger.Warning("Large packet detected! %d bytes", len(packet.Data()))
	}

	if ipLayer != nil && (allowedProtocols["IPv4"] || allowedProtocols["IPv6"]) {
		ip := ipLayer.(*layers.IPv4)
		msg += "IP %s -> %s "
		args = append(args, ip.SrcIP, ip.DstIP)
	}

	if tcpLayer != nil && allowedProtocols["TCP"] {
		tcp := tcpLayer.(*layers.TCP)
		msg += "| TCP %d -> %d"
		args = append(args, tcp.SrcPort, tcp.DstPort)
	}

	if udpLayer != nil && allowedProtocols["UDP"] {
		udp := udpLayer.(*layers.UDP)
		msg += "| UDP %d -> %d"
		args = append(args, udp.SrcPort, udp.DstPort)
	}

	logger.Info(msg, args...)
}

func FindDevices() ([]pcap.Interface, error) {
	return pcap.FindAllDevs()
}
