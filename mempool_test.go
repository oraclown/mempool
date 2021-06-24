package main

import (
	"reflect"
	"testing"
	"time"
)

func TestTableMempool(t *testing.T) {
	testCases := []struct {
		name             string
		currentTimestamp int64
		startTime        int64
		mempool          Mempool
		transactions     []Transaction
		expected         []Transaction
	}{
		{
			"Add three different transactions.",
			time.Now().Unix(),
			time.Now().AddDate(0, 0, 1).Unix(),
			Mempool{
				transactions: make([]Transaction, 0),
				txHashes:     make(map[string]int),
			},
			[]Transaction{
				{
					txHash:    "txhash1",
					gas:       100.0,
					feePerGas: 5.0,
					feePaid:   500.0,
					signature: "sig1",
				},
				{
					txHash:    "txhash2",
					gas:       100.0,
					feePerGas: 3.0,
					feePaid:   300.0,
					signature: "sig2",
				},
				{
					txHash:    "txhash3",
					gas:       100.0,
					feePerGas: 6.0,
					feePaid:   600.0,
					signature: "sig3",
				},
			},
			[]Transaction{
				{
					txHash:    "txhash3",
					gas:       100.0,
					feePerGas: 6.0,
					feePaid:   600.0,
					signature: "sig3",
				},
				{
					txHash:    "txhash1",
					gas:       100.0,
					feePerGas: 5.0,
					feePaid:   500.0,
					signature: "sig1",
				},
				{
					txHash:    "txhash2",
					gas:       100.0,
					feePerGas: 3.0,
					feePaid:   300.0,
					signature: "sig2",
				},
			},
		},
		{
			"Add duplicate transactions.",
			time.Now().Unix(),
			time.Now().AddDate(0, 0, 1).Unix(),
			Mempool{
				transactions: make([]Transaction, 0),
				txHashes:     make(map[string]int),
			},
			[]Transaction{
				{
					txHash:    "txhash1",
					gas:       100.0,
					feePerGas: 5.0,
					feePaid:   500.0,
					signature: "sig1",
				},
				{
					txHash:    "txhash1",
					gas:       100.0,
					feePerGas: 5.0,
					feePaid:   500.0,
					signature: "sig1",
				},
			},
			[]Transaction{
				{
					txHash:    "txhash1",
					gas:       100.0,
					feePerGas: 5.0,
					feePaid:   500.0,
					signature: "sig1",
				},
			},
		},
	}

	for idx, testCase := range testCases {
		mempool := Mempool{
			transactions: make([]Transaction, 0),
			txHashes:     make(map[string]int),
		}

		for _, tx := range testCase.transactions {
			mempool.addTransaction(tx)
		}

		if !reflect.DeepEqual(mempool.transactions, testCase.expected) {
			t.Errorf(
				"Test %d: %s failed",
				idx,
				testCase.name,
			)
		}
	}
}
