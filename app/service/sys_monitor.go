package service

import (
	"encoding/json"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"gsadmin/app/model"
	"gsadmin/core/baseservice"
	"gsadmin/core/cache"
	"gsadmin/core/config"
	"gsadmin/core/log"
	"gsadmin/core/utils/assertion"
	"gsadmin/core/utils/str"
	"gsadmin/global/e"
	"os"
	"runtime"
	"time"
)

type SysMonitor struct {
	baseservice.Service
}

func (s *SysMonitor) Server() (model.ServerInfo, error) {
	res, err := cache.Instance().Get(e.SysBase, e.ServerMonitor)
	if err == nil {
		var serverInfo model.ServerInfo
		err = json.Unmarshal([]byte(assertion.AnyToString(res)), &serverInfo)
		if err == nil {
			return serverInfo, err
		} else {
			log.Instance().Error("Unmarshal ServerInfo Failed..." + err.Error())
		}
	}

	return getServerInfo()
}

func getServerInfo() (model.ServerInfo, error) {
	var serverInfo = model.ServerInfo{
		CpuInfo:  getCpuInfo(),
		MemInfo:  getMemInfo(),
		SysInfo:  getSysInfo(),
		DiskInfo: getDiskInfo(),
		RunInfo:  getRunInfo(),
	}
	res, err := json.Marshal(serverInfo)
	if err != nil {
		log.Instance().Error("Marshal ServerInfo Failed..." + err.Error())
		return serverInfo, err
	}

	_ = cache.Instance().Set(e.SysBase, e.ServerMonitor, string(res), e.MonCacheTime)
	return serverInfo, nil
}

// CPU信息
func getCpuInfo() model.CpuInfo {
	var cpuinfo model.CpuInfo
	infos, _ := cpu.Info()
	totalPercent, _ := cpu.Percent(100*time.Millisecond, false)

	cpuinfo.ModelName = infos[0].ModelName
	cpuinfo.PhysicalId = infos[0].PhysicalID
	cpuinfo.Speed = infos[0].Mhz
	cpuinfo.UsedPercent = str.Decimal(totalPercent[0])
	cpuinfo.PhysicalCore, _ = cpu.Counts(false)
	cpuinfo.LogicalCore, _ = cpu.Counts(true)

	return cpuinfo
}

// 内存信息
func getMemInfo() model.MemInfo {
	var meminfo model.MemInfo
	v, _ := mem.VirtualMemory()

	meminfo.TotalMem = str.Decimal(float64(v.Total) / 1024 / 1024 / 1024)
	meminfo.AvailableMem = str.Decimal(float64(v.Available) / 1024 / 1024 / 1024)
	meminfo.UsedPercent = str.Decimal(v.UsedPercent)

	return meminfo
}

// 获取操作系统信息
func getSysInfo() model.SysInfo {
	//获取OS信息
	var osinfo model.SysInfo
	osinfo.HostName, _ = os.Hostname()
	timestamp, _ := host.BootTime()
	t := time.Unix(int64(timestamp), 0)
	osinfo.StartTime = t.Local().Format("2006-01-02 15:04:05")
	osinfo.Kernel, _ = host.KernelVersion()
	osinfo.Platform, osinfo.Family, osinfo.Version, _ = host.PlatformInformation()
	return osinfo
}

// 获取磁盘信息
func getDiskInfo() []model.DiskInfo {
	var di []model.DiskInfo

	infos, _ := disk.Partitions(false)
	for _, info := range infos {
		uinfo, _ := disk.Usage(info.Mountpoint)
		di = append(di, model.DiskInfo{
			Device:      info.Device,
			Path:        uinfo.Path,
			FsType:      info.Fstype,
			TotalCap:    str.Decimal(float64(uinfo.Total) / 1024 / 1024 / 1024),
			FreeCap:     str.Decimal(float64(uinfo.Free) / 1024 / 1024 / 1024),
			UsedCap:     str.Decimal(float64(uinfo.Used) / 1024 / 1024 / 1024),
			UsedPercent: str.Decimal(uinfo.UsedPercent),
		})
	}

	return di
}

// 获取程序运行信息
func getRunInfo() model.RunInfo {
	path, _ := os.Executable()
	return model.RunInfo{
		SdkVersion: runtime.Version(),
		AppVersion: config.Instance().App.Name + "@" + config.Instance().App.Version,
		Path:       path,
	}
}
