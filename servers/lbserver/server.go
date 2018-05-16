package lbserver

import (
	"github.com/SemenchenkoVitaliy/GDC-CP/balancer"
	"github.com/SemenchenkoVitaliy/GDC-CP/netutils"
	"github.com/SemenchenkoVitaliy/GDC-CP/utils"

	dbDriver "github.com/SemenchenkoVitaliy/GDC-CP/dbDriver/mongo"
)

var (
	httpServers = balancer.NewRoundRobin()
	apiServers  = balancer.NewRoundRobin()
	fileServers = balancer.NewDistributed()

	db = dbDriver.NewDatabaseFS()
)

func Start() {
	db.Connect(utils.Config.DB.User, utils.Config.DB.Password, utils.Config.DB.DbName, utils.Config.DB.IP, utils.Config.DB.Port)

	domainRouter := netutils.NewDomainRouter(root)
	domainRouter.AddSubdomain("api", api)
	domainRouter.AddSubdomain("img", file)

	go domainRouter.Listen(utils.Config.IP, utils.Config.Port)

	cTCP := netutils.NewTCPServer()
	cTCP.Cert(utils.Config.TCP.CertPath, utils.Config.TCP.KeyPath)
	cTCP.Listen(utils.Config.TCP.IP, utils.Config.TCP.Port, tcpHandler)
}
