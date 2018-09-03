package util

const (
	LOG_LEVEL = "Logger.Level"
)

type Error struct {
	msg string
}

func (err *Error) Error() string {
	return err.msg
}
