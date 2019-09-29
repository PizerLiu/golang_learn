package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type CfgFileConfig struct {
	configMap map[string]map[string]string
}

// 初始化配置
func InitConfig(filePath string) (*CfgFileConfig, error) {
	confFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer confFile.Close()

	reader := bufio.NewReader(confFile)

	config := new(CfgFileConfig)
	config.configMap = make(map[string]map[string]string)

	section := ""
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		trimedLine := strings.TrimSpace(string(line))
		if strings.Index(trimedLine, "#") == 0 {
			continue
		}

		n1 := strings.Index(trimedLine, "[")
		n2 := strings.LastIndex(trimedLine, "]")
		if n1 > -1 && n2 > -1 && n2 > n1+1 {
			section = strings.TrimSpace(trimedLine[n1+1 : n2])
			config.configMap[section] = make(map[string]string)
			continue
		}

		if len(section) == 0 {
			continue
		}

		index := strings.Index(trimedLine, " #")
		if index > -1 {
			trimedLine = strings.TrimSpace(trimedLine[0:index])
		}

		index = strings.Index(trimedLine, "=")
		if index < 0 {
			continue
		}

		key := strings.TrimSpace(trimedLine[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(trimedLine[index+1:])

		if len(value) == 0 {
			continue
		}

		config.configMap[section][key] = value
	}

	return config, nil

}

// 从config中取配置数据
// section和key都存在时，返回对应value
// section或key不存在时，返回空字符串
func (c CfgFileConfig) Get(section string, key string) string {
	return c.configMap[section][key]
}

func (c CfgFileConfig) GetSection(section string) map[string]string {
	return c.configMap[section]
}

func (c CfgFileConfig) GetFloat64(section string, key string) (float64, error) {
	strValue := c.configMap[section][key]
	floatValue, err := strconv.ParseFloat(strValue, 64)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return floatValue, nil
}
