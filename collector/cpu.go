package goniplus

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

// LocalCPUMetric keeps total, idle value for calculating cpu usage.
type localCPUMetric struct {
	total, idle uint64
}

// calcCPUUsage takes fields which parsed from /proc/stat,
// and returns calculated CPU usage. (float64)
func calcCPUUsage(fields []string) float64 {
	var idle, total uint64
	for i := 1; i < len(fields); i++ {
		u, _ := strconv.ParseUint(fields[i], 10, 64)
		if i == 4 || i == 5 {
			idle += u
		}
		total += u
	}
	prevCPUMetric := client.tMetric.prevCPUMetric
	v := float64((total-prevCPUMetric.total)-(idle-prevCPUMetric.idle)) / float64(total-prevCPUMetric.total)
	client.tMetric.prevCPUMetric.idle = idle
	client.tMetric.prevCPUMetric.total = total
	return v
}

// GetCPUUsage returns CPU usage(float64). If system doesn't support /proc/stat
// or failed to parse cpu data will return 0.0 with Error.
func GetCPUUsage() (float64, error) {
	d, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return 0.0, errors.New("Cannot read CPU data")
	}
	data := string(d)
	lines := strings.Split(data, "\n")
	fields := strings.Fields(lines[0])
	if fields[0][:3] == "cpu" {
		usage := calcCPUUsage(fields)
		return usage, nil
	}
	return float64(0), errors.New("Cannot parse CPU data")
}
