package Benchmark

import (
	"github.com/go-ping/ping"
)

// UpdatePing Updates Ping by icmp req to host 5 reqs each time
func (b Benchmark) UpdatePing() error {
	pinger, err := ping.NewPinger(b.url)
	if err != nil {
		return err
	}
	pinger.Count = 5
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		return err
	}
	stats := pinger.Statistics()
	b.Ping = stats.StdDevRtt
	return err
}
