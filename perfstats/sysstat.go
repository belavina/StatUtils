package perfstats

// SysStat Performance statistics
type SysStat struct {
	Date  string // date when stat item was grabbed
	Key   string // type of sys stats (memory, cpu etc.)
	Value string // current value of sys stat (cpu usage %, free space)
}
