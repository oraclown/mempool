# Simple Mempool
- Here's an implementation of a simple mempool with a maximum capacity of 5k transactions. Transactions are prioritized by the fee paid (feePerGas * gas).

### Next steps
- Additional tests (empty files, empty transactions, more realistic transactions, duplicate transactions, transactions with different formatting, transactions with incorrect/different data types)
- Better user feedback messages when their transactions are rejected for any reason
- Function that takes a proposed transaction and returns to the user what its priority (index) in the mempool would be if added
- Account for malicious actors (front-running, sandwhiching, etc.)?
- In addTransaction, don't need to pop last element, just check if insertIndex is last and capacity already 5000
