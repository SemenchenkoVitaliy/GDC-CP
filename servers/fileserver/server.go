package fileserver

import (
	"github.com/SemenchenkoVitaliy/GDC-CP/netutils"
	"github.com/SemenchenkoVitaliy/GDC-CP/utils"
)

func Start() {
	sHTTP := netutils.NewHTTPServer()

	sHTTP.Handle("/", root)

	go sHTTP.Listen(utils.Config.IP, utils.Config.Port)

	cTCP := netutils.NewTCPClient()

	cTCP.Cert(utils.Config.TCP.CertPath, utils.Config.TCP.KeyPath)
	cTCP.Connect(utils.Config.TCP.IP, utils.Config.TCP.Port, tcpHandler)
}

func StartHashed() {
	sHTTP := netutils.NewHTTPServer()

	sHTTP.Handle("/", rootHashed)

	go sHTTP.Listen(utils.Config.IP, utils.Config.Port)

	cTCP := netutils.NewTCPClient()

	cTCP.Cert(utils.Config.TCP.CertPath, utils.Config.TCP.KeyPath)
	cTCP.Connect(utils.Config.TCP.IP, utils.Config.TCP.Port, tcpHandlerHashed)
}
