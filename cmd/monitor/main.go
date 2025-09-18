package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	devices, err := pcap.FindAllDevs()

	if err != nil {
		log.Fatal(err)
	}

	if len(devices) == 0 {
		log.Fatal("No interfaces found. Try running with sudo?")
	}

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)

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
				if networkLayer == nil {
					continue
				}

				src, _ := networkLayer.NetworkFlow().Endpoints()
				if src.String() != localAddr.IP.String() {
					continue
				}

				ipLayer := packet.Layer(layers.LayerTypeIPv4)
				tcpLayer := packet.Layer(layers.LayerTypeTCP)
				udpLayer := packet.Layer(layers.LayerTypeUDP)

				fmt.Printf("[%s] %s | Len=%d | ", deviceName, time.Now().Format("19:16:05"), len(packet.Data()))

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
		}(device.Name)
	}

	select {}

}
