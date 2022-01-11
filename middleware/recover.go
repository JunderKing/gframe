package middleware

import (
	"fmt"
	"gframe/helper/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Recover 捕获所有panic，并且返回错误信息
func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 正常响应
				if reflect.TypeOf(err).String() == "*http.RespStruct" {
					c.JSON(200, err)
					return
				}
				// 服务器异常
				if viper.GetBool("debug") {
					resp := &http.RespStruct{Code: http.CodeException, Msg: fmt.Sprint(err)}
					c.JSON(200, resp)
				} else {
					resp := &http.RespStruct{Code: http.CodeException, Msg: "内部错误"}
					c.JSON(200, resp)
				}
			}
		}()
		c.Next()
	}
}
