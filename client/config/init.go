package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		TimestampFormat:           "2006-01-02 15:04:05", //时间格式
		FullTimestamp:             true,
		DisableLevelTruncation:    true,
	})
	//log.SetReportCaller(true)
}

func InitConfig(namespace, configName string) *viper.Viper {
	configInstance := viper.New()

	// 配置文件名称
	configInstance.SetConfigName(configName)
	configInstance.SetConfigType("toml")

	// 配置文件查找路径
	configInstance.AddConfigPath("/etc/eureka-go/")
	configInstance.AddConfigPath("$HOME/.eureka-go/")
	configInstance.AddConfigPath(".")

	if err := configInstance.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Warn("未发现配置文件")
		} else {
			// Config file was found but another error was produced
			log.Error("配置文件解析失败", err)
		}
	}

	//// watch配置
	//viper.WatchConfig()
	//viper.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("Config file changed:", e.Name)
	//})

	//  配置文件环境 TODO
	// 将配置文件加载到环境变量 TODO
	return configInstance
}
