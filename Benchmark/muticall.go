package Benchmark

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"time"
)

// UpdateMulticall :)
func (b Benchmark) excMultiCall(callCount int32) error {
	ctx, cnl := context.WithTimeout(*b.context, b.reqTimeout)
	defer cnl()
	_, err := b.multicall.Aggregate3(
		&bind.CallOpts{Context: ctx},
		balanceOfCalls[:callCount],
	)
	return err
}

// UpdateMulticallStats :)
func (b Benchmark) UpdateMulticallStats() error {
	if b.MS == nil {
		b.MS = &MulticallStats{}
	}
	var err error
	var s time.Time
	// 1 - 2000 calls

	s = time.Now()
	err = b.excMultiCall(1)
	if err != nil {
		b.MS.QT1 = -1
		b.MS.QT10 = -1
		b.MS.QT100 = -1
		b.MS.QT1000 = -1
		b.MS.QT2000 = -1
		return err
	}
	b.MS.QT1 = time.Since(s)

	s = time.Now()
	err = b.excMultiCall(10)
	if err != nil {
		b.MS.QT10 = -1
		b.MS.QT100 = -1
		b.MS.QT1000 = -1
		b.MS.QT2000 = -1
		return err
	}
	b.MS.QT10 = time.Since(s)

	s = time.Now()
	err = b.excMultiCall(100)
	if err != nil {
		b.MS.QT100 = -1
		b.MS.QT1000 = -1
		b.MS.QT2000 = -1
		return err
	}
	b.MS.QT100 = time.Since(s)

	s = time.Now()
	err = b.excMultiCall(1000)
	if err != nil {
		b.MS.QT1000 = -1
		b.MS.QT2000 = -1
		return err
	}
	b.MS.QT1000 = time.Since(s)

	s = time.Now()
	err = b.excMultiCall(2000)
	if err != nil {
		b.MS.QT2000 = -1
		return err
	}
	b.MS.QT2000 = time.Since(s)

	return err
}
