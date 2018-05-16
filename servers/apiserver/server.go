package apiserver

import (
	"github.com/SemenchenkoVitaliy/GDC-CP/netutils"
	"github.com/SemenchenkoVitaliy/GDC-CP/utils"
)

func Start() {
	db.Connect(utils.Config.DB.User, utils.Config.DB.Password, utils.Config.DB.DbName, utils.Config.DB.IP, utils.Config.DB.Port)
	fsPublicUrl = utils.Config.PublicUrl
	publicUrl = utils.Config.PublicUrl

	sHTTP := netutils.NewHTTPServer()

	sHTTP.AddRoute("/", "GET", rootGET)
	sHTTP.AddRoute("/{lection}", "GET", lectionGET)

	sHTTP.AddRoute("/", "POST", rootPOST)
	sHTTP.AddRoute("/{lection}", "POST", lectionPOST)

	go sHTTP.Listen(utils.Config.IP, utils.Config.Port)

	cTCP := netutils.NewTCPClient()

	cTCP.Cert(utils.Config.TCP.CertPath, utils.Config.TCP.KeyPath)
	cTCP.Connect(utils.Config.TCP.IP, utils.Config.TCP.Port, tcpHandler)

}
