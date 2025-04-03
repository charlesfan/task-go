package resp

import (
	"github.com/gin-gonic/gin"

	"github.com/charlesfan/task-go/entity/errcode"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func WriteResponse(c *gin.Context, err error, data interface{}) {
	coder := errcode.ParseError(errcode.New(errcode.ErrorCodeSuccess))
	if err != nil {
		coder = errcode.ParseError(err)
	}
	c.JSON(coder.HTTPStatus(), &Response{
		Code: coder.Code(),
		Msg:  coder.Text(),
		Data: data,
	})

}
