package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/JojiiOfficial/SystemdGoService"
)

func install() {
	if SystemdGoService.SystemfileExists(serviceName) {
		fmt.Println("Service already exists")
		return
	}

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	service := SystemdGoService.NewDefaultService(serviceName, "Show DBus messages in cli", ex+" server")
	service.Service.User = "root"
	service.Service.Group = "root"
	service.Service.SuccessExitStatus = "2"
	service.Service.Restart = SystemdGoService.OnSuccess
	cpath, _ := filepath.Abs(ex)
	cpath, _ = path.Split(cpath)
	service.Service.WorkingDirectory = cpath
	service.Service.RestartSec = "20"
	err = service.Create()
	if err == nil {
		err := service.Enable()
		if err != nil {
			fmt.Println("Couldn't enable service: " + err.Error())
			return
		}
		err = service.Start()
		if err != nil {
			fmt.Println("Couldn't start service: " + err.Error())
			return
		}
		fmt.Println("Service installed and started")
	} else {
		fmt.Println("An error occured installitg the service: ", err.Error())
	}
}
