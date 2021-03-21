package discovery

import (
	"eureka-go-client/appinfo"
	"eureka-go-client/config"
	"fmt"
	"github.com/spf13/viper"
	"strings"
	"sync"
)

type ClientConfig interface {
	// 使用dns获取server url相关
	/* 是否使用dns获取serverurl地址 */
	ShouldUseDnsForFetchingServiceUrls() bool

	/* server的dns名 */
	GetEurekaServerDNSName() string

	/* server的端口 */
	GetEurekaServerPort() string

	/* server的url context */
	GetEurekaServerURLContext() string

	/* 轮训获取server地址变化更新频率, 单位秒. 默认300s */
	GetEurekaServiceUrlPollIntervalSeconds() int

	// 直接配合 Eureka-Server URL 相关
	/* Server 的 URL 集合 */
	GetEurekaServerServiceUrls() []string

	// 发现：从 Eureka-Server 获取注册信息相关
	/* 是否从 Eureka-Server 拉取注册信息 */
	ShouldFetchRegistry() bool

	/* 从 Eureka-Server 拉取注册信息频率，单位：秒。默认：30 秒。 */
	GetRegistryFetchIntervalSeconds() int

	/* 是否过滤，只获取状态为开启( Up )的应用实例集合 */
	ShouldFilterOnlyUpInstances() bool

	/* 注册信息缓存刷新线程池大小 */
	GetCacheRefreshExecutorThreadPoolSize() int

	/* 注册信息缓存刷新执行超时后的延迟重试的时间 */
	GetCacheRefreshExecutorExponentialBackOffBound() int

	// 注册：向 Eureka-Server 注册自身服务
	/* 是否向 Eureka-Server 注册自身服务 */
	ShouldRegisterWithEureka() bool

	/* 是否向 Eureka-Server 取消注册自身服务，当进程关闭时 */
	ShouldUnregisterOnShutdown() bool

	/* 向 Eureka-Server 同步应用实例信息变化频率，单位：秒 */
	GetInstanceInfoReplicationIntervalSeconds() int

	/* 向 Eureka-Server 同步应用信息变化初始化延迟，单位：秒 */
	GetInitialInstanceInfoReplicationIntervalSeconds() int

	/* 获取备份注册中心实现类。当 Eureka-Client 启动时，无法从 Eureka-Server 读取注册信息（可能挂了）
	，从备份注册中心读取注册信息。目前 Eureka-Client 未提供合适的实现。*/
	GetBackupRegistryImpl() string

	/* 心跳执行线程池大小 */
	GetHeartbeatExecutorThreadPoolSize() int

	/* 心跳执行超时后的延迟重试的时间 */
	GetHeartbeatExecutorExponentialBackOffBound() int
}

type DefaultClientConfig struct {
	// 命名空间
	namespace string
	// 配置文件对象
	configInstance *viper.Viper
	// HTTP传输配置
	transportConfig TransportConfig
}

func newDefaultClientConfig(namespace string) ClientConfig {
	if namespace == "" {
		namespace = appinfo.DEFAULT_NAMESPACE
	}
	if !strings.HasSuffix(namespace, ".") {
		namespace = fmt.Sprintf("%s.", namespace)
	}
	configInstance := config.InitConfig(namespace, appinfo.CONFIG_FILE_NAME)

	// 默认配置
	configInstance.SetDefault(namespace+INITIAL_REGISTRATION_REPLICATION_DELAY_KEY, 40)
	configInstance.SetDefault(namespace+HEARTBEAT_THREADPOOL_SIZE_KEY, 5)
	configInstance.SetDefault(namespace+HEARTBEAT_BACKOFF_BOUND_KEY, DEFAULT_EXECUTOR_THREAD_POOL_BACKOFF_BOUND)

	configInstance.SetDefault(namespace+REGISTRATION_REPLICATION_INTERVAL_KEY, 30)

	configInstance.SetDefault(namespace+SHOULD_UNREGISTER_ON_SHUTDOWN_KEY, false)
	configInstance.SetDefault(namespace+SHOULD_USE_DNS_KEY, false)
	configInstance.SetDefault(namespace+EUREKA_SERVER_URL_POLL_INTERVAL_KEY, 300_000)
	configInstance.SetDefault(namespace+FETCH_REGISTRY_ENABLED_KEY, true)
	configInstance.SetDefault(namespace+REGISTRATION_ENABLED_KEY, true)

	configInstance.SetDefault(namespace+REGISTRY_REFRESH_INTERVAL_KEY, 30)
	configInstance.SetDefault(namespace+SHOULD_FILTER_ONLY_UP_INSTANCES_KEY, true)
	configInstance.SetDefault(namespace+CACHEREFRESH_THREADPOOL_SIZE_KEY, DEFAULT_EXECUTOR_THREAD_POOL_SIZE)
	configInstance.SetDefault(namespace+CACHEREFRESH_BACKOFF_BOUND_KEY, DEFAULT_EXECUTOR_THREAD_POOL_BACKOFF_BOUND)

	return &DefaultClientConfig{
		namespace:       namespace,
		configInstance:  configInstance,
		transportConfig: NewDefaultTransportConfig(namespace, configInstance),
	}
}

