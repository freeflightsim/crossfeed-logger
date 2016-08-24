

package main

import (
	"flag"


	"github.com/freeflightsim/crossfeed-logger/server"
	"github.com/freeflightsim/crossfeed-logger/cfdb"
)





func main() {

	file_path := flag.String("-c", "./config.yaml-skel", "Path to yaml file")

	conf, err := server.LoadConfig(*file_path)
	if err != nil {

		panic(err)
	}

	err = cfdb.Init(conf.Db.User, conf.Db.Password, conf.Db.Database )
	if err != nil {
		panic(err)
	}
	server.Run()
}
