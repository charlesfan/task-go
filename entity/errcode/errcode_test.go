package errcode_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/charlesfan/task-go/entity/errcode"
)

// test error status code
func Test_ErrorStatusCode(t *testing.T) {
	err := errcode.New(errcode.ErrorCodeNotFound)
	statusCode := errcode.ParseError(err).HTTPStatus()
	assert.Equal(t, statusCode, 404)
}

// test error message
func Test_ErrorMsg(t *testing.T) {
	err := errcode.New(errcode.ErrorCodeForbidden)
	errMsg := errcode.ParseError(err).Text()
	assert.Equal(t, errMsg, "Forbidden")
}
