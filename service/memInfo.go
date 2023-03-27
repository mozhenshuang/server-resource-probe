/**
 * @Author root$
 * @Date 2023/3/27$
 * @Note 获取内存信息
 **/

package service

import (
	"github.com/shirou/gopsutil/mem"
	"server-resource-probe/model"
)

func MemInfo() model.Mem {
	//获取内存数据
	var m model.Mem
	v, _ := mem.VirtualMemory()
	m.MemTotal = v.Total
	m.MemAvailable = v.Available
	m.MemUsedPercent = v.UsedPercent
	return m
}
