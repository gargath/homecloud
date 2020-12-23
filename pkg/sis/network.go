package sis

import "net"

func gatherNetwork() (*Network, error) {
	netinfo := &Network{
		Interfaces: []*Interface{},
	}

	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err //TODO: format better error
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		ifaceInfo := &Interface{
			Name:         iface.Name,
			MTU:          iface.MTU,
			HardwareAddr: iface.HardwareAddr.String(),
			Flags:        iface.Flags.String(),
			Addresses:    []*Address{},
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err // TODO: format better error
		}
		for _, a := range addrs {
			addrinfo := &Address{
				Network: a.Network(),
				Address: a.String(),
			}
			ifaceInfo.Addresses = append(ifaceInfo.Addresses, addrinfo)
		}

		netinfo.Interfaces = append(netinfo.Interfaces, ifaceInfo)
	}

	return netinfo, nil

}
