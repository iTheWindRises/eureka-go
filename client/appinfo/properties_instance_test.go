package appinfo

import (
	"fmt"
	"testing"
)

func TestNewPropertiesInstanceConfig(t *testing.T) {
	config := NewPropertiesInstanceConfig(DEFAULT_NAMESPACE)
	fmt.Println(config.InstanceId())
	fmt.Println(config.AppName())
	fmt.Println(config.AppGroupName())
	fmt.Println(config.Port())
	fmt.Println(config.EnabledOnInit())
	fmt.Println(config.LeaseRenewalIntervalInSeconds())
	fmt.Println(config.LeaseExpirationDurationInSeconds())
	fmt.Println(config.EnabledNonSecurePort())
	fmt.Println(config.MetaDataMap())
}
