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
	r.GET("/gethostinfo", service.GetHostInfo)
	r.Run(":" + port)
}
