package transaction

type Pool struct {
	IdTransactionMap      map[string]*Transaction `json:"transactions"`
	AddressTransactionMap map[string]*Transaction `json:"address_transaction_map"`
}

var mainPool = Pool{}

func GetPool() *Pool {
	if len(mainPool.IdTransactionMap) == 0 {
		mainPool.IdTransactionMap = make(map[string]*Transaction)
		mainPool.AddressTransactionMap = make(map[string]*Transaction)
	}
	return &mainPool
}

func UpdateOrAddTransaction(transaction *Transaction) {

	GetPool().IdTransactionMap[transaction.Id] = transaction
	GetPool().AddressTransactionMap[transaction.Input.Address] = transaction

}

func GetExistingTransaction(key string) *Transaction {
	return GetPool().AddressTransactionMap[key]
}
