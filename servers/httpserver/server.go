package httpserver

import (
	"github.com/SemenchenkoVitaliy/GDC-CP/netutils"
	"github.com/SemenchenkoVitaliy/GDC-CP/utils"
)

func Start() {
	db.Connect(utils.Config.DB.User, utils.Config.DB.Password, utils.Config.DB.DbName, utils.Config.DB.IP, utils.Config.DB.Port)
	fsPublicUrl = utils.Config.PublicUrl
	publicUrl = utils.Config.PublicUrl

	loadTemplates()

	sHTTP := netutils.NewHTTPServer()

	sHTTP.AddRoute("/", "GET", rootGET)
	sHTTP.AddRoute("/search", "GET", searchGET)
	sHTTP.AddRoute("/admin", "GET", adminGET)
	sHTTP.AddRoute("/admin/{lection}", "GET", adminLectionGET)
	sHTTP.AddRoute("/{lection}", "GET", lectionGET)

	sHTTP.AddDir("/static/", "./static/")

	go sHTTP.Listen(utils.Config.IP, utils.Config.Port)

	cTCP := netutils.NewTCPClient()

	cTCP.Cert(utils.Config.TCP.CertPath, utils.Config.TCP.KeyPath)
	cTCP.Connect(utils.Config.TCP.IP, utils.Config.TCP.Port, tcpHandler)
}
