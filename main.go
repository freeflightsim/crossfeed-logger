

package main

import (
	"flag"
	"fmt"

	"github.com/freeflightsim/crossfeed-logger/server"
	"github.com/freeflightsim/crossfeed-logger/cfdb"
)





func main() {

	file_path := flag.String("-c", "./config.yaml-skel", "Path to yaml file")

	conf, err := server.LoadConfig(*file_path)
	if err != nil {
		fmt.Println("HTTP FATAL: " + err.Error())
		return
	}

	err = cfdb.Init(conf.Db.User, conf.Db.Password, conf.Db.Database )
	if err != nil {
		fmt.Println("POSTGRES FATAL: " + err.Error())
		return
	}
	server.Run()
}
