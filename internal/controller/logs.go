package controller

import (
	"tikuAdapter/internal/dao"

	"github.com/gin-gonic/gin"
)

// LogList - 日志列表
func LogList(c *gin.Context) {
	find, err := dao.Log.Where(dao.Log.UserID.Eq(0)).Order(dao.Log.ID.Desc()).Find()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
		})
		return
	}
	c.JSON(200, find)
}
