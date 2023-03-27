/**
 * @Author root$
 * @Date 2023/3/26$
 * @Note 路由文件
 **/

package router

import (
	"github.com/gin-gonic/gin"
	"server-resource-probe/middleware"
	"server-resource-probe/service"
)

func SetupRouter(port string) {
	r := gin.Default()
	r.Use(middleware.Cors()) //配置跨域
	r.GET("/cpu", service.GetCPUInfo)
	r.GET("/mem", service.GetMemInfo)
	r.GET("/disk", service.GetDiskInfo)
	r.GET("/net", service.GetNetInfo)
	r.GET("/host", service.GetHostInfo)
	r.Run(":" + port)
}
