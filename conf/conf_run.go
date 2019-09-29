package conf

import "pizer_project/utils"

var baseConf *utils.CfgFileConfig

func Run() {
	//加载db配置文件
	baseConf = Config("env", "base")
}

func GetBaseConf() *utils.CfgFileConfig {
	return baseConf
}
