package appinfo

import (
	"fmt"
	"testing"
)

func TestNewLeaseInfo(t *testing.T) {
	info := NewLeaseInfo(
		WithDurationInSecs(-1),
		WithRenewalIntervalInSecs(199))
	fmt.Println(info)

	info2 := NewLeaseInfo(
		WithDurationInSecs(2341),
		WithRenewalIntervalInSecs(424))
	fmt.Println(info2)
}
