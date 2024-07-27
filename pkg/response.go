package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Respon struct {
	Code    int `json:"-"`
	Data    any `json:"data,omitempty"`
	Message any `json:"message,omitempty"`
}

func Response(res *Respon) *Respon {
	respon := new(Respon)

	respon.Code = http.StatusOK
	if res.Code != 0 {
		respon.Code = res.Code
	}

	if res.Data != nil {
		respon.Data = res.Data
	}
	if res.Message != nil {
		respon.Message = res.Message
	}

	return respon
}

func (r *Respon) Send(c *gin.Context) {
	c.JSON(r.Code, r)
}
