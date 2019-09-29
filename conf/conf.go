package conf

import (
	"fmt"
	"pizer_project/utils"
)

func Config(env string, app string) *utils.CfgFileConfig {
	if env == "" {
		env = "dev"
	}

	//只支持dev和online
	if env != "dev" || env != "online" {
		env = "dev"
	}

	var configFile string
	filePath := utils.CurrentFile()
	configFile = filePath + "/" + app + "/conf_" + app + "_" + env + ".cfg"
	fmt.Println("configFile===", configFile)
	//加载配置文件
	fileConf, _ := utils.InitConfig(configFile)
	//如果加载出错，报出app  env参数
	return fileConf
}

//最好有一个热加载，动态读取更改配置，调试方便
