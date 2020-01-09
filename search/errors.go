package search

type KsError struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
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
