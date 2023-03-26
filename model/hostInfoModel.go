/**
 * @Author root$
 * @Date 2023/3/26$
 * @Note
 **/

package model

import (
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/net"
)

//
// Disk
// @Description: 硬盘信息结构体
//
type Disk struct {
	Driver string              `json:"driver"`
	Usage  *disk.UsageStat     `json:"disk_usage"`
	DiskIO disk.IOCountersStat `json:"disk_io"`
}

//
// HostInfo
// @Description: 主机信息结构体
//
type HostInfo struct {
	CpuPhysicalCount int                  `json:"cpu_physical_count"` //物理核心数
	CpuLogicalCount  int                  `json:"cpu_logical_count"`  //逻辑核心数
	CpuTotalPercent  []float64            `json:"cpu_total_percent"`  //CPU总占用
	CpuPerPercent    []float64            `json:"cpu_per_percent"`    //每一个逻辑核心的占用
	MemTotal         uint64               `json:"mem_total"`          //内存总量Gb
	MenAvailable     uint64               `json:"men_available"`      //内存可用容量Gb
	MemUsedPercent   float64              `json:"mem_used_percent"`   //内存已使用百分比
	DiskList         []Disk               `json:"disk_list"`          //磁盘分区信息
	NetIO            []net.IOCountersStat `json:"net_io"`             //网络IO
}
