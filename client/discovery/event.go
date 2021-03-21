package discovery

import "eureka-go-client/appinfo"

type StatusChangeEvent struct {
	Current  appinfo.InstanceStatus
	Previous appinfo.InstanceStatus
}

func (event *StatusChangeEvent) IsUp() bool {
	return event.Current == appinfo.UP
}

func (event *StatusChangeEvent) GetStatus() appinfo.InstanceStatus {
	return event.Current
}

func (event *StatusChangeEvent) GetPreviousStatus() appinfo.InstanceStatus {
	return event.Previous
}
