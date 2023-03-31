package consts

type ResponseCodes int

const (
	Success = iota
	ValidationError
	UnexpectedError
)
