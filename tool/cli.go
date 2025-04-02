// main.go
package tool

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"sync"
)

func CliInit() *cli.App {
	app := &cli.App{
		//参数定义
		Flags: []cli.Flag{
			/*&cli.StringFlag{
				Name:     "key",
				Usage:    "FOFA API Key",
				EnvVars:  []string{"FOFA_API_KEY"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "email",
				Usage:   "FOFA Email",
				EnvVars: []string{"FOFA_EMAIL"},
			},*/
			&cli.StringFlag{
				Name:    "query",
				Aliases: []string{"q"},
				Usage:   "Search query",
				//Required: true, // 标记为必填参数
			},
		},
		//cli.exe触发
		Action: func(c *cli.Context) error {
			//调用FOFA
			resBody := Fofa("8b9524c1ec8db699ae7b3803ac1ea19d")
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
					//todo select截图和3秒定时任务
					filePath, err := TakeScreenshot(response.Results[i])
					if err != nil {
						fmt.Printf("%s    %v\n", response.Results[i], err)
					} else {
						fmt.Printf("%s    %s\n", response.Results[i], filePath)
					}
				}, &wg)
			}
			wg.Wait()
			/*for true {

			}*/
			return nil
		},
	}
	return app
}
