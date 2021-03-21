package appinfo

import (
	"eureka-go-client/config"
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

// 基于配置文件的应用实例配置
type PropertiesInstanceConfig struct {
	namespace      string
	configInstance *viper.Viper
}

func NewPropertiesInstanceConfig(namespace string) *PropertiesInstanceConfig {
	if namespace == "" {
		namespace = DEFAULT_NAMESPACE
	}
	if !strings.HasSuffix(namespace, ".") {
		namespace = fmt.Sprintf("%s.", namespace)
	}

	return &PropertiesInstanceConfig{namespace: namespace, configInstance: initConfig(namespace, CONFIG_FILE_NAME)}
}

func (pic *PropertiesInstanceConfig) InstanceId() string {
	return pic.configInstance.GetString(fmt.Sprintf("%s%s", pic.namespace, INSTANCE_ID_KEY))
}

func (pic *PropertiesInstanceConfig) AppName() string {
	return pic.configInstance.GetString(fmt.Sprintf("%s%s", pic.namespace, APP_NAME_KEY))
}

func (pic *PropertiesInstanceConfig) AppGroupName() string {
	return pic.configInstance.GetString(fmt.Sprintf("%s%s", pic.namespace, APP_GROUP_KEY))
}

func (pic *PropertiesInstanceConfig) EnabledOnInit() bool {
	return pic.configInstance.GetBool(fmt.Sprintf("%s%s", pic.namespace, TRAFFIC_ENABLED_ON_INIT_KEY))
}

func (pic *PropertiesInstanceConfig) HostName(refresh bool) string {
	return defaultInstanceConfig.HostName(refresh)
}

func (pic *PropertiesInstanceConfig) IpAddress() string {
	return defaultInstanceConfig.IpAddress()
}

func (pic *PropertiesInstanceConfig) Port() int {
	return pic.configInstance.GetInt(fmt.Sprintf("%s%s", pic.namespace, PORT_KEY))
}

func (pic *PropertiesInstanceConfig) EnabledNonSecurePort() bool {
	return pic.configInstance.GetBool(fmt.Sprintf("%s%s", pic.namespace, PORT_ENABLED_KEY))
}

func (pic *PropertiesInstanceConfig) LeaseRenewalIntervalInSeconds() int {
	return pic.configInstance.GetInt(fmt.Sprintf("%s%s", pic.namespace, LEASE_RENEWAL_INTERVAL_KEY))
}

func (pic *PropertiesInstanceConfig) LeaseExpirationDurationInSeconds() int {
	return pic.configInstance.GetInt(fmt.Sprintf("%s%s", pic.namespace, LEASE_EXPIRATION_DURATION_KEY))
}

func (pic *PropertiesInstanceConfig) MetaDataMap() map[string]interface{} {
	return pic.configInstance.GetStringMap(pic.namespace + INSTANCE_METADATA_PREFIX)
}

func (pic *PropertiesInstanceConfig) NameSpace() string {
	return pic.namespace
}

func (pic *PropertiesInstanceConfig) HomeUrl() string {
	return pic.configInstance.GetString(pic.namespace + HOME_PAGE_URL_KEY)
}

func (pic *PropertiesInstanceConfig) StatusUrl() string {
	return pic.configInstance.GetString(pic.namespace + STATUS_PAGE_URL_KEY)
}

func (pic *PropertiesInstanceConfig) HealthCheckUrl() string {
	return pic.configInstance.GetString(pic.namespace + HEALTHCHECK_URL_KEY)
}

func initConfig(namespace, configName string) *viper.Viper {
	configInstance := config.InitConfig(namespace, configName)

	// 设置默认值
	configInstance.SetDefault(namespace+LEASE_RENEWAL_INTERVAL_KEY, LEASE_RENEWAL_INTERVAL_SECONDS)
	configInstance.SetDefault(namespace+LEASE_EXPIRATION_DURATION_KEY, LEASE_EXPIRATION_DURATION_SECONDS)
	configInstance.SetDefault(namespace+PORT_KEY, PORT)
	configInstance.SetDefault(namespace+PORT_ENABLED_KEY, NON_SECURE_PORT_ENABLED)
	configInstance.SetDefault(namespace+TRAFFIC_ENABLED_ON_INIT_KEY, INSTANCE_ENABLED_ON_INIT)
	return configInstance
}
