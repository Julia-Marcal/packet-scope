package capture

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func StartCapture(devices []pcap.Interface, localIP string) {
	for _, device := range devices {
		go func(deviceName string) {
			handle, err := pcap.OpenLive(deviceName, 65536, true, pcap.BlockForever)

			if err != nil {
				log.Printf("Error opening device %s: %v", deviceName, err)
				return
			}

			defer handle.Close()

			packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

			for packet := range packetSource.Packets() {
				networkLayer := packet.NetworkLayer()
				src, _ := networkLayer.NetworkFlow().Endpoints()

				if ignorePacket(src.String(), localIP, networkLayer) {
					continue
				}

				logPacketInfo(deviceName, packet)
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

func logPacketInfo(deviceName string, packet gopacket.Packet) {
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	udpLayer := packet.Layer(layers.LayerTypeUDP)

	fmt.Printf("[%s] %s | Len=%d | ", deviceName, time.Now().Format("15:04:05"), len(packet.Data()))

	if ipLayer != nil {
		ip := ipLayer.(*layers.IPv4)
		fmt.Printf("IP %s -> %s ", ip.SrcIP, ip.DstIP)
	}

	if tcpLayer != nil {
		tcp := tcpLayer.(*layers.TCP)
		fmt.Printf("| TCP %d -> %d\n", tcp.SrcPort, tcp.DstPort)
	} else if udpLayer != nil {
		udp := udpLayer.(*layers.UDP)
		fmt.Printf("| UDP %d -> %d\n", udp.SrcPort, udp.DstPort)
	} else {
		fmt.Printf("| Other protocol\n")
	}
}
