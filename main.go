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
	//8b9524c1ec8db699ae7b3803ac1ea19d
}

/*
}*/
