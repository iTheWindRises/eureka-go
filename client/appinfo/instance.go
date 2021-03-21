package appinfo

import (
	"os"
)

/*
提供连接到server的实例的配置信息的能力, 如命名空间、实例id、实例名称等
应用实例:Application Consumer 和 Application Provider
*/
type InstanceConfig interface {
	// 获取实例id
	InstanceId() string

	// 获取实例名称
	AppName() string

	// 获取应用组名称
	AppGroupName() string

	// 设置实例是否连接成功后立即可用
	EnabledOnInit() bool

	// 获取该实例的hostname
	// refresh 获取时是否刷新
	HostName(refresh bool) string

	// 获取实例的ip
	IpAddress() string

	// 实例接收流量的端口
	Port() int

	// 是否启用非安全接口
	EnabledNonSecurePort() bool

	// 心跳间隔时间
	LeaseRenewalIntervalInSeconds() int

	// 心跳最长接收时间, 大于 LeaseRenewalIntervalInSeconds
	LeaseExpirationDurationInSeconds() int

	// 获取实例的元配置信息
	MetaDataMap() map[string]interface{}

	NameSpace() string

	HomeUrl() string

	StatusUrl() string

	HealthCheckUrl() string
}

/**
一个抽象的实例信息配置，默认情况下可以使用户快速入门。
用户只需重写一些方法即可在eureka服务器上注册其实例
*/
const (
	LEASE_EXPIRATION_DURATION_SECONDS = 90
	LEASE_RENEWAL_INTERVAL_SECONDS    = 30
	PORT                              = 9090
	INSTANCE_ENABLED_ON_INIT          = false
	NON_SECURE_PORT_ENABLED           = true
	DEFAULT_NAMESPACE                 = "eureka"
	CONFIG_FILE_NAME                  = "eureka-client"
	UNKNOWN                           = "unknown"
)

var defaultInstanceConfig *DefaultInstanceConfig = &DefaultInstanceConfig{}

type DefaultInstanceConfig struct {
	instanceId   string
	appName      string
	appGroupName string
	metaDataMap  map[string]interface{}
	hostName     string
}

func (dic *DefaultInstanceConfig) InstanceId() string {
	return dic.instanceId
}

func (dic *DefaultInstanceConfig) AppName() string {
	return dic.appName
}

func (dic *DefaultInstanceConfig) AppGroupName() string {
	return dic.appGroupName
}

func (dic *DefaultInstanceConfig) EnabledOnInit() bool {
	return INSTANCE_ENABLED_ON_INIT
}

func (dic *DefaultInstanceConfig) HostName(refresh bool) string {
	if refresh || dic.hostName == "" {
		dic.hostName, _ = os.Hostname()
	}
	return dic.hostName
}

func (dic *DefaultInstanceConfig) IpAddress() string {
	return ""
}

func (dic *DefaultInstanceConfig) Port() int {
	return PORT
}

func (dic *DefaultInstanceConfig) EnabledNonSecurePort() bool {
	return NON_SECURE_PORT_ENABLED
}

func (dic *DefaultInstanceConfig) LeaseRenewalIntervalInSeconds() int {
	return LEASE_RENEWAL_INTERVAL_SECONDS
}

func (dic *DefaultInstanceConfig) LeaseExpirationDurationInSeconds() int {
	return LEASE_EXPIRATION_DURATION_SECONDS
}

func (dic *DefaultInstanceConfig) MetaDataMap() map[string]interface{} {
	return dic.metaDataMap
}

func (dic *DefaultInstanceConfig) NameSpace() string {
	return DEFAULT_NAMESPACE
}

func (dic *DefaultInstanceConfig) HomeUrl() string {
	return ""
}

func (dic *DefaultInstanceConfig) StatusUrl() string {
	return ""
}

func (dic *DefaultInstanceConfig) HealthCheckUrl() string {
	return ""
}
