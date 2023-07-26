// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SystemSettingsAPIHeaders struct {
}

type SystemSettingsAPISsl struct {
	CaPath      *string `json:"caPath,omitempty"`
	CertPath    string  `json:"certPath"`
	Disabled    bool    `json:"disabled"`
	Passphrase  string  `json:"passphrase"`
	PrivKeyPath string  `json:"privKeyPath"`
}

func (o *SystemSettingsAPISsl) GetCaPath() *string {
	if o == nil {
		return nil
	}
	return o.CaPath
}

func (o *SystemSettingsAPISsl) GetCertPath() string {
	if o == nil {
		return ""
	}
	return o.CertPath
}

func (o *SystemSettingsAPISsl) GetDisabled() bool {
	if o == nil {
		return false
	}
	return o.Disabled
}

func (o *SystemSettingsAPISsl) GetPassphrase() string {
	if o == nil {
		return ""
	}
	return o.Passphrase
}

func (o *SystemSettingsAPISsl) GetPrivKeyPath() string {
	if o == nil {
		return ""
	}
	return o.PrivKeyPath
}

type SystemSettingsAPI struct {
	BaseURL            *string                   `json:"baseUrl,omitempty"`
	DisableAPICache    *bool                     `json:"disableApiCache,omitempty"`
	Disabled           bool                      `json:"disabled"`
	Headers            *SystemSettingsAPIHeaders `json:"headers,omitempty"`
	Host               string                    `json:"host"`
	IdleSessionTTL     *string                   `json:"idleSessionTTL,omitempty"`
	LoginRateLimit     *int64                    `json:"loginRateLimit,omitempty"`
	Port               int64                     `json:"port"`
	Ssl                SystemSettingsAPISsl      `json:"ssl"`
	SsoRateLimit       *string                   `json:"ssoRateLimit,omitempty"`
	WorkerRemoteAccess bool                      `json:"workerRemoteAccess"`
}

func (o *SystemSettingsAPI) GetBaseURL() *string {
	if o == nil {
		return nil
	}
	return o.BaseURL
}

func (o *SystemSettingsAPI) GetDisableAPICache() *bool {
	if o == nil {
		return nil
	}
	return o.DisableAPICache
}

func (o *SystemSettingsAPI) GetDisabled() bool {
	if o == nil {
		return false
	}
	return o.Disabled
}

func (o *SystemSettingsAPI) GetHeaders() *SystemSettingsAPIHeaders {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *SystemSettingsAPI) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *SystemSettingsAPI) GetIdleSessionTTL() *string {
	if o == nil {
		return nil
	}
	return o.IdleSessionTTL
}

func (o *SystemSettingsAPI) GetLoginRateLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.LoginRateLimit
}

func (o *SystemSettingsAPI) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *SystemSettingsAPI) GetSsl() SystemSettingsAPISsl {
	if o == nil {
		return SystemSettingsAPISsl{}
	}
	return o.Ssl
}

func (o *SystemSettingsAPI) GetSsoRateLimit() *string {
	if o == nil {
		return nil
	}
	return o.SsoRateLimit
}

func (o *SystemSettingsAPI) GetWorkerRemoteAccess() bool {
	if o == nil {
		return false
	}
	return o.WorkerRemoteAccess
}

type SystemSettingsBackups struct {
	BackupPersistence string `json:"backupPersistence"`
	BackupsDirectory  string `json:"backupsDirectory"`
}

func (o *SystemSettingsBackups) GetBackupPersistence() string {
	if o == nil {
		return ""
	}
	return o.BackupPersistence
}

func (o *SystemSettingsBackups) GetBackupsDirectory() string {
	if o == nil {
		return ""
	}
	return o.BackupsDirectory
}

type SystemSettingsCustomLogo struct {
	Enabled         bool   `json:"enabled"`
	LogoDescription string `json:"logoDescription"`
	LogoImage       string `json:"logoImage"`
}

func (o *SystemSettingsCustomLogo) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *SystemSettingsCustomLogo) GetLogoDescription() string {
	if o == nil {
		return ""
	}
	return o.LogoDescription
}

func (o *SystemSettingsCustomLogo) GetLogoImage() string {
	if o == nil {
		return ""
	}
	return o.LogoImage
}

type SystemSettingsDistributed struct {
	Mode AppMode `json:"mode"`
}

