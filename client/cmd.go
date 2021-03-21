package main

import (
	"eureka-go-client/appinfo"
	_ "eureka-go-client/config"
	"fmt"
)

func main() {
	config := appinfo.NewPropertiesInstanceConfig(appinfo.DEFAULT_NAMESPACE)
	fmt.Println(config.InstanceId())
}
