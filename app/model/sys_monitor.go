package model

// ServerInfo 监控信息不入库
type ServerInfo struct {
	CpuInfo  CpuInfo    `json:"cpuInfo"`
	MemInfo  MemInfo    `json:"memInfo"`
	SysInfo  SysInfo    `json:"sysInfo"`
	RunInfo  RunInfo    `json:"runInfo"`
	DiskInfo []DiskInfo `json:"diskInfo"`
}

type CpuInfo struct {
	PhysicalId   string  `json:"physicalId"`
	ModelName    string  `json:"modelName"`
	Speed        float64 `json:"speed"`
	PhysicalCore int     `json:"physicalCore"`
	LogicalCore  int     `json:"logicalCore"`
	UsedPercent  float64 `json:"usedPercent"`
}

type MemInfo struct {
	TotalMem     float64 `json:"totalMem"`
	AvailableMem float64 `json:"availableMem"`
	UsedPercent  float64 `json:"usedPercent"`
}

type RunInfo struct {
	SdkVersion string `json:"sdkVersion"`
	AppVersion string `json:"appVersion"`
	Path       string `json:"path"`
}

type SysInfo struct {
	HostName  string `json:"hostName"`
	Kernel    string `json:"kernel"`
	Platform  string `json:"platform"`
	Family    string `json:"family"`
	Version   string `json:"version"`
	StartTime string `json:"startTime"`
}

type DiskInfo struct {
	Device      string  `json:"device"`
	Path        string  `json:"path"`
	FsType      string  `json:"fsType"`
	TotalCap    float64 `json:"totalCap"`
	FreeCap     float64 `json:"freeCap"`
	UsedCap     float64 `json:"usedCap"`
	UsedPercent float64 `json:"usedPercent"`
}
