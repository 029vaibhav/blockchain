package util

const (
	LogLevel         = "Logger.Level"
	Blocks           = "block"
	Transactions     = "transaction"
	ClearTransaction = "clearTransaction"
)

type Error struct {
	msg string
}

func (err *Error) Error() string {
	return err.msg
}
