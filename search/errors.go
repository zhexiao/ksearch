package search

type KsError struct {
	ErrorCode int
	ErrorMsg  string
}

func NewKsError(code int, msg string) *KsError {
	return &KsError{
		ErrorCode: code,
		ErrorMsg:  msg,
	}
}

func (k *KsError) Error() string {
	return k.ErrorMsg
}
