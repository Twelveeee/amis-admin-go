package main

import (
	"flag"

	"github.com/twelveeee/amis-admin-go/conf"
	"github.com/twelveeee/amis-admin-go/dao/mdb"
	"github.com/twelveeee/amis-admin-go/router"
)

var configPath = flag.String("conf", "./conf/app.yaml", "app config file")

func main() {
	flag.Parse()
	err := conf.InitConfig(*configPath)
	if err != nil {
		panic(err)
	}
	mdb.InitMdb()

	router.Router()
}
