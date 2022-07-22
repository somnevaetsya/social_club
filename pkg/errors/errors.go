package customErrors

import (
	"errors"
	"net/http"
)

var (
	ErrBadInputData = errors.New("bad input data")
)

var errorToCode = map[error]int{
	ErrBadInputData: http.StatusBadRequest,
}

func ConvertErrorToCode(err error) (code int) {
	code, isErrorExist := errorToCode[err]
	if !isErrorExist {
		code = http.StatusInternalServerError
	}
	return
}
