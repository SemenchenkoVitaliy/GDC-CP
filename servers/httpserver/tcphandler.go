package httpserver

import (
	"github.com/SemenchenkoVitaliy/GDC-CP/netutils"
	"github.com/SemenchenkoVitaliy/GDC-CP/utils"
)

func tcpHandler(server netutils.Server) {
	err := server.Auth(netutils.AuthData{
		IP:   utils.Config.IP,
		Port: utils.Config.Port,
		Type: "http",
	})
	if err != nil {
		utils.LogCritical(err, "unable to authentifacate")
		return
	}
	for {
		server.Recieve()
	}

}
