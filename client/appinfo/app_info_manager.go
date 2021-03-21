package appinfo

import (
	"context"
	"eureka-go-client/discovery"
	"sync"
)

/**
应用信息管理器
*/
type ApplicationInfoManager struct {
	context.Context
	Listeners            map[string]StatusChangeListener
	InstanceInfo         *InstanceInfo
	Config               InstanceConfig
	InstanceStatusMapper InstanceStatusMapper
	mutex                sync.Mutex
}

func NewApplicationInfoManager(ctx context.Context, config InstanceConfig, instanceInfo *InstanceInfo, mapper InstanceStatusMapper) *ApplicationInfoManager {
	return &ApplicationInfoManager{
		Context:              ctx,
		Listeners:            nil,
		InstanceInfo:         instanceInfo,
		Config:               config,
		InstanceStatusMapper: mapper,
	}
}

func (manager *ApplicationInfoManager) initComponent(config InstanceConfig) {
	manager.Config = config
	manager.InstanceInfo = GetInstanceInfo(config)
}

/**
注册用户特定的实例元数据。
*/
func (manager *ApplicationInfoManager) RegisterAppMetadata(metadata map[string]interface{}) {
	manager.mutex.Lock()
	if metadata != nil {
		for k, v := range metadata {
			manager.InstanceInfo.Metadata[k] = v
		}
	}
	manager.mutex.Unlock()

}

/**
设置实例状态
*/
func (manager *ApplicationInfoManager) SetInstanceStatus(status InstanceStatus) {
	manager.mutex.Lock()
	next := manager.InstanceStatusMapper.GetStatus(status)
	if next == nil {
		return
	}
	prev := manager.InstanceInfo.Status
	manager.InstanceInfo.Status = *next

	if prev != 0 {
		for _, listener := range manager.Listeners {
			listener.Notify(&discovery.StatusChangeEvent{Previous: prev, Current: *next})
		}
	}
	manager.mutex.Unlock()
}

/**
刷新实例租约信息
*/
func (manager *ApplicationInfoManager) RefreshLeaseInfoIfRequired() {
	leaseInfo := manager.InstanceInfo.LeaseInfo
	if leaseInfo == nil {
		return
	}

	curDuration := manager.Config.LeaseExpirationDurationInSeconds()
	curRenewal := manager.Config.LeaseRenewalIntervalInSeconds()

	if leaseInfo.DurationInSecs != curDuration || leaseInfo.RenewalIntervalInSecs != curRenewal {

		manager.InstanceInfo.LeaseInfo = NewLeaseInfo(
			WithDurationInSecs(curDuration),
			WithRenewalIntervalInSecs(curRenewal))
		manager.InstanceInfo.SetIsDirty()
	}
}

/**
状态变化的监听器
*/
type StatusChangeListener interface {
	GetId() string
	Notify(event *discovery.StatusChangeEvent)
}

type InstanceStatusMapper interface {
	/**
	根据给定的状态返回新状态, 无则返回nil
	*/
	GetStatus(status InstanceStatus) *InstanceStatus
}
