package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/atareversei/ids/internal/capture"
	"github.com/google/gopacket/pcap"
)

func main() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	d := make(map[int]string)

	fmt.Println("Available devices:")
	for i, device := range devices {
		fmt.Printf("%d- %s: %s\n", i, device.Name, device.Description)
		d[i] = device.Name
		for _, addr := range device.Addresses {
			fmt.Printf("\t- %s\n", addr)
		}
	}

	fmt.Println("Choose a device (either the number or the full device name): ")
	reader := bufio.NewReader(os.Stdin)
	dev, err := reader.ReadString('\n')
	dev, _ = strings.CutSuffix(dev, "\r\n")
	dev, _ = strings.CutSuffix(dev, "\n")
	if err != nil {
		log.Fatal(err)
	}
	opt, err := strconv.Atoi(dev)
	if err == nil {
		device, ok := d[opt]
		if ok {
			dev = device
		}
	}

	fmt.Printf("trying to open the device: %q\n", dev)
	c := capture.New(dev)
	c.Start()
}
