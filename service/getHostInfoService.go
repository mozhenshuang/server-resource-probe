/**
 * @Author root$
 * @Date 2023/3/26$
 * @Note 获取主机信息（gopsutil）
 **/

package service

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"net/http"
	"server-resource-probe/model"
	"time"
)

func GetHostInfo(c *gin.Context) {
	var hostInfo model.HostInfo
	//获取CPU数据
	hostInfo.CpuPhysicalCount, _ = cpu.Counts(false)
	hostInfo.CpuLogicalCount, _ = cpu.Counts(true)
	hostInfo.CpuTotalPercent, _ = cpu.Percent(1*time.Second, false)
	hostInfo.CpuPerPercent, _ = cpu.Percent(1*time.Second, true)
	//获取内存数据
	v, _ := mem.VirtualMemory()
	hostInfo.MemTotal = v.Total
	hostInfo.MenAvailable = v.Available
	hostInfo.MemUsedPercent = v.UsedPercent
	//获取硬盘数据
	info, _ := disk.Partitions(true) //所有分区
	info3, _ := disk.IOCounters()    //所有硬盘的io信息
	for _, device := range info {
		var diskTemp model.Disk
		diskTemp.Driver = device.Device
		info2, _ := disk.Usage(device.Device) //指定某路径的硬盘使用情况
		diskTemp.Usage = info2
		diskTemp.DiskIO = info3[device.Device]
		hostInfo.DiskList = append(hostInfo.DiskList, diskTemp)
	}
	//获取网络数据
	net_info, _ := net.IOCounters(false)
	hostInfo.NetIO = net_info
	c.JSON(http.StatusOK, hostInfo)
}
