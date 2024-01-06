package main

import (
	"flag"

	"github.com/twelveeee/amis-admin-go/conf"
	"github.com/twelveeee/amis-admin-go/router"
)

var config = flag.String("conf", "./conf/app.yaml", "app config file")

func main() {
	flag.Parse()
	_, err := conf.ParseConfig(*config)
	if err != nil {
		panic(err)
	}

	router.Router()
}
