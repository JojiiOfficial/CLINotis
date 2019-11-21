package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
)

//Config config file
type Config struct {
	LastCheck int64 `json:"lastCheck"`
}

func getHome() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return usr.HomeDir
}

func getConf() *Config {
	file := getHome() + "/.clinotis/conf.json"
	_, e := os.Stat(file)
	if e != nil {
		conf := Config{}
		dir, _ := path.Split(file)
		os.MkdirAll(dir, os.ModePerm)
		f, err := os.Create(file)
		if err != nil {
			fmt.Println("Error creating config: " + err.Error())
			os.Exit(1)
		}
		d, err := json.Marshal(conf)
		if err != nil {
			fmt.Println("Error creating config: " + err.Error())
			os.Exit(1)
		}
		f.Write(d)
		f.Close()
		return &Config{}
	}
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading config!" + err.Error())
		os.Exit(1)
	}
	var conf Config
	err = json.Unmarshal(b, &conf)
	if err != nil {
		fmt.Println("Error parsing json: " + err.Error())
		os.Exit(1)
	}
	return &conf
}

func (c *Config) save() {
	file := getHome() + "/.clinotis/conf.json"
	b, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Error saving json: " + err.Error())
		os.Exit(1)
	}
	err = ioutil.WriteFile(file, b, 0600)
	if err != nil {
		fmt.Println("Error saving json: " + err.Error())
		os.Exit(1)
	}
}
