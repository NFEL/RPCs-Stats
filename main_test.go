package main

import (
	"github.com/stretchr/testify/assert"
	bench "rpc-stats/Benchmark"
	"rpc-stats/core"
	"testing"
	"time"
)

type bRunner func(bench.Benchmark) error
type resChecker func(bench.Benchmark, error)

func runOnAllRPCs(fun bRunner, fun2 resChecker) error {
	for _, rpcs := range core.Chains {
		for _, rpc := range rpcs {
			b, err := bench.New(rpc)
			if err == nil {
				err = fun(*b)
			}
			if fun2 != nil {
				fun2(*b, err)
			}
		}
	}
	return nil
}

func TestGen(t *testing.T) {
	b, err := bench.New("https://singapore.rpc.blxrbdn.com")
	assert.NoError(t, err, "https://singapore.rpc.blxrbdn.com")
	assert.NotNil(t, b)
	if b != nil {
		assert.NotNil(t, b.Provider)
		if b.Provider != nil {
			assert.Equal(t, b.Provider.ChainId, int64(1), "ChainId")
		}
	}
}

func TestChainId(t *testing.T) {

}

func TestMC(t *testing.T) {
	assert.Error(t, runOnAllRPCs(bench.Benchmark.UpdateMulticallStats, nil))
}
func TestPing(t *testing.T) {
	assert.Error(
		t,
		runOnAllRPCs(bench.Benchmark.UpdatePing,
			func(b bench.Benchmark, err error) {
				t.Error(err)
				if b.Ping > 5*time.Second {
					t.Error(b.Ping, "RPC ping took more than 5s")
				}
			},
		))

}
