package errs

type ErrZeroDivision struct {
	message string
}

func (e ErrZeroDivision) Error() string {
	return e.message
}

func NewDivideByZeroError(message string) ErrZeroDivision {
	return ErrZeroDivision{
		message: message,
	}
}

func IsDivideByZeroError(err error) bool {
	_, ok := err.(ErrZeroDivision)
	return ok
}
