

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/freeflightsim/crossfeed-logger/cfdb"
	"github.com/freeflightsim/crossfeed-logger/server"

)



func main() {


	var conf server.ConfigOpts

	// The http server address
	conf.HTTPAddress = *flag.String("listen", "0.0.0.0:55667", "HTTP server address and port")

	// Local of csv dailies
	conf.CSVDir = *flag.String("csv_dir", "/home/ffs/crossfeed-dailies/csv", "Path to `csv` dir")
	if _, err := os.Stat(conf.CSVDir); os.IsNotExist(err) {
		fmt.Println("cvs dir `" + conf.CSVDir + "` does not exist")
		return
	}





	server.Run( conf )
}
