package sis

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func getUptime() (*Uptime, error) {
	uptime, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		return nil, err //TODO: Print better error
	}
	cpuinfo, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		return nil, err //TODO: Print better error
	}

	fields := strings.Fields(string(uptime))
	upsecs, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return nil, err // TODO: better error - should never happen
	}
	idlesecs, err := strconv.ParseFloat(fields[1], 64)
	if err != nil {
		return nil, err // TODO: better error - should never happen
	}

	procPattern := regexp.MustCompile("processor")

	matches := procPattern.FindAllStringIndex(string(cpuinfo), -1)
	cpus := len(matches)

	var idlePerCore float64 = idlesecs / float64(cpus)

	var idlepercs float64 = idlePerCore / upsecs

	uptimeinfo := &Uptime{
		UpSeconds:      upsecs,
		NumCores:       cpus,
		IdleSeconds:    idlesecs,
		IdlePercentage: idlepercs,
	}

	return uptimeinfo, nil
}
