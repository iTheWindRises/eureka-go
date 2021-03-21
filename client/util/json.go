package util

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func ModelToJsonStr(model interface{}) string {
	if bytes, err := json.Marshal(model); err != nil {
		log.Error("json转换异常: 对象转json字符串失败", err)
	} else {
		return string(bytes)
	}
	return ""
}
