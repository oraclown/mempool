package main

func main() {
	mempool := Mempool{
		transactions: make([]Transaction, 0),
		txHashes:     make(map[string]int),
	}
	mempool.ingestFile("transactions.txt")
	mempool.dumps("prioritized-transactions.txt")
}
