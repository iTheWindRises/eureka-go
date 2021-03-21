package discovery

const (
	CLIENT_REGION_FALLBACK_KEY                = "eureka.region"
	CLIENT_REGION_KEY                         = "region"
	REGISTRATION_ENABLED_KEY                  = "registration.enabled"
	FETCH_REGISTRY_ENABLED_KEY                = "shouldFetchRegistry"
	SHOULD_ENFORCE_FETCH_REGISTRY_AT_INIT_KEY = "shouldEnforceFetchRegistryAtInit"

	REGISTRY_REFRESH_INTERVAL_KEY              = "client.refresh.interval"
	REGISTRATION_REPLICATION_INTERVAL_KEY      = "appinfo.replicate.interval"
	INITIAL_REGISTRATION_REPLICATION_DELAY_KEY = "appinfo.initial.replicate.time"
	HEARTBEAT_THREADPOOL_SIZE_KEY              = "client.heartbeat.threadPoolSize"
	HEARTBEAT_BACKOFF_BOUND_KEY                = "client.heartbeat.exponentialBackOffBound"
	CACHEREFRESH_THREADPOOL_SIZE_KEY           = "client.cacheRefresh.threadPoolSize"
	CACHEREFRESH_BACKOFF_BOUND_KEY             = "client.cacheRefresh.exponentialBackOffBound"

	SHOULD_UNREGISTER_ON_SHUTDOWN_KEY   = "shouldUnregisterOnShutdown"
	SHOULD_ONDEMAND_UPDATE_STATUS_KEY   = "shouldOnDemandUpdateStatusChange"
	SHOULD_ENFORCE_REGISTRATION_AT_INIT = "shouldEnforceRegistrationAtInit"
	SHOULD_DISABLE_DELTA_KEY            = "disableDelta"
	SHOULD_FETCH_REMOTE_REGION_KEY      = "fetchRemoteRegionsRegistry"
	SHOULD_FILTER_ONLY_UP_INSTANCES_KEY = "shouldFilterOnlyUpInstances"
	FETCH_SINGLE_VIP_ONLY_KEY           = "registryRefreshSingleVipAddress"
	CLIENT_ENCODER_NAME_KEY             = "encoderName"
	CLIENT_DECODER_NAME_KEY             = "decoderName"
	CLIENT_DATA_ACCEPT_KEY              = "clientDataAccept"

	BACKUP_REGISTRY_CLASSNAME_KEY = "backupregistry"

	SHOULD_PREFER_SAME_ZONE_SERVER_KEY = "preferSameZone"
	SHOULD_ALLOW_REDIRECTS_KEY         = "allowRedirects"
	SHOULD_USE_DNS_KEY                 = "shouldUseDns"

	EUREKA_SERVER_URL_POLL_INTERVAL_KEY    = "serviceUrlPollIntervalMs"
	EUREKA_SERVER_URL_CONTEXT_KEY          = "eurekaServer.context"
	EUREKA_SERVER_FALLBACK_URL_CONTEXT_KEY = "context"
	EUREKA_SERVER_PORT_KEY                 = "eurekaServer.port"
	EUREKA_SERVER_FALLBACK_PORT_KEY        = "port"
	EUREKA_SERVER_DNS_NAME_KEY             = "eurekaServer.domainName"
	EUREKA_SERVER_FALLBACK_DNS_NAME_KEY    = "domainName"

	EUREKA_SERVER_PROXY_HOST_KEY     = "eurekaServer.proxyHost"
	EUREKA_SERVER_PROXY_PORT_KEY     = "eurekaServer.proxyPort"
	EUREKA_SERVER_PROXY_USERNAME_KEY = "eurekaServer.proxyUserName"
	EUREKA_SERVER_PROXY_PASSWORD_KEY = "eurekaServer.proxyPassword"

	EUREKA_SERVER_GZIP_CONTENT_KEY             = "eurekaServer.gzipContent"
	EUREKA_SERVER_READ_TIMEOUT_KEY             = "eurekaServer.readTimeout"
	EUREKA_SERVER_CONNECT_TIMEOUT_KEY          = "eurekaServer.connectTimeout"
	EUREKA_SERVER_MAX_CONNECTIONS_KEY          = "eurekaServer.maxTotalConnections"
	EUREKA_SERVER_MAX_CONNECTIONS_PER_HOST_KEY = "eurekaServer.maxConnectionsPerHost"
	// yeah the case on eurekaserver is different, backwards compatibility requirements :(
	EUREKA_SERVER_CONNECTION_IDLE_TIMEOUT_KEY = "eurekaserver.connectionIdleTimeoutInSeconds"

	SHOULD_LOG_DELTA_DIFF_KEY = "printDeltaFullDiff"

	CONFIG_DOLLAR_REPLACEMENT_KEY      = "dollarReplacement"
	CONFIG_ESCAPE_CHAR_REPLACEMENT_KEY = "escapeCharReplacement"

	// additional namespaces
	CONFIG_EXPERIMENTAL_PREFIX              = "experimental"
	CONFIG_AVAILABILITY_ZONE_PREFIX         = "availabilityZones"
	CONFIG_EUREKA_SERVER_SERVICE_URL_PREFIX = "serviceUrls"

	CONFIG_DOLLAR_REPLACEMENT      = "_-"
	CONFIG_ESCAPE_CHAR_REPLACEMENT = "__"

	DEFAULT_CLIENT_REGION = "us-east-1"

	DEFAULT_EXECUTOR_THREAD_POOL_SIZE          = 5
	DEFAULT_EXECUTOR_THREAD_POOL_BACKOFF_BOUND = 10

	// Transport 配置
	SESSION_RECONNECT_INTERVAL_KEY      = "sessionedClientReconnectIntervalSeconds"
	QUARANTINE_REFRESH_PERCENTAGE_KEY   = "retryableClientQuarantineRefreshPercentage"
	ASYNC_RESOLVER_REFRESH_INTERVAL_KEY = "asyncResolverRefreshIntervalMs"
	ASYNC_RESOLVER_WARMUP_TIMEOUT_KEY   = "asyncResolverWarmupTimeoutMs"
	ASYNC_EXECUTOR_THREADPOOL_SIZE_KEY  = "asyncExecutorThreadPoolSize"

	TRANSPORT_CONFIG_SUB_NAMESPACE = "transport"

	SESSION_RECONNECT_INTERVAL      = 20 * 60
	QUARANTINE_REFRESH_PERCENTAGE   = 0.66
	ASYNC_RESOLVER_REFRESH_INTERVAL = 5 * 60 * 1000
	ASYNC_RESOLVER_WARMUP_TIMEOUT   = 5000
	ASYNC_EXECUTOR_THREADPOOL_SIZE  = 5
)
