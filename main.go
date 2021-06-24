package main

func main() {
	mempool := Mempool{
		transactions: make([]Transaction, 0),
		txHashes:     make(map[string]int),
	}
	mempool.ingestFile("transactions.txt")

	// for i, t := range mempool.transactions[:10] {
	// 	if i == 10 {
	// 		break
	// 	}
	// 	fmt.Println(t.feePaid)
	// 	fmt.Println(t.feePerGas)
	// }

	mempool.dumps()
}
