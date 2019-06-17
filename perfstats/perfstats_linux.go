package perfstats

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	procCPUUser    = iota
	procCPUNice    = iota
	procCPUSystem  = iota
	procCPUIdle    = iota
	procCPUIowait  = iota
	procCPUirq     = iota
	procCPUSoftirq = iota
	procCPUSteal   = iota
	procCPUGuest   = iota
)

const (
	memoryTotal     = iota
	memoryUsed      = iota
	memoryFree      = iota
	memoryShared    = iota
	memoryBuffCache = iota
	memoryAvailable = iota
)

// convert proc time samples to float
func parseProcTokens(tokens []string) []float64 {
	var cpuRead []float64
	for _, token := range tokens {
		if s, err := strconv.ParseFloat(token, 32); err == nil {
			cpuRead = append(cpuRead, s)
		}
	}

	return cpuRead
}

// read & parse /proc/stat
func parseProcStat() (map[string][]float64, error) {

	file, err := os.Open("/proc/stat")
	defer file.Close()

	if err != nil {
		return nil, err
	}

	// parse text & store it in a format {"cpu": [12, 15734...]}
	cpuStats := make(map[string][]float64)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()

		if !strings.HasPrefix(line, "cpu") {
			break
		}

		tokens := strings.Fields(line)
		// first field is the cpu name, the rest are cpu time spent doing stuff
		// see http://www.linuxhowtos.org/System/procstat.htm
		cpuStats[tokens[0]] = parseProcTokens(tokens[1:])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return cpuStats, nil
}

func getDateFormatted() string {
	// match date format returned by win
	dt := time.Now()
	return dt.Format("1/2/2006 3:04:05 PM")
}

// compute active and total CPU utilizations from proc/stat readings
func computeActiveTotalCPU(procStats map[string][]float64) (map[string]float64, map[string]float64) {

	active := make(map[string]float64)
	total := make(map[string]float64)

	for cpuName, cpu := range procStats {
		active[cpuName] = cpu[procCPUUser] +
			cpu[procCPUSystem] +
			cpu[procCPUNice] +
			cpu[procCPUSoftirq] +
			cpu[procCPUSteal]

		total[cpuName] = active[cpuName] +
			cpu[procCPUIdle] +
			cpu[procCPUIowait]

	}

	return active, total
}

// compute CPU utilization by getting 2 samples and calculating delta between them
func getCPUStats() ([]SysStat, error) {

	// based on:
	// https://stackoverflow.com/questions/26791240/how-to-get-percentage-of-processor-use-with-bash

	var sysStats []SysStat
	var statEntry SysStat

	dateFormatted := getDateFormatted()
	timeBetweenSamples := 2 * time.Second

	// sample 2 stats with a time-delay in between
	procStatOne, err := parseProcStat()
	if err != nil {
		return sysStats, err
	}
	activeOne, totalOne := computeActiveTotalCPU(procStatOne)
	time.Sleep(timeBetweenSamples)
	procStatTwo, _ := parseProcStat()
	activeTwo, totalTwo := computeActiveTotalCPU(procStatTwo)

	// compute delta for all cpus (cpu % utilization)
	for cpuName := range activeOne {
		cpuUtilization := (100 * (activeTwo[cpuName] - activeOne[cpuName]) /
			(totalTwo[cpuName] - totalOne[cpuName]))

		// Populate returned cpu performance stats
		statEntry.Date = dateFormatted
		statEntry.Key = cpuName
		statEntry.Value = strconv.FormatFloat(cpuUtilization, 'f', 6, 64)

		sysStats = append(sysStats, statEntry)
	}

	return sysStats, nil
}

func getMemoryStats() (SysStat, error) {

	var memStat SysStat

	const (
		header = iota
		memory = iota
		swap   = iota
	)

	cmdResult := exec.Command("free", "--bytes")

	out, err := cmdResult.Output()
	if err != nil {
		return memStat, err
	}

	// get total, used, free etc. (discard first token sicne it's a label "Mem:")
	memInfo := strings.Fields(strings.Split(string(out[:]), "\n")[memory])[1:]

	memStat.Date = getDateFormatted()
	memStat.Key = "Memory Available"
	memStat.Value = memInfo[memoryAvailable]

	return memStat, nil
}

// PlatformSysStats Query performance stats on linux platform
func PlatformSysStats() (interface{}, error) {

	var stats []SysStat

	memInfo, err := getMemoryStats()
	if err != nil {
		return nil, fmt.Errorf("Cannot get memory details: %s", err)
	}
	cpuInfo, err := getCPUStats()
	if err != nil {
		return nil, fmt.Errorf("Cannot get CPU details: %s", err)
	}

	stats = append(stats, memInfo)
	stats = append(stats, cpuInfo...)

	return stats, nil
}
