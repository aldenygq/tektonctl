package main

import (
	"os"
	
	"tektonctl/cmd"
	"tektonctl/server"
)

func main() {
	server.DbInit()
	err := server.InitMaster()
	if err != nil {
		os.Exit(-1)
	}
	cmd.Execute()
}
