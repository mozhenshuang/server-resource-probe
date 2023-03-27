/**
 * @Author root$
 * @Date 2023/3/27$
 * @Note 获取网络IO信息
 **/

package service

import (
	"github.com/shirou/gopsutil/net"
	"server-resource-probe/model"
)

//
// NetIO
//  @Description: 获取网络IO信息
//
func NetIO() model.Net {
	//获取网络数据
	var netIO model.Net
	netInfo, _ := net.IOCounters(false)
	netIO.NetIO = netInfo
	return netIO
}
