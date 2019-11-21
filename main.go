package main

import (
	"fmt"
	"io/ioutil"
	"os"

	dbus "github.com/godbus/dbus"
)

var messageFile = "/tmp/snaOTSPrdc"

func main() {
	conf := getConf()
	if len(os.Args) > 1 {
		arg := os.Args[1]
		if arg == "server" {
			msgServer(conf)
		} else if arg == "check" {
			check(conf)
		}
	} else {
		showMessages(conf)
	}
}
func msgServer(conf *Config) {
	conn, err := dbus.SessionBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}

	var rules = []string{
		"type='signal',member='Notify',path='/org/freedesktop/Notifications',interface='org.freedesktop.Notifications'",
		"type='method_call',member='Notify',path='/org/freedesktop/Notifications',interface='org.freedesktop.Notifications'",
		"type='method_return',member='Notify',path='/org/freedesktop/Notifications',interface='org.freedesktop.Notifications'",
		"type='error',member='Notify',path='/org/freedesktop/Notifications',interface='org.freedesktop.Notifications'",
	}
	var flag uint = 0

	call := conn.BusObject().Call("org.freedesktop.DBus.Monitoring.BecomeMonitor", 0, rules, flag)
	if call.Err != nil {
		fmt.Println("Error calling bus!" + err.Error())
		os.Exit(1)
	}

	c := make(chan *dbus.Message, 10)
	conn.Eavesdrop(c)
	for v := range c {
		if len(v.Body) < 4 {
			continue
		}
		title := v.Body[3].(string)
		message := v.Body[4].(string)
		writeMessage(title+": "+message, messageFile)
	}
}

func showMessages(conf *Config) {
	b, err := ioutil.ReadFile(messageFile)
	if err != nil {
		fmt.Println("Error reading file: " + err.Error())
		return
	}
	fmt.Print(string(b))
	os.Truncate(messageFile, 0)
}

func writeMessage(message, file string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(message + "\n"); err != nil {
		panic(err)
	}
}

func check(conf *Config) {
	d, err := os.Stat(messageFile)
	if err == nil && d.Size() > 0 && conf.LastCheck < d.ModTime().Unix() {
		fmt.Println(1)
		conf.LastCheck = d.ModTime().Unix()
		conf.save()
	}
}
