package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Transaction struct {
	txHash    string
	gas       float64
	feePerGas float64
	feePaid   float64
	signature string
}

type Mempool struct {
	transactions []Transaction
	txHashes     map[string]int
}

func (mempool *Mempool) addTransaction(transaction Transaction) {
	// Skip duplicates.
	if _, ok := mempool.txHashes[transaction.txHash]; ok {
		return
	}

	// Append transaction if mempool under capacity.
	if len(mempool.transactions) > 0 &&
		transaction.feePaid < mempool.transactions[len(mempool.transactions)-1].feePaid {

		if len(mempool.transactions) < 5000 {
			mempool.transactions = append(mempool.transactions, transaction)
		}
		return
	}

	// Find where transaction should be inserted in the mempool.
	insertIndex := sort.Search(
		len(mempool.transactions),
		func(i int) bool { return mempool.transactions[i].feePaid < transaction.feePaid },
	)

	// Insert transaction in mempool.
	mempool.transactions = append(mempool.transactions, Transaction{})
	copy(mempool.transactions[insertIndex+1:], mempool.transactions[insertIndex:])
	mempool.transactions[insertIndex] = transaction
	mempool.txHashes[transaction.txHash] = 0

	// Pop last transaction if mempool over max capacity.
	if len(mempool.transactions) > 5000 {
		delete(mempool.txHashes, mempool.transactions[len(mempool.transactions)-1].txHash)
		mempool.transactions = mempool.transactions[:len(mempool.transactions)-1]
	}
}

// Add transactions to mempool from text file.
func (mempool *Mempool) ingestFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var gas float64
	var feePerGas float64
	var feePaid float64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := strings.Fields(scanner.Text())

		txHash := string([]rune(t[0])[7:])
		if gas, err = strconv.ParseFloat(t[1][4:], 64); err != nil {
			log.Fatal(err)
			continue
		}
		if feePerGas, err = strconv.ParseFloat(t[2][10:], 64); err != nil {
			log.Fatal(err)
			continue
		}
		feePaid = gas * feePerGas
		signature := string([]rune(t[3])[10:])

		temp := Transaction{
			txHash:    txHash,
			gas:       gas,
			feePerGas: feePerGas,
			feePaid:   feePaid,
			signature: signature,
		}

		mempool.addTransaction(temp)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// Write mempool transactions to text file.
func (mempool *Mempool) dumps(path string) {
	f, _ := os.Create(path)
	w := bufio.NewWriter(f)

	for i, tx := range mempool.transactions {
		line := "TxHash=" + tx.txHash + " " +
			"Gas=" + strconv.Itoa(int(tx.gas)) + " " +
			"FeePerGas=" + fmt.Sprintf("%.16f", tx.feePerGas) + " " +
			"Signature=" + tx.signature

		if i != 4999 {
			line = line + "\n"
		}
		w.WriteString(line)
	}
	w.Flush()
}
