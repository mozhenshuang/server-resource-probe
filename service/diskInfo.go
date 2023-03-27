/**
 * @Author root$
 * @Date 2023/3/27$
 * @Note 获取硬盘信息列表
 **/

package service

import (
	"github.com/shirou/gopsutil/disk"
	"server-resource-probe/model"
)

//
// DiskInfo
//  @Description: 获取硬盘信息列表
//
func DiskInfo() []model.Disk {
	//获取硬盘数据
	var diskList []model.Disk
	info, _ := disk.Partitions(true) //所有分区
	info3, _ := disk.IOCounters()    //所有硬盘的io信息
	for _, device := range info {
		var diskTemp model.Disk
		diskTemp.Driver = device.Device
		info2, _ := disk.Usage(device.Device) //指定某路径的硬盘使用情况
		diskTemp.Usage = info2
		diskTemp.DiskIO = info3[device.Device]
		diskList = append(diskList, diskTemp)
	}
	return diskList
}
