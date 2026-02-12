package bu_errors

type BusinessError struct {
	msg string
}

func (this *BusinessError) Error() string {
	return this.msg
}

func NewBusinessErrorMsg(msg string) error {
	return &BusinessError{
		msg: msg,
	}
}

func NewBusinessError(err error) error {
	return &BusinessError{
		msg: err.Error(),
	}
}