var once sync.Once
var clientConfig ClientConfig

func GetClientConfig(namespace string) ClientConfig {
	once.Do(func() {
		clientConfig = newDefaultClientConfig(namespace)
	})
	return clientConfig
}

func (dcc *DefaultClientConfig) ShouldUseDnsForFetchingServiceUrls() bool {
	return dcc.configInstance.GetBool(dcc.namespace + SHOULD_USE_DNS_KEY)
}

func (dcc *DefaultClientConfig) GetEurekaServerDNSName() string {
	DNSName := dcc.configInstance.GetString(dcc.namespace + EUREKA_SERVER_DNS_NAME_KEY)
	if DNSName == "" {
		DNSName = dcc.configInstance.GetString(dcc.namespace + EUREKA_SERVER_FALLBACK_DNS_NAME_KEY)
	}
	return DNSName
}

func (dcc *DefaultClientConfig) GetEurekaServerPort() string {
	port := dcc.configInstance.GetString(dcc.namespace + EUREKA_SERVER_PORT_KEY)
	if port == "" {
		port = dcc.configInstance.GetString(dcc.namespace + EUREKA_SERVER_FALLBACK_PORT_KEY)
	}
	return port
}

func (dcc *DefaultClientConfig) GetEurekaServerURLContext() string {
	context := dcc.configInstance.GetString(dcc.namespace + EUREKA_SERVER_URL_CONTEXT_KEY)
	if context == "" {
		context = dcc.configInstance.GetString(dcc.namespace + EUREKA_SERVER_FALLBACK_URL_CONTEXT_KEY)
	}
	return context
}

func (dcc *DefaultClientConfig) GetEurekaServiceUrlPollIntervalSeconds() int {
	return dcc.configInstance.GetInt(dcc.namespace+EUREKA_SERVER_URL_POLL_INTERVAL_KEY) / 1000
}

func (dcc *DefaultClientConfig) GetEurekaServerServiceUrls() []string {
	return dcc.configInstance.GetStringSlice(dcc.namespace + CONFIG_EUREKA_SERVER_SERVICE_URL_PREFIX)
}

func (dcc *DefaultClientConfig) ShouldFetchRegistry() bool {
	return dcc.configInstance.GetBool(dcc.namespace + FETCH_REGISTRY_ENABLED_KEY)
}

func (dcc *DefaultClientConfig) GetRegistryFetchIntervalSeconds() int {
	return dcc.configInstance.GetInt(dcc.namespace + REGISTRY_REFRESH_INTERVAL_KEY)
}

func (dcc *DefaultClientConfig) ShouldFilterOnlyUpInstances() bool {
	return dcc.configInstance.GetBool(dcc.namespace + SHOULD_FILTER_ONLY_UP_INSTANCES_KEY)
}

func (dcc *DefaultClientConfig) GetCacheRefreshExecutorThreadPoolSize() int {
	return dcc.configInstance.GetInt(dcc.namespace + CACHEREFRESH_THREADPOOL_SIZE_KEY)
}

func (dcc *DefaultClientConfig) GetCacheRefreshExecutorExponentialBackOffBound() int {
	return dcc.configInstance.GetInt(dcc.namespace + CACHEREFRESH_BACKOFF_BOUND_KEY)
}

func (dcc *DefaultClientConfig) ShouldRegisterWithEureka() bool {
	return dcc.configInstance.GetBool(dcc.namespace + REGISTRATION_ENABLED_KEY)
}

func (dcc *DefaultClientConfig) ShouldUnregisterOnShutdown() bool {
	return dcc.configInstance.GetBool(dcc.namespace + SHOULD_UNREGISTER_ON_SHUTDOWN_KEY)
}

func (dcc *DefaultClientConfig) GetInstanceInfoReplicationIntervalSeconds() int {
	return dcc.configInstance.GetInt(dcc.namespace + REGISTRATION_REPLICATION_INTERVAL_KEY)
}

func (dcc *DefaultClientConfig) GetInitialInstanceInfoReplicationIntervalSeconds() int {
	return dcc.configInstance.GetInt(dcc.namespace + INITIAL_REGISTRATION_REPLICATION_DELAY_KEY)
}

func (dcc *DefaultClientConfig) GetBackupRegistryImpl() string {
	return dcc.configInstance.GetString(dcc.namespace + BACKUP_REGISTRY_CLASSNAME_KEY)
}

func (dcc *DefaultClientConfig) GetHeartbeatExecutorThreadPoolSize() int {
	return dcc.configInstance.GetInt(dcc.namespace + HEARTBEAT_THREADPOOL_SIZE_KEY)
}

func (dcc *DefaultClientConfig) GetHeartbeatExecutorExponentialBackOffBound() int {
	return dcc.configInstance.GetInt(dcc.namespace + HEARTBEAT_BACKOFF_BOUND_KEY)
}
