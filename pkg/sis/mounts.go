package sis

import (
	"bufio"
	"os"
	"strings"
)

func parseMounts() ([]*Mount, error) {
	mountinfo, err := os.Open("/proc/mounts")
	if err != nil {
		return nil, err //TODO: Print better error
	}
	defer mountinfo.Close()

	mounts := []*Mount{}

	scanner := bufio.NewScanner(mountinfo)
	for scanner.Scan() {
		mount, err := parseMount(scanner.Text())
		if err != nil {
			return nil, err //TODO: Print better error
		}
		mounts = append(mounts, mount)
	}

	return mounts, nil
}

func parseMount(line string) (*Mount, error) {
	fields := strings.Fields(line)
	mount := &Mount{
		Device:     fields[0],
		MountPoint: fields[1],
		FsType:     fields[2],
	}
	if strings.HasPrefix(fields[3], "ro") {
		mount.ReadOnly = true
	}
	return mount, nil
}
