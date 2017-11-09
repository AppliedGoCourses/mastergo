package temperr

// TempError is an error type for temporary errors
type TempError struct {
	msg string
}

func New(m string) error {
	return &TempError{msg: m}
}

func (e *TempError) Error() string {
	return e.msg
}

func (e *TempError) IsTemporary() bool {
	return true
}
