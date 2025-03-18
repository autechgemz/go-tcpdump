package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func convertSize(size int) string {
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
	)

	switch {
	case size >= GB:
		return fmt.Sprintf("%.2f GB", float64(size)/GB)
	case size >= MB:
		return fmt.Sprintf("%.2f MB", float64(size)/MB)
	case size >= KB:
		return fmt.Sprintf("%.2f KB", float64(size)/KB)
	default:
		return fmt.Sprintf("%d B", size)
	}
}

func main() {
	iface := flag.String("i", "", "Network interface to capture packets from")
	filter := flag.String("f", "", "BPF filter for packet capture")
	payload_debug := flag.Bool("d", false, "Enable payload debug output")
	flag.Parse()

	if *iface == "" {
		fmt.Println("Error: Network interface is required")
		flag.Usage()
		os.Exit(1)
	}

	if *filter == "" {
		fmt.Println("Error: BPF filter is required")
		flag.Usage()
		os.Exit(1)
	}

	handle, err := pcap.OpenLive(*iface, 65536, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	if err := handle.SetBPFFilter(*filter); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Starting packet capture on", *iface, "with filter:", *filter)

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		timestamp := packet.Metadata().Timestamp
		var srcIP, dstIP string

		if ipLayer := packet.Layer(layers.LayerTypeIPv4); ipLayer != nil {
			ip := ipLayer.(*layers.IPv4)
			srcIP, dstIP = ip.SrcIP.String(), ip.DstIP.String()
		}

		if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
			tcp := tcpLayer.(*layers.TCP)
			flags := []string{"---", "---", "---", "---", "---", "---"}
			if tcp.SYN {
				flags[0] = "SYN"
			}
			if tcp.ACK {
				flags[1] = "ACK"
			}
			if tcp.PSH {
				flags[2] = "PSH"
			}
			if tcp.FIN {
				flags[3] = "FIN"
			}
			if tcp.RST {
				flags[4] = "RST"
			}
			if tcp.URG {
				flags[5] = "URG"
			}

			packetSize := convertSize(packet.Metadata().Length)

			fmt.Printf("Time: %-30s, IPv4: %-15s -> %-15s, TCP: %5d -> %5d, Size: %10s, Seq: %10d, Flags: [%s]\n",
				timestamp.Format("2006-01-02T15:04:05.000000Z07:00"),
				srcIP,
				dstIP,
				tcp.SrcPort,
				tcp.DstPort,
				packetSize,
				tcp.Seq,
				strings.Join(flags, " "),
			)

			if *payload_debug && len(tcp.Payload) > 0 {
				fmt.Println("TCP Payload:", string(tcp.Payload))
			}
		}
	}
}
