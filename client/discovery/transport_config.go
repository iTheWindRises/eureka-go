package discovery

import (
	"github.com/spf13/viper"
)

type TransportConfig interface {

	/* client会话周期性重连时间, 单位: 秒 */
	GetSessionedClientReconnectIntervalSeconds() int

	/* 重试 EurekaHttpClient ，请求失败的 Eureka-Server 隔离集合占比 Eureka-Server 全量集合占比
	，超过该比例，进行清空 */
	GetRetryableClientQuarantineRefreshPercentage() float64

	/* 异步解析 EndPoint 集群频率，单位：毫秒 */
	GetAsyncResolverRefreshIntervalMs() int

	/* 异步解析器预热解析 EndPoint 集群超时时间，单位：毫秒 */
	GetAsyncResolverWarmUpTimeoutMs() int

	/* 异步解析器线程池大小 */
	GetAsyncExecutorThreadPoolSize() int
}

type DefaultTransportConfig struct {
	// 命名空间
	namespace string
	// 配置文件对象
	configInstance *viper.Viper
}

func NewDefaultTransportConfig(parentNamespace string, configInstance *viper.Viper) *DefaultTransportConfig {
	namespace := parentNamespace + TRANSPORT_CONFIG_SUB_NAMESPACE + "."

	// 设置默认值
	configInstance.SetDefault(namespace+SESSION_RECONNECT_INTERVAL_KEY, SESSION_RECONNECT_INTERVAL)
	configInstance.SetDefault(namespace+QUARANTINE_REFRESH_PERCENTAGE_KEY, QUARANTINE_REFRESH_PERCENTAGE)
	configInstance.SetDefault(namespace+ASYNC_RESOLVER_REFRESH_INTERVAL_KEY, ASYNC_RESOLVER_REFRESH_INTERVAL)
	configInstance.SetDefault(namespace+ASYNC_RESOLVER_WARMUP_TIMEOUT_KEY, ASYNC_RESOLVER_WARMUP_TIMEOUT)
	configInstance.SetDefault(namespace+ASYNC_EXECUTOR_THREADPOOL_SIZE_KEY, ASYNC_EXECUTOR_THREADPOOL_SIZE)

	return &DefaultTransportConfig{namespace: namespace, configInstance: configInstance}
}

func (tc *DefaultTransportConfig) GetSessionedClientReconnectIntervalSeconds() int {
	return tc.configInstance.GetInt(tc.namespace + SESSION_RECONNECT_INTERVAL_KEY)
}

func (tc *DefaultTransportConfig) GetRetryableClientQuarantineRefreshPercentage() float64 {
	return tc.configInstance.GetFloat64(tc.namespace + QUARANTINE_REFRESH_PERCENTAGE_KEY)
}

func (tc *DefaultTransportConfig) GetAsyncResolverRefreshIntervalMs() int {
	return tc.configInstance.GetInt(tc.namespace + ASYNC_RESOLVER_REFRESH_INTERVAL_KEY)
}

func (tc *DefaultTransportConfig) GetAsyncResolverWarmUpTimeoutMs() int {
	return tc.configInstance.GetInt(tc.namespace + ASYNC_RESOLVER_WARMUP_TIMEOUT_KEY)
}

func (tc *DefaultTransportConfig) GetAsyncExecutorThreadPoolSize() int {
	return tc.configInstance.GetInt(tc.namespace + ASYNC_EXECUTOR_THREADPOOL_SIZE_KEY)
}
