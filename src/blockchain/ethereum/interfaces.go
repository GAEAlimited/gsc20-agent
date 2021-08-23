package ethereum

import "swap.io-agent/src/blockchain/ethereum/api"

type IGeth interface {
	GetBlockByIndex(
		index int,
	) (*api.Block, int)
	GetTransactionByHash(
		hash string,
	) (*api.BlockTransaction, int)
	GetTransactionLogs(
		hash string,
	) (*api.TransactionLogs, int)
	GetInternalTransaction(
		hash string,
	) (*api.InteranlTransaction, int)
	GetBlockTraceByIndex(
		index string,
	) (interface{}, int)
}
