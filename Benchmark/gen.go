package Benchmark

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"net/url"
	"rpc-stats/Benchmark/contracts/MulticallContract"
	"rpc-stats/core"
	"time"
)

const (
	DEFAULT_RPC_TIMEOUT = 5 * time.Second
)

var (
	MULTICALL_V3_ADDRESS   = common.HexToAddress("0xca11bde05977b3631167028862be2a173976ca11")
	createdBenchmarks      = make([]Benchmark, 0)
	BenchmarkContext       context.Context
	BenchmarkContextCancel context.CancelFunc
)

func init() {
	ctx, cnl := context.WithCancel(context.Background())
	BenchmarkContextCancel, BenchmarkContext = cnl, ctx
}

// New by reading chain from RPC
func New(rpc string) (*Benchmark, error) {
	rpcUrl, err := url.Parse(rpc)
	if err != nil {
		return nil, err
	}
	client, err := ethclient.Dial(rpc)
	if err != nil {
		//log.Errorf("Client Connection Error : %s  @ chainId: %d @ RPC: %s", err, chainId, RPC)
		return nil, err
	} else {
		ctx, cnl := context.WithTimeout(BenchmarkContext, DEFAULT_RPC_TIMEOUT)
		defer cnl()
		chainId, _err := client.ChainID(ctx)
		if _err != nil {
			return nil, _err
		}

		mcc, err := Multicall.NewMulticallCaller(MULTICALL_V3_ADDRESS, client)
		if err != nil {
			return nil, err
		}

		p := Provider{
			W3:      client,
			ID:      core.MD5Hash(fmt.Sprintf("%s%d", rpc, chainId.Int64())),
			RPC:     rpc,
			ChainId: chainId.Int64(),
		}
		_ctx, _cnl := context.WithCancel(BenchmarkContext)

		b := Benchmark{
			Provider:      &p,
			url:           rpcUrl.Host,
			multicall:     mcc,
			context:       &_ctx,
			contextCancel: &_cnl,
		}
		createdBenchmarks = append(createdBenchmarks, b)
		return &b, nil
	}
}
