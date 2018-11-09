package errs

type ErrInvalidArgs struct {
	message string
}

func (e ErrInvalidArgs) Error() string {
	return e.message
}

func NewErrInvalidArgs(message string) ErrInvalidArgs {
	return ErrInvalidArgs{
		message: message,
	}
}

func IsErrInvalidArgs(err error) bool {
	_, ok := err.(ErrInvalidArgs)
	return ok
}
