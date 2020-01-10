package main

type LogicError struct {
	msg string
}

func NewLogicError(msg string) *LogicError {

	return &LogicError{
		msg: msg,
	}
}

func (le *LogicError) Error() string {
	return le.msg
}
