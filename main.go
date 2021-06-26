package main

func main() {
	mempool := Mempool{
		transactions: make([]Transaction, 0),
		txHashes:     make(map[string]int),
	}
	mempool.IngestFile("transactions.txt")
	mempool.Dumps("prioritized-transactions.txt")
}
