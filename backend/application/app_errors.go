package application

type ErrInputValidationFailed struct {
	Msg string
}

func (e *ErrInputValidationFailed) Error() string {
	if e.Msg != "" {
		return e.Msg
	}
	return "[ERROR] failed to validate input"
}

type ErrLoginFailed struct {
	Msg string
}

func (e *ErrLoginFailed) Error() string {
	if e.Msg != "" {
		return e.Msg
	}
	return "[ERROR] failed to login"
}

type ErrInternal struct {
	Msg string
}

func (e *ErrInternal) Error() string {
	if e.Msg != "" {
		return e.Msg
	}
	return "[ERROR] internal error"
}
