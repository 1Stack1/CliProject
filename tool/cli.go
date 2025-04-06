// main.go
package tool

import (
	"encoding/base64"
	"fmt"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
	"strconv"
	"sync"
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:    "key",
		Usage:   "FOFA API Key",
		EnvVars: []string{"FOFA_API_KEY"},
		//Required: true,
	},
	&cli.StringFlag{
		Name:    "query",
		Aliases: []string{"q"},
		Usage:   "Search query",
		//Required: true, // 标记为必填参数
	},
}

var configPath, configName, configType = "./config", "config", "yml"

func CliInit() *cli.App {
	app := &cli.App{
		//参数定义
		Flags: flags,
		//cli.exe触发
		Action: func(c *cli.Context) error {
			//调用FOFA
			query := c.String("query")
			apiKey := c.String("key")
			//读取配置文件
			v, err := ConfigInit(configPath, configName, configType)
			if err != nil {
				return err
			}
			count, numberConcurrency, err := readFromConfig(v)
			if err != nil {
				return err
			}
			transport, err := proxyConfig(v)
			if err != nil {
				return err
			}
			//循环调用fofa
			next := ""
			page := 0
			for true {
				var resBody string
				var size int
				if (count - 10*page) >= 10 {
					size = 10
				} else if count <= 10*page {
					break
				} else {
					size = count - 10*page
				}
				if next == "" {
					url := fofaUrlInit(apiKey, query, "", size)
					resBody = Fofa(url, transport)
				} else {
					pageContent := "&next=" + next
					url := fofaUrlInit(apiKey, query, pageContent, size)
					resBody = Fofa(url, transport)
				}
				//解析resBody
				response, err2 := FofaResJsonDes(resBody)
				if err2 != nil {
				}
				//fmt.Println("config_number", numberConcurrency)
				//根据url并发截图
				var wg sync.WaitGroup
				NewThreadPool(numberConcurrency)
				for i := 0; i < len(response.Results); i++ {
					AppendJob(func() {
						filePath, err := TakeScreenshot(response.Results[i][0])
						if err != nil {
							fmt.Printf("%s   %s   %s   %v\n", response.Results[i][0], response.Results[i][1], response.Results[i][2], err)
						} else {
							fmt.Printf("%s   %s   %s   %s\n", response.Results[i][0], response.Results[i][1], response.Results[i][2], filePath)
						}
					}, &wg)
				}
				wg.Wait()
				next = response.Next
				page++
			}
			return nil
		},
	}
	return app
}

func base64QueryArg(QueryContent string) string {
	QueryContent = "title=\"" + QueryContent + "\""
	encoded := base64.StdEncoding.EncodeToString([]byte(QueryContent))
	encoded = "&qbase64=" + encoded
	return encoded
}

/*
return 查询记录数，并发下载线程数，error
*/
func readFromConfig(v *viper.Viper) (int, int, error) {
	count, err1 := ConfigReadCount(v)
	if err1 != nil {
		return 0, 0, err1
	}
	numberConcurrency, err := ConfigReadConcurrency(v)
	if err != nil {
		return 0, 0, err
	}
	if numberConcurrency <= 0 {
		numberConcurrency = 1
	}
	return count, numberConcurrency, nil
}

func fofaUrlInit(userApi string, queryContent string, nextContent string, size int) string {
	base64QueryRes := base64QueryArg(queryContent)

	url := "https://fofa.info/api/v1/search/next?&fields=link%2Ctitle%2Cstatus_code&key="
	url += userApi
	url += base64QueryRes
	url += nextContent
	url += "&size=" + strconv.Itoa(size)
	return url
}
