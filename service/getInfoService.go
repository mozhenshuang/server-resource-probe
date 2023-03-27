/**
 * @Author root$
 * @Date 2023/3/26$
 * @Note 获取主机信息（gopsutil）
 **/

package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//
// GetCPUInfo
//  @Description: 获取CPU信息
//  @param c:
//
func GetCPUInfo(c *gin.Context) {
	cpuInfo := CPUInfo()
	c.JSON(http.StatusOK, cpuInfo)
}

//
// GetMemInfo
//  @Description: 获取内存信息
//  @param c:
//
func GetMemInfo(c *gin.Context) {
	memInfo := MemInfo()
	c.JSON(http.StatusOK, memInfo)
}

//
// GetDiskInfo
//  @Description: 获取硬盘信息
//  @param c:
//
func GetDiskInfo(c *gin.Context) {
	diskInfo := DiskInfo()
	c.JSON(http.StatusOK, diskInfo)
}

//
// GetNetInfo
//  @Description: 获取网络信息
//  @param c:
//
func GetNetInfo(c *gin.Context) {
	netInfo := NetIO()
	c.JSON(http.StatusOK, netInfo)
}

//
// GetHostInfo
//  @Description: 响应HTTP以Get方式获取服务器硬件资源信息请求
//  @param c:
//
func GetHostInfo(c *gin.Context) {
	hostInfo := HostInfo()
	c.JSON(http.StatusOK, hostInfo)
}
