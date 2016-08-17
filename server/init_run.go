

package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

)

// Lets go ;-))
func Run(conf ConfigOpts) {

	Config = conf

	// Setup www router
	router := mux.NewRouter()

	router.HandleFunc("/", H_Home)

	router.HandleFunc("/ajax/info", AX_Info)

	router.HandleFunc("/ajax/csvlogs", AX_CSVLogFiles)

	//router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Serving on " + conf.HTTPAddress)
	http.Handle("/", router)
	err := http.ListenAndServe(conf.HTTPAddress , nil)
	if err != nil {
		fmt.Println("HTTP Error: ", err)
	}
}