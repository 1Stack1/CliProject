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
	//"https://fofa.info/api/v1/search/next?&fields=link%2Ctitle%2Cstatus_code&key=8b9524c1ec8db699ae7b3803ac1ea19d&qbase64=dGl0bGU9IiI=&next=sVrlrWJ7W4pqfF9vuQCm6Rcm9aSZHvsA+tuk0wPgsiral7uLCt6ATg==&size=7"
}
