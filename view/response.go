package view

import "github.com/gin-gonic/gin"

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

func Respond(c *gin.Context, code int, data interface{}, err error) {
	respond := Response{Code: code, Data: data, Error: nil}
	if err != nil {
		respond.Error = err.Error()
	}

	c.JSON(code, respond)
}
