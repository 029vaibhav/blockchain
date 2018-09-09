package transaction

import log "github.com/sirupsen/logrus"

type Pool struct {
	IdTransactionMap      map[string]*Transaction `json:"transactions"`
	AddressTransactionMap map[string]*Transaction `json:"address_transaction_map"`
}

var mainPool = Pool{}

func getPool() *Pool {
	if len(mainPool.IdTransactionMap) == 0 {
		mainPool.IdTransactionMap = make(map[string]*Transaction)
		mainPool.AddressTransactionMap = make(map[string]*Transaction)
	}
	return &mainPool
}

func GetTransactions() map[string]*Transaction {
	return getPool().IdTransactionMap
}

func UpdateOrAddTransaction(transaction *Transaction) {

	getPool().IdTransactionMap[transaction.Id] = transaction
	getPool().AddressTransactionMap[transaction.Input.Address] = transaction

}

func GetExistingTransaction(key string) *Transaction {
	return getPool().AddressTransactionMap[key]
}

func GetValidTransaction() map[string]*Transaction {

	transactions := GetTransactions()

	validTransacions := make(map[string]*Transaction)
	for _, v := range transactions {
		var total float64
		for _, v := range v.Outputs {
			total = total + v.Amount
		}
		if v.Input.Amount != total {
			log.Info("invalid transaction for address ", v.Input.Address)
		} else if !v.VerifyTransaction() {
			log.Info("invalid signature transaction for address ", v.Input.Address)
		} else {
			validTransacions[v.Id] = v
		}
	}
	return validTransacions

}

func Clear() {
	getPool().IdTransactionMap = nil
	getPool().AddressTransactionMap = nil
}