func (o *SystemSettingsDistributed) GetMode() AppMode {
	if o == nil {
		return AppMode("")
	}
	return o.Mode
}

type SystemSettingsProxy struct {
	UseEnvVars bool `json:"useEnvVars"`
}

func (o *SystemSettingsProxy) GetUseEnvVars() bool {
	if o == nil {
		return false
	}
	return o.UseEnvVars
}

type SystemSettingsRollback struct {
	RollbackEnabled bool   `json:"rollbackEnabled"`
	RollbackRetries *int64 `json:"rollbackRetries,omitempty"`
	RollbackTimeout *int64 `json:"rollbackTimeout,omitempty"`
}

func (o *SystemSettingsRollback) GetRollbackEnabled() bool {
	if o == nil {
		return false
	}
	return o.RollbackEnabled
}

func (o *SystemSettingsRollback) GetRollbackRetries() *int64 {
	if o == nil {
		return nil
	}
	return o.RollbackRetries
}

func (o *SystemSettingsRollback) GetRollbackTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.RollbackTimeout
}

type SystemSettingsShutdown struct {
	DrainTimeout int64 `json:"drainTimeout"`
}

func (o *SystemSettingsShutdown) GetDrainTimeout() int64 {
	if o == nil {
		return 0
	}
	return o.DrainTimeout
}

type SystemSettingsSystem struct {
	Intercom bool `json:"intercom"`
}

func (o *SystemSettingsSystem) GetIntercom() bool {
	if o == nil {
		return false
	}
	return o.Intercom
}

type SystemSettingsTLS struct {
	DefaultCipherList  string `json:"defaultCipherList"`
	DefaultEcdhCurve   string `json:"defaultEcdhCurve"`
	MaxVersion         string `json:"maxVersion"`
	MinVersion         string `json:"minVersion"`
	RejectUnauthorized bool   `json:"rejectUnauthorized"`
}

func (o *SystemSettingsTLS) GetDefaultCipherList() string {
	if o == nil {
		return ""
	}
	return o.DefaultCipherList
}

func (o *SystemSettingsTLS) GetDefaultEcdhCurve() string {
	if o == nil {
		return ""
	}
	return o.DefaultEcdhCurve
}

func (o *SystemSettingsTLS) GetMaxVersion() string {
	if o == nil {
		return ""
	}
	return o.MaxVersion
}

func (o *SystemSettingsTLS) GetMinVersion() string {
	if o == nil {
		return ""
	}
	return o.MinVersion
}

func (o *SystemSettingsTLS) GetRejectUnauthorized() bool {
	if o == nil {
		return false
	}
	return o.RejectUnauthorized
}

type SystemSettingsWorkers struct {
	Count                  int64  `json:"count"`
	LoadThrottlePerc       *int64 `json:"loadThrottlePerc,omitempty"`
	Memory                 int64  `json:"memory"`
	Minimum                int64  `json:"minimum"`
	StartupMaxConns        *int64 `json:"startupMaxConns,omitempty"`
	StartupThrottleTimeout *int64 `json:"startupThrottleTimeout,omitempty"`
	V8SingleThread         *bool  `json:"v8SingleThread,omitempty"`
}

func (o *SystemSettingsWorkers) GetCount() int64 {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *SystemSettingsWorkers) GetLoadThrottlePerc() *int64 {
	if o == nil {
		return nil
	}
	return o.LoadThrottlePerc
}

func (o *SystemSettingsWorkers) GetMemory() int64 {
	if o == nil {
		return 0
	}
	return o.Memory
}

func (o *SystemSettingsWorkers) GetMinimum() int64 {
	if o == nil {
		return 0
	}
	return o.Minimum
}

func (o *SystemSettingsWorkers) GetStartupMaxConns() *int64 {
	if o == nil {
		return nil
	}
	return o.StartupMaxConns
}

func (o *SystemSettingsWorkers) GetStartupThrottleTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.StartupThrottleTimeout
}

func (o *SystemSettingsWorkers) GetV8SingleThread() *bool {
	if o == nil {
		return nil
	}
	return o.V8SingleThread
}

