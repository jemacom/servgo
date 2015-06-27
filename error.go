package main

type UnparsableRequestError struct {
	s string
}

func (err *UnparsableRequestError) Error() string {
	return err.s
}

type NotAllowedMethodError struct {
	s string
}

func (err *NotAllowedMethodError) Error() string {
	return err.s
}
