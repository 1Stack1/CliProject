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
			//todo 调用FOFA

			//todo 根据url并发截图
			configPath := "./config"
			configName := "config"
			configType := "yml"
			v, err := ConfigInit(configPath, configName, configType)
			var wg sync.WaitGroup
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
			fmt.Println(number)
			NewThreadPool(number)
			for i := 0; i < 20; i++ {

				AppendJob(func() {
					filePath, err := TakeScreenshot("https://www.baidu.com")
					fmt.Println(filePath)
					if err != nil {
						fmt.Println(err)
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
