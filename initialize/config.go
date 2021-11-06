package initialize

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/jettjia/go-micro-frame-cli/global"
)

func InitConfig() {
	configFileName := fmt.Sprintf("config.yaml")

	// 读取文件配置内容
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	// 把内容设置到全局变量的 ServerConfig中
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}
}
