// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SystemInfoConf struct {
	ConfVersion *string `json:"confVersion,omitempty"`
	Inputs      int64   `json:"inputs"`
	Outputs     int64   `json:"outputs"`
	Pipelines   int64   `json:"pipelines"`
	Routes      int64   `json:"routes"`
	Rules       int64   `json:"rules"`
}

func (o *SystemInfoConf) GetConfVersion() *string {
	if o == nil {
		return nil
	}
	return o.ConfVersion
}

func (o *SystemInfoConf) GetInputs() int64 {
	if o == nil {
		return 0
	}
	return o.Inputs
}

func (o *SystemInfoConf) GetOutputs() int64 {
	if o == nil {
		return 0
	}
	return o.Outputs
}

func (o *SystemInfoConf) GetPipelines() int64 {
	if o == nil {
		return 0
	}
	return o.Pipelines
}

func (o *SystemInfoConf) GetRoutes() int64 {
	if o == nil {
		return 0
	}
	return o.Routes
}

func (o *SystemInfoConf) GetRules() int64 {
	if o == nil {
		return 0
	}
	return o.Rules
}

type SystemInfoCpus struct {
	Model string                 `json:"model"`
	Speed int64                  `json:"speed"`
	Times map[string]interface{} `json:"times"`
}

func (o *SystemInfoCpus) GetModel() string {
	if o == nil {
		return ""
	}
	return o.Model
}

func (o *SystemInfoCpus) GetSpeed() int64 {
	if o == nil {
		return 0
	}
	return o.Speed
}

func (o *SystemInfoCpus) GetTimes() map[string]interface{} {
	if o == nil {
		return map[string]interface{}{}
	}
	return o.Times
}

type SystemInfoDiskUsage struct {
	BytesAvailable int64  `json:"bytesAvailable"`
	BytesUsed      int64  `json:"bytesUsed"`
	DiskPath       string `json:"diskPath"`
	TotalDiskSize  int64  `json:"totalDiskSize"`
}

func (o *SystemInfoDiskUsage) GetBytesAvailable() int64 {
	if o == nil {
		return 0
	}
	return o.BytesAvailable
}

func (o *SystemInfoDiskUsage) GetBytesUsed() int64 {
	if o == nil {
		return 0
	}
	return o.BytesUsed
}

func (o *SystemInfoDiskUsage) GetDiskPath() string {
	if o == nil {
		return ""
	}
	return o.DiskPath
}

func (o *SystemInfoDiskUsage) GetTotalDiskSize() int64 {
	if o == nil {
		return 0
	}
	return o.TotalDiskSize
}

type SystemInfoEnv struct {
}

type SystemInfoLimitsSamples struct {
	MaxSize string `json:"maxSize"`
}

func (o *SystemInfoLimitsSamples) GetMaxSize() string {
	if o == nil {
		return ""
	}
	return o.MaxSize
}

type SystemInfoLimits struct {
	Samples SystemInfoLimitsSamples `json:"samples"`
}

func (o *SystemInfoLimits) GetSamples() SystemInfoLimitsSamples {
	if o == nil {
		return SystemInfoLimitsSamples{}
	}
	return o.Samples
}

type SystemInfoMemory struct {
	Free  int64 `json:"free"`
	Total int64 `json:"total"`
}

func (o *SystemInfoMemory) GetFree() int64 {
	if o == nil {
		return 0
	}
	return o.Free
}

func (o *SystemInfoMemory) GetTotal() int64 {
	if o == nil {
		return 0
	}
	return o.Total
}

type SystemInfoOs struct {
	Arch       string `json:"arch"`
	Endianness string `json:"endianness"`
	Platform   string `json:"platform"`
	Release    string `json:"release"`
	Type       string `json:"type"`
}

func (o *SystemInfoOs) GetArch() string {
	if o == nil {
		return ""
	}
	return o.Arch
}

func (o *SystemInfoOs) GetEndianness() string {
	if o == nil {
		return ""
	}
	return o.Endianness
}

func (o *SystemInfoOs) GetPlatform() string {
	if o == nil {
		return ""
	}
	return o.Platform
}

func (o *SystemInfoOs) GetRelease() string {
	if o == nil {
		return ""
	}
	return o.Release
}

func (o *SystemInfoOs) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

