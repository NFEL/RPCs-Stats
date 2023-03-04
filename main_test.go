package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
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
			if fun2 != nil && b != nil {
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
	assert.NoError(
		t,
		runOnAllRPCs(bench.Benchmark.UpdateMulticallStats, func(b bench.Benchmark, err error) {
			if err != nil {
				logrus.Error(err)
			} else {
				fmt.Printf("<MC> QT1 %d QT10 %d QT100 %d QT1000 %d QT2000 %d => [%s] \n", b.MS.QT1, b.MS.QT10, b.MS.QT100, b.MS.QT1000, b.MS.QT2000, b.Provider.RPC)
			}
		}),
	)
}
func TestPing(t *testing.T) {
	assert.NoError(
		t,
		runOnAllRPCs(bench.Benchmark.UpdatePing,
			func(b bench.Benchmark, err error) {
				if err != nil {
					t.Error(err)
					if b.Ping > 5*time.Second {
						t.Error(b.Ping, "RPC ping took more than 5s")
					}
				} else {
					fmt.Printf("<Ping>  %d => [%s]", b.Ping, b.Provider.RPC)
				}
			},
		))

}
