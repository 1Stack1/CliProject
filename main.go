package main

import (
	"CliProject/tool"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "key",
				Usage:   "FOFA API Key",
				EnvVars: []string{"FOFA_API_KEY"},
			},
			&cli.StringFlag{
				Name:    "email",
				Usage:   "FOFA Email",
				EnvVars: []string{"FOFA_EMAIL"},
			},
			&cli.StringFlag{
				Name:    "query",
				Usage:   "Search query",
				Aliases: []string{"q"},
			},
			// 其他参数...
		},
		Action: func(c *cli.Context) error {
			tool.TakeScreenshot("https://www.baidu.com")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		//log.Fatal(err)
	}
}
