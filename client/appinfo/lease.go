package appinfo

/**
实例的租约信息, 设置的持续时间，将实例从其视图中删除。 租约还跟踪上次续约。
*/
type LeaseInfo struct {

	// Client settings
	// 续约间隔
	RenewalIntervalInSecs int `json:"renewalIntervalInSecs"`
	// 最长持续时间
	DurationInSecs int `json:"durationInSecs"`

	// Server populated
	// 注册时间戳
	RegistrationTimestamp int64 `json:"registrationTimestamp"`
	// 上次续约时间戳
	LastRenewalTimestamp int64 `json:"lastRenewalTimestamp"`
	// 逐出时间戳
	EvictionTimestamp int64 `json:"evictionTimestamp"`
	// 服务up时间戳
	ServiceUpTimestamp int64 `json:"serviceUpTimestamp"`
}

type OptionLeaseInfo func(*LeaseInfo)

func NewLeaseInfo(opts ...OptionLeaseInfo) *LeaseInfo {

	li := &LeaseInfo{
		RenewalIntervalInSecs: LEASE_RENEWAL_INTERVAL_SECONDS,
		DurationInSecs:        LEASE_EXPIRATION_DURATION_SECONDS,
		RegistrationTimestamp: 0,
		LastRenewalTimestamp:  0,
		EvictionTimestamp:     0,
		ServiceUpTimestamp:    0,
	}

	for _, opt := range opts {
		opt(li)
	}
	return li
}

func WithRegistrationTimestamp(registrationTimestamp int64) OptionLeaseInfo {
	return func(li *LeaseInfo) {
		li.RegistrationTimestamp = registrationTimestamp
	}
}

func WithLastRenewalTimestamp(lastRenewalTimestamp int64) OptionLeaseInfo {
	return func(li *LeaseInfo) {
		li.LastRenewalTimestamp = lastRenewalTimestamp
	}
}

func WithEvictionTimestamp(evictionTimestamp int64) OptionLeaseInfo {
	return func(li *LeaseInfo) {
		li.EvictionTimestamp = evictionTimestamp
	}
}

func WithServiceUpTimestamp(serviceUpTimestamp int64) OptionLeaseInfo {
	return func(li *LeaseInfo) {
		li.ServiceUpTimestamp = serviceUpTimestamp
	}
}

func WithRenewalIntervalInSecs(renewalIntervalInSecs int) OptionLeaseInfo {
	return func(li *LeaseInfo) {
		li.RenewalIntervalInSecs = renewalIntervalInSecs
	}
}

func WithDurationInSecs(durationInSecs int) OptionLeaseInfo {
	return func(li *LeaseInfo) {
		li.DurationInSecs = durationInSecs
	}
}
