package main

import (
	"github.com/owenburton/mempool/mempool"
)

func main() {
	mempool := mempool.Mempool{
		transactions: make([]mempool.Transaction, 0),
		txHashes:     make(map[string]int, 0),
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
