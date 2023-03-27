/**
 * @Author root$
 * @Date 2023/3/27$
 * @Note 获取主机信息
 **/

package service

import (
	"server-resource-probe/model"
)

//
// HostInfo
//  @Description: 获取服务器硬件资源信息函数
//  @param c:
//
func HostInfo() model.HostInfo {
	var hostInfo model.HostInfo
	//获取CPU数据
	hostInfo.CPU = CPUInfo()
	//获取内存数据
	hostInfo.Mem = MemInfo()
	//获取硬盘数据
	hostInfo.DiskList = DiskInfo()
	//获取网络数据
	hostInfo.Net = NetIO()
	return hostInfo

}
