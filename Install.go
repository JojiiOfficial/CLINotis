package main

import (
	"bufio"
	"fmt"
	"os"

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

	reader := bufio.NewReader(os.Stdin)
	i, t := WaitForMessage("Which user is using the Desktop environment? ", reader)
	if i != 1 {
		os.Exit(1)
		return
	}

	service := SystemdGoService.NewDefaultService(serviceName, "Show DBus messages in cli", ex+" server")
	service.Service.User = t
	service.Service.Group = t
	service.Service.SuccessExitStatus = "2"
	service.Service.Restart = SystemdGoService.OnSuccess
	service.Service.RestartSec = "20"
	service.Install.Also = "dbus.service"
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
