package Benchmark

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"rpc-stats/Benchmark/contracts/MulticallContract"
	"time"
)

type MulticallStats struct {
	MaxSize int32
	QT1     time.Duration
	QT10    time.Duration
	QT100   time.Duration
	QT1000  time.Duration
	QT2000  time.Duration
}
type ThroughputStats struct {
	Second int
	Hour   int
	Day    int
	Week   int
	Month  int
	// TimeFrames
	startSecond time.Time
	endSecond   time.Time

	startHour time.Time
	endHour   time.Time

	startDay time.Time
	endDay   time.Time

	startWeek time.Time
	endWeek   time.Time

	startMonth time.Time
	endMonth   time.Time
}
type Provider struct {
	W3      *ethclient.Client `json:"-"`
	RPC     string            `json:"RPC"`
	ID      string            `json:"id"` // md5(url[with-schema],chainId)
	Index   int               `json:"index"`
	ChainId int64             `json:"chainId"`
}

type Benchmark struct {
	Provider    *Provider
	Ping        time.Duration
	BlockNumber uint64
	MS          *MulticallStats
	TS          *ThroughputStats

	LastBlockNumberRQT time.Duration

	startedAt      time.Time
	lastAvlCheckAt time.Time
	lastRelCheckAt time.Time
	reqTimeout     time.Duration
	// For Avoidance of reallocation
	multicall *Multicall.MulticallCaller
	url       string
	// Misc.
	context       *context.Context
	contextCancel *context.CancelFunc

	// TODO Availability float64 //  Percentage over a one-week period
	// TODO Reliability  float64 //  Percentage of stats changing
	// TODO - Security     float64 // Comparing this RPC data to others
}
