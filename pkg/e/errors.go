package e

type Apperror struct {
	Code int
	Msg  string
}

func (err *Apperror) Error() string {
	return err.Msg
}

func NewError(code int) *Apperror {
	return &Apperror{
		Code: code,
		Msg:  GetMsg(code),
	}
}
