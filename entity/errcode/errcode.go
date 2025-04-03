/*
http status code + category + serial number
xxx 00 xx General
xxx 01 xx Task
*/
package errcode

import (
	"fmt"
	"net/http"
)

type Coder interface {
	Code() int
	Text() string
	SetText(s string)
	HTTPStatus() int
}

type xcode struct {
	code       int
	text       string
	httpStatus int
}

func (x xcode) Code() int {
	return x.code
}

func (x xcode) Text() string {
	return x.text
}

func (x *xcode) SetText(s string) {
	x.text = s
}

func (x xcode) HTTPStatus() int {
	return x.httpStatus
}

func (x xcode) Error() string {
	return x.text
}

// 200 00
const (
	ErrorCodeSuccess = iota + 20000000
	ErrorCodeSuccessButNotFund
)

// 400 00
const (
	ErrorCodeBadRequest = iota + 40000000
)

// 403 00
const (
	ErrorCodeForbidden = iota + 40300000
)

// 404 00
const (
	ErrorCodeNotFound = iota + 40400000
)

// 500 00
const (
	ErrorCodeServerErr = iota + 50000000
)

// 500 01
const (
	ErrorCodeTaskErr = iota + 50001000
)

func New(code int) error {
	if v, ok := xcodes[code]; ok {
		return v
	}

	return &xcode{
		-1,
		fmt.Sprintf("unknown code %d", code),
		http.StatusOK,
	}
}

func NewWithBind(code int, params ...interface{}) error {
	if v, ok := xcodes[code]; ok {
		if len(params) > 0 {
			return &xcode{
				code:       v.code,
				text:       fmt.Sprintf(v.text, params...),
				httpStatus: v.httpStatus,
			}
		}
		return v
	}

	return &xcode{
		-1,
		fmt.Sprintf("unknown code %d", code),
		http.StatusOK,
	}
}

var xcodes = map[int]*xcode{
	ErrorCodeSuccess:           &xcode{ErrorCodeSuccess, "success", http.StatusOK},
	ErrorCodeSuccessButNotFund: &xcode{ErrorCodeSuccessButNotFund, "success but not found", http.StatusOK},
	ErrorCodeBadRequest:        {ErrorCodeBadRequest, "bad request", http.StatusBadRequest},
	ErrorCodeForbidden:         {ErrorCodeForbidden, "Forbidden", http.StatusForbidden},
	ErrorCodeNotFound:          {ErrorCodeNotFound, "not found", http.StatusNotFound},
	ErrorCodeServerErr:         {ErrorCodeServerErr, "internal server error", http.StatusInternalServerError},
	ErrorCodeTaskErr:           {ErrorCodeTaskErr, "task service error", http.StatusInternalServerError},
}

func ParseError(err error) Coder {
	if err == nil {
		return nil
	}

	if v, ok := err.(*xcode); ok {
		return v
	}

	return &xcode{
		-1,
		fmt.Sprintf("unknown code %s", err),
		http.StatusOK,
	}
}
