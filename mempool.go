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

// added extra field feePaid so i don't have to recalculate each time another transaction is added
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
		} else {
			// Provide more helpful message to user submitting the transaction.
			return
		}
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
func (mempool *Mempool) dumps() {
	f, _ := os.Create("prioritized-transactions.txt")
	w := bufio.NewWriter(f)

	for i, tx := range mempool.transactions {
		line := "TxHash=" + tx.txHash + " " +
			"Gas=" + strconv.Itoa(int(tx.gas)) + " " +
			"FeePerGas=" + fmt.Sprintf("%.16f", tx.feePerGas) + " " +
			"Signature=" + tx.signature

		if i != 4999 {
			w.WriteString(line + "\n")
		} else {
			w.WriteString(line)
		}
	}
	w.Flush()
}

// func main() {
// 	mempool := Mempool{
// 		transactions: make([]Transaction, 0),
// 		txHashes:     make(map[string]int, 0),
// 	}
// 	mempool.ingestFile("transactions.txt")

// 	// for i, t := range mempool.transactions[:10] {
// 	// 	if i == 10 {
// 	// 		break
// 	// 	}
// 	// 	fmt.Println(t.feePaid)
// 	// 	fmt.Println(t.feePerGas)
// 	// }

// 	mempool.dumps()

// }

// better user feedback messages when user submits duplicate transactions or their tx not added in general
// is gas always going to be an integer?
//

// for tests
// test for empty file
// test for empty transaction
// test file more than 5000 transactions
// file with duplicates
//

// bonus is add in sending back whether it's high to low priority fee
// or just get what a high to low priority fee could be
// function takes in the fee that you would pay and it would return the index you'd be at in the mempool

// don't need to pop last element, just check if insertIndex is last and capacity already 5000

// write tests
// write good docs/README
// make it according to code style of Tellor
// push to github
