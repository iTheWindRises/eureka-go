package appinfo

import (
	log "github.com/sirupsen/logrus"
	"strings"
	"sync"
	"time"
)

type InstanceInfo struct {
	LeaseInfo                    *LeaseInfo             `json:"leaseInfo"`
	Namespace                    string                 `json:"namespace"`
	InstanceId                   string                 `json:"instanceId"`
	AppName                      string                 `json:"app"`
	AppGroupName                 string                 `json:"appGroupName"`
	IpAddr                       string                 `json:"ipAddr"`
	Port                         int32                  `json:"port"`
	HostName                     string                 `json:"hostName"`
	EnabledNonSecurePort         bool                   `json:"enabledNonSecurePort"`
	Status                       InstanceStatus         `json:"status"`
	Metadata                     map[string]interface{} `json:"metadata"`
	HomeUrl                      string                 `json:"homeUrl"`
	StatusUrl                    string                 `json:"statusUrl"`
	HealthCheckUrl               string                 `json:"healthCheckUrl"`
	IsCoordinatingDiscoverServer bool                   `json:"isCoordinatingDiscoveryServer"`
	// 最后更新时间戳
	LastUpdatedTimestamp int64 `json:"lastUpdatedTimestamp"`
	// 最后xx时间戳
	LastDirtyTimestamp int64 `json:"lastDirtyTimestamp"`
	// 类型
	ActionType          string `json:"actionType"`
	IsInstanceInfoDirty bool   `json:"isInstanceInfoDirty"`
	mutex               sync.Mutex
}

type InstanceStatus int8

const (
	_ InstanceStatus = iota + 1
	UP
	DOWN
	STARTING
	OUT_OF_SERVICE
	UNKNOWN_STATUS
)

var once sync.Once
var info *InstanceInfo

func (info *InstanceInfo) SetIsDirty() {
	info.mutex.Lock()
	info.LastDirtyTimestamp = time.Now().Unix()
	info.IsInstanceInfoDirty = true
	info.mutex.Unlock()
}

func GetInstanceInfo(config InstanceConfig) *InstanceInfo {
	once.Do(func() {
		// 创建租约信息
		leaseInfo := NewLeaseInfo(
			WithRenewalIntervalInSecs(config.LeaseRenewalIntervalInSeconds()),
			WithDurationInSecs(config.LeaseExpirationDurationInSeconds()))

		// 设置instanceId
		instanceId := config.InstanceId()
		if instanceId == "" {
			instanceId = config.HostName(false)
		}

		// 获取主机名
		defaultAddress := config.HostName(false)

		info = newInstanceInfo(
			WithLeaseInfo(leaseInfo),
			WithNamespace(config.NameSpace()),
			WithInstanceId(instanceId),
			WithAppName(config.AppName()),
			WithAppGroupName(config.AppGroupName()),
			WithIpAddr(config.IpAddress()),
			WithHostName(defaultAddress),
			WithPort(int32(config.Port())),
			WithEnabledNonSecurePort(config.EnabledNonSecurePort()),
			WithHomeUrl(config.HomeUrl()),
			WithStatusUrl(config.StatusUrl()),
			WithHealthCheckUrl(config.HealthCheckUrl()),
			WithStatus(config.EnabledOnInit()),
			WithMetadata(config.MetaDataMap()),
		)
	})
	return info
}

type OptionInstanceInfo func(info *InstanceInfo)

func newInstanceInfo(opts ...OptionInstanceInfo) *InstanceInfo {
	info := &InstanceInfo{
		LeaseInfo:                    nil,
		Namespace:                    "",
		InstanceId:                   "",
		AppName:                      "",
		AppGroupName:                 "",
		IpAddr:                       "",
		Port:                         0,
		HostName:                     "",
		EnabledNonSecurePort:         false,
		IsCoordinatingDiscoverServer: false,
		Status:                       UP,
		Metadata:                     nil,
		HomeUrl:                      "",
		StatusUrl:                    "",
		HealthCheckUrl:               "",
		LastUpdatedTimestamp:         0,
		LastDirtyTimestamp:           0,
		ActionType:                   "",
	}
	for _, opt := range opts {
		opt(info)
	}
	return info
}

func WithNamespace(namespace string) OptionInstanceInfo {
	return func(info *InstanceInfo) {
		info.Namespace = namespace
	}
}

func WithInstanceId(instanceId string) OptionInstanceInfo {
	return func(info *InstanceInfo) {
		info.InstanceId = instanceId
	}
}

func WithAppName(appName string) OptionInstanceInfo {
	return func(info *InstanceInfo) {
		info.AppName = strings.ToLower(appName)
	}
}

func WithAppGroupName(appGroupName string) OptionInstanceInfo {
	return func(info *InstanceInfo) {
		info.AppGroupName = strings.ToLower(appGroupName)
	}
}

func WithIpAddr(ipAddr string) OptionInstanceInfo {
	return func(info *InstanceInfo) {
		info.IpAddr = ipAddr
	}
}

func WithHostName(hostName string) OptionInstanceInfo {
	return func(info *InstanceInfo) {
		info.HostName = hostName
		if hostName == "" {
			log.Warn("传入的主机名为空，未设置")
		}
	}
}

func WithPort(port int32) OptionInstanceInfo {
	return func(info *InstanceInfo) {
		info.Port = port
	}
}

func WithEnabledNonSecurePort(enabledNonSecurePort bool) OptionInstanceInfo {
	return func(info *InstanceInfo) {
		info.EnabledNonSecurePort = enabledNonSecurePort
	}
}
func WithHomeUrl(url string) OptionInstanceInfo {
	return func(info *InstanceInfo) {
		info.HomeUrl = url
	}
}

func WithStatusUrl(url string) OptionInstanceInfo {
	return func(info *InstanceInfo) {
		info.StatusUrl = url
	}
}
func WithHealthCheckUrl(url string) OptionInstanceInfo {
	return func(info *InstanceInfo) {
		info.HealthCheckUrl = url
	}
}

func WithLeaseInfo(lease *LeaseInfo) OptionInstanceInfo {
	return func(info *InstanceInfo) {
		info.LeaseInfo = lease
	}
}

func WithStatus(onInit bool) OptionInstanceInfo {
	return func(info *InstanceInfo) {
		if !onInit {
			info.Status = STARTING
			log.Info("实例初始化状态被设置为: STARTING")

			return
		}
		log.Info("实例初始化状态被设置为: UP")
	}
}

func WithMetadata(metaMap map[string]interface{}) OptionInstanceInfo {
	return func(info *InstanceInfo) {
		info.Metadata = metaMap
	}
}
