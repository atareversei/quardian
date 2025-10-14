package capture

import (
	"github.com/google/gopacket/pcap"
	"log"
	"net"
)

type Device struct {
	Name        string
	Description string
	Flags       uint32
	Addresses   []DeviceAddress
}

type DeviceAddress struct {
	IP        net.IP
	Netmask   net.IPMask
	Broadaddr net.IP
	P2P       net.IP
}

func GetAllDevices() ([]Device, error) {
	rawDevices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	devices := make([]Device, 0)
	for _, dev := range rawDevices {
		device := Device{
			Name:        dev.Name,
			Description: dev.Description,
			Flags:       dev.Flags,
			Addresses:   make([]DeviceAddress, 0),
		}
		for _, addr := range dev.Addresses {
			address := DeviceAddress{
				IP:        addr.IP,
				Netmask:   addr.Netmask,
				Broadaddr: addr.Broadaddr,
				P2P:       addr.P2P,
			}
			device.Addresses = append(device.Addresses, address)
		}
		devices = append(devices, device)
	}

	return devices, nil
}
