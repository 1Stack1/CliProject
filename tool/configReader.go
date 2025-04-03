package tool

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Number string
}

var (
	lastUpdate   time.Time
	ignoredFiles = []string{".tmp", "~"} // 需要忽略的文件后缀
)

func ConfigInit(configPath string, configName string, configType string) (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName(configName)
	v.SetConfigType(configType)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		// 忽略临时文件
		for _, suffix := range ignoredFiles {
			if strings.HasSuffix(e.Name, suffix) {
				return
			}
		}

		// 防抖：500ms 内仅处理一次
		if time.Since(lastUpdate) < 500*time.Millisecond {
			return
		}
		lastUpdate = time.Now()

		fmt.Println("config file changed:", e.Name)
		if err := v.ReadInConfig(); err != nil {
			panic(err)
		}
	})
	return v, nil
}

/*
*
读取并发携程数
*/
func ConfigReadConcurrency(v *viper.Viper) (int, error) {
	numberStr := v.GetString("number_concurrency")
	if numberStr == "" {
		return -1, nil
	}
	numberConcurrency, err := strconv.Atoi(numberStr)
	if err != nil {
		return 0, fmt.Errorf("number转换为整数错误: %w", err)
	}
	return numberConcurrency, nil
}

/*
*
读取搜索数量
*/
func ConfigReadCount(v *viper.Viper) (int, error) {
	countStr := v.GetString("count")
	if countStr == "" {
		return 10, nil
	}
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return 0, fmt.Errorf("number转换为整数错误: %w", err)
	}
	return count, nil
}
