package Benchmark

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"time"
)

func BlockNumber(w3 *ethclient.Client, timeout time.Duration) (uint64, error) {
	contx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return w3.BlockNumber(contx)
}

// UpdateBlockNumber RPC call for new block number
func (b Benchmark) UpdateBlockNumber() error {
	s := time.Now()
	bno, err := BlockNumber(b.Provider.W3, b.reqTimeout)
	if err == nil {
		b.BlockNumber = bno
	}
	b.LastBlockNumberRQT = time.Now().Sub(s)
	return err
}
