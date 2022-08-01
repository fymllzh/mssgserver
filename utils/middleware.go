package utils

import (
	"github.com/gin-gonic/gin"
	"time"
)

// StatCost 是一个统计耗时请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		//c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		Logger.Info("请求的耗时是:" + cost.String())
	}
}

//ip白名单
//func IpWhiteList() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		ip := GetRequestIP(c)
//		Logger.Info("ip 是:" + ip)
//		var ipList server.Ip
//		row,_ := ipList.Allow(ip)
//		if row.Id <= 0 {
//			Logger.Info("ip_denny:", ip)
//			c.JSON(http.StatusOK, gin.H{
//				"errno":   403,
//				"message": "Access Denny.",
//				"ip":      ip,
//			})
//			c.Abort()
//			return
//		}
//		c.Next()
//	}
//}
