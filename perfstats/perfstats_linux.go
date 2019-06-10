package perfstats

import (
	"encoding/json"
	"fmt"
)

func SayHi() {
	fmt.Println("hi unix")
}

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

	return jsonData
}
