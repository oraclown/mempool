# Simple Mempool
- Here's an implementation of a simple mempool with a maximum capacity of 5k transactions. Transactions are prioritized by the fee paid (feePerGas * gas).

### Run tests
- `cd src && go test`

### Next steps
- Additional tests (empty files, empty transactions, transactions with different formatting, transactions with incorrect/different data types)
- Better user feedback messages when their transactions are rejected for any reason
- Function that takes a proposed transaction and returns to the user what its priority (index) in the mempool would be if added
- Account for malicious actors (front-running, sandwhiching, etc.)?
- In addTransaction, don't need to pop last element, just check if insertIndex is last and capacity already 5000
- Add remove transaction function
- Remove transactions from mempool once added to block
