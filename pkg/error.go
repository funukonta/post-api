package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ErrorWriter struct {
	Code    int `json:"code"`
	Message any `json:"message"`
}

func ErrorRes(code int, err error) *ErrorWriter {
	eW := new(ErrorWriter)

	eW.Code = http.StatusBadRequest
	if code != 0 {
		eW.Code = code
	}

	switch err {
	case gorm.ErrRecordNotFound:
		eW.Code = http.StatusNotFound
		eW.Message = "Record not found!"
	default:
		eW.Message = err.Error()
	}

	return eW
}

func (r *ErrorWriter) Send(c *gin.Context) {
	c.JSON(r.Code, map[string]interface{}{"error": r})
}
