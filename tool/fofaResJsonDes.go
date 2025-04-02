package tool

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Response struct {
	Error            bool
	Consumed_fpoint  int
	Required_fpoints int
	Size             int
	Tip              string
	Page             int
	Mode             string
	Query            string
	Results          []string
}

func FofaResJsonDes(resJson string) (*Response, error) {
	v := viper.New()

	// 设置配置类型为JSON
	v.SetConfigType("json")

	// 从字符串读取配置
	if err := v.ReadConfig(strings.NewReader(resJson)); err != nil {
		return nil, fmt.Errorf("读取配置失败: %w", err)
	}

	// 解析到结构体
	var response Response
	if err := v.Unmarshal(&response); err != nil {
		return nil, fmt.Errorf("解析配置失败: %w", err)
	}

	return &response, nil
}
