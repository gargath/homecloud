package sis

import (
	"fmt"
	"os"
)

// Collect gathers system information and returns it
func Collect(version string) (*SysInfo, error) {
	info := &SysInfo{}
	info.SisVersion = version
	hostname, err := os.Hostname()
	if err != nil {
		return nil, fmt.Errorf("Failed to identify hostname: %v", err)
	}
	info.Hostname = hostname

	mounts, err := parseMounts()
	if err != nil {
		return nil, fmt.Errorf("Failed to parse mounts: %v", err)
	}
	info.Mounts = mounts

	net, err := gatherNetwork()
	if err != nil {
		return nil, fmt.Errorf("failed to enumerate network interfaces: %v", err)
	}
	info.Network = net

	up, err := getUptime()
	if err != nil {
		return nil, fmt.Errorf("failed to parse uptime: %v", err)
	}
	info.Uptime = up

	return info, nil
}
