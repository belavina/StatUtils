package perfstats

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func SayHi() {
	fmt.Println("hi unix")
}

const (
	procCPUName    = iota
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

func parseProcStat() [][]string {
	file, err := os.Open("/proc/stat")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var cpuStats [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()

		if !strings.HasPrefix(line, "cpu") {
			break
		}

		cpuStats = append(cpuStats, strings.Fields(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return cpuStats
}

func getCPUStats() {
	sampleOne := parseProcStat()
	time.Sleep(2 * time.Second)
	sampleTwo := parseProcStat()
	fmt.Println(sampleOne)
	fmt.Println(sampleTwo)

}

// https://stackoverflow.com/questions/26791240/how-to-get-percentage-of-processor-use-with-bash
// Query performance stats on linux platform
func PlatformSysStats() []byte {

	var statEntry SysStat
	var stats []SysStat

	csvData := [][]string{
		{"a", "b", "c"},
		{"1", "2", "3"},
	}

	for _, each := range csvData[1:] {
		statEntry.Date = each[0]
		statEntry.Key = each[1]
		statEntry.Value = each[2]

		stats = append(stats, statEntry)
	}

	jsonData, err := json.Marshal(stats)
	if err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
	}

	getCPUStats()
	return jsonData
}
