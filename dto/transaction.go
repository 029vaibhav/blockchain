package dto

type TransactionReq struct {
	RecipientAdd string  `json:"address"`
	Amount       float64 `json:"amount"`
}