type SystemInfo struct {
	Build           map[string]interface{} `json:"BUILD"`
	APIPort         int64                  `json:"apiPort"`
	Conf            SystemInfoConf         `json:"conf"`
	ConfigPath      string                 `json:"configPath"`
	Cpus            []SystemInfoCpus       `json:"cpus"`
	DiskUsage       SystemInfoDiskUsage    `json:"diskUsage"`
	DistMode        AppMode                `json:"distMode"`
	Env             SystemInfoEnv          `json:"env"`
	GUID            string                 `json:"guid"`
	Hostname        string                 `json:"hostname"`
	InstallPath     string                 `json:"installPath"`
	License         LicenseInfo            `json:"license"`
	Limits          SystemInfoLimits       `json:"limits"`
	Loadavg         []int64                `json:"loadavg"`
	Memory          SystemInfoMemory       `json:"memory"`
	Messages        []BulletinMessage      `json:"messages"`
	Net             map[string]interface{} `json:"net"`
	Os              SystemInfoOs           `json:"os"`
	SystemConf      SystemConf             `json:"systemConf"`
	Uptime          int64                  `json:"uptime"`
	Version         string                 `json:"version"`
	WorkerProcesses int64                  `json:"workerProcesses"`
}

func (o *SystemInfo) GetBuild() map[string]interface{} {
	if o == nil {
		return map[string]interface{}{}
	}
	return o.Build
}

func (o *SystemInfo) GetAPIPort() int64 {
	if o == nil {
		return 0
	}
	return o.APIPort
}

func (o *SystemInfo) GetConf() SystemInfoConf {
	if o == nil {
		return SystemInfoConf{}
	}
	return o.Conf
}

func (o *SystemInfo) GetConfigPath() string {
	if o == nil {
		return ""
	}
	return o.ConfigPath
}

func (o *SystemInfo) GetCpus() []SystemInfoCpus {
	if o == nil {
		return []SystemInfoCpus{}
	}
	return o.Cpus
}

func (o *SystemInfo) GetDiskUsage() SystemInfoDiskUsage {
	if o == nil {
		return SystemInfoDiskUsage{}
	}
	return o.DiskUsage
}

func (o *SystemInfo) GetDistMode() AppMode {
	if o == nil {
		return AppMode("")
	}
	return o.DistMode
}

func (o *SystemInfo) GetEnv() SystemInfoEnv {
	if o == nil {
		return SystemInfoEnv{}
	}
	return o.Env
}

func (o *SystemInfo) GetGUID() string {
	if o == nil {
		return ""
	}
	return o.GUID
}

func (o *SystemInfo) GetHostname() string {
	if o == nil {
		return ""
	}
	return o.Hostname
}

func (o *SystemInfo) GetInstallPath() string {
	if o == nil {
		return ""
	}
	return o.InstallPath
}

func (o *SystemInfo) GetLicense() LicenseInfo {
	if o == nil {
		return LicenseInfo{}
	}
	return o.License
}

func (o *SystemInfo) GetLimits() SystemInfoLimits {
	if o == nil {
		return SystemInfoLimits{}
	}
	return o.Limits
}

func (o *SystemInfo) GetLoadavg() []int64 {
	if o == nil {
		return []int64{}
	}
	return o.Loadavg
}

func (o *SystemInfo) GetMemory() SystemInfoMemory {
	if o == nil {
		return SystemInfoMemory{}
	}
	return o.Memory
}

func (o *SystemInfo) GetMessages() []BulletinMessage {
	if o == nil {
		return []BulletinMessage{}
	}
	return o.Messages
}

func (o *SystemInfo) GetNet() map[string]interface{} {
	if o == nil {
		return map[string]interface{}{}
	}
	return o.Net
}

func (o *SystemInfo) GetOs() SystemInfoOs {
	if o == nil {
		return SystemInfoOs{}
	}
	return o.Os
}

func (o *SystemInfo) GetSystemConf() SystemConf {
	if o == nil {
		return SystemConf{}
	}
	return o.SystemConf
}

func (o *SystemInfo) GetUptime() int64 {
	if o == nil {
		return 0
	}
	return o.Uptime
}

func (o *SystemInfo) GetVersion() string {
	if o == nil {
		return ""
	}
	return o.Version
}

func (o *SystemInfo) GetWorkerProcesses() int64 {
	if o == nil {
		return 0
	}
	return o.WorkerProcesses
}
