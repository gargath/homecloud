package sis

// SysInfo contains the system information gathered by the service
type SysInfo struct {
	SisVersion string   `json:"sisVersion"`
	Hostname   string   `json:"hostname"`
	Uptime     *Uptime  `json:"uptime"`
	Mounts     []*Mount `json:"mounts"`
	Network    *Network `json:"network"`
}

// Mount contains information on a single filesystem mount
type Mount struct {
	Device     string `json:"device"`
	MountPoint string `json:"mountPoint"`
	FsType     string `json:"fsType"`
	ReadOnly   bool   `json:"readOnlyMount"`
}

// Network contains information on network interfaces
type Network struct {
	Interfaces []*Interface `json:"interfaces"`
}

// Interface contains information on a single network interface
type Interface struct {
	Name         string     `json:"name"`
	MTU          int        `json:"mtu"`
	HardwareAddr string     `json:"hardwareAddr"`
	Flags        string     `json:"flags"`
	Addresses    []*Address `json:"addresses"`
}

// Address contains information on an IP address the system holds
type Address struct {
	Network string `json:"network"`
	Address string `json:"address"`
}

// Uptime contains information on the system uptime
type Uptime struct {
	NumCores       int     `json:"cpus"`
	UpSeconds      float64 `json:"upSeconds"`
	IdleSeconds    float64 `json:"idleSeconds"`
	IdlePercentage float64 `json:"idlePercentage"`
}