// SystemSettings - a list of SystemSettings objects
type SystemSettings struct {
	API                  SystemSettingsAPI         `json:"api"`
	Auth                 AuthConfig                `json:"auth"`
	Backups              SystemSettingsBackups     `json:"backups"`
	CustomLogo           SystemSettingsCustomLogo  `json:"customLogo"`
	Distributed          SystemSettingsDistributed `json:"distributed"`
	Git                  GitSettings               `json:"git"`
	JobLimits            JobSettings               `json:"jobLimits"`
	Limits               Limits                    `json:"limits"`
	Proxy                SystemSettingsProxy       `json:"proxy"`
	RedisCacheLimits     RedisCacheLimits          `json:"redisCacheLimits"`
	Rollback             SystemSettingsRollback    `json:"rollback"`
	SearchLimits         SearchSettings            `json:"searchLimits"`
	ServicesLimits       ServicesLimits            `json:"servicesLimits"`
	Shutdown             SystemSettingsShutdown    `json:"shutdown"`
	System               SystemSettingsSystem      `json:"system"`
	TLS                  SystemSettingsTLS         `json:"tls"`
	UpgradeGroupSettings UpgradeGroupSettings      `json:"upgradeGroupSettings"`
	UpgradeSettings      UpgradeSettings           `json:"upgradeSettings"`
	Workers              SystemSettingsWorkers     `json:"workers"`
}

func (o *SystemSettings) GetAPI() SystemSettingsAPI {
	if o == nil {
		return SystemSettingsAPI{}
	}
	return o.API
}

func (o *SystemSettings) GetAuth() AuthConfig {
	if o == nil {
		return AuthConfig{}
	}
	return o.Auth
}

func (o *SystemSettings) GetBackups() SystemSettingsBackups {
	if o == nil {
		return SystemSettingsBackups{}
	}
	return o.Backups
}

func (o *SystemSettings) GetCustomLogo() SystemSettingsCustomLogo {
	if o == nil {
		return SystemSettingsCustomLogo{}
	}
	return o.CustomLogo
}

func (o *SystemSettings) GetDistributed() SystemSettingsDistributed {
	if o == nil {
		return SystemSettingsDistributed{}
	}
	return o.Distributed
}

func (o *SystemSettings) GetGit() GitSettings {
	if o == nil {
		return GitSettings{}
	}
	return o.Git
}

func (o *SystemSettings) GetJobLimits() JobSettings {
	if o == nil {
		return JobSettings{}
	}
	return o.JobLimits
}

func (o *SystemSettings) GetLimits() Limits {
	if o == nil {
		return Limits{}
	}
	return o.Limits
}

func (o *SystemSettings) GetProxy() SystemSettingsProxy {
	if o == nil {
		return SystemSettingsProxy{}
	}
	return o.Proxy
}

func (o *SystemSettings) GetRedisCacheLimits() RedisCacheLimits {
	if o == nil {
		return RedisCacheLimits{}
	}
	return o.RedisCacheLimits
}

func (o *SystemSettings) GetRollback() SystemSettingsRollback {
	if o == nil {
		return SystemSettingsRollback{}
	}
	return o.Rollback
}

func (o *SystemSettings) GetSearchLimits() SearchSettings {
	if o == nil {
		return SearchSettings{}
	}
	return o.SearchLimits
}

func (o *SystemSettings) GetServicesLimits() ServicesLimits {
	if o == nil {
		return ServicesLimits{}
	}
	return o.ServicesLimits
}

func (o *SystemSettings) GetShutdown() SystemSettingsShutdown {
	if o == nil {
		return SystemSettingsShutdown{}
	}
	return o.Shutdown
}

func (o *SystemSettings) GetSystem() SystemSettingsSystem {
	if o == nil {
		return SystemSettingsSystem{}
	}
	return o.System
}

func (o *SystemSettings) GetTLS() SystemSettingsTLS {
	if o == nil {
		return SystemSettingsTLS{}
	}
	return o.TLS
}

func (o *SystemSettings) GetUpgradeGroupSettings() UpgradeGroupSettings {
	if o == nil {
		return UpgradeGroupSettings{}
	}
	return o.UpgradeGroupSettings
}

func (o *SystemSettings) GetUpgradeSettings() UpgradeSettings {
	if o == nil {
		return UpgradeSettings{}
	}
	return o.UpgradeSettings
}

func (o *SystemSettings) GetWorkers() SystemSettingsWorkers {
	if o == nil {
		return SystemSettingsWorkers{}
	}
	return o.Workers
}