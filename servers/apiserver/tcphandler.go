package apiserver

import (
	"github.com/SemenchenkoVitaliy/GDC-CP/netutils"
	"github.com/SemenchenkoVitaliy/GDC-CP/utils"
)

var mainServer netutils.Server

func tcpHandler(server netutils.Server) {
	mainServer = server

	err := server.Auth(netutils.AuthData{
		IP:   utils.Config.IP,
		Port: utils.Config.Port,
		Type: "api",
	})
	if err != nil {
		utils.LogCritical(err, "unable to authentifacate")
		return
	}
	for {
		server.Recieve()
	}
}
