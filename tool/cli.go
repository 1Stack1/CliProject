// main.go
package tool

import (
	"encoding/base64"
	"fmt"
	"github.com/urfave/cli/v2"
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

func CliInit() *cli.App {
	app := &cli.App{
		//参数定义
		Flags: flags,
		//cli.exe触发
		Action: func(c *cli.Context) error {
			//调用FOFA
			query := c.String("query")
			base64QueryRes := base64QueryArg(query)
			apiKey := c.String("key")
			next := ""
			for true {
				var resBody string
				if next == "" {
					resBody = Fofa(apiKey, base64QueryRes, "")
				} else {
					pageContent := "&next=" + next
					resBody = Fofa(apiKey, base64QueryRes, pageContent)
				}
				//解析resBody
				response, err2 := FofaResJsonDes(resBody)
				if err2 != nil {
				}
				//读取配置文件
				configPath, configName, configType := "./config", "config", "yml"
				v, err := ConfigInit(configPath, configName, configType)
				if err != nil {
					return err
				}
				number, err := ConfigRead(v)
				if err != nil {
					return err
				}
				if number <= 0 {
					number = 1
				}
				fmt.Println("config_number", number)
				//根据url并发截图
				var wg sync.WaitGroup
				NewThreadPool(number)
				for i := 0; i < len(response.Results); i++ {
					AppendJob(func() {
						filePath, err := TakeScreenshot(response.Results[i][0])
						if err != nil {
							fmt.Printf("%-70s   %-100s   %s   %v\n", response.Results[i][0], response.Results[i][1], response.Results[i][2], err)
						} else {
							fmt.Printf("%-70s   %-100s   %s   %s\n", response.Results[i][0], response.Results[i][1], response.Results[i][2], filePath)
						}
					}, &wg)
				}
				wg.Wait()
				next = response.Next
				//todo 判断结束
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
