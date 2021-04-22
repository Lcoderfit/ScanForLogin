package controller

import (
	"ScanForLogin/constant"
	"github.com/gin-gonic/gin"
	"net/http"
)

// success 请求成功
func success(c *gin.Context, code int) {
	c.JSON(http.StatusOK, gin.H{
		"code":  constant.Success,
		"message": constant.CodeMsg[constant.Success],
		"data":  nil,
	})
}

// successWithData 请求成功，可以自定义返回的数据
func successWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":  constant.Success,
		"message": constant.CodeMsg[constant.Success],
		"data":  data,
	})
}

// successWithStatusCode 自定义响应的http状态码和数据
func successWithStatusCode(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"code":    constant.Success,
		"message": constant.CodeMsg[constant.Success],
		"data":    data,
	})
}

// fail 请求失败
func fail(c *gin.Context, code int) {
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"error": constant.CodeMsg[code],
	})
}

// failWithStatusCode 自定义响应的http状态码和数据
