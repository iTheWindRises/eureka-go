package appinfo

import (
	_ "eureka-go-client/config"
	"eureka-go-client/util"
	"fmt"
	"testing"
)

func TestGetInstanceInfo(t *testing.T) {
	config := NewPropertiesInstanceConfig("eureka")
	info := GetInstanceInfo(config)
	fmt.Println(util.ModelToJsonStr(info))
	info = GetInstanceInfo(config)
	fmt.Println(util.ModelToJsonStr(info))
}
