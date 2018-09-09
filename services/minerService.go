package services

import (
	"bitbucket.org/blockchain/block"
	"bitbucket.org/blockchain/transaction"
	"bitbucket.org/blockchain/wallet"
)

func MineTransaction() block.Block {
	validTransaction := transaction.GetValidTransaction()
	rewardTransaction := transaction.RewardTransaction(WalletAddress, wallet.BlockChainWallet())
	validTransaction[rewardTransaction.Id] = rewardTransaction

	var transactions []transaction.Transaction
	for _, v := range validTransaction {
		transactions = append(transactions, *v)
	}
	newBlock := AddBlockToBlockChain(transactions)
	ClearTransaction()
	return newBlock
}
