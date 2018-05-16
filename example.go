package main

import (
	"fmt"
	"github.com/SemenchenkoVitaliy/GDC-CP/servers/apiserver"
	"github.com/SemenchenkoVitaliy/GDC-CP/servers/fileserver"
	"github.com/SemenchenkoVitaliy/GDC-CP/servers/httpserver"
	"github.com/SemenchenkoVitaliy/GDC-CP/servers/lbserver"
	"github.com/SemenchenkoVitaliy/GDC-CP/utils"
)

func main() {
	utils.LoadConfig("./configs/config.json")
	switch utils.ServerType {
	case "api":
		apiserver.Start()
	case "file":
		fileserver.Start()
	case "http":
		httpserver.Start()
	case "lb":
		lbserver.Start()
	default:
		utils.LogCritical(fmt.Errorf("Incorrect server type: "+utils.ServerType), "Launch server")
	}
}
