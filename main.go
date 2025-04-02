package main

import (
	"CliProject/tool"
	"log"
	"os"
)

func main() {

	app := tool.CliInit()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	//tool.TakeScreenshot("https://www.baidu.com")
}
